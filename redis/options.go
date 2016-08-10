/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package redis

// Option configures the etcd client
type Option func(*Client)

// WithNodes sets the redis backend nodes
func WithNodes(nodes ...string) Option {
	return func(o *Client) {
		o.machines = nodes
	}
}

// WithPassword sets the redis password
func WithPassword(pw string) Option {
	return func(o *Client) {
		o.password = pw
	}
}
