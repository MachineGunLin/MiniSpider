/*
	test for seedfile_load.go
*/

package mini_spider

import (
	"testing"
)

// test for LoadSeedFile()
func TestLoadSeedFile(t *testing.T) {
	filePath := "./test_data/seedfile/url.data"
	seeds, err := LoadSeedFile(filePath)
	if err != nil {
		t.Errorf("err in seedFileLoad(%s):%s", filePath, err.Error())
	}

	if len(seeds) != 2 {
		t.Errorf("len(seeds) should be 2, now it's %d", len(seeds))
	}

	if seeds[0] != "http://www.baidu.com" || seeds[1] != "http://www.sina.com.cn" {
		t.Errorf("err in seeds, %s", seeds)
	}
}
