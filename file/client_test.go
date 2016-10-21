/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/HeavyHorst/easyKV/testutils"

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
      - 192.168.0.1: test1
      - 192.168.0.2: test2

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
					"192.168.0.1": "test1",
				    "192.168.0.2": "test2"
		        },
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

func testGetVal(file, data string, t *C) {
	// write testfile
	err := ioutil.WriteFile(file, []byte(data), 0666)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file)

	c, _ := New(file)
	testutils.GetValues(t, c)
}

func (s *FilterSuite) TestGetValuesYML(t *C) {
	testGetVal(filepathYML, testfileYML, t)
}

func (s *FilterSuite) TestGetValuesJSON(t *C) {
	testGetVal(filepathJSON, testfileJSON, t)
}
