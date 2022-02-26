/*
 * This file is part of easyKV.
 * © 2022 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package nats

// Options contains all values that are needed to connect to nats.
type Options struct {
	Nodes []string
	Auth  BasicAuthOptions
	Token string
	Creds string
}

// BasicAuthOptions contains options regarding to basic authentication.
type BasicAuthOptions struct {
	Username string
	Password string
}

// Option configures the nats client.
type Option func(*Options)

// WithBasicAuth enables the basic authentication and sets the username and password.
func WithBasicAuth(b BasicAuthOptions) Option {
	return func(o *Options) {
		o.Auth = b
	}
}

// WithCredentials enables the NATS 2.0 and NATS NGS compatible user credentials and sets the path to the credentials file
func WithCredentials(c string) Option {
	return func(o *Options) {
		o.Creds = c
	}
}

// WithToken enables the token authentication and sets the token.
func WithToken(t string) Option {
	return func(o *Options) {
		o.Token = t
	}
}
