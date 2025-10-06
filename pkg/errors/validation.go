package errors

import (
	"fmt"
	"strings"
)

// ValidationError is an error with a field and a list of messages.
type ValidationError struct {
	Code     int      `json:"code"`
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

func NewValidationError(code int, field string, messages ...string) *ValidationError {
	return &ValidationError{
		Code:     code,
		Field:    field,
		Messages: messages,
	}
}

// Error returns the error message.
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, strings.Join(e.Messages, ", "))
}

type ValidationErrorCollector struct {
	errors []*ValidationError
}

func NewValidationErrorCollector() *ValidationErrorCollector {
	return &ValidationErrorCollector{}
}

// Add adds a new validation error to the collector.
func (c *ValidationErrorCollector) Add(err *ValidationError) ValidationErrorCollector {
	c.errors = append(c.errors, err)
	return *c
}

// HasError returns true if the collector has any error.
func (c ValidationErrorCollector) HasError() bool {
	return len(c.errors) > 0
}

// Errors returns the list of errors.
func (c ValidationErrorCollector) Errors() []*ValidationError {
	return c.errors
}

// Error returns the error message.
func (c ValidationErrorCollector) Error() string {
	var errors []string
	for _, err := range c.errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, ", ")
}
