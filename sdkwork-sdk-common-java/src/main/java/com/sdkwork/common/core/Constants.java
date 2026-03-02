package com.sdkwork.common.core;

import java.util.List;

public final class Constants {

    private Constants() {
    }

    public static final class HttpMethod {
        public static final String GET = "GET";
        public static final String POST = "POST";
        public static final String PUT = "PUT";
        public static final String DELETE = "DELETE";
        public static final String PATCH = "PATCH";
        public static final String HEAD = "HEAD";
        public static final String OPTIONS = "OPTIONS";

        private HttpMethod() {
        }
    }

    public static final class LogLevel {
        public static final String DEBUG = "debug";
        public static final String INFO = "info";
        public static final String WARN = "warn";
        public static final String ERROR = "error";
        public static final String SILENT = "silent";

        private LogLevel() {
        }
    }

    public static final class RetryBackoff {
        public static final String FIXED = "fixed";
        public static final String LINEAR = "linear";
        public static final String EXPONENTIAL = "exponential";

        private RetryBackoff() {
        }
    }

    public static final class HttpStatus {
        public static final int OK = 200;
        public static final int CREATED = 201;
        public static final int NO_CONTENT = 204;
        public static final int BAD_REQUEST = 400;
        public static final int UNAUTHORIZED = 401;
        public static final int FORBIDDEN = 403;
        public static final int NOT_FOUND = 404;
        public static final int METHOD_NOT_ALLOWED = 405;
        public static final int CONFLICT = 409;
        public static final int UNPROCESSABLE_ENTITY = 422;
        public static final int TOO_MANY_REQUESTS = 429;
        public static final int INTERNAL_SERVER_ERROR = 500;
        public static final int BAD_GATEWAY = 502;
        public static final int SERVICE_UNAVAILABLE = 503;
        public static final int GATEWAY_TIMEOUT = 504;

        private HttpStatus() {
        }
    }

    public static final class MimeTypes {
        public static final String JSON = "application/json";
        public static final String FORM_DATA = "multipart/form-data";
        public static final String URL_ENCODED = "application/x-www-form-urlencoded";
        public static final String OCTET_STREAM = "application/octet-stream";
        public static final String TEXT_PLAIN = "text/plain";
        public static final String TEXT_HTML = "text/html";

        private MimeTypes() {
        }
    }

    public static final class DefaultValues {
        public static final int DEFAULT_TIMEOUT = 30000;
        public static final int DEFAULT_MAX_RETRIES = 3;
        public static final int DEFAULT_RETRY_DELAY = 1000;
        public static final int DEFAULT_CACHE_TTL = 300000;
        public static final int DEFAULT_CACHE_MAX_SIZE = 100;

        private DefaultValues() {
        }
    }

    public static final class SuccessCodes {
        public static final List<Object> CODES = List.of(0, 200, 2000, "0", "200", "2000");

        private SuccessCodes() {
        }
    }
}
