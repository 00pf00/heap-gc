package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"runtime/trace"
	"time"
)

type Msg struct {
	Number uint64
	Uid    string
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	ch := make(chan *Msg,100)
	go Input(ch)

	go Output(ch)

	time.Sleep(100*time.Second)


}

func Input(input chan *Msg) {
	var num uint64
	num = 0
	for true {
		input <- &Msg{
			Number: num,
			Uid:    uuid.Must(uuid.NewV4()).String(),
		}
		num += 1
		time.Sleep(1 * time.Second)
	}
}

func Output(output chan *Msg) {
	for {
		msg := <-output
		fmt.Printf("number = %v uuid = %v\n", msg.Number, msg.Uid)

	}
}
