// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/data/defi.proto

package data

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

// Validate checks the field values on GetDefiPriceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDefiPriceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDefiPriceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDefiPriceRequestMultiError, or nil if none found.
func (m *GetDefiPriceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDefiPriceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SourceToken

	// no validation rules for TargetToken

	if len(errors) > 0 {
		return GetDefiPriceRequestMultiError(errors)
	}

	return nil
}

// GetDefiPriceRequestMultiError is an error wrapping multiple validation
// errors returned by GetDefiPriceRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDefiPriceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDefiPriceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDefiPriceRequestMultiError) AllErrors() []error { return m }

// GetDefiPriceRequestValidationError is the validation error returned by
// GetDefiPriceRequest.Validate if the designated constraints aren't met.
type GetDefiPriceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDefiPriceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDefiPriceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDefiPriceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDefiPriceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDefiPriceRequestValidationError) ErrorName() string {
	return "GetDefiPriceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDefiPriceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDefiPriceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDefiPriceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDefiPriceRequestValidationError{}

// Validate checks the field values on GetDefiPriceReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetDefiPriceReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDefiPriceReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDefiPriceReplyMultiError, or nil if none found.
func (m *GetDefiPriceReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDefiPriceReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDefiPriceReplyValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDefiPriceReplyValidationError{
					field:  "Item",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDefiPriceReplyValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetDefiPriceReplyMultiError(errors)
	}

	return nil
}

// GetDefiPriceReplyMultiError is an error wrapping multiple validation errors
// returned by GetDefiPriceReply.ValidateAll() if the designated constraints
// aren't met.
type GetDefiPriceReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDefiPriceReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDefiPriceReplyMultiError) AllErrors() []error { return m }

// GetDefiPriceReplyValidationError is the validation error returned by
// GetDefiPriceReply.Validate if the designated constraints aren't met.
type GetDefiPriceReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDefiPriceReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDefiPriceReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDefiPriceReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDefiPriceReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDefiPriceReplyValidationError) ErrorName() string {
	return "GetDefiPriceReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetDefiPriceReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDefiPriceReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDefiPriceReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDefiPriceReplyValidationError{}
