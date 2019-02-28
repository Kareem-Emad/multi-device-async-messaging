package main

import (
	"fmt"
	"strconv"
	"strings"

	zmq "github.com/alecthomas/gozmq"
)

const connectionURL string = "tcp://127.0.0.1:"

func QueryUser() {

}

func ServerSetup() {
	var port int
	println("[Server] What port to listen to?")
	fmt.Scanf("%d", &port)

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	portString := strconv.Itoa(port)
	socket.Bind(connectionURL + portString)

	for {
		msg, _ := socket.Recv(0)
		recString := string(msg)

		if recString != "" {
			//println("Here")
			query := strings.Fields(recString)
			l, _ := strconv.Atoi(query[1])
			r, _ := strconv.Atoi(query[2])
			sum := ((l + r) / 2) * (r - l + 1)
			res := fmt.Sprintf("Sum = %d", sum)
			fmt.Printf("Sum of numbers in [%d, %d] = %d\n", l, r, sum)
			socket.Send([]byte(res), 0)
		} else {
			socket.Send(msg, 0)
		}
	}
}

func ClinetSetup() {
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

	//socket.Send([]byte(""), 0)

	var idxStart, idxEnd int

	for {
		fmt.Println("[Client] What indicies to send to process?")
		fmt.Scanf("%d %d", &idxStart, &idxEnd)
		println("[Client] Preparing")

		msg := fmt.Sprintf("msg %d %d", idxStart, idxEnd)
		socket.Send([]byte(msg), 0)
		println("[Client] Sending", msg)

		recievedMsg, _ := socket.Recv(0)

		fmt.Println(string(recievedMsg))
	}
}

func main() {

	var cs string
	fmt.Println("[] Do you want to start a client/server?(c/s)")
	fmt.Scanf("%s", &cs)

	if cs == "s" {
		ServerSetup()
	}

	if cs == "c" {
		ClinetSetup()
	}

	fmt.Println("Msh na2sa habal 3al sobh ya 3m enta :V")

}
