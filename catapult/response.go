package catapult

import (
	"bytes"
	"io"
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

type Response struct {
	valid      bool
	statusCode int

	ContentLength int64
	Body          io.ReadCloser

	started     bool
	Context     *Ctx
	rawResponse *http.Response
	buffer      *bytes.Buffer
}

var AlreadyWrittenError = errors.New("already written")

func (r *Response) Valid() bool {
	return r.Context.err == nil
}

func (r *Response) Status() int {
	return r.statusCode
}

func (r *Response) Write(p []byte) (n int, err error) {
	if r.started {
		return 0, AlreadyWrittenError
	}
	r.started = true
	return r.buffer.Write(p)
}

func (r *Response) Read(p []byte) (n int, err error) {
	return r.buffer.Read(p)
}

func (r *Response) Bytes() []byte {
	r.populateResponseByteBuffer()
	return r.buffer.Bytes()
}

func (r *Response) String() string {
	r.populateResponseByteBuffer()
	return r.buffer.String()
}

func (r *Response) Decode() (interface{}, error) {
	r.populateResponseByteBuffer()
	return r.Context.decodeResFunc(r)
}

func (r *Response) populateResponseByteBuffer() {
	if r.Context.err != nil {
		pp.Println(r.Context.err.Error())
		return
	}

	if r.buffer == nil {
		r.buffer = bytes.NewBuffer([]byte{})
	}

	// Have I done this already?
	if r.buffer.Len() != 0 {
		return
	}

	defer r.Body.Close()

	// Is there any content?
	if r.ContentLength == 0 {
		return
	}

	// Did the server tell us how big the response is going to be?
	if r.ContentLength > 0 {
		r.buffer.Grow(int(r.rawResponse.ContentLength))
	}

	_, err := io.Copy(r.buffer, r.Body)
	if err != nil && err != io.EOF {
		r.Context.err = err
	}

}
