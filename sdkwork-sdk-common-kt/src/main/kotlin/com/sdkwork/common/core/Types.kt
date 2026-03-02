package com.sdkwork.common.core

object HttpMethod {
    const val GET = "GET"
    const val POST = "POST"
    const val PUT = "PUT"
    const val DELETE = "DELETE"
    const val PATCH = "PATCH"
    const val HEAD = "HEAD"
    const val OPTIONS = "OPTIONS"
}

object LogLevel {
    const val DEBUG = "debug"
    const val INFO = "info"
    const val WARN = "warn"
    const val ERROR = "error"
    const val SILENT = "silent"
}

object RetryBackoff {
    const val FIXED = "fixed"
    const val LINEAR = "linear"
    const val EXPONENTIAL = "exponential"
}

object HttpStatus {
    const val OK = 200
    const val CREATED = 201
    const val NO_CONTENT = 204
    const val BAD_REQUEST = 400
    const val UNAUTHORIZED = 401
    const val FORBIDDEN = 403
    const val NOT_FOUND = 404
    const val METHOD_NOT_ALLOWED = 405
    const val CONFLICT = 409
    const val UNPROCESSABLE_ENTITY = 422
    const val TOO_MANY_REQUESTS = 429
    const val INTERNAL_SERVER_ERROR = 500
    const val BAD_GATEWAY = 502
    const val SERVICE_UNAVAILABLE = 503
    const val GATEWAY_TIMEOUT = 504
}

object MimeTypes {
    const val JSON = "application/json"
    const val FORM_DATA = "multipart/form-data"
    const val URL_ENCODED = "application/x-www-form-urlencoded"
    const val OCTET_STREAM = "application/octet-stream"
    const val TEXT_PLAIN = "text/plain"
    const val TEXT_HTML = "text/html"
}

object DefaultValues {
    const val DEFAULT_TIMEOUT = 30000
    const val DEFAULT_MAX_RETRIES = 3
    const val DEFAULT_RETRY_DELAY = 1000
    const val DEFAULT_CACHE_TTL = 300000
    const val DEFAULT_CACHE_MAX_SIZE = 100
}

data class ApiResult(
    val code: Any?,
    val data: Any?,
    val msg: String?,
    val message: String?,
    val timestamp: Long?,
    val traceId: String?
)

data class PageResult<T>(
    val content: List<T>?,
    val list: List<T>?,
    val total: Int,
    val totalElements: Int?,
    val page: Int,
    val pageSize: Int,
    val size: Int?,
    val totalPages: Int,
    val hasMore: Boolean,
    val first: Boolean?,
    val last: Boolean?,
    val empty: Boolean?,
    val number: Int?
)

data class Pageable(
    val page: Int? = 1,
    val pageSize: Int? = 10,
    val size: Int? = null,
    val sort: String? = null,
    val order: String? = null
)

data class RetryConfig(
    val maxRetries: Int = 3,
    val retryDelay: Long = 1000L,
    val retryBackoff: String = RetryBackoff.EXPONENTIAL,
    val maxRetryDelay: Long = 30000L
)

data class CacheConfig(
    val enabled: Boolean = false,
    val ttl: Long = 300000L,
    val maxSize: Int = 100
)

data class LoggerConfig(
    val level: String = LogLevel.INFO,
    val prefix: String = "[SDK]",
    val timestamp: Boolean = true,
    val colors: Boolean = true
)

open class HttpClientConfig(
    open val baseUrl: String,
    open val timeout: Int? = null,
    open val headers: Map<String, String>? = null,
    open val retry: RetryConfig? = null,
    open val cache: CacheConfig? = null,
    open val logger: LoggerConfig? = null
)

data class SdkConfig(
    override val baseUrl: String,
    override val timeout: Int? = null,
    override val headers: Map<String, String>? = null,
    override val retry: RetryConfig? = null,
    override val cache: CacheConfig? = null,
    override val logger: LoggerConfig? = null,
    val tenantId: String? = null,
    val organizationId: String? = null,
    val platform: String? = null,
    val userId: Any? = null
) : HttpClientConfig(baseUrl, timeout, headers, retry, cache, logger)

val successCodes = listOf(0, 200, 2000, "0", "200", "2000")
