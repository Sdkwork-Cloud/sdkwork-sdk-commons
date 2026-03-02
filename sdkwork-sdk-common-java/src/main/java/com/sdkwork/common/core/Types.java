package com.sdkwork.common.core;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public final class Types {

    private Types() {
    }

    public record ApiResult(
        Object code,
        Object data,
        String msg,
        String message,
        Long timestamp,
        String traceId
    ) {
    }

    public record PageResult(
        List<Object> content,
        List<Object> list,
        int total,
        Integer totalElements,
        int page,
        int pageSize,
        Integer size,
        int totalPages,
        boolean hasMore,
        Boolean first,
        Boolean last,
        Boolean empty,
        Integer number
    ) {
    }

    public record Pageable(
        Integer page,
        Integer pageSize,
        Integer size,
        String sort,
        String order
    ) {
        public Pageable() {
            this(1, 10, null, null, null);
        }
    }

    public record RetryConfig(
        int maxRetries,
        int retryDelay,
        String retryBackoff,
        int maxRetryDelay
    ) {
        public static RetryConfig defaults() {
            return new RetryConfig(3, 1000, "exponential", 30000);
        }
    }

    public record CacheConfig(
        boolean enabled,
        int ttl,
        int maxSize
    ) {
        public static CacheConfig defaults() {
            return new CacheConfig(false, 300000, 100);
        }
    }

    public record LoggerConfig(
        String level,
        String prefix,
        boolean timestamp,
        boolean colors
    ) {
        public static LoggerConfig defaults() {
            return new LoggerConfig("info", "[SDK]", true, true);
        }
    }

    public record HttpClientConfig(
        String baseUrl,
        Integer timeout,
        Map<String, String> headers,
        RetryConfig retry,
        CacheConfig cache,
        LoggerConfig logger
    ) {
        public HttpClientConfig {
            headers = headers != null ? headers : new HashMap<>();
        }
    }

    public record SdkConfig(
        String baseUrl,
        Integer timeout,
        Map<String, String> headers,
        RetryConfig retry,
        CacheConfig cache,
        LoggerConfig logger,
        String apiKey,
        String authToken,
        String accessToken,
        String tenantId,
        String organizationId,
        String platform,
        Object userId
    ) {
        public SdkConfig {
            headers = headers != null ? headers : new HashMap<>();
        }

        public SdkConfig(String baseUrl) {
            this(baseUrl, null, new HashMap<>(), null, null, null, null, null, null, null, null, null, null);
        }
    }
}
