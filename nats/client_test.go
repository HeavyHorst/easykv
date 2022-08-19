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
	"sync"
	"testing"
	"time"

	"github.com/HeavyHorst/easykv/testutils"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FilterSuite struct{}

var _ = Suite(&FilterSuite{})

func init() {
	opts := &server.Options{
		JetStream: true,
		Port:      4223,
	}

	// Initialize new server with options
	ns, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	go ns.Start()
	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	nc, err := nats.Connect("nats://127.0.0.1:4223")
	if err != nil {
		panic(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	buckets := []string{"config"}
	for _, bucket := range buckets {
		_, err = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:  bucket,
			History: 1,
			Storage: nats.MemoryStorage,
		})
		if err != nil {
			panic(err)
		}
	}
}

func (s *FilterSuite) TestGetValues(t *C) {
	c, err := New([]string{"nats://127.0.0.1:4223"}, "config")
	if err != nil {
		t.Fatal(err)
	}

	c.kv.PutString("premtest.database.url", "www.google.de")
	c.kv.PutString("premtest.database.user", "Boris")
	c.kv.PutString("remtest.database.hosts.0.name", "test1")
	c.kv.PutString("remtest.database.hosts.0.ip", "192.168.0.1")
	c.kv.PutString("remtest.database.hosts.0.size", "60")
	c.kv.PutString("remtest.database.hosts.1.name", "test2")
	c.kv.PutString("remtest.database.hosts.1.ip", "192.168.0.2")
	c.kv.PutString("remtest.database.hosts.1.size", "80")

	testutils.GetValues(t, c)
}

func (s *FilterSuite) TestWatchEmptyPrefix(t *C) {
	c, err := New([]string{"nats://127.0.0.1:4223"}, "config")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		testutils.WatchPrefix(context.Background(), t, c, "", []string{"/"})
	}()

	time.Sleep(100 * time.Millisecond)
	c.kv.Put("remtest.database.hosts.192.168.0.3", []byte("test3"))
	c.kv.Delete("remtest.database.hosts.192.168.0.3")
	wg.Wait()
}

func (s *FilterSuite) TestWatchPrefix(t *C) {
	c, err := New([]string{"nats://127.0.0.1:4223"}, "config")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		testutils.WatchPrefix(context.Background(), t, c, "/remtest/database", []string{"/remtest/database/hosts"})
	}()

	time.Sleep(100 * time.Millisecond)
	c.kv.Put("remtest.database.hosts.192.168.0.3", []byte("test3"))
	c.kv.Delete("remtest.database.hosts.192.168.0.3")
	wg.Wait()
}

func (s *FilterSuite) TestWatchPrefixCancel(t *C) {
	c, err := New([]string{"nats://127.0.0.1:4223"}, "config")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		testutils.WatchPrefix(ctx, t, c, "", []string{"/"})
	}()

	cancel()
	wg.Wait()
}
