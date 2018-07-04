package main

import (
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Success *statusConfig
	Warning *statusConfig
	Error   *statusConfig
}

type statusConfig struct {
	Regex    string
	Messages []string
}

func parseConfig(str string) (*config, error) {
	var c config
	err := yaml.Unmarshal([]byte(str), &c)
	return &c, err
}

func (c *config) checkTitle(title string) (status string, messages []string) {

	if c.Warning != nil && c.Warning.matchTitle(title, false) {
		status = "pending"
		messages = c.Warning.Messages
	} else if c.Success != nil && c.Success.matchTitle(title, false) {
		status = "success"
		messages = c.Success.Messages
	} else {
		status = "error"
		messages = c.Error.Messages
	}

	return
}

func (sc *statusConfig) matchTitle(title string, def bool) bool {
	re, err := regexp.Compile(sc.Regex)
	if err != nil {
		log.Printf("cannot compile regex %s", sc.Regex)
		return def
	}
	return re.MatchString(title)
}
