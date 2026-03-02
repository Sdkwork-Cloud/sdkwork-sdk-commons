namespace SDKwork.Common.Core;

public enum HttpMethod
{
    GET,
    POST,
    PUT,
    DELETE,
    PATCH,
    HEAD,
    OPTIONS
}

public enum LogLevel
{
    Debug,
    Info,
    Warn,
    Error,
    Silent
}

public enum RetryBackoff
{
    Fixed,
    Linear,
    Exponential
}

public class HttpStatus
{
    public const int OK = 200;
    public const int CREATED = 201;
    public const int NO_CONTENT = 204;
    public const int BAD_REQUEST = 400;
    public const int UNAUTHORIZED = 401;
    public const int FORBIDDEN = 403;
    public const int NOT_FOUND = 404;
    public const int METHOD_NOT_ALLOWED = 405;
    public const int CONFLICT = 409;
    public const int UNPROCESSABLE_ENTITY = 422;
    public const int TOO_MANY_REQUESTS = 429;
    public const int INTERNAL_SERVER_ERROR = 500;
    public const int BAD_GATEWAY = 502;
    public const int SERVICE_UNAVAILABLE = 503;
    public const int GATEWAY_TIMEOUT = 504;
}

public class MimeTypes
{
    public const string JSON = "application/json";
    public const string FORM_DATA = "multipart/form-data";
    public const string URL_ENCODED = "application/x-www-form-urlencoded";
    public const string OCTET_STREAM = "application/octet-stream";
    public const string TEXT_PLAIN = "text/plain";
    public const string TEXT_HTML = "text/html";
}

public class DefaultValues
{
    public const int DEFAULT_TIMEOUT = 30000;
    public const int DEFAULT_MAX_RETRIES = 3;
    public const int DEFAULT_RETRY_DELAY = 1000;
    public const int DEFAULT_CACHE_TTL = 300000;
    public const int DEFAULT_CACHE_MAX_SIZE = 100;
}

public static class SuccessCodes
{
    public static readonly object[] Codes = { 0, 200, 2000, "0", "200", "2000" };
}

public record ApiResult(
    object? Code,
    object? Data,
    string? Msg,
    string? Message,
    long? Timestamp,
    string? TraceId
);

public record PageResult<T>(
    List<T>? Content,
    List<T>? List,
    int Total,
    int? TotalElements,
    int Page,
    int PageSize,
    int? Size,
    int TotalPages,
    bool HasMore,
    bool? First,
    bool? Last,
    bool? Empty,
    int? Number
);

public record Pageable(
    int? Page = 1,
    int? PageSize = 10,
    int? Size = null,
    string? Sort = null,
    string? Order = null
);

public record RetryConfig(
    int MaxRetries = 3,
    int RetryDelay = 1000,
    RetryBackoff RetryBackoff = RetryBackoff.Exponential,
    int MaxRetryDelay = 30000
);

public record CacheConfig(
    bool Enabled = false,
    int Ttl = 300000,
    int MaxSize = 100
);

public record LoggerConfig(
    LogLevel Level = LogLevel.Info,
    string Prefix = "[SDK]",
    bool Timestamp = true,
    bool Colors = true
);

public record HttpClientConfig(
    string BaseUrl,
    int? Timeout = null,
    Dictionary<string, string>? Headers = null,
    RetryConfig? Retry = null,
    CacheConfig? Cache = null,
    LoggerConfig? Logger = null
);

public record SdkConfig(
    string BaseUrl,
    int? Timeout = null,
    Dictionary<string, string>? Headers = null,
    RetryConfig? Retry = null,
    CacheConfig? Cache = null,
    LoggerConfig? Logger = null,
    string? TenantId = null,
    string? OrganizationId = null,
    string? Platform = null,
    object? UserId = null
) : HttpClientConfig(BaseUrl, Timeout, Headers, Retry, Cache, Logger);
