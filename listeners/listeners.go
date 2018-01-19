package listeners

import (
	"fmt"
	"net"
	//"github.com/sirupsen/logrus"
)

// Init creates new listeners for the server.
// TODO: Clean up the fact that socketGroup and tlsConfig aren't always used.
func Init(proto, addr string) ([]net.Listener, error) {
	ls := []net.Listener{}

	switch proto {
	case "tcp":
		l, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	case "unix":
		l, err := net.Listen("unix", addr)
		if err != nil {
			return nil, fmt.Errorf("can't create unix socket %s: %v", addr, err)
		}
		ls = append(ls, l)
	default:
		return nil, fmt.Errorf("invalid protocol format: %q", proto)
	}

	return ls, nil
}
