package conn

import (
	"github.com/streadway/amqp"
	"log"
	"time"
	"fmt"
)

type RabbitConn struct {
	host      string
	port      string
	username  string
	password  string

	conn      *amqp.Connection
	connErr   chan *amqp.Error

	reconnect chan bool
}

func NewRabbitConn(host, port, username, password string) *RabbitConn {
	return &RabbitConn{
		host: host, port: port,
		username: username, password: password,
		connErr: make(chan *amqp.Error),
		reconnect: make(chan bool, 1),
	}
}

func (c *RabbitConn) Connect() error {
	c.conn = connectToRabbit(c.url())
	c.conn.NotifyClose(c.connErr)
	c.recovery()
	return nil
}

// 每秒尝试连接一次rabbit
func connectToRabbit(url string) *amqp.Connection {
	for {
		conn, err := amqp.Dial(url)
		if err == nil {
			return conn
		}
		log.Println(err)
		log.Printf("Trying to reconnect to RabbitMQ at %s\n", url)
		time.Sleep(1 * time.Second)
	}
}

func (c *RabbitConn) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.username, c.password, c.host, c.port)
}

func (c *RabbitConn) recovery() {
	go func() {
		for {
			<-c.connErr
			log.Printf("Connecting to %s\n", c.url())
			c.conn = connectToRabbit(c.url())
			c.connErr = make(chan *amqp.Error)
			c.conn.NotifyClose(c.connErr)
			c.reconnect <- true
		}
	}()
}

func (c *RabbitConn) Conn() *amqp.Connection {
	return c.conn
}
