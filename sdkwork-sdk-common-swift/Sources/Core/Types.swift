import Foundation

public enum HttpMethod: String {
    case get = "GET"
    case post = "POST"
    case put = "PUT"
    case delete = "DELETE"
    case patch = "PATCH"
    case head = "HEAD"
    case options = "OPTIONS"
}

public enum LogLevel: String {
    case debug
    case info
    case warn
    case error
    case silent
}

public typealias QueryParams = [String: Any?]
public typealias HttpHeaders = [String: String]

public struct ApiResult {
    public let code: Any?
    public let data: Any?
    public let msg: String?
    public let message: String?
    public let timestamp: Int64?
    public let traceId: String?

    public init(code: Any?, data: Any?, msg: String?, message: String?, timestamp: Int64?, traceId: String?) {
        self.code = code
        self.data = data
        self.msg = msg
        self.message = message
        self.timestamp = timestamp
        self.traceId = traceId
    }
}

public struct PageResult<T> {
    public let content: [T]?
    public let list: [T]?
    public let total: Int
    public let totalElements: Int?
    public let page: Int
    public let pageSize: Int
    public let size: Int?
    public let totalPages: Int
    public let hasMore: Bool
    public let first: Bool?
    public let last: Bool?
    public let empty: Bool?
    public let number: Int?

    public init(content: [T]?, list: [T]?, total: Int, totalElements: Int?, page: Int, pageSize: Int, size: Int?, totalPages: Int, hasMore: Bool, first: Bool?, last: Bool?, empty: Bool?, number: Int?) {
        self.content = content
        self.list = list
        self.total = total
        self.totalElements = totalElements
        self.page = page
        self.pageSize = pageSize
        self.size = size
        self.totalPages = totalPages
        self.hasMore = hasMore
        self.first = first
        self.last = last
        self.empty = empty
        self.number = number
    }
}

public struct Pageable {
    public var page: Int?
    public var pageSize: Int?
    public var size: Int?
    public var sort: String?
    public var order: String?

    public init(page: Int? = 1, pageSize: Int? = 10, size: Int? = nil, sort: String? = nil, order: String? = nil) {
        self.page = page
        self.pageSize = pageSize
        self.size = size
        self.sort = sort
        self.order = order
    }
}

public enum RetryBackoff: String {
    case fixed
    case linear
    case exponential
}

public struct RetryConfig {
    public var maxRetries: Int
    public var retryDelay: Double
    public var retryBackoff: RetryBackoff
    public var maxRetryDelay: Double

    public static let `default` = RetryConfig(
        maxRetries: 3,
        retryDelay: 1.0,
        retryBackoff: .exponential,
        maxRetryDelay: 30.0
    )

    public init(maxRetries: Int = 3, retryDelay: Double = 1.0, retryBackoff: RetryBackoff = .exponential, maxRetryDelay: Double = 30.0) {
        self.maxRetries = maxRetries
        self.retryDelay = retryDelay
        self.retryBackoff = retryBackoff
        self.maxRetryDelay = maxRetryDelay
    }
}

public struct CacheConfig {
    public var enabled: Bool
    public var ttl: Int
    public var maxSize: Int

    public static let `default` = CacheConfig(enabled: false, ttl: 300000, maxSize: 100)

    public init(enabled: Bool = false, ttl: Int = 300000, maxSize: Int = 100) {
        self.enabled = enabled
        self.ttl = ttl
        self.maxSize = maxSize
    }
}

public struct LoggerConfig {
    public var level: LogLevel
    public var prefix: String
    public var timestamp: Bool
    public var colors: Bool

    public static let `default` = LoggerConfig(level: .info, prefix: "[SDK]", timestamp: true, colors: true)

    public init(level: LogLevel = .info, prefix: String = "[SDK]", timestamp: Bool = true, colors: Bool = true) {
        self.level = level
        self.prefix = prefix
        self.timestamp = timestamp
        self.colors = colors
    }
}

public protocol HttpClientConfigProtocol {
    var baseUrl: String { get }
    var timeout: Int? { get }
    var headers: HttpHeaders? { get }
    var retry: RetryConfig? { get }
    var cache: CacheConfig? { get }
    var logger: LoggerConfig? { get }
}

public struct HttpClientConfig: HttpClientConfigProtocol {
    public var baseUrl: String
    public var timeout: Int?
    public var headers: HttpHeaders?
    public var retry: RetryConfig?
    public var cache: CacheConfig?
    public var logger: LoggerConfig?

    public init(baseUrl: String, timeout: Int? = nil, headers: HttpHeaders? = nil, retry: RetryConfig? = nil, cache: CacheConfig? = nil, logger: LoggerConfig? = nil) {
        self.baseUrl = baseUrl
        self.timeout = timeout
        self.headers = headers
        self.retry = retry
        self.cache = cache
        self.logger = logger
    }
}

public struct SdkConfig: HttpClientConfigProtocol {
    public var baseUrl: String
    public var timeout: Int?
    public var headers: HttpHeaders?
    public var retry: RetryConfig?
    public var cache: CacheConfig?
    public var logger: LoggerConfig?
    public var apiKey: String?
    public var authToken: String?
    public var accessToken: String?
    public var tenantId: String?
    public var organizationId: String?
    public var platform: String?
    public var userId: Any?

    public init(baseUrl: String, timeout: Int? = nil, headers: HttpHeaders? = nil, retry: RetryConfig? = nil, cache: CacheConfig? = nil, logger: LoggerConfig? = nil, apiKey: String? = nil, authToken: String? = nil, accessToken: String? = nil, tenantId: String? = nil, organizationId: String? = nil, platform: String? = nil, userId: Any? = nil) {
        self.baseUrl = baseUrl
        self.timeout = timeout
        self.headers = headers
        self.retry = retry
        self.cache = cache
        self.logger = logger
        self.apiKey = apiKey
        self.authToken = authToken
        self.accessToken = accessToken
        self.tenantId = tenantId
        self.organizationId = organizationId
        self.platform = platform
        self.userId = userId
    }
}

public enum HttpStatus {
    public static let ok = 200
    public static let created = 201
    public static let noContent = 204
    public static let badRequest = 400
    public static let unauthorized = 401
    public static let forbidden = 403
    public static let notFound = 404
    public static let methodNotAllowed = 405
    public static let conflict = 409
    public static let unprocessableEntity = 422
    public static let tooManyRequests = 429
    public static let internalServerError = 500
    public static let badGateway = 502
    public static let serviceUnavailable = 503
    public static let gatewayTimeout = 504
}

public enum MimeTypes {
    public static let json = "application/json"
    public static let formData = "multipart/form-data"
    public static let urlEncoded = "application/x-www-form-urlencoded"
    public static let octetStream = "application/octet-stream"
    public static let textPlain = "text/plain"
    public static let textHtml = "text/html"
}

public let defaultTimeout = 30000
public let successCodes: [Any] = [0, 200, 2000, "0", "200", "2000"]
