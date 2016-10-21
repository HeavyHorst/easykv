/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package env

import (
	"os"
	"testing"

	"github.com/HeavyHorst/easyKV"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FilterSuite struct{}

var _ = Suite(&FilterSuite{})

func (s *FilterSuite) TestTransform(t *C) {
	dat := transform("/foo/bar/test")
	t.Check(dat, Equals, "FOO_BAR_TEST")
}

func (s *FilterSuite) TestClean(t *C) {
	dat := clean("FOO_BAR_TEST")
	t.Check(dat, Equals, "/foo/bar/test")
}

func (s *FilterSuite) TestWatchPrefix(t *C) {
	c, _ := New()
	stop := make(chan bool)
	num, err := c.WatchPrefix("", stop)
	t.Check(num, Equals, uint64(0))
	t.Check(err, Equals, easyKV.ErrWatchNotSupported)
}

func (s *FilterSuite) TestGetValues(t *C) {
	//set some env vars
	os.Setenv("ENVTEST_FOO_BAR", "some_data")
	os.Setenv("ENVTEST_BAR_FOO", "data_some")

	c, _ := New()
	m, err := c.GetValues([]string{"/envtest"})
	if err != nil {
		t.Error(err)
	}

	t.Check(len(m), Equals, 2)
	t.Check(m["/envtest/foo/bar"], Equals, "some_data")
	t.Check(m["/envtest/bar/foo"], Equals, "data_some")

}
