package polysocket

import (
	"net/http"
	"time"
)

func Timeout(req *http.Request, param string) (timeout time.Duration, err error) {
	timeout_str := req.URL.Query().Get(param)
	timeout, err = time.ParseDuration(timeout_str)
	return
}
