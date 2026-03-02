from typing import Any, Optional, Callable, Dict
import time
import logging
from functools import wraps
from datetime import datetime, timedelta

from .types import (
    RetryConfig,
    RetryBackoff,
    CacheConfig,
    LoggerConfig,
    DEFAULT_RETRY_CONFIG,
    DEFAULT_CACHE_CONFIG,
    DEFAULT_LOGGER_CONFIG,
)
from .errors import SdkError, is_retryable_error


def with_retry(
    func: Callable[..., Any],
    config: Optional[RetryConfig] = None
) -> Callable[..., Any]:
    cfg = config or DEFAULT_RETRY_CONFIG
    
    @wraps(func)
    def wrapper(*args, **kwargs):
        last_exception = None
        for attempt in range(cfg.max_retries + 1):
            try:
                return func(*args, **kwargs)
            except Exception as e:
                last_exception = e
                if attempt >= cfg.max_retries:
                    break
                if cfg.retry_condition and not cfg.retry_condition(e, attempt):
                    break
                delay = calculate_delay(attempt, cfg)
                time.sleep(delay / 1000)
        raise last_exception
    
    return wrapper


def calculate_delay(
    retry_count: int,
    config: RetryConfig
) -> float:
    if config.retry_backoff == RetryBackoff.FIXED:
        return config.retry_delay
    elif config.retry_backoff == RetryBackoff.LINEAR:
        return config.retry_delay * (retry_count + 1)
    else:
        delay = config.retry_delay * (2 ** retry_count)
        return min(delay, config.max_retry_delay)


def create_retry_config(
    max_retries: int = 3,
    retry_delay: float = 1.0,
    retry_backoff: RetryBackoff = RetryBackoff.EXPONENTIAL,
    max_retry_delay: float = 30.0,
    retry_condition: Optional[Callable[[Exception, int], bool]] = None
) -> RetryConfig:
    return RetryConfig(
        max_retries=max_retries,
        retry_delay=retry_delay,
        retry_backoff=retry_backoff,
        max_retry_delay=max_retry_delay,
        retry_condition=retry_condition,
    )


def sleep(ms: int) -> None:
    time.sleep(ms / 1000)


class CacheStore:
    def __init__(self, config: Optional[CacheConfig] = None):
        self.config = config or DEFAULT_CACHE_CONFIG
        self._cache: Dict[str, tuple] = {}

    def get(self, key: str) -> Optional[Any]:
        if not self.config.enabled:
            return None
        if key in self._cache:
            value, timestamp = self._cache[key]
            if int(time.time() * 1000) - timestamp < self.config.ttl:
                return value
            del self._cache[key]
        return None

    def set(self, key: str, value: Any) -> None:
        if not self.config.enabled:
            return
        if len(self._cache) >= self.config.max_size:
            oldest_key = min(self._cache.keys(), key=lambda k: self._cache[k][1])
            del self._cache[oldest_key]
        self._cache[key] = (value, int(time.time() * 1000))

    def delete(self, key: str) -> None:
        if key in self._cache:
            del self._cache[key]

    def clear(self) -> None:
        self._cache.clear()


def generate_cache_key(*args, **kwargs) -> str:
    import hashlib
    import json
    key_data = {'args': args, 'kwargs': kwargs}
    return hashlib.md5(json.dumps(key_data, sort_keys=True).encode()).hexdigest()


class Logger:
    def __init__(self, config: Optional[LoggerConfig] = None):
        self.config = config or DEFAULT_LOGGER_CONFIG
        self._logger = logging.getLogger('sdkwork')
        self._logger.setLevel(getattr(logging, self.config.level.upper()))

    def debug(self, message: str, **kwargs) -> None:
        self._logger.debug(self._format(message, kwargs))

    def info(self, message: str, **kwargs) -> None:
        self._logger.info(self._format(message, kwargs))

    def warn(self, message: str, **kwargs) -> None:
        self._logger.warning(self._format(message, kwargs))

    def error(self, message: str, **kwargs) -> None:
        self._logger.error(self._format(message, kwargs))

    def _format(self, message: str, kwargs: Dict) -> str:
        parts = [message]
        if kwargs:
            parts.append(str(kwargs))
        return f'{self.config.prefix} {" ".join(parts)}'


def create_logger(config: Optional[LoggerConfig] = None) -> Logger:
    return Logger(config)


def noop_logger() -> Logger:
    return Logger(LoggerConfig(level='silent'))


__all__ = [
    'with_retry',
    'calculate_delay',
    'create_retry_config',
    'sleep',
    'CacheStore',
    'generate_cache_key',
    'Logger',
    'create_logger',
    'noop_logger',
]
