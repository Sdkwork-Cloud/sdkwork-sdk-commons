namespace SDKwork.Common.Errors;

public enum ErrorCode
{
    Unknown,
    NetworkError,
    Timeout,
    Cancelled,
    Unauthorized,
    Forbidden,
    NotFound,
    ValidationError,
    RateLimit,
    ServerError,
    TokenExpired,
    TokenInvalid,
    BusinessError,
    Conflict,
    ServiceUnavailable,
    BadGateway,
    GatewayTimeout
}

public record ErrorDetail(
    string? Field,
    string? Message,
    object? Value,
    string? Code,
    string? Constraint
);

public class SdkException : Exception
{
    public ErrorCode Code { get; }
    public int? HttpStatus { get; }
    public List<ErrorDetail>? Details { get; }
    public long Timestamp { get; }
    public string? TraceId { get; }
    public Dictionary<string, object>? Metadata { get; }

    public SdkException(
        string message,
        ErrorCode code = ErrorCode.Unknown,
        int? httpStatus = null,
        List<ErrorDetail>? details = null,
        string? traceId = null,
        Dictionary<string, object>? metadata = null,
        Exception? innerException = null
    ) : base(message, innerException)
    {
        Code = code;
        HttpStatus = httpStatus;
        Details = details;
        Timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();
        TraceId = traceId;
        Metadata = metadata;
    }

    public bool IsRetryable() => this is NetworkException || this is TimeoutException || this is ServerException;

    public bool IsAuthError() => Code == ErrorCode.Unauthorized || Code == ErrorCode.TokenExpired || Code == ErrorCode.TokenInvalid;

    public bool IsClientError() => HttpStatus is >= 400 and < 500;

    public bool IsServerError() => HttpStatus is >= 500;
}

public class NetworkException : SdkException
{
    public NetworkException(string message = "Network error")
        : base(message, ErrorCode.NetworkError) { }
}

public class TimeoutException : SdkException
{
    public int? Timeout { get; }

    public TimeoutException(string message = "Request timeout", int? timeout = null)
        : base(message, ErrorCode.Timeout)
    {
        Timeout = timeout;
    }
}

public class CancelledException : SdkException
{
    public CancelledException(string message = "Request cancelled")
        : base(message, ErrorCode.Cancelled) { }
}

public class AuthenticationException : SdkException
{
    public AuthenticationException(string message = "Authentication failed")
        : base(message, ErrorCode.Unauthorized, 401) { }
}

public class TokenExpiredException : AuthenticationException
{
    public TokenExpiredException(string message = "Token expired")
        : base(message) { }
}

public class TokenInvalidException : AuthenticationException
{
    public TokenInvalidException(string message = "Invalid token")
        : base(message) { }
}

public class ForbiddenException : SdkException
{
    public ForbiddenException(string message = "Access forbidden")
        : base(message, ErrorCode.Forbidden, 403) { }
}

public class NotFoundException : SdkException
{
    public NotFoundException(string message = "Resource not found")
        : base(message, ErrorCode.NotFound, 404) { }
}

public class ValidationException : SdkException
{
    public ValidationException(string message = "Validation error", List<ErrorDetail>? details = null)
        : base(message, ErrorCode.ValidationError, 400, details) { }
}

public class ConflictException : SdkException
{
    public ConflictException(string message = "Resource conflict")
        : base(message, ErrorCode.Conflict, 409) { }
}

public class RateLimitException : SdkException
{
    public int? RetryAfter { get; }

    public RateLimitException(string message = "Rate limit exceeded", int? retryAfter = null)
        : base(message, ErrorCode.RateLimit, 429)
    {
        RetryAfter = retryAfter;
    }
}

public class ServerException : SdkException
{
    public ServerException(string message = "Server error", int httpStatus = 500)
        : base(message, ErrorCode.ServerError, httpStatus) { }
}

public class BadGatewayException : ServerException
{
    public BadGatewayException(string message = "Bad gateway")
        : base(message, 502) { }
}

public class ServiceUnavailableException : ServerException
{
    public ServiceUnavailableException(string message = "Service unavailable")
        : base(message, 503) { }
}

public class GatewayTimeoutException : ServerException
{
    public GatewayTimeoutException(string message = "Gateway timeout")
        : base(message, 504) { }
}

public class BusinessException : SdkException
{
    public object? BusinessCode { get; }
    public object? Data { get; }

    public BusinessException(string message, object? businessCode = null, object? data = null)
        : base(message, ErrorCode.BusinessError)
    {
        BusinessCode = businessCode;
        Data = data;
    }
}

public static class ErrorHelper
{
    public static bool IsSdkError(Exception ex) => ex is SdkException;

    public static bool IsNetworkError(Exception ex) => ex is NetworkException;

    public static bool IsTimeoutError(Exception ex) => ex is TimeoutException;

    public static bool IsAuthError(Exception ex) => ex is AuthenticationException;

    public static bool IsValidationError(Exception ex) => ex is ValidationException;

    public static bool IsRateLimitError(Exception ex) => ex is RateLimitException;

    public static bool IsServerError(Exception ex) => ex is ServerException;

    public static bool IsRetryableError(Exception ex) => ex is NetworkException or TimeoutException or ServerException;
}
