package main

import (
	"net"
	"bytes"
	"math/rand"
	"log"
	"bufio"
)

const (
	debug = true
)

type Server struct {
}

type Client struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListenAndServe(host string) {

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("unable to listen on address ",host, ": ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting client connection: ", err)
			continue
		}

		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {

	client := NewClient()

	client.handshake(&conn)

}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) handshake(conn *net.Conn) {
	connr := bufio.NewReader(*conn)
	connw := bufio.NewWriter(*conn)

	// lets read the version number
	version, _ := connr.ReadByte()

	if version != 3 {
		log.Println("invalid client version number detected: ", version)
	}

	if debug {
		log.Println("Received Client C0 Handshake")
	}

	// read client gibberish
	cgibberish := make([]byte, 1536)

	_, err := connr.Read(cgibberish)
	if err != nil {
		log.Println("unable to read client gibberish: ", err)
	}

	if debug {
		log.Println("Received Client C1 Gibberish")
	}

	// generate our own gibberish
	sgibberish := make([]byte, 1536)
	for i:=0; i<1536; i++ {
		sgibberish[i] = byte(rand.Intn(256))
	}

	// send our S0 packet
	tmp := make([]byte, 1)
	tmp[0] = 3
	connw.Write(tmp)


	// now send our S1 and S2 packets
	connw.Write(sgibberish)
	connw.Write(cgibberish)
	connw.Flush()

	// receive the clients final C2 packet
	_, err = connr.Read(cgibberish)
	if err != nil {
		log.Println("invalid client C2 handshake received")
	}

	if bytes.Compare(cgibberish, sgibberish) != 0 {
		log.Println("client returned invalid gibberish in handshake")
	}

	if debug {
		log.Println("Received Client C2 Confirmation")
	}

	if debug {
		log.Println("Client/Server Handshake Complete")
	}
}
