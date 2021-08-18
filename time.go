package config

import (
	"strings"
	"time"

	"github.com/haiyiyun/log"
)

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if d.Duration, err = time.ParseDuration(s); err != nil {
		log.Fatal("ParseDuration error:", err)
	}

	return
}
