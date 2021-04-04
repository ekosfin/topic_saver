package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Data struct {
	XMLName xml.Name `xml:"data"`
	Topic   []Topic  `xml:"topic"`
}

type Topic struct {
	XMLName xml.Name `xml:"topic"`
	Name    string   `xml:"name,attr"`
	Note    []Note   `xml:"note"`
}

type Note struct {
	XMLName   xml.Name `xml:"note"`
	Name      string   `xml:"name,attr"`
	Text      string   `xml:"text"`
	Timestamp string   `xml:"timestamp"`
}

//Reading the DB
func OpenXMLfile() Data {
	xmlFile, err := os.Open("db2.xml")
	if err != nil {
		log.Fatal("XML File open:", err)
	}
	fmt.Println("Successfully Opened db.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var data Data
	err = xml.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal("XML unmarshal:", err)
	}
	return data
}

//Writing out the db
func WriteXMLfile(data Data) {
	file, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal("XML marshal:", err)
	}
	err = ioutil.WriteFile("db2.xml", file, 0644)
	if err != nil {
		log.Fatal("XML file write:", err)
	}
}

//For appending to the db
func AppendToData(data Data, message Message) Data {
	timeNow := time.Now().Format(time.RFC822)
	note := Note{
		Timestamp: timeNow,
		Text:      message.Text,
		Name:      message.Note,
	}
	for i := range data.Topic {
		if data.Topic[i].Name == message.Topic {
			data.Topic[i].Note = append(data.Topic[i].Note, note)
			return data
		}
	}
	var noteArr []Note
	noteArr = append(noteArr, note)
	temp := Topic{
		Name: message.Topic,
		Note: noteArr,
	}
	data.Topic = append(data.Topic, temp)

	return data
}

// For finding the topic
func FindTopic(data Data, topic string, reply *GetTopic) error {

	for i := range data.Topic {
		if data.Topic[i].Name == topic {
			var tempNotes []GetNote
			for j := range data.Topic[i].Note {
				tempNote := GetNote{
					Name:      data.Topic[i].Note[j].Name,
					Text:      data.Topic[i].Note[j].Text,
					Timestamp: data.Topic[i].Note[j].Timestamp,
				}
				tempNotes = append(tempNotes, tempNote)
			}
			reply.Name = data.Topic[i].Name
			reply.Notes = tempNotes
			return nil
		}
	}
	reply.Name = "No topic found"
	return nil
}
