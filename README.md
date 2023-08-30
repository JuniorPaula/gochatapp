# Go WebSocket Chat Application

The **Go WebSocket Chat Application** is a project that demonstrates the implementation of a chat application using WebSockets in the Go programming language. It leverages channels to manage socket connections and facilitate real-time communication between users.

## Features

- **Real-Time Chat**: Users can connect to the chat server using a WebSocket connection and send messages in real time.

- **Message Broadcasting**: Messages sent by one user are broadcasted to all connected users, enabling group conversations.

- **Channel Management**: Go channels are used to manage WebSocket connections, allowing seamless communication between the server and clients.

## Technologies Used

- Go: The backend is developed using the Go programming language, known for its simplicity and efficiency.

- WebSockets: WebSockets provide a full-duplex communication channel over a single TCP connection.

## How to Run

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Run `go run ./cmd/web/*.go` to start the chat server.
4. Access the application in a web browser using the provided URL.

## Usage

1. Open the chat application in multiple browser tabs or windows.
2. Enter a username to identify yourself.
3. Start sending and receiving messages in real time.
4. Observe how messages are broadcasted to all connected users.

## Contribution

Contributions to enhance and improve the project are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Contact

If you have any questions or suggestions, feel free to contact us at [luke.junnior@icloud.com](mailto:luke.junnior@icloud.com).

Thank you for using the Go WebSocket Chat Application to explore real-time communication using WebSockets and channels!