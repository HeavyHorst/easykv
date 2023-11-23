/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package etcdv2

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/HeavyHorst/easykv/testutils"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FilterSuite struct{}

var _ = Suite(&FilterSuite{})

func (s *FilterSuite) TestGetValues(t *C) {
	c, err := NewEtcdClient([]string{"http://localhost:2379"}, "", "", "", false, "", "", false)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	c.client.Set(context.Background(), "/premtest/database/url", "www.google.de", nil)
	c.client.Set(context.Background(), "/premtest/database/user", "Boris", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/0/name", "test1", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/0/ip", "192.168.0.1", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/0/size", "60", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/1/name", "test2", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/1/ip", "192.168.0.2", nil)
	c.client.Set(context.Background(), "/remtest/database/hosts/1/size", "80", nil)

	testutils.GetValues(t, c)
}

func (s *FilterSuite) TestWatchPrefix(t *C) {
	c, err := NewEtcdClient([]string{"http://localhost:2379"}, "", "", "", false, "", "", false)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		testutils.WatchPrefix(context.Background(), t, c, "/", []string{"/"})
	}()

	time.Sleep(100 * time.Millisecond)
	c.client.Set(context.Background(), "/remtest/database/hosts/192.168.0.3", "test3", nil)
	c.client.Delete(context.Background(), "remtest/database/hosts/192.168.0.3", nil)
	wg.Wait()
}

func (s *FilterSuite) TestWatchPrefixCancel(t *C) {
	c, err := NewEtcdClient([]string{"http://localhost:2379"}, "", "", "", false, "", "", false)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		testutils.WatchPrefix(ctx, t, c, "/", []string{"/"})
	}()

	cancel()
	wg.Wait()
}
