package gate

import (
	"net"
)

//	
func NewListener(addressAndPort string) (net.Listener, error) {
	if addressAndPort == "" {
		return nil, nil
	}

	l, e := net.Listen("tcp", addressAndPort)
	if e != nil {
		return nil, e
	}

	return l, nil
}