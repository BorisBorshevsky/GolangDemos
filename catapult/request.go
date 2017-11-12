package catapult

import (
	"net/http"

	"time"

	"net/url"

	"net"

	"bytes"

	"github.com/k0kubun/pp"
	"golang.org/x/net/context"
	"gopkg.in/h2non/gentleman.v1/utils"
)

type Request struct {
	timeout    time.Duration
	Context    *Ctx
	rawRequest *http.Request
}

func (r *Request) populateRawRequest() (*http.Request, func()) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, r.timeout)

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
		},
		ProtoMajor: 1,
		ProtoMinor: 1,
		Proto:      "HTTP/1.1",
		Header:     make(http.Header),
		Body:       utils.NopCloser(),
	}

	return req.WithContext(timeoutCtx), cancel
}

func (r *Request) Send() (*Response, error) {
	for _, fn := range r.Context.middlwares {
		fn.Register(r.Context)
	}

	rawRequest, cancel := r.populateRawRequest()
	r.rawRequest = rawRequest
	defer cancel()

	for _, fn := range r.Context.before {
		fn(r)
	}

	if r.Context.encodeReqFunc == nil {
		r.Context.encodeReqFunc = GenericEncoder
	}

	if err := r.Context.encodeReqFunc(r.rawRequest, r.Context.requestPayload); err != nil {
		r.Context.SetError(err)
		r.Context.skipDispatch = true
	}

	for _, fn := range r.Context.justBefore {
		fn(r)
	}

	if r.Context.skipDispatch || r.Context.err != nil {
		resp := &Response{
			buffer:  &bytes.Buffer{},
			Context: r.Context,
			Body:    utils.NopCloser(),
		}

		r.Context.Response = resp
	} else {
		rawResponse, err := r.Context.client.Do(r.rawRequest)

		if e, ok := err.(net.Error); ok && e.Timeout() {
			r.Context.err = TimeoutError
		} else if err != nil {
			r.Context.err = err
		}

		resp := &Response{
			Context: r.Context,
			buffer:  &bytes.Buffer{},
		}

		if err == nil {
			resp.rawResponse = rawResponse
			resp.Body = rawResponse.Body
			resp.ContentLength = rawResponse.ContentLength
			resp.valid = rawResponse.StatusCode/100 != 4 && rawResponse.StatusCode/100 != 5
		}

		r.Context.Response = resp

	}

	for _, fn := range r.Context.justAfter {
		fn(r.Context.Response)
	}

	for _, fn := range r.Context.after {
		fn(r.Context.Response)
	}

	pp.Println(r.Context.err)

	return r.Context.Response, r.Context.err
}

//func (r *Request) buildResponse(Context *Ctx, rawResponse *http.Response) *Response {
//	r
//}

func (r *Request) Raw() *http.Request {
	return r.rawRequest
}

func (r *Request) Method() string {
	return r.rawRequest.Method
}

func (r *Request) Timeout(dur time.Duration) {
	r.timeout = dur
}

func (r *Request) Body(payload interface{}) {
	r.Context.requestPayload = payload
}

func (r *Request) SetDecoder(fn func(response *Response) (interface{}, error)) {
	r.Context.SetDecoder(fn)
}

//func (r *Request) SetHost(host string) *Request {
//	r.Context.SetHost(host)
//	return r
//}
//
//func (r *Request) SetPath(path string) *Request {
//	r.Context.SetPath(path)
//	return r
//}
//
//func (r *Request) SetParams(params url.Values) *Request {
//	r.Context.params = params
//	return r
//}
//
//func (r *Request) AddParam(key, val string) *Request {
//	r.Context.AddParam(key, val)
//	return r
//}
//
//func (r *Request) SetParam(key, val string) *Request {
//	r.Context.SetParam(key, val)
//	return r
//}
//
//func (r *Request) DelParam(key string) *Request {
//	r.Context.DelParam(key)
//	return r
//}
//
func (r *Request) Wrap(feature ClientFeature) *Request {
	r.Context.middlwares = append(r.Context.middlwares, feature)
	return r
}

//
//func (r *Request) method() string {
//	if r.Context.method != "" {
//		return r.Context.method
//	}
//
//	if r.Context.requestPayload == nil {
//		return http.MethodGet
//	}
//
//	return http.MethodPost
//}
//
//func (r *Request) Get() (*Response, error) {
//	r.Method(http.MethodGet)
//	return r.Send()
//}
//
//func (r *Request) Post() (*Response, error) {
//	r.Method(http.MethodPost)
//	return r.Send()
//}
//
//func (r *Request) Put() (*Response, error) {
//	r.Method(http.MethodPut)
//	return r.Send()
//}
//
//func (r *Request) Delete() (*Response, error) {
//	r.Method(http.MethodDelete)
//	return r.Send()
//}
