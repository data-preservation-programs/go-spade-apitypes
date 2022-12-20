// Package apitypes contains complete, dependency-free definitions of all
// possible responses of the ♠️ API. This includes the "root" ResponseEnvelope
// struct and all of its dependencies.
package apitypes

import (
	"time"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=APIErrorCode -output=types_err.go

//nolint:revive
type APIErrorCode int

//nolint:revive
const (
	ErrInvalidRequest            APIErrorCode = 4400
	ErrUnauthorizedAccess        APIErrorCode = 4401
	ErrSystemTemporarilyDisabled APIErrorCode = 4503
)

// ResponseEnvelope is the structure wrapping all responses from the deal engine
type ResponseEnvelope struct {
	RequestID          string          `json:"request_id,omitempty"`
	ResponseTime       time.Time       `json:"response_timestamp"`
	ResponseStateEpoch int64           `json:"response_state_epoch,omitempty"`
	ResponseCode       int             `json:"response_code"`
	ErrCode            int             `json:"error_code,omitempty"`
	ErrSlug            string          `json:"error_slug,omitempty"`
	ErrLines           []string        `json:"error_lines,omitempty"`
	InfoLines          []string        `json:"info_lines,omitempty"`
	ResponseEntries    *int            `json:"response_entries,omitempty"`
	Response           ResponsePayload `json:"response"`
}
type isResponsePayload struct{}

//nolint:revive
type ResponsePayload interface {
	is() isResponsePayload
}

const (
	cidTrimPrefix = 6
	cidTrimSuffix = 8
)

//nolint:revive
func TrimCidString(cs string) string {
	if len(cs) <= cidTrimPrefix+cidTrimSuffix+2 {
		return cs
	}
	return cs[0:cidTrimPrefix] + "~" + cs[len(cs)-cidTrimSuffix:]
}
