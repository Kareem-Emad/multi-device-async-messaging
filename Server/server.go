package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	zmq "github.com/alecthomas/gozmq"
)

const connectionURL string = "tcp://127.0.0.1:"

var arr [100100]int
var size int

//ReadFile Reada a file specified by user
func ReadFile(fileName string) {
	inputFile, err := os.Open(fileName + ".txt")
	defer inputFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inputFile)
	size = 0

	for scanner.Scan() {
		num := scanner.Text()
		arr[size], _ = strconv.Atoi(num)
		size++
	}
}

//SumRange Sums a range specified by [l, r]
func SumRange(l int, r int) int64 {
	var sum int64

	for i := l; i <= r; i++ {
		sum += int64(arr[i])
	}

	return sum
}

//ServerSetup Sets up the server and does server logic
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
			query := strings.Fields(recString)
			l, _ := strconv.Atoi(query[0])
			r, _ := strconv.Atoi(query[1])

			sum := SumRange(l, r)

			res := fmt.Sprintf("%d", sum)

			fmt.Printf("Sum of numbers in [%d, %d] = %d\n", l, r, sum)

			socket.Send([]byte(res), 0)
		} else {
			socket.Send(msg, 0)
		}
	}
}

func main() {
	var fileName string
	fmt.Println("Enter input file name")
	fmt.Scanf("%s", &fileName)

	ReadFile(fileName)

	ServerSetup()
}
