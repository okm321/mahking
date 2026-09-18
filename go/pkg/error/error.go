package error

import (
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

type ClientError struct {
	Description string
	Reason      string
}

func (e *ClientError) Error() string {
	return e.Description
}

func NewClientError(desc string) error {
	err := &ClientError{
		Description: desc,
	}
	return WithStack(err)
}

func NewClientErrorf(desc string, a ...any) error {
	err := &ClientError{
		Description: fmt.Sprintf(desc, a...),
	}
	return WithStack(err)
}

// ErrorNotFound は interceptor でハンドリングされ、CodeNotFound を返します。
type ErrorNotFound ClientError

func NewErrorNotFound(desc string, code ErrCode) error {
	err := &ErrorNotFound{
		Description: desc,
		Reason:      string(code),
	}
	return WithStack(err)
}

func (e *ErrorNotFound) Error() string {
	return e.Description
}

var ErrNotFound error = &ErrorNotFound{
	Description: "存在しないデータです",
	Reason:      string(ErrCodeNotFound),
}

// ErrorResponse はHTTPエラーレスポンス用の構造体
type ErrorResponse struct {
	Message string `json:"message"`
	Reason  string `json:"reason,omitzero"`
}

////////////////////////////////////////////////////////////////
// 以下は github.com/pkg/errors のWrapper
////////////////////////////////////////////////////////////////

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(msg string) error {
	return pkgerrors.New(msg)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...any) error {
	err := fmt.Errorf(format, args...)
	return pkgerrors.WithStack(err)
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return pkgerrors.WithStack(err)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return pkgerrors.Wrap(err, message)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...any) error {
	return pkgerrors.Wrapf(err, format, args...)
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	return pkgerrors.WithMessage(err, message)
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...any) error {
	return pkgerrors.WithMessagef(err, format, args...)
}
