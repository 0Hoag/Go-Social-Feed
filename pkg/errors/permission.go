package errors

import (
	"fmt"
	"strings"
)

// PermissionError is an error with a field and a list of messages.
type PermissionError struct {
	Code     int      `json:"code"`
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

func NewPermissionError(code int, field string, messages ...string) *PermissionError {
	return &PermissionError{
		Code:     code,
		Field:    field,
		Messages: messages,
	}
}

// Error returns the error message.
func (e PermissionError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, strings.Join(e.Messages, ", "))
}

type PermissionErrorCollector struct {
	errors []*PermissionError
}

func NewPermissionErrorCollector() *PermissionErrorCollector {
	return &PermissionErrorCollector{}
}

// Add adds a new Permission error to the collector.
func (c *PermissionErrorCollector) Add(err *PermissionError) PermissionErrorCollector {
	c.errors = append(c.errors, err)
	return *c
}

// HasError returns true if the collector has any error.
func (c PermissionErrorCollector) HasError() bool {
	return len(c.errors) > 0
}

// Errors returns the list of errors.
func (c PermissionErrorCollector) Errors() []*PermissionError {
	return c.errors
}

// Error returns the error message.
func (c PermissionErrorCollector) Error() string {
	var errors []string
	for _, err := range c.errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, ", ")
}
