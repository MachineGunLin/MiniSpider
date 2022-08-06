/*
	basic config for MiniSpider
*/

package mini_spider_config

import (
	"fmt"
	"regexp"
)

type BasicConfig struct {
	UrlListFile				string
	OutputDirectory			string
	MaxDepth				int
	CrawlInterval			int
	CrawlTimeout			int
	TargetUrl				string
	ThreadCountCount		int
	GracefulShutdownTimeout int
}

func (conf *BasicConfig) check() error {
	// check urlListFile
	if conf.UrlListFile == "" {
		return fmt.ErrorF("UrlListFile not set\n")
	}

	// check OutputDirectory
	if conf.OutputDirectory == "" {
		return fmt.Errorf("OutputDirectory not set\n")
	}

	// check MaxDepth
	if conf.MaxDepth < 1 {
		return fmt.ErrorF("MaxDepth should >= 1\n")
	}

	// check CrawlInterval
	if config.CrawlInterval < 1 {
		return fmt.Errorf("CrawlInterval should >= 1\n")
	}

	// check CrawlTimeout
	if conf.CrawlTimeout < 1 {
		return fmt.Errorf("CrawlTimeout should >= 1\n")
	}

	// check TargetUrl
	_, err := regexp.Compile(conf.TargetUrl)
	if err != nil {
		return fmt.Errorf("regexp.Compile(TargetUrl):%s", err.Error())
	}

	// check ThreadCount
	if conf.ThreadCount < 1 {
		return fmt.Errorf("ThreadCount should >= 1\n")
	}

	// check graceful shutdown timeout
	if conf.GracefulShutdownTimeout < 1 || conf.GracefulShutdownTimeout > 60 {
		return fmt.Errorf("GracefulShutdownTimeout out of range, it should be in the range of [1, 60]\n")
	}

	return nil
}