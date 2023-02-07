// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pay/interface/v1/pay.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GoPayReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GoPayReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GoPayReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GoPayReqMultiError, or nil
// if none found.
func (m *GoPayReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GoPayReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Channel

	// no validation rules for TransactionId

	if len(errors) > 0 {
		return GoPayReqMultiError(errors)
	}

	return nil
}

// GoPayReqMultiError is an error wrapping multiple validation errors returned
// by GoPayReq.ValidateAll() if the designated constraints aren't met.
type GoPayReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GoPayReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GoPayReqMultiError) AllErrors() []error { return m }

// GoPayReqValidationError is the validation error returned by
// GoPayReq.Validate if the designated constraints aren't met.
type GoPayReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoPayReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoPayReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoPayReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoPayReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoPayReqValidationError) ErrorName() string { return "GoPayReqValidationError" }

// Error satisfies the builtin error interface
func (e GoPayReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoPayReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoPayReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoPayReqValidationError{}

// Validate checks the field values on GoPayReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GoPayReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GoPayReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GoPayReplyMultiError, or
// nil if none found.
func (m *GoPayReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GoPayReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GoPayReplyMultiError(errors)
	}

	return nil
}

// GoPayReplyMultiError is an error wrapping multiple validation errors
// returned by GoPayReply.ValidateAll() if the designated constraints aren't met.
type GoPayReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GoPayReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GoPayReplyMultiError) AllErrors() []error { return m }

// GoPayReplyValidationError is the validation error returned by
// GoPayReply.Validate if the designated constraints aren't met.
type GoPayReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoPayReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoPayReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoPayReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoPayReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoPayReplyValidationError) ErrorName() string { return "GoPayReplyValidationError" }

// Error satisfies the builtin error interface
func (e GoPayReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoPayReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoPayReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoPayReplyValidationError{}