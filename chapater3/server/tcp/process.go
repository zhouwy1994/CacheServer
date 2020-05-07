package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k, v, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, s.Set(k, v), conn)
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}

	v, err := s.Get(k)
	return sendResponse(v, err, conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, s.Del(k), conn)
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close connection due to error:", err)
			}
			return
		}

		if op == 'S' {
			err = s.set(conn, r)
		} else if op == 'G' {
			err = s.get(conn, r)
		} else if op == 'D' {
			err = s.del(conn, r)
		} else {
			log.Println("close connection due to invalid op:", op)
			return
		}

		if err != nil {
			log.Println("close connection due to:", err)
			return
		}
	}
}
