package glc

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func BindProxy(frontend string, backend string) {
	// create XSUB for publishers to connect to
	xSub, _ := zmq4.NewSocket(zmq4.XSUB)
	defer xSub.Close()
	err := xSub.Bind(frontend)
	if err != nil {
        log.Fatal(err)
    }
	
	// create XPUB for subscribers to connect to
	xPub, _ := zmq4.NewSocket(zmq4.XPUB)
	defer xPub.Close()
	err = xPub.Bind(backend)
	if err != nil {
        log.Fatal(err)
    }

	// create the proxy
	go proxy(xSub, xPub)
}

func proxy(frontend *zmq.Socket, backend *zmq.Socket) {
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}

func BindPublisher(endpoint string) *zmq.Socket {
	socket, err := zmq.NewSocket(zmq.PUB)
	defer socket.Close()
	
	if err != nil {
        log.Fatal(err)
    }
	
	err = socket.Connect(endpoint)
	
	if err != nil {
        log.Fatal(err)
    }
	
	return socket
}

func BindSubscriber(endpoint string, filter string, callback func) *zmq.Socket {
	socket, err := zmq.NewSocket(zmq.PUB)
	defer socket.Close()
	
	socket.SetSubscribe(filter)
	
	if err != nil {
        log.Fatal(err)
    }
	
	err = socket.Connect(endpoint)
	
	if err != nil {
        log.Fatal(err)
    }
	
	for {
		msg, _ := socket.Recv(zmq.DONTWAIT)

		if msg != "" {
			callback(msg)
		}
	}
	
	return socket
}

func PublishMessage(message string, socket *zmq.Socket) {
	socket.Send(message, zmq.DONTWAIT)
}