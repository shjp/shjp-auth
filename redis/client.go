package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	auth "github.com/shjp/shjp-auth"
)

const subsetName = "userSessions"

// Client is a wrapper around redis connection
type Client struct {
	conn     redis.Conn
	reusable bool
}

// Options contains the option parameters for the redis client
type Options struct {
	Network  string
	Address  string
	Reusable bool
}

// NewClient instantiates a new redis client from the options
func (o *Options) NewClient() (auth.SessionClient, error) {
	log.Println("Dialing redis server | network:", o.Network, "|", o.Address)
	conn, err := redis.Dial(o.Network, o.Address, redis.DialKeepAlive(10*time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "Error dialing redis server")
	}

	return &Client{conn: conn, reusable: o.Reusable}, nil
}

// String stringifies the options object
func (o *Options) String() string {
	return fmt.Sprintf("{ network: %s, address: %s, reusable: %t }", o.Network, o.Address, o.Reusable)
}

// Get executes GET command on the redis server
func (c *Client) Get(key string) ([]byte, error) {
	defer func() {
		if !c.reusable {
			c.Close()
		}
	}()
	val, err := redis.Bytes(c.conn.Do("HGET", subsetName, key))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error executing GET for key: %s", key))
	}

	return val, nil
}

// Set executes SET command on the redis server
func (c *Client) Set(key string, val []byte) error {
	defer func() {
		if !c.reusable {
			c.Close()
		}
	}()

	if _, err := c.conn.Do("HSET", subsetName, key, val); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error executing SET for key: %s | value: %s", key, val))
	}

	return nil
}

// Close closes the connection
func (c *Client) Close() error {
	return c.conn.Close()
}
