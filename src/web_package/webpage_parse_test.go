/* 
	test for webpage_parse.go 
*/

package web_package

import (
	"testing"
)

// test for parseWebPage()
func TestParseWebPage(t *testing.T) {
	s := []byte(`<p>Links:</p><ul><li><a href="test">Test</a><li><a href="/test1/test2">Test1_est2</a></ul>`)

	links, err := parseWebPage(s, "http://www.baidu.com/a/b.html")
	if err != nil {
		t.Errorf("err in parseWebPage():%s", err.Error())
		return
	}

	if len(links) != 2 {
		t.Errorf("len(links) should be 2, now it's %d", len(links))
		return
	}

	if links[0] != "http://www.baidu.com/a/test" || links[1] != "http://www.baidu.com/test1/test2" {
		t.Errorf("links:%s", links)
	}
}
