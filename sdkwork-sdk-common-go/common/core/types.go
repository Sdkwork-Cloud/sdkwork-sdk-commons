package common

type HttpMethod string

const (
	MethodGET    HttpMethod = "GET"
	MethodPOST   HttpMethod = "POST"
	MethodPUT    HttpMethod = "PUT"
	MethodDELETE HttpMethod = "DELETE"
	MethodPATCH  HttpMethod = "PATCH"
	MethodHEAD   HttpMethod = "HEAD"
	MethodOPTIONS HttpMethod = "OPTIONS"
)

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
	LevelSilent LogLevel = "silent"
)

type QueryParams map[string]interface{}

type HttpHeaders map[string]string

type ApiResult struct {
	Code      interface{} `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
	TraceId   string      `json:"traceId"`
}

type PageResult struct {
	Content       []interface{} `json:"content"`
	List          []interface{} `json:"list"`
	Total         int           `json:"total"`
	TotalElements int           `json:"totalElements"`
	Page          int           `json:"page"`
	PageSize      int           `json:"pageSize"`
	Size          int           `json:"size"`
	TotalPages    int           `json:"totalPages"`
	HasMore       bool          `json:"hasMore"`
	First         bool          `json:"first"`
	Last          bool          `json:"last"`
	Empty         bool          `json:"empty"`
	Number        int           `json:"number"`
}

type Pageable struct {
	Page    *int    `json:"page,omitempty"`
	PageSize *int   `json:"pageSize,omitempty"`
	Size    *int    `json:"size,omitempty"`
	Sort    *string `json:"sort,omitempty"`
	Order   *string `json:"order,omitempty"`
}

type RetryBackoff string

const (
	BackoffFixed      RetryBackoff = "fixed"
	BackoffLinear    RetryBackoff = "linear"
	BackoffExponential RetryBackoff = "exponential"
)

type RetryConfig struct {
	MaxRetries     int
	RetryDelay     float64
	RetryBackoff   RetryBackoff
	MaxRetryDelay  float64
	RetryCondition func(error error, retryCount int) bool
}

type CacheConfig struct {
	Enabled bool
	TTL     int
	MaxSize int
}

type LoggerConfig struct {
	Level     LogLevel
	Prefix    string
	Timestamp bool
	Colors    bool
}

type HttpClientConfig struct {
	BaseURL  string
	Timeout  *int
	Headers  HttpHeaders
	Retry    *RetryConfig
	Cache    *CacheConfig
	Logger   *LoggerConfig
}

type SdkConfig struct {
	HttpClientConfig
	ApiKey         string
	AuthToken      string
	AccessToken    string
	TenantId       string
	OrganizationId string
	Platform       string
	UserId         interface{}
}

var DefaultRetryConfig = RetryConfig{
	MaxRetries:    3,
	RetryDelay:    1.0,
	RetryBackoff:  BackoffExponential,
	MaxRetryDelay: 30.0,
}

var DefaultCacheConfig = CacheConfig{
	Enabled: false,
	TTL:     300000,
	MaxSize: 100,
}

var DefaultLoggerConfig = LoggerConfig{
	Level:     LevelInfo,
	Prefix:    "[SDK]",
	Timestamp: true,
	Colors:    true,
}

const DefaultTimeout = 30000

var SuccessCodes = []interface{}{0, 200, 2000, "0", "200", "2000"}

var HttpStatus = map[string]int{
	"OK":                   200,
	"CREATED":              201,
	"NO_CONTENT":           204,
	"BAD_REQUEST":          400,
	"UNAUTHORIZED":         401,
	"FORBIDDEN":            403,
	"NOT_FOUND":            404,
	"METHOD_NOT_ALLOWED":   405,
	"CONFLICT":             409,
	"UNPROCESSABLE_ENTITY": 422,
	"TOO_MANY_REQUESTS":    429,
	"INTERNAL_SERVER_ERROR": 500,
	"BAD_GATEWAY":          502,
	"SERVICE_UNAVAILABLE":  503,
	"GATEWAY_TIMEOUT":      504,
}

var MimeTypes = map[string]string{
	"JSON":        "application/json",
	"FORM_DATA":   "multipart/form-data",
	"URL_ENCODED": "application/x-www-form-urlencoded",
	"OCTET_STREAM": "application/octet-stream",
	"TEXT_PLAIN":  "text/plain",
	"TEXT_HTML":   "text/html",
}
