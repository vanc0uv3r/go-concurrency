package main

import (
	"fmt"
	"net"
	"github.com/vanc0uv3r/go-concurrency/cmd/storage/lexer"
	"github.com/vanc0uv3r/go-concurrency/cmd/storage/engine"
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
		fmt.Printf("Cant listent server: %s", err.Error())
	}

	server := NewServer(listener, serverConfig)
	server.serve(listener)
}

func (s *Server) serve(l net.Listener) {
	fmt.Println("Ready to serve")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Cant accept client: %s", err.Error())
		}

		if s.checkActiveConnectionsLimit() {
			conn.Write([]byte("Too much conn\n"))
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
	lex := lexer.NewLex()
	engine := engine.NewEngine()

	fmt.Println("Got client")
	
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("REQUEST FROM CLIENT: %s", string(buff[:n]))

		lex.Analyze(buff[:n])
		engine.SetLexemes(lex.GetLexemes())
		res, err := engine.Execute()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("result of %s operation: %s\n", engine.GetCommandName(), res)
		}
		lex.ClearLexer()
	}
}

func (s *Server) closeConn(conn net.Conn) {
	s.activeConnections--
	conn.Close()
}