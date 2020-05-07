package json

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

type myRC struct {
	err bool
	r   io.Reader
}

func (m *myRC) Read(b []byte) (n int, err error) {
	if m.err {
		return 0, errors.New("Boom")
	}
	return m.r.Read(b)
}

func (m *myRC) Close() error {
	return nil
}

func newReaderCloser(b []byte, hasErr bool) io.ReadCloser {
	return &myRC{r: bytes.NewReader(b), err: hasErr}
}

type testStruct struct {
	A int      `json:"a"`
	B string   `json:"b"`
	C []string `json:"C"`
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		reqBody io.ReadCloser
		t       interface{}
	}
	tests := []struct {
		name    string
		args    func() args
		want    interface{}
		wantErr bool
	}{
		{
			name: "OK",
			args: func() args {
				return args{
					reqBody: newReaderCloser([]byte("{\"a\":1}"), false),
					t:       &testStruct{},
				}
			},
			want: &testStruct{A: 1},
		},
		{
			name: "Incorrectly formatted data",
			args: func() args {
				return args{
					reqBody: newReaderCloser([]byte("[{\"a\":1}]"), false),
					t:       &testStruct{},
				}
			},
			wantErr: true,
		},
		{
			name: "Array OK",
			args: func() args {
				return args{
					reqBody: newReaderCloser([]byte("[{\"a\":1}]"), false),
					t:       &[]testStruct{},
				}
			},
			want: &[]testStruct{{A: 1}},
		},
		{
			name: "Badly formatted bytes",
			args: func() args {
				return args{
					reqBody: newReaderCloser([]byte("[\"a\":1}]"), false),
					t:       &[]testStruct{},
				}
			},
			wantErr: true,
		},
		{
			name: "Empty byte array - reader ok",
			args: func() args {
				return args{
					reqBody: newReaderCloser(nil, false),
					t:       &[]testStruct{},
				}
			},
			wantErr: true,
		},
		{
			name: "Empty byte array - reader not ok",
			args: func() args {
				return args{
					reqBody: newReaderCloser(nil, true),
					t:       &testStruct{},
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.args()
			if err := Unmarshal(a.reqBody, a.t); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(a.t, tt.want) {
				t.Errorf("Unmarshal() got = %v, want %v", a.t, tt.want)
				return
			}
		})
	}
}
