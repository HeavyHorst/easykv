/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package testutils

import (
	"github.com/HeavyHorst/easyKV"

	. "gopkg.in/check.v1"
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

// GetValues is a util function to test the easyKV.ReadWatcher.GetValues Interface
func GetValues(t *C, c easyKV.ReadWatcher) {
	m, err := c.GetValues([]string{"/remtest", "/premtest"})
	if err != nil {
		t.Error(err)
	}
	t.Check(m, DeepEquals, expected)

	m2, err := c.GetValues([]string{"/premtest"})
	if err != nil {
		t.Error(err)
	}
	t.Check(m2, DeepEquals, expectedPrefix)
}
