/*
	test for mini_spider.go
*/

package mini_spider

import (
	"testing"
)

import (
	"mini_spider_config"
)

func TestNewMiniSpider(t *testing.T) {
	conf, _ := mini_spider_config.LoadConfig("../mini_spider_config/test_data/spider.config")
	seeds, _ := LoadSeedFile(conf.Basic.UrlListFile)
	_, err := NewMiniSpider(&conf, seeds)
	if err != nil {
		t.Errorf("err happen in NewMiniSpider: %s", err.Error())
	}
}
