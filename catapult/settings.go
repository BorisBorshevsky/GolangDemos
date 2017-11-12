package catapult

import "time"

type RequestSettings struct {
	RequestTimeout time.Duration
	decoder        func(response *Response) (interface{}, error)
}
