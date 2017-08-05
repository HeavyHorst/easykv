/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package testutils

import (
	"context"

	"github.com/HeavyHorst/easykv"
	"gopkg.in/check.v1"
)

var expected = map[string]string{
	"/premtest/database/url":              "www.google.de",
	"/premtest/database/user":             "Boris",
	"/remtest/database/hosts/192.168.0.1": "test1",
	"/remtest/database/hosts/192.168.0.2": "test2",
}

var expectedPrefix = map[string]string{
	"/premtest/database/url":  "www.google.de",
	"/premtest/database/user": "Boris",
}

// GetValues is a util function to test the easykv.ReadWatcher.GetValues Method
func GetValues(t *check.C, c easykv.ReadWatcher) error {
	m, err := c.GetValues([]string{"/remtest", "/premtest"})
	if err != nil {
		return err
	}
	t.Check(m, check.DeepEquals, expected)

	m2, err := c.GetValues([]string{"/premtest"})
	if err != nil {
		return err
	}
	t.Check(m2, check.DeepEquals, expectedPrefix)
	return nil
}

// WatchPrefix is a util function to test the easykv.ReadWatcher.WatchPrefix Method
func WatchPrefix(ctx context.Context, t *check.C, c easykv.ReadWatcher, prefix string, keys []string) uint64 {
	n, err := c.WatchPrefix(ctx, prefix, easykv.WithWaitIndex(0), easykv.WithKeys(keys))
	if err != nil {
		if err != easykv.ErrWatchCanceled {
			t.Error(err)
		}
	}
	return n
}

// WatchPrefixError is a util function to test if the easykv.ReadWatcher.WatchPrefix returns easykv.ErrWatchNotSupported
func WatchPrefixError(t *check.C, c easykv.ReadWatcher) {
	num, err := c.WatchPrefix(context.Background(), "")
	t.Check(num, check.Equals, uint64(0))
	t.Check(err, check.Equals, easykv.ErrWatchNotSupported)
}
