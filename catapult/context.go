package catapult

import "net/http"

type Ctx struct {
	skipDispatch bool

	Request  *Request
	Response *Response
	Client   *Client

	client *http.Client

	Data map[interface{}]interface{}

	middlwares []ClientFeature

	before     []func(request *Request)
	justBefore []func(request *Request) //after we have a body...
	after      []func(response *Response)
	justAfter  []func(response *Response) //after we docoded the body...

	encodeReqFunc func(r *http.Request, payload interface{}) error
	decodeResFunc func(response *Response) (interface{}, error)

	requestPayload interface{}

	err error
}

func (c *Ctx) cloneData() map[interface{}]interface{} {
	cloned := make(map[interface{}]interface{})
	for k, v := range c.Data {
		cloned[k] = v
	}
	return cloned
}

func (c *Ctx) SkipDispatch() {
	c.skipDispatch = true
}

func (c *Ctx) WasDispatchSkipped() bool {
	return c.skipDispatch
}

//func (c *Ctx) cloneParams() (cloned url.Values) {
//	for k, vals := range c.params {
//		for _, v := range vals {
//			cloned.Add(k, v)
//		}
//
//	}
//	return
//}

func (c *Ctx) SetError(err error) {
	//if c.err == nil {
	c.err = err
	//}
}

func NewContext() *Ctx {
	ctx := &Ctx{
		before: make([]func(request *Request), 0),
		//justBefore: make([]func(request *Request), 0),
		after: make([]func(request *Response), 0),
		Data:  make(map[interface{}]interface{}),
		//justAfter:  make([]func(request *Response), 0),
		//url: url.URL{
		//	Scheme: "http",
		//},
		client:        http.DefaultClient, //todo
		decodeResFunc: GenericDecoder,
		encodeReqFunc: GenericEncoder,
	}

	ctx.client.Transport = DefaultTransport
	return ctx

}

//
//func (c *Ctx) SetHost(host string) *Ctx {
//	c.url.Host = host
//	return c
//}
//
//func (c *Ctx) SetPath(path string) *Ctx {
//	c.url.Path = path
//	return c
//}
//
//func (c *Ctx) SetParams(params url.Values) *Ctx {
//	c.params = params
//	return c
//}
//
//func (c *Ctx) AddParam(key, val string) *Ctx {
//	c.params.Add(key, val)
//	return c
//}
//
//func (c *Ctx) SetParam(key, val string) *Ctx {
//	c.params.Set(key, val)
//	return c
//}
//
//func (c *Ctx) DelParam(key string) *Ctx {
//	c.params.Del(key)
//	return c
//}

func (c *Ctx) Clone() *Ctx {
	return &Ctx{
		before:        append([]func(*Request){}, c.before...),
		justBefore:    append([]func(*Request){}, c.justBefore...),
		after:         append([]func(*Response){}, c.after...),
		justAfter:     append([]func(*Response){}, c.justAfter...),
		Data:          c.cloneData(),
		middlwares:    append([]ClientFeature{}, c.middlwares...),
		encodeReqFunc: c.encodeReqFunc,
		decodeResFunc: c.decodeResFunc,
		client:        c.client,
	}
}

func (c *Ctx) SetDecoder(fn func(response *Response) (interface{}, error)) {
	c.decodeResFunc = fn
}

func (c *Ctx) AddBefore(fn func(request *Request)) {
	c.before = append(c.before, fn)
}

func (c *Ctx) AddAfter(fn func(*Response)) {
	c.after = append([]func(*Response){fn}, c.after...)
}

//func (c *Ctx) AddJustAfter(fn func(*Response)) {
//c.justAfter = append(c.justAfter, fn)
//}

func (c *Ctx) AddJustBefore(fn func(request *Request)) {
	c.justBefore = append(c.justBefore, fn)
}
