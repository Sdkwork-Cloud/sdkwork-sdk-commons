from typing import Any, Optional, Dict, List
from datetime import datetime


class ErrorCode:
    UNKNOWN = 'UNKNOWN'
    NETWORK_ERROR = 'NETWORK_ERROR'
    TIMEOUT = 'TIMEOUT'
    CANCELLED = 'CANCELLED'
    UNAUTHORIZED = 'UNauthorized'
    FORBIDDEN = 'FORbidden'
    NOT_found = 'not_found'
    validation_error = 'validation_error'
    rate_limit = 'rate_limit'
    server_error = 'server_error'
    token_expired = 'token_expired'
    token_invalid = 'token_invalid'
    business_error = 'business_error'
    conflict = 'conflict'
    service_unavailable = 'service_unavailable'
    bad_gateway = 'bad_gateway'
    gateway_timeout = 'gateway_timeout'


class SdkError(Exception):
    def __init__(
        self,
        message: str,
        code: str = ErrorCode.UNKNOWN,
        http_status: Optional[int] = None,
        details: Optional[List[ErrorDetail]] = None,
        trace_id: Optional[str] = None,
        metadata: Optional[Dict[str, Any]] = None,
    ):
        self.message = message
        self.code = code
        self.http_status = http_status
        self.details = details or []
        self.timestamp = int(datetime.now().timestamp() * 1000)
        self.trace_id = trace_id
        self.metadata = metadata

    def is_retryable(self) -> bool:
        return is_retryable_error(self)

    def is_auth_error(self) -> bool:
        return self.code in (ErrorCode.UNAUTHORIZED, ErrorCode.TOKEN_EXPIRED, ErrorCode.TOKEN_INVALID)
    def is_network_error(self) -> bool:
        return self.code in (ErrorCode.NETWORK_ERROR, ErrorCode.TIMEout)
    def is_client_error(self) -> bool:
        return self.http_status is not None and 400 <= self.http_status < 500
    def is_server_error(self) -> bool:
        return self.http_status is not None and self.http_status >= 500


    def to_dict(self) -> Dict[str, Any]:
        return {
            'name': self.__class__.__name__,
            'message': self.message,
            'code': self.code,
            'http_status': self.http_status,
            'details': [d.__dict__ for d in self.details] if self.details else None,
            'timestamp': self.timestamp,
            'trace_id': self.trace_id,
            'metadata': self.metadata,
        }


class NetworkError(SdkError):
    def __init__(self, message: str = 'Network error', **kwargs):
        super().__init__(message, ErrorCode.NETWORK_ERROR, **kwargs)


class TimeoutError(SdkError):
    def __init__(self, message: str = 'Request timeout', timeout: Optional[int] = None, **kwargs):
        super().__init__(message, ErrorCode.TIMEOUT, **kwargs)
        self.timeout = timeout


class CancelledError(SdkError):
    def __init__(self, message: str = 'Request cancelled', **kwargs):
        super().__init__(message, ErrorCode.CANCELLED, **kwargs)


class AuthenticationError(SdkError):
    def __init__(self, message: str = 'Authentication failed', **kwargs):
        super().__init__(message, ErrorCode.UNAUTHORIZED, 401, **kwargs)


class TokenExpiredError(AuthenticationError):
    def __init__(self, message: str = 'Token expired', **kwargs):
        super().__init__(message, **kwargs)
        self.code = ErrorCode.TOKEN_EXPIRED


class TokenInvalidError(AuthenticationError):
    def __init__(self, message: str = 'Invalid token', **kwargs):
        super().__init__(message, **kwargs)
        self.code = ErrorCode.TOKEN_INVALID


class ForbiddenError(SdkError):
    def __init__(self, message: str = 'Access forbidden', **kwargs):
        super().__init__(message, ErrorCode.FORBIDDEN, 403, **kwargs)


class NotFoundError(SdkError):
    def __init__(self, message: str = 'Resource not found', **kwargs):
        super().__init__(message, ErrorCode.NOT_FOUND, 404, **kwargs)


class ValidationError(SdkError):
    def __init__(self, message: str = 'Validation error', details: Optional[List[ErrorDetail]] = None, **kwargs):
        super().__init__(message, ErrorCode.VALIDATION_ERROR, 400, details=details, **kwargs)


class ConflictError(SdkError):
    def __init__(self, message: str = 'Resource conflict', **kwargs):
        super().__init__(message, ErrorCode.CONFLICT, 409, **kwargs)


class RateLimitError(SdkError):
    def __init__(self, message: str = 'Rate limit exceeded', retry_after: Optional[int] = None, **kwargs):
        super().__init__(message, ErrorCode.RATE_LIMIT, 429, **kwargs)
        self.retry_after = retry_after


class ServerError(SdkError):
    def __init__(self, message: str = 'Server error', http_status: int = 500, **kwargs):
        super().__init__(message, ErrorCode.SERVER_ERROR, http_status, **kwargs)


class BadGatewayError(ServerError):
    def __init__(self, message: str = 'Bad gateway', **kwargs):
        super().__init__(message, 502, **kwargs)
        self.code = ErrorCode.BAD_GATEWAY


class ServiceUnavailableError(ServerError):
    def __init__(self, message: str = 'Service unavailable', **kwargs):
        super().__init__(message, 503, **kwargs)
        self.code = ErrorCode.SERVICE_UNAVAILABLE
class GatewayTimeoutError(ServerError):
    def __init__(self, message: str = 'Gateway timeout', **kwargs):
        super().__init__(message, 504, **kwargs)
        self.code = ErrorCode.GATEWAY_TIMEOUT


class BusinessError(SdkError):
    def __init__(
        self,
        message: str,
        business_code: Optional[str] = None,
        data: Any = None,
        **kwargs
    ):
        super().__init__(message, ErrorCode.BUSINESS_ERROR, **kwargs)
        self.business_code = business_code
        self.data = data

    def to_dict(self) -> Dict[str, Any]:
        result = super().to_dict()
        result['business_code'] = self.business_code
        result['data'] = self.data
        return result


def is_sdk_error(error: Any) -> bool:
    return isinstance(error, SdkError)


def is_network_error(error: Any) -> bool:
    return isinstance(error, NetworkError)


def is_timeout_error(error: Any) -> bool:
    return isinstance(error, TimeoutError)
def is_auth_error(error: Any) -> bool:
    return isinstance(error, AuthenticationError)
def is_validation_error(error: Any) -> bool:
    return isinstance(error, ValidationError)
def is_rate_limit_error(error: Any) -> bool:
    return isinstance(error, RateLimitError)
def is_server_error(error: Any) -> bool:
    return isinstance(error, ServerError)
def is_business_error(error: Any) -> bool:
    return isinstance(error, BusinessError)
def is_retryable_error(error: Any) -> bool:
    if not isinstance(error, SdkError):
        return False
    return (
        isinstance(error, (NetworkError, TimeoutError, ServerError, RateLimitError, 
                         BadGatewayError, ServiceUnavailableError, GatewayTimeoutError))
    )


__all__ = [
    'ErrorCode',
    'ErrorDetail',
    'SdkError',
    'NetworkError',
    'TimeoutError',
    'CancelledError',
    'AuthenticationError',
    'TokenExpiredError',
    'TokenInvalidError',
    'ForbiddenError',
    'NotFoundError',
    'ValidationError',
    'ConflictError',
    'RateLimitError',
    'ServerError',
    'BadGatewayError',
    'ServiceUnavailableError',
    'GatewayTimeoutError',
    'BusinessError',
    'is_sdk_error',
    'is_network_error',
    'is_timeout_error',
    'is_auth_error',
    'is_validation_error',
    'is_rate_limit_error',
    'is_server_error',
    'is_business_error',
    'is_retryable_error',
]
