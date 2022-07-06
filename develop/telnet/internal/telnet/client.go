package telnet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Client struct {
	host    string
	port    int
	timeout time.Duration
	conn    net.Conn
}

func NewClient(host string, port int, timeout time.Duration) *Client {
	return &Client{
		host: host, port: port, timeout: timeout,
	}
}

func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(c.host, strconv.Itoa(c.port)), c.timeout)
	if err != nil {
		return fmt.Errorf("error connection: %s\n", err.Error())
	}
	c.conn = conn
	log.Println("telnet client connected successfully")
	return nil
}

func (c *Client) Send(msg string) error {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		if err == io.EOF {
			err = errors.New("error: closed connection")
		}
		return err
	}

	return nil
}

func (c *Client) Receive() error {
	reader := bufio.NewReader(c.conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			err = errors.New("error: closed connection")
		}
		return err
	}

	_, err = fmt.Print(msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("error close conn: %s", err.Error())
	}
	return nil
}
