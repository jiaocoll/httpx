package httpx

import "time"

type HTTPOptions struct {
	Timeout          time.Duration
	RetryMax         int
	FollowRedirects  bool
	HTTPProxy        string
	Unsafe           bool
	DefaultUserAgent string
}
