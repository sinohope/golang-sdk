package http

import (
	"fmt"
	"time"
)

func timestamp() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	return fmt.Sprintf("%d", now.UnixNano()/int64(time.Millisecond))
}
