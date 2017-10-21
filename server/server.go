package server

import (
	"database/sql"
	"fmt"
	"net"
	"runtime"
	"sync"

	"github.com/rfyiamcool/go_pubsub/config"
)

const (
	KeyRecordTableName = "__idgo__"
)

type Server struct {
	cfg *config.Config

	listener net.Listener
	db       *sql.DB
	//keyGeneratorMap map[string]*MySQLIdGenerator
	sync.RWMutex
	running bool
}

func NewServer(c *config.Config) (*Server, error) {
	s := new(Server)
	s.cfg = c

	var err error

	netProto := "tcp"
	s.listener, err = net.Listen(netProto, s.cfg.Addr)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) Init() error {
	return nil
}

func (s *Server) Serve() error {
	s.running = true
	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		go s.onConn(conn)
	}
	return nil
}

func (s *Server) onConn(conn net.Conn) error {
	defer func() {
		clientAddr := conn.RemoteAddr().String()
		r := recover()
		if err, ok := r.(error); ok {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			fmt.Println("server", "onConn", "error", 0,
				"remoteAddr", clientAddr,
				"stack", string(buf),
				"err", err.Error(),
			)
			reply := &ErrorReply{
				message: err.Error(),
			}
			reply.WriteTo(conn)
		}
		conn.Close()
	}()

	for {
		request, err := NewRequest(conn)
		if err != nil {
			return err
		}

		reply := s.ServeRequest(request)
		if _, err := reply.WriteTo(conn); err != nil {
			return err
		}

	}
	return nil
}

func (s *Server) ServeRequest(request *Request) Reply {
	switch request.Command {
	case "PUBLISH":
		return s.handlePub(request)
	case "SUBSCRIBE":
		return s.handleSub(request)
	case "DEL":
		return s.handleDel(request)
	case "SELECT":
		return s.handleSelect(request)
	default:
		return ErrMethodNotSupported
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
	fmt.Println("server", "close", "server closed!", 0)
}

func (s *Server) IsKeyExist(key string) (bool, error) {
	return true, nil
}

func (s *Server) GetKey(key string) (string, error) {
	return "", nil
}

func (s *Server) SetKey(key string) error {
	return nil
}

func (s *Server) DelKey(key string) error {
	return nil
}
