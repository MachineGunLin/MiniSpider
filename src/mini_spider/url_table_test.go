/*
	test for url_table.go
*/

package mini_spider

import (
	"testing"
)

func TestUrlTable_case_1(t *testing.T) {
	// create table
	ut := NewUrlTable()

	// check whether exist
	if ut.Exist("www.baidu.com") {
		t.Errorf("www.baidu.com should not exist")
	}

	// add to table
	ut.Add("www.baidu.com")

	// check whether exist
	if !ut.Exist("www.baidu.com") {
		t.Errorf("www.baidu.com should exist")
	}

	if ut.Exist("www.sina.com.cn") {
		t.Errorf("www.sina.com.cn should not exist")
	}
}
