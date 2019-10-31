/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package file

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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

const filepathYML string = "/tmp/easyKV_filetest.yml"
const testfileYML string = `
remtest:
  database:
    hosts:
    - name: test1
      ip: 192.168.0.1
      size: 60
    - name: test2
      ip: 192.168.0.2
      size: 80

premtest:
  database: {url: www.google.de, user: Boris}
`

const filepathJSON string = "/tmp/easyKV_filetest.json"
const testfileJSON string = `
{
	"remtest": {
		"database": {
			"hosts": [
				{
					"name": "test1",
					"ip": "192.168.0.1",
					"size": 60
				},
				{
					"name": "test2",
					"ip": "192.168.0.2",
					"size": 80
				}
			]
		}
	},
	"premtest": {
		"database": {
			"url": "www.google.de",
			"user": "Boris"
		}
	}
}
`

const filepathJSON2 string = "/tmp/easyKV_filetest2.json"
const testfileJSON2 string = `
{
	"remtest": [
		1,
		true,
		null
	],
	"premtest": {
		"database": {
			"url": 100,
			"user": false
		}
	}
}
`

func testGetVal(file, data string, expected map[string]string, t *C) {
	// write testfile
	err := ioutil.WriteFile(file, []byte(data), 0666)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file)

	c, _ := New(file)

	if expected == nil {
		err = testutils.GetValues(t, c)
		if err != nil {
			t.Error(err)
		}
	} else {
		m, err := c.GetValues([]string{"/remtest", "/premtest"})
		if err != nil {
			t.Error(err)
		}
		t.Check(m, DeepEquals, expected)
	}
}

func (s *FilterSuite) TestGetValuesYML(t *C) {
	testGetVal(filepathYML, testfileYML, nil, t)
}

func (s *FilterSuite) TestGetValuesJSON(t *C) {
	testGetVal(filepathJSON, testfileJSON, nil, t)
}

func (s *FilterSuite) TestGetValuesJSON2(t *C) {
	testGetVal(filepathJSON2, testfileJSON2, map[string]string{
		"/remtest/0":              "1",
		"/remtest/1":              "true",
		"/remtest/2":              "<nil>",
		"/premtest/database/url":  "100",
		"/premtest/database/user": "false",
	}, t)
}

func (s *FilterSuite) TestWatchPrefix(t *C) {
	err := ioutil.WriteFile(filepathYML, []byte(testfileYML), 0666)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filepathYML)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c, _ := New(filepathYML)
		testutils.WatchPrefix(context.Background(), t, c, "/", []string{})
	}()

	time.Sleep(100 * time.Millisecond)
	err = ioutil.WriteFile(filepathYML, []byte(testfileJSON), 0666)
	if err != nil {
		t.Error(err)
	}
	wg.Wait()
}

func (s *FilterSuite) TestWatchPrefixCancel(t *C) {
	err := ioutil.WriteFile(filepathYML, []byte(testfileYML), 0666)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filepathYML)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c, _ := New(filepathYML)
		testutils.WatchPrefix(ctx, t, c, "/", []string{})
	}()

	cancel()
	wg.Wait()
}

func (s *FilterSuite) TestHTTPWithHeaders(t *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "---")
		for name, value := range r.Header {
			fmt.Fprintf(w, "%s: %q\n", name, value[0])
		}
	}))
	defer ts.Close()

	c, _ := New(ts.URL, WithHeaders(map[string]string{
		"X-Test-Token": "Hi",
		"Content-Type": "application/json",
	}))
	vals, err := c.GetValues([]string{"/"})
	if err != nil {
		t.Fatal(err)
	}

	testHeader := func(header, expected string) {
		if vals[header] != expected {
			t.Errorf("Expected %q to equal %q", header, expected)
		}
	}
	testHeader("/X-Test-Token", "Hi")
	testHeader("/Content-Type", "application/json")
	testHeader("/X-Nonexistent", "")
}
