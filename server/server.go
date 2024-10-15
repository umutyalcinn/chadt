package server

import (
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln net.Listener

	quitch chan struct{}
}

type Message struct {

}

func NewServer (listenAddr string) *Server {
	return &Server {
		listenAddr: listenAddr,
			quitch: make(chan struct{}),
	}
}

func (s* Server) Start() {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		log.Fatal("Failed to create tcp socket!")
	}

	defer ln.Close()
	s.ln = ln

	log.Println("Server started. Listening", s.listenAddr)

	go s.acceptLoop()

	<- s.quitch
	close(s.quitch)
}

func (s* Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("connection received from %s", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buff := make([]byte, 2048)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("read error: ", err)
		}

		msg := buff[:n]
		log.Println(string(msg))
	}
}
