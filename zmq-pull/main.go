//
//  Binds PULL socket to tcp://localhost:5000
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	//	"time"
)

func main() {

	mpull, _ := zmq.NewSocket(zmq.PULL)
	defer mpull.Close()
	mcon := mpull.Connect("tcp://localhost:5000")
	if mcon != nil {
		fmt.Printf("connect fail.")
	}

	for {

		recv, _ := mpull.Recv(0)
		fmt.Printf("Received [%s]\n", recv)

	}

}
