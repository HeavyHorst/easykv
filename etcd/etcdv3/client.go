/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package etcdv3

import (
	"strings"
	"time"

	"context"

	"github.com/HeavyHorst/easykv"
	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Client is a wrapper around the etcd client
type Client struct {
	client         *clientv3.Client
	serializable   bool
	requestTimeout time.Duration
}

// NewEtcdClient returns an *etcdv3.Client with a connection to named machines.
func NewEtcdClient(machines []string, cert, key, caCert string, basicAuth bool, username string, password string, serializable bool, requestTimeout time.Duration) (*Client, error) {
	var cli *clientv3.Client
	cfg := clientv3.Config{
		Endpoints:   machines,
		DialTimeout: 5 * time.Second,
	}
	tlsInfo := &transport.TLSInfo{}

	if basicAuth {
		cfg.Username = username
		cfg.Password = password
	}

	tls := false
	if caCert != "" {
		tlsInfo.TrustedCAFile = caCert
		tls = true
	}
	if cert != "" && key != "" {
		tlsInfo.CertFile = cert
		tlsInfo.KeyFile = key
		tls = true
	}

	if tls {
		clientConf, err := tlsInfo.ClientConfig()
		if err != nil {
			return &Client{cli, serializable, requestTimeout}, err
		}
		cfg.TLS = clientConf
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		return &Client{cli, serializable, requestTimeout}, err
	}
	return &Client{cli, serializable, requestTimeout}, nil
}

// Close closes the etcdv3 client connection.
func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

// GetValues is used to lookup all keys with a prefix.
// Several prefixes can be specified in the keys array.
func (c *Client) GetValues(keys []string) (map[string]string, error) {
	vars := make(map[string]string)
	for _, key := range keys {
		ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
		opts := []clientv3.OpOption{clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend)}
		if c.serializable {
			opts = append(opts, clientv3.WithSerializable())
		}
		resp, err := c.client.Get(ctx, key, opts...)
		cancel()
		if err != nil {
			return vars, err
		}
		for _, ev := range resp.Kvs {
			vars[string(ev.Key)] = string(ev.Value)
		}
	}
	return vars, nil
}

// WatchPrefix watches a specific prefix for changes.
func (c *Client) WatchPrefix(ctx context.Context, prefix string, opts ...easykv.WatchOption) (uint64, error) {
	var options easykv.WatchOptions
	for _, o := range opts {
		o(&options)
	}

	etcdctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var err error

	rch := c.client.Watch(etcdctx, prefix, clientv3.WithPrefix())
	for wresp := range rch {
		if wresp.Err() != nil {
			return options.WaitIndex, wresp.Err()
		}
		for _, ev := range wresp.Events {
			// Only return if we have a key prefix we care about.
			// This is not an exact match on the key so there is a chance
			// we will still pickup on false positives. The net win here
			// is reducing the scope of keys that can trigger updates.
			for _, k := range options.Keys {
				if strings.HasPrefix(string(ev.Kv.Key), k) {
					return uint64(ev.Kv.Version), err
				}
			}
		}
	}
	if ctx.Err() == context.Canceled {
		return options.WaitIndex, easykv.ErrWatchCanceled
	}
	return 0, err
}
