package util

import (
	"mime/multipart"
	"net/http"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func FromFile(r *http.Request, name string) (multipart.File, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, _, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return f, err
}
