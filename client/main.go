package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

//connecting settings
const (
	PORT = "8888"
	TYPE = "tcp"
)

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

func printMenu() {
	fmt.Println("-----Main menu-----")
	fmt.Println("exit: to exit")
	fmt.Println("send: for sending topic")
	fmt.Println("view: for viewing topic")
	fmt.Print("> ")
}

func main() {

	fmt.Println("-----Topic saver-----")

	client, err := rpc.DialHTTP(TYPE, "localhost:"+PORT)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	log.Println("Connected to server.")

	scanner := bufio.NewScanner(os.Stdin)
	printMenu()
	for scanner.Scan() {
		text := scanner.Text()
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("send", text) == 0 {
			fmt.Println("-----Sending topic-----")
			fmt.Printf("Topic name: ")
			scanner.Scan()
			topic := scanner.Text()
			fmt.Printf("Note name: ")
			scanner.Scan()
			note := scanner.Text()
			fmt.Printf("Topic text: ")
			scanner.Scan()
			ttext := scanner.Text()
			msg := Message{
				Topic: topic,
				Note:  note,
				Text:  ttext,
			}
			fmt.Println("Sending topic....")
			err := client.Call("API.AddNote", msg, &msg)
			if err != nil {
				fmt.Println("Error sending topic:", err)
				printMenu()
				continue
			}
			fmt.Println("Topic Sent.")
		} else if strings.Compare("view", text) == 0 {
			fmt.Println("-----Get topic-----")
			fmt.Printf("Topic name: ")
			scanner.Scan()
			topic := scanner.Text()
			var getTopic GetTopic
			fmt.Println("Fetching topic")
			err := client.Call("API.GetNote", topic, &getTopic)
			if err != nil {
				fmt.Println("Error fetching topic:", err)
				printMenu()
				continue
			}
			fmt.Println("Fetch complete.")
			if getTopic.Name == "No topic found" {
				fmt.Println("No such topics exist")
				printMenu()
				continue
			} else {
				fmt.Println("Topic found:")
				fmt.Printf("Topic: %s\n", getTopic.Name)
				for i := range getTopic.Notes {
					fmt.Println("----")
					fmt.Printf("-Note: %s\n", getTopic.Notes[i].Name)
					fmt.Printf("-Text: %s\n", getTopic.Notes[i].Text)
					fmt.Printf("-Timestamp: %s\n", getTopic.Notes[i].Timestamp)
				}
			}
		} else if strings.Compare("exit", text) == 0 {
			fmt.Println("Exiting")
			break
		} else {
			fmt.Println("Unkown command")
		}
		printMenu()

	}

}
