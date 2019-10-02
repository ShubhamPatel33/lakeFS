package indexer

import "golang.org/x/xerrors"

var (
	ErrNotFound       = xerrors.New("")
	ErrIndexMalformed = xerrors.New("")
	ErrBadBlock       = xerrors.New("")
)