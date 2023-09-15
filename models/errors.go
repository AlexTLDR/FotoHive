package models

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrEmailTaken = errors.New("models: email already taken")
	ErrNotFound   = errors.New("models: resource not found")
)

type FileError struct {
	Issue string
}

func (fe FileError) Error() string {
	return fmt.Sprintf("models: invalid file: %s", fe.Issue)
}

func checkContentType(r io.ReadSeeker, allowed []string) error {
	testBytes := make([]byte, 512)
	_, err := r.Read(testBytes)
	if err != nil {
		return fmt.Errorf("models: checking content type: %w", err)
	}
	_, err = r.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("models: checking content type: %w", err)
	}
	contentType := http.DetectContentType(testBytes)
	for _, a := range allowed {
		if contentType == a {
			return nil
		}
	}
	return FileError{Issue: fmt.Sprintf("file type not allowed: %s", contentType)}
}

func checkExtension(filename string, allowed []string) error {
	if hasExtension(filename, allowed) {
		return nil
	}
	return FileError{Issue: fmt.Sprintf("file extension not allowed: %s", filename)}
}
