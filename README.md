# Topic saver

Distributed Systems assignment 2

This assignments purpose is to create an distributed system of an client-server relationship where the client can evoke RPC to send the server topics with text that the server then save to the mockup xml database. The requirements demand that the server is able to handle multiple clients and detect if the topic is already in the db and decide if it should append or create a new topic.

### Sources for the code:

- https://tutorialedge.net/golang/parsing-xml-with-golang/
- https://golang.org/pkg/net/rpc/
- https://github.com/tensor-programming/go-basic-rpc
- https://tutorialedge.net/golang/reading-console-input-golang/

### Instructions to run

This project consists of two parts: the client and the server. To compile them you can call go build in each directory and get an executable that you can run.
