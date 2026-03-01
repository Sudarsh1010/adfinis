package milvus

import "errors"

var (
	ErrConnectionFailed   = errors.New("failed to connect to Milvus")
	ErrCollectionNotFound = errors.New("collection not found")
)
