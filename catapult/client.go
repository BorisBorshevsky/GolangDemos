package catapult

import (
	"net"
	"net/http"
	"time"
)

type Client struct {
	RequestTimeout time.Duration
	context        *Ctx
	rawClient      *http.Client
}

var (
	DefaultRequestTimeout        = time.Second * 20
	DefaultExpectContinueTimeout = time.Second
	DefaultTLSHandshakeTimeout   = time.Second
	DefaultDialTimeout           = time.Second
	DefaultKeepAlive             = time.Second * 90
	DefaultIdleConnTimeout       = time.Second * 90
	DefaultMaxIdleConnections    = 100
	DefaultTransport             = buildDefaultTransport()
)

func buildDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   DefaultDialTimeout,
			KeepAlive: DefaultKeepAlive,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          DefaultMaxIdleConnections,
		IdleConnTimeout:       DefaultIdleConnTimeout,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ExpectContinueTimeout: DefaultExpectContinueTimeout,
	}
}

func NewClient() *Client {
	client := &Client{
		RequestTimeout: DefaultRequestTimeout,
		context:        NewContext(),
		rawClient:      http.DefaultClient,
	}

	client.context.Client = client
	client.rawClient.Transport = DefaultTransport

	return client
}

//func (c *Client) SetHost(host string) {
//	c.Context.SetHost(host)
//}
//
//func (c *Client) SetTimeout(dur time.Duration) {
//	c.RequestTimeout = dur
//}

func (c *Client) NewRequest() *Request {
	ctx := c.context.Clone()

	req := &Request{
		timeout: c.RequestTimeout,
		Context: ctx,
	}

	return req
}

func (c *Client) Wrap(feature ClientFeature) {
	c.context.middlwares = append(c.context.middlwares, feature)
}

//
//func (c *Client) SetDecodeFunc(fn func(response *Response) (interface{}, error)) {
//	c.Context.SetDecoder(fn)
//}
