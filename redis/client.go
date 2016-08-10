/*
 * This file is part of easyKV.
 * Based on code from confd. https://github.com/kelseyhightower/confd
 * © 2013 Kelsey Hightower
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package redis

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/HeavyHorst/easyKV"
	"github.com/garyburd/redigo/redis"
)

// Client is a wrapper around the redis client
type Client struct {
	client   redis.Conn
	machines []string
	password string
}

// Iterate through `machines`, trying to connect to each in turn.
// Returns the first successful connection or the last error encountered.
// Assumes that `machines` is non-empty.
func tryConnect(machines []string, password string) (redis.Conn, error) {
	var err error
	for _, address := range machines {
		var conn redis.Conn
		network := "tcp"
		if _, err = os.Stat(address); err == nil {
			network = "unix"
		}

		dialops := []redis.DialOption{
			redis.DialConnectTimeout(time.Second),
			redis.DialReadTimeout(time.Second),
			redis.DialWriteTimeout(time.Second),
		}

		if password != "" {
			dialops = append(dialops, redis.DialPassword(password))
		}

		conn, err = redis.Dial(network, address, dialops...)

		if err != nil {
			continue
		}
		return conn, nil
	}
	return nil, err
}

// Retrieves a connected redis client from the client wrapper.
// Existing connections will be tested with a PING command before being returned. Tries to reconnect once if necessary.
// Returns the established redis connection or the error encountered.
func (c *Client) connectedClient() (redis.Conn, error) {
	if c.client != nil {
		resp, err := c.client.Do("PING")
		if (err != nil && err == redis.ErrNil) || resp != "PONG" {
			c.client = nil
		}
	}

	// Existing client could have been deleted by previous block
	if c.client == nil {
		var err error
		c.client, err = tryConnect(c.machines, c.password)
		if err != nil {
			return nil, err
		}
	}

	return c.client, nil
}

// NewRedisClient returns an *redis.Client with a connection to named machines.
// It returns an error if a connection to the cluster cannot be made.
func NewRedisClient(opts ...Option) (*Client, error) {
	var err error
	var c Client
	for _, o := range opts {
		o(&c)
	}

	c.client, err = tryConnect(c.machines, c.password)
	return &c, err
}

// GetValues queries redis for keys prefixed by prefix.
func (c *Client) GetValues(keys []string) (map[string]string, error) {
	// Ensure we have a connected redis client
	rClient, err := c.connectedClient()
	if err != nil && err != redis.ErrNil {
		return nil, err
	}

	vars := make(map[string]string)
	for _, key := range keys {
		key = strings.Replace(key, "/*", "", -1)
		value, err := redis.String(rClient.Do("GET", key))
		if err == nil {
			vars[key] = value
			continue
		}

		if err != redis.ErrNil {
			return vars, err
		}

		if key == "/" {
			key = "/*"
		} else {
			key = fmt.Sprintf("%s/*", key)
		}

		idx := 0
		for {
			values, err := redis.Values(rClient.Do("SCAN", idx, "MATCH", key, "COUNT", "1000"))
			if err != nil && err != redis.ErrNil {
				return vars, err
			}
			idx, _ = redis.Int(values[0], nil)
			items, _ := redis.Strings(values[1], nil)
			for _, item := range items {
				var newKey string
				if newKey, err = redis.String(item, nil); err != nil {
					return vars, err
				}
				if value, err = redis.String(rClient.Do("GET", newKey)); err == nil {
					vars[newKey] = value
				}
			}
			if idx == 0 {
				break
			}
		}
	}
	return vars, nil
}

// WatchPrefix is not yet implemented.
func (c *Client) WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	return 0, easyKV.ErrWatchNotSupported
}