package xim

import (
	"errors"
)

var (
	errBadRequest  = errors.New("iid: bad request")
	errNilReceiver = errors.New("iid: nil receiver")
	errNothing     = errors.New("iid: nothing")
)
