package catapult

import (
	"encoding/json"
	"strings"
)

var GenericDecoder = func(resp *Response) (interface{}, error) {
	if strings.Contains(resp.rawResponse.Header.Get("Content-Type"), "application/json") {
		res := map[string]interface{}{}
		err := json.NewDecoder(resp.buffer).Decode(&res)
		return res, err
	}

	return resp.buffer.Bytes(), nil
}
