package main

import (
	"log"
	"time"
)

var reqCh = make(chan *Req)
var respCh = make(chan *Resp)

type Req struct {
	question string
	respChan chan *Resp
}

type Resp struct {
	answer string
	responseCode int
}

func start(handler func(*Req), reqCh <- chan *Req) {
	for {
		go handler(<-reqCh)
	}
}

func handler(req *Req) {
	req.respChan <- &Resp{"\n"+req.question + " \n Mama: I don’t have money ... Empty pocket $ 0.00 :( ",404}
}

func createRequestChannels(){
	reqCh <- &Req{"Baby: Mama ... Can you buy me a Cola ... ", respCh}
	time.Sleep(2 * time.Second)
	reqCh <- &Req{"Baby: Mama ... Can you buy me a Cola ... ", respCh}
	time.Sleep(2 * time.Second)
	reqCh <- &Req{"Baby: Mama ... I need a Cola ... ", respCh}
}

func main() {

	go createRequestChannels()

	go start(handler,reqCh)

	quitCh := time.After(6 * time.Second)

	for {
		select {
			case resp := <- respCh:
				log.Println(resp.answer)
			case <- quitCh:
				log.Println("TIME OUT -> :)  Do not cry ... Will buy you a Cola ...")
				return
		}
	}
	
}
