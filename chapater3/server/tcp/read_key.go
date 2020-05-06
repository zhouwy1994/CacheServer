package tcp

import (
	"bufio"
	"io"
)

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", err
	}

	k := make([]byte, klen)
	_,err = io.ReadFull(r, k)
	if err != nil {
		return "", err
	}

	return string(k), nil
}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}

	vlen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}

	k := make([]byte, klen)
	_,err = io.ReadFull(r, k)
	if err != nil {
		return "", nil, err
	}

	v := make([]byte, vlen)
	_,err = io.ReadFull(r, v)
	if err != nil {
		return "", nil, err
	}

	return string(k), v, nil
}
