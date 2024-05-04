package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

type Client struct {
	host   string
	hostIp net.IP
	port   int
}

func New(host string, port int) Client {
	return Client{host: host, port: port, hostIp: nil}
}

func (c *Client) Connect() error {
	ips, err := net.LookupIP(c.host)
	if err != nil {
		return nil
	}
	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 != nil {
			c.hostIp = ipv4
			break
		}
	}
	if c.hostIp == nil {
		return errors.New("failed to lookup ip from host")
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.hostIp.String(), c.port))
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		// TODO: error handling
		text, err := reader.ReadString('\n')
		if text == "" || err != nil {
			return conn.Close()
		}
		_, err = conn.Write([]byte(text[:len(text)-1]))
		if err != nil {
			conn.Close()
			return err
		}
	}
}
