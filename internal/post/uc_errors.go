package post

import "errors"

var wantErrors = []error{
	ErrPostNotFound,
	ErrRequiredField,
}

var (
	// Post
	ErrPostNotFound                = errors.New("post not found")
	ErrTypeNotFound                = errors.New("type not found")
	ErrPermissionNotFound          = errors.New("permission not found")
	ErrDepartmentNotBelongToBranch = errors.New("department not belong to branch")
	ErrAssignNotBelongToBranch     = errors.New("assign not belong to branch")
	ErrPermissionDenied            = errors.New("permission denied")
	ErrPostNotPending              = errors.New("post not pending")
	ErrSelfPostTagged              = errors.New("self posts can only tag users")

	// version
	ErrPostVersionNotFound = errors.New("post version not found")

	// emotion
	ErrReactionNotFound = errors.New("reaction not found")

	// comment
	ErrCommentNotFound = errors.New("Comment not found")

	ErrRequiredField = errors.New("required field")
)
