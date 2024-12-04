package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vanc0uv3r/go-concurrency/cmd/storage/engine"
	"github.com/vanc0uv3r/go-concurrency/cmd/storage/lexer"
)

type Server struct {
	listener net.Listener
	activeConnections uint16
	config Config
}

func NewServer(l net.Listener, c Config) *Server{
	return &Server{
		listener: l,
		activeConnections: 0,
		config: c,
	}
}


func main() {
	serverConfig := importConfig()
	listenerAddr := serverConfig.Addr + ":" + serverConfig.Port

	listener, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		log.Fatal("Cant listent server: ", err.Error())
	}

	server := NewServer(listener, serverConfig)
	log.Println("Ready to serve")
	server.serve(listener)
}

func (s *Server) serve(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Cant accept client: %s", err.Error())
		}

		if s.checkActiveConnectionsLimit() {
			responseClient(conn, "Too much conn")
			conn.Close()
		}

		go s.handleUser(conn)
	}
}

func (s *Server) checkActiveConnectionsLimit() bool {
	return s.activeConnections + 1 > s.config.MaxConnections
}


func (s *Server) handleUser(conn net.Conn) {
	s.activeConnections++
	defer s.closeConn(conn)
	log.Println("Got client")
	
	serveClient(conn)
}

func serveClient(conn net.Conn) {
	var msgClient string
	lex := lexer.NewLex()
	engine := engine.NewEngine()

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Println(err.Error())
			break
		}
		log.Printf("Request from client: %s\n", string(buff[:n]))

		lex.Analyze(buff[:n])
		lexemes := lex.GetLexemes()
		engine.SetLexemes(lexemes)
		res, err := engine.Execute()

		if err != nil {
			msgClient = fmt.Sprintf("Error while executing command: %s", err.Error())
		} else {
			msgClient = fmt.Sprintf("result of %s operation: %s\n", engine.GetCommandName(), res)
		}
		responseClient(conn, msgClient)
		lex.ClearLexer()
	}
}

func (s *Server) closeConn(conn net.Conn) {
	s.activeConnections--
	conn.Close()
}

func responseClient(conn net.Conn, msg string) {
	bytesMsg := []byte(msg + "\n")
	_, err := conn.Write(bytesMsg)
	if err != nil {
		log.Println("Error while sending response to client", err.Error())
	}
}