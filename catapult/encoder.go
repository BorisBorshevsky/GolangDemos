package catapult

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var GenericEncoder = func(r *http.Request, payload interface{}) error {
	if payload == nil {
		return nil
	}

	buf := &bytes.Buffer{}

	switch payload.(type) {
	case string:
		buf.WriteString(payload.(string))
	case []byte:
		buf.Write(payload.([]byte))
	default:
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			return err
		}
		r.Header.Set("Content-Type", "application/json")
	}

	r.Body = ioutil.NopCloser(buf)
	r.ContentLength = int64(buf.Len())
	return nil
}
