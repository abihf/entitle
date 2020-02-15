package entitle

import (
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Valid   *statusConfig `json:"valid"`
	Invalid *statusConfig `json:"invalid"`
	WIP     *statusConfig `json:"wip"`
}

type statusConfig struct {
	Regex    string   `json:"regex"`
	Messages []string `json:"messages"`
}

func parseConfig(str string) (*config, error) {
	var c config
	err := yaml.Unmarshal([]byte(str), &c)
	return &c, err
}

func (c *config) checkTitle(title string) (status string, messages []string) {

	if c.WIP != nil && c.WIP.matchTitle(title, false) {
		status = "pending"
		messages = c.WIP.Messages
	} else if c.Valid != nil && c.Valid.matchTitle(title, false) {
		status = "success"
		messages = c.Valid.Messages
	} else {
		status = "error"
		messages = c.Invalid.Messages
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
