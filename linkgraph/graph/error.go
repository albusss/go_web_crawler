package graph

import "golang.org/x/xerrors"

var (
	ErrNotFound = xerrors.New("Not found")

	ErrUnknownEdgeLinks = xerrors.New("unknown source and/or destination for edge")
)
