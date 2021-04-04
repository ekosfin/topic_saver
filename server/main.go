package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//Server port settings
const (
	PORT = "8888"
	TYPE = "tcp"
)

var data Data

type Message struct {
	Topic string
	Note  string
	Text  string
}

type GetTopic struct {
	Name  string
	Notes []GetNote
}

type GetNote struct {
	Name      string
	Text      string
	Timestamp string
}

type API int

func (a *API) AddNote(send Message, reply *Message) error {
	fmt.Printf("Client appends data for topic: %s, note: %s and text: %s", send.Topic, send.Note, send.Text)
	data = AppendToData(data, send)
	WriteXMLfile(data)
	return nil
}

func (a *API) GetNote(topic string, reply *GetTopic) error {
	fmt.Println("Client gets note for", topic)
	err := FindTopic(data, topic, reply)
	return err
}

func main() {
	data = OpenXMLfile()

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error in API registering:", err)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen(TYPE, ":"+PORT)

	if err != nil {
		log.Fatal("Listener error", err)
	}
	defer listener.Close()
	log.Printf("Serving rpc on port %s", PORT)
	http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}

}
