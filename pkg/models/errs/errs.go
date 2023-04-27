package errs

import "errors"

var (
	ErrWarehouseConnect    = errors.New("warehouse connect")
	ErrWarehouseNotHasGood = errors.New("warehouse connect")
	ErrRecordNotFound      = errors.New("record not found")
)
