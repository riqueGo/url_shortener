package contants

import "errors"

var ErrCodeNotFound = errors.New("code not found")
var ErrMaxAttemptsReached = errors.New("maximum attempts reached")
