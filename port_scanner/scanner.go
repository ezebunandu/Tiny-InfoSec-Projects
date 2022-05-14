package scanner

import (
	"errors"
	"net"
	"strconv"
	"time"
)

var ErrInvalidHost = errors.New("invalid hostname")

func scanPort(hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func GetOpenPorts(hostname string, portRange []int) ([]int, error) {
	openPorts := []int{}
	for i := portRange[0]; i <= portRange[1]; i++ {
		connected := scanPort(hostname, i)
		if connected {
			openPorts = append(openPorts, i)
		}
	}
	if len(openPorts) == 0 {
		return nil, ErrInvalidHost
	}
	return openPorts, nil
}
