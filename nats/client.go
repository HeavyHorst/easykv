/*
 * This file is part of easyKV.
 * © 2022 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package nats

import (
	"context"
	"fmt"
	"strings"

	"github.com/HeavyHorst/easykv"
	"github.com/nats-io/nats.go"
)

var cleanReplacer = strings.NewReplacer(".", "/")

// Client provides a shell for the env client
type Client struct {
	nc *nats.Conn
	kv nats.KeyValue

	revisionMap map[string]uint64
}

// New returns a new client
func New(nodes []string, bucket string, opts ...Option) (*Client, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	if len(nodes) == 0 {
		nodes = append(nodes, nats.DefaultURL)
	}

	natsOptions := []nats.Option{nats.MaxReconnects(-1)}

	// override authentication, if any was specified
	if options.Auth.Username != "" && options.Auth.Password != "" {
		natsOptions = append(natsOptions, nats.UserInfo(options.Auth.Username, options.Auth.Password))
	}

	if options.Token != "" {
		natsOptions = append(natsOptions, nats.Token(options.Token))
	}

	if options.Creds != "" {
		natsOptions = append(natsOptions, nats.UserCredentials(options.Creds))
	}

	nc, err := nats.Connect(strings.Join(nodes, ","), natsOptions...)
	if err != nil {
		return nil, fmt.Errorf("could't connect to nats: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("could't initialize jetstream: %w", err)
	}

	kv, err := js.KeyValue(bucket)
	if err != nil {
		return nil, fmt.Errorf("could't open kv bucket: %w", err)
	}

	return &Client{
		nc:          nc,
		kv:          kv,
		revisionMap: make(map[string]uint64),
	}, nil
}

// Close is only meant to fulfill the easykv.ReadWatcher interface.
// Does nothing.
func (c *Client) Close() {
	c.nc.Close()
}

func clean(key string) string {
	newKey := "/" + key
	return cleanReplacer.Replace(newKey)
}

func getWatchKey(prefix string) string {
	prefix = strings.ReplaceAll(prefix, "/", ".")
	prefix = strings.TrimPrefix(prefix, ".")

	if prefix == "" {
		return ">"
	}

	return prefix + ".>"
}

// GetValues is used to lookup all keys with a prefix.
// Several prefixes can be specified in the keys array.
func (c *Client) GetValues(keys []string) (map[string]string, error) {
	allKeys, err := c.kv.Keys()
	if err != nil {
		return nil, fmt.Errorf("couldn't get keys: %w", err)
	}

	// filter keys
	var filteredKeys []string
	for _, key := range keys {
		for _, k := range allKeys {
			if strings.HasPrefix(clean(k), key) {
				filteredKeys = append(filteredKeys, k)
			}
		}
	}

	vars := make(map[string]string)
	for _, key := range filteredKeys {
		val, err := c.kv.Get(key)
		if err != nil {
			return nil, fmt.Errorf("couldn't get key: %v %w", key, err)
		}
		vars[clean(key)] = string(val.Value())
	}

	return vars, nil
}

// WatchPrefix
func (c *Client) WatchPrefix(ctx context.Context, prefix string, opts ...easykv.WatchOption) (uint64, error) {
	var (
		options easykv.WatchOptions
		watcher nats.KeyWatcher
		err     error
	)
	for _, o := range opts {
		o(&options)
	}

	watcher, err = c.kv.Watch(getWatchKey(prefix), nats.Context(ctx), nats.MetaOnly())
	if err != nil {
		return 0, fmt.Errorf("couldn't create nats watcher: %w", err)
	}

	defer watcher.Stop()

	for v := range watcher.Updates() {
		if v == nil {
			break
		}
		c.revisionMap[v.Key()] = v.Revision()
	}

	for {
		select {
		case v := <-watcher.Updates():
			if v == nil {
				break
			}

			for _, k := range options.Keys {
				if strings.HasPrefix(clean(string(v.Key())), k) {
					if v.Revision() != c.revisionMap[v.Key()] {
						return v.Revision(), nil
					}
				}
			}
		case <-ctx.Done():
			return 0, easykv.ErrWatchCanceled
		}
	}
}
