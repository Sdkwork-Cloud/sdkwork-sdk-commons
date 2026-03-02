from typing import Any, Optional, Dict, Callable
from dataclasses import dataclass, field
from datetime import datetime, timedelta
from enum import Enum
import time
import threading


class AuthMode(str, Enum):
    NONE = 'none'
    API_KEY = 'api_key'
    BEARER = 'bearer'
    OAUTH2 = 'oauth2'


@dataclass
class AuthTokens:
    access_token: str
    refresh_token: Optional[str] = None
    expires_at: Optional[int] = None
    token_type: str = 'Bearer'
    scope: Optional[str] = None
    extra: Dict[str, Any] = field(default_factory=dict)


class AuthTokenManager:
    def __init__(self):
        self._tokens: Optional[AuthTokens] = None
        self._lock = threading.Lock()
        self._refresh_callback: Optional[Callable[[], AuthTokens]] = None

    def set_tokens(self, tokens: AuthTokens) -> None:
        with self._lock:
            self._tokens = tokens

    def get_tokens(self) -> Optional[AuthTokens]:
        with self._lock:
            return self._tokens

    def get_access_token(self) -> Optional[str]:
        tokens = self.get_tokens()
        if tokens:
            return tokens.access_token
        return None

    def is_token_valid(self) -> bool:
        tokens = self.get_tokens()
        if not tokens:
            return False
        if not tokens.expires_at:
            return True
        return int(time.time() * 1000) < tokens.expires_at

    def requires_refresh(self) -> bool:
        tokens = self.get_tokens()
        if not tokens or not tokens.refresh_token:
            return False
        if not tokens.expires_at:
            return False
        buffer = 5 * 60 * 1000
        return int(time.time() * 1000) >= (tokens.expires_at - buffer)

    async def refresh_if_needed(self) -> Optional[AuthTokens]:
        if self.requires_refresh() and self._refresh_callback:
            tokens = await self._refresh_callback()
            self.set_tokens(tokens)
            return tokens
        return None

    def set_refresh_callback(self, callback: Callable[[], AuthTokens]) -> None:
        self._refresh_callback = callback

    def clear(self) -> None:
        with self._lock:
            self._tokens = None


def build_auth_headers(token: str, mode: AuthMode = AuthMode.BEARER) -> Dict[str, str]:
    if mode == AuthMode.BEARER:
        return {'Authorization': f'Bearer {token}'}
    elif mode == AuthMode.API_KEY:
        return {'Authorization': f'Bearer {token}'}
    else:
        return {}


def is_token_valid(tokens: Optional[AuthTokens]) -> bool:
    if not tokens:
        return False
    if not tokens.expires_at:
        return True
    return int(time.time() * 1000) < tokens.expires_at


def requires_refresh(tokens: Optional[AuthTokens]) -> bool:
    if not tokens or not tokens.refresh_token:
        return False
    if not tokens.expires_at:
        return False
    buffer = 5 * 60 * 1000
    return int(time.time() * 1000) >= (tokens.expires_at - buffer)


def create_token_manager(
    tokens: Optional[AuthTokens] = None,
    refresh_callback: Optional[Callable[[], AuthTokens]] = None
) -> AuthTokenManager:
    manager = AuthTokenManager()
    if tokens:
        manager.set_tokens(tokens)
    if refresh_callback:
        manager.set_refresh_callback(refresh_callback)
    return manager


__all__ = [
    'AuthMode',
    'AuthTokens',
    'AuthTokenManager',
    'build_auth_headers',
    'is_token_valid',
    'requires_refresh',
    'create_token_manager',
]
