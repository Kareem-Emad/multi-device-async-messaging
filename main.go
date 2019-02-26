package main

import (
	"fmt"
	"strconv"

	zmq "github.com/alecthomas/gozmq"
)

const connectionURL string = "tcp://127.0.0.1:"

func main() {

	var cs string
	fmt.Println("[] Do you want to start a client/server?(c/s)")
	fmt.Scanf("%s", &cs)

	if cs == "s" {
		var port int
		println("[Server] What port to listen at ?")
		fmt.Scanf("%d", &port)
		context, _ := zmq.NewContext()
		socket, _ := context.NewSocket(zmq.REQ)
		portString := strconv.Itoa(port)
		socket.Bind(connectionURL + portString)
		for {
			msg, _ := socket.Recv(0)
			println("[Server] Got", string(msg))
			socket.Send(msg, 0)
		}
	}

	if cs == "c" {
		context, _ := zmq.NewContext()
		socket, _ := context.NewSocket(zmq.REQ)

		var port int
		println("[Client] What port(s) to listen at ?(Enter -1 to skip after entering first port)")

		fmt.Scanf("%d", &port)
		portString := strconv.Itoa(port)
		socket.Connect(connectionURL + portString)

		for port != -1 {
			fmt.Scanf("%d", &port)
			portString = strconv.Itoa(port)
			socket.Connect(connectionURL + portString)
		}

		var idxStart, idxEnd int

		for {
			fmt.Println("[Client] What indicies to send to process ?")
			fmt.Scanf("%d %d", &idxStart, &idxEnd)
			println("[Client] Preparing")

			msg := fmt.Sprintf("msg %d %d", idxStart, idxEnd)
			socket.Send([]byte(msg), 0)
			println("[Client] Sending", msg)
			socket.Recv(0)
		}

	}

	fmt.Println("Msh na2sa habal 3al sobh ya 3m enta :V")

}
