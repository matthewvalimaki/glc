package glc

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

type GlcZmq struct {
	Socket *zmq.Socket
}

func BindProxy(frontend string, backend string, filter string) {
	// create XSUB for publishers to connect to
	xSub, _ := zmq.NewSocket(zmq.XSUB)
	defer xSub.Close()
	
	xSub.SetSubscribe(filter)
	
	err := xSub.Bind(frontend)
	if err != nil {
        log.Fatal(err)
    }
	
	// create XPUB for subscribers to connect to
	xPub, _ := zmq.NewSocket(zmq.XPUB)
	defer xPub.Close()
	err = xPub.Bind(backend)
	if err != nil {
        log.Fatal(err)
    }

	err = zmq.Proxy(xSub, xPub, nil)
	log.Fatalln("Proxy interrupted:", err)
}

func BindPublisher(endpoint string) *GlcZmq {
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
	
	return &GlcZmq{socket}
}

func BindSubscriber(endpoint string, filter string, callback func(message string)) *GlcZmq {
	socket, err := zmq.NewSocket(zmq.PUB)
	defer socket.Close()
	
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
	
	return &GlcZmq{socket}
}

func PublishMessage(message string, glcZqm *GlcZmq) {
	glcZqm.Socket.Send(message, zmq.DONTWAIT)
}