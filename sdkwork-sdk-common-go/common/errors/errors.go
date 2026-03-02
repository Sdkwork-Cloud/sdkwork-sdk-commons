package common

import (
	"fmt"
	"time"
)

type ErrorCode string

const (
	ErrCodeUnknown           ErrorCode = "UNKNOWN"
	ErrCodeNetworkError     ErrorCode = "NETWORK_ERROR"
	ErrCodeTimeout          ErrorCode = "TIMEOUT"
	ErrCodeCancelled        ErrorCode = "CANCELLED"
	ErrCodeUnauthorized     ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden        ErrorCode = "FORBIDDEN"
	ErrCodeNotFound         ErrorCode = "NOT_FOUND"
	ErrCodeValidationError  ErrorCode = "VALIDATION_ERROR"
	ErrCodeRateLimit        ErrorCode = "RATE_LIMIT"
	ErrCodeServerError      ErrorCode = "SERVER_ERROR"
	ErrCodeTokenExpired     ErrorCode = "TOKEN_EXPIRED"
	ErrCodeTokenInvalid     ErrorCode = "TOKEN_INVALID"
	ErrCodeBusinessError   ErrorCode = "BUSINESS_ERROR"
	ErrCodeConflict         ErrorCode = "CONFLICT"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrCodeBadGateway      ErrorCode = "BAD_GATEWAY"
	ErrCodeGatewayTimeout ErrorCode = "GATEWAY_TIMEOUT"
)

type ErrorDetail struct {
	Field      string      `json:"field,omitempty"`
	Message    string      `json:"message,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Code       string      `json:"code,omitempty"`
	Constraint string      `json:"constraint,omitempty"`
}

type SdkError struct {
	Message    string               `json:"message"`
	Code       ErrorCode            `json:"code"`
	HttpStatus *int                 `json:"httpStatus,omitempty"`
	Details    []ErrorDetail        `json:"details,omitempty"`
	Timestamp  int64                `json:"timestamp"`
	TraceId    string               `json:"traceId,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

func (e *SdkError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *SdkError) IsRetryable() bool {
	return IsRetryableError(e)
}

func (e *SdkError) IsAuthError() bool {
	return e.Code == ErrCodeUnauthorized || e.Code == ErrCodeTokenExpired || e.Code == ErrCodeTokenInvalid
}

func (e *SdkError) IsNetworkError() bool {
	return e.Code == ErrCodeNetworkError || e.Code == ErrCodeTimeout
}

func (e *SdkError) IsClientError() bool {
	return e.HttpStatus != nil && *e.HttpStatus >= 400 && *e.HttpStatus < 500
}

func (e *SdkError) IsServerError() bool {
	return e.HttpStatus != nil && *e.HttpStatus >= 500
}

type NetworkError struct {
	SdkError
}

func NewNetworkError(message string) *NetworkError {
	return &NetworkError{SdkError: SdkError{
		Message:   message,
		Code:      ErrCodeNetworkError,
		Timestamp: Now(),
	}}
}

type TimeoutError struct {
	SdkError
	Timeout *int
}

func NewTimeoutError(message string, timeout *int) *TimeoutError {
	return &TimeoutError{
		SdkError: SdkError{
			Message:   message,
			Code:      ErrCodeTimeout,
			Timestamp: Now(),
		},
		Timeout: timeout,
	}
}

type CancelledError struct {
	SdkError
}

func NewCancelledError(message string) *CancelledError {
	return &CancelledError{SdkError: SdkError{
		Message:   message,
		Code:      ErrCodeCancelled,
		Timestamp: Now(),
	}}
}

type AuthenticationError struct {
	SdkError
}

func NewAuthenticationError(message string) *AuthenticationError {
	status := 401
	return &AuthenticationError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeUnauthorized,
		HttpStatus: &status,
		Timestamp:  Now(),
	}}
}

type TokenExpiredError struct {
	AuthenticationError
}

func NewTokenExpiredError(message string) *TokenExpiredError {
	return &TokenExpiredError{AuthenticationError: AuthenticationError{
		SdkError: SdkError{
			Message:   message,
			Code:      ErrCodeTokenExpired,
			Timestamp: Now(),
		},
	}}
}

type TokenInvalidError struct {
	AuthenticationError
}

func NewTokenInvalidError(message string) *TokenInvalidError {
	return &TokenInvalidError{AuthenticationError: AuthenticationError{
		SdkError: SdkError{
			Message:   message,
			Code:      ErrCodeTokenInvalid,
			Timestamp: Now(),
		},
	}}
}

type ForbiddenError struct {
	SdkError
}

func NewForbiddenError(message string) *ForbiddenError {
	status := 403
	return &ForbiddenError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeForbidden,
		HttpStatus: &status,
		Timestamp:  Now(),
	}}
}

type NotFoundError struct {
	SdkError
}

func NewNotFoundError(message string) *NotFoundError {
	status := 404
	return &NotFoundError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeNotFound,
		HttpStatus: &status,
		Timestamp:  Now(),
	}}
}

type ValidationError struct {
	SdkError
}

func NewValidationError(message string, details []ErrorDetail) *ValidationError {
	status := 400
	return &ValidationError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeValidationError,
		HttpStatus: &status,
		Details:    details,
		Timestamp:  Now(),
	}}
}

type ConflictError struct {
	SdkError
}

func NewConflictError(message string) *ConflictError {
	status := 409
	return &ConflictError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeConflict,
		HttpStatus: &status,
		Timestamp:  Now(),
	}}
}

type RateLimitError struct {
	SdkError
	RetryAfter *int
}

func NewRateLimitError(message string, retryAfter *int) *RateLimitError {
	status := 429
	return &RateLimitError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeRateLimit,
		HttpStatus: &status,
		Timestamp:  Now(),
	}, RetryAfter: retryAfter}
}

type ServerError struct {
	SdkError
}

func NewServerError(message string, httpStatus int) *ServerError {
	return &ServerError{SdkError: SdkError{
		Message:    message,
		Code:       ErrCodeServerError,
		HttpStatus: &httpStatus,
		Timestamp:  Now(),
	}}
}

type BadGatewayError struct {
	ServerError
}

func NewBadGatewayError(message string) *BadGatewayError {
	status := 502
	return &BadGatewayError{ServerError: ServerError{
		SdkError: SdkError{
			Message:    message,
			Code:       ErrCodeBadGateway,
			HttpStatus: &status,
			Timestamp:  Now(),
		},
	}}
}

type ServiceUnavailableError struct {
	ServerError
}

func NewServiceUnavailableError(message string) *ServiceUnavailableError {
	status := 503
	return &ServiceUnavailableError{ServerError: ServerError{
		SdkError: SdkError{
			Message:    message,
			Code:       ErrCodeServiceUnavailable,
			HttpStatus: &status,
			Timestamp:  Now(),
		},
	}}
}

type GatewayTimeoutError struct {
	ServerError
}

func NewGatewayTimeoutError(message string) *GatewayTimeoutError {
	status := 504
	return &GatewayTimeoutError{ServerError: ServerError{
		SdkError: SdkError{
			Message:    message,
			Code:       ErrCodeGatewayTimeout,
			HttpStatus: &status,
			Timestamp:  Now(),
		},
	}}
}

type BusinessError struct {
	SdkError
	BusinessCode interface{}
	Data         interface{}
}

func NewBusinessError(message string, businessCode interface{}, data interface{}) *BusinessError {
	return &BusinessError{
		SdkError: SdkError{
			Message:   message,
			Code:      ErrCodeBusinessError,
			Timestamp: Now(),
		},
		BusinessCode: businessCode,
		Data:         data,
	}
}

func IsSdkError(err error) bool {
	_, ok := err.(*SdkError)
	return ok
}

func IsNetworkError(err error) bool {
	_, ok := err.(*NetworkError)
	return ok
}

func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}

func IsAuthError(err error) bool {
	_, ok := err.(*AuthenticationError)
	return ok
}

func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

func IsServerError(err error) bool {
	_, ok := err.(*ServerError)
	return ok
}

func IsBusinessError(err error) bool {
	_, ok := err.(*BusinessError)
	return ok
}

func IsRetryableError(err error) bool {
	if sdkErr, ok := err.(*SdkError); ok {
		switch sdkErr.Code {
		case ErrCodeNetworkError, ErrCodeTimeout, ErrCodeServerError,
			ErrCodeRateLimit, ErrCodeBadGateway,
			ErrCodeServiceUnavailable, ErrCodeGatewayTimeout:
			return true
		}
	}
	return false
}

func Now() int64 {
	return time.Now().UnixMilli()
}
