package main

import (
	"fmt"
	"strconv"

	zmq "github.com/alecthomas/gozmq"
)

const connectionURL string = "tcp://127.0.0.1:"

//ClientSetup Sets up the client and does logic
func ClientSetup() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)

	var port int
	println("[Client] What port(s) to listen to?")

	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &port)
		portString := strconv.Itoa(port)
		socket.Connect(connectionURL + portString)
	}

	var idxStart, idxEnd int

	for {
		fmt.Println("[Client] Enter interval?")
		fmt.Scanf("%d %d", &idxStart, &idxEnd)

		msg := fmt.Sprintf("%d %d", idxStart, idxEnd)
		socket.Send([]byte(msg), 0)
		println("[Client] Message sent.")

		recievedMsg, _ := socket.Recv(0)

		fmt.Println("Sum = " + string(recievedMsg))
	}
}

func main() {
	ClientSetup()
}
