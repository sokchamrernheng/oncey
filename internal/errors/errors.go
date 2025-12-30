package errors

import "errors"

var ErrNotFound = errors.New("key not found")
var ErrKeyNotSet = errors.New("missing idempotency key")
