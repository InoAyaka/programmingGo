package io

import "io"

/*
	p.201
	インタフェースの埋め込みについて
	２段階埋め込み可能かをチェック
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	io.Writer
}

type ReadWriteCloser interface {
	ReadWriter
	Closer
}
