//
//  Binds PUSH socket to tcp://localhost:5000

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	//  Socket to send messages on
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5000")

	//  Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	//  Send 100 tasks
	total_msec := 0
	for task_nbr := 0; task_nbr < 100; task_nbr++ {
		workload := task_nbr //rand.Intn(100) + 1
		total_msec += workload
		s := fmt.Sprintf("push %d", workload)
		sender.Send(s, 0)
		time.Sleep(time.Second)
	}
	fmt.Println("Total expected cost:", time.Duration(total_msec)*time.Millisecond)
	time.Sleep(time.Second) //  Give 0MQ time to deliver

}
