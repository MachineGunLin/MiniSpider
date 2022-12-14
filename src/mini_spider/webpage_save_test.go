/* 
	test for webpage_save.go 
*/

package mini_spider

import (
	"testing"
)

// test for genFilePath()
func TestGenFilePath(t *testing.T) {
	rootPath := "./output"
	url := "www.baidu.com"

	filePath := genFilePath(url, rootPath)
	if filePath != "./output/www.baidu.com" {
		t.Errorf("err in genFilePath(), filePath=%s", filePath)
	}
}

// test for saveWebPage()
func TestSaveWebPage(t *testing.T) {
	rootPath := "./output"
	url := "www.baidu.com"
	data := []byte("this is a test")

	err := saveWebPage(rootPath, url, data)
	if err != nil {
		t.Errorf("err in saveWebPage(%s, %s):%s", rootPath, url, err.Error())
	}
}

