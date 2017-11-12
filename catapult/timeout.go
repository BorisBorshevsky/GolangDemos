package catapult

import "time"

type Timeouts struct {
	Request   time.Duration
	TLS       time.Duration
	Dial      time.Duration
	KeepAlive time.Duration
}
