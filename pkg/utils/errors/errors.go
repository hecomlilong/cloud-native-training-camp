package errors

import "fmt"

var (
	ErrorParseConfigNil         = fmt.Errorf("parse config got nil")
	ErrorConfigSchemaNotSupport = fmt.Errorf("config schema not support")
)
