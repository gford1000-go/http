package json

import (
	js "encoding/json"
	"io"
	"io/ioutil"
)

// Unmarshal attempts to populate t with the contents of a io.ReadCloser
func Unmarshal(r io.ReadCloser, t interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = js.Unmarshal(body, t)
	if err != nil {
		return err
	}
	return nil
}
