package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

// views is a set of jet templates.
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default.
		return true
	},
}

// Home is the handler for the home page route.
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

// WebSocketConnection is a wrapper around the websocket.Conn type.
type WebSocketConnection struct {
	*websocket.Conn
}

// WsJSONResponse is the response sent to the client when a websocket event occurs.
type WsJSONResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

// WsPayload is the payload sent to the server when a websocket event occurs.
type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"conn"`
}

// WsEndpoint upgrades connection to websocket.
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected to Endpoint")

	var response WsJSONResponse
	response.Message = `<em><small>Connected to Server</small></em>`

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)
}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToWsChannel() {
	var response WsJSONResponse

	for {
		evt := <-wsChan

		switch evt.Action {
		case "username":
			clients[evt.Conn] = evt.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(clients, evt.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf(`
				<p><span class='text-secondary username'>%s: </span><span class='message'>%s</span></p>
				`,
				evt.Username, evt.Message,
			)
			broadcastToAll(response)
		}
	}
}

func getUserList() []string {
	var userList []string

	for _, v := range clients {
		if v != "" {
			userList = append(userList, v)
		}
	}

	sort.Strings(userList)
	return userList
}

func broadcastToAll(response WsJSONResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket error")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

// renderPage renders a jet template.
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
