package common

import "io"

type NopCloser struct {
	io.Writer
}

func (NopCloser) Close() error { return nil }
