package catapult

import "net/http"

type Ctx struct {
	skipDispatch bool

	Request  *Request
	Response *Response
	Client   *Client

	//client *http.Client

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

func (c *Ctx) SetError(err error) {
	c.err = err
}

func NewContext() *Ctx {
	ctx := &Ctx{
		before:        make([]func(request *Request), 0),
		after:         make([]func(request *Response), 0),
		Data:          make(map[interface{}]interface{}),
		decodeResFunc: GenericDecoder,
		encodeReqFunc: GenericEncoder,
	}

	return ctx
}

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

func (c *Ctx) AddJustBefore(fn func(request *Request)) {
	c.justBefore = append(c.justBefore, fn)
}
