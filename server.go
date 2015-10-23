package main

import (
	"net"
"log"
	"bufio"
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

}
