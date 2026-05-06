from sdkwork.common.core.types import SdkConfig
from sdkwork.common.errors import ErrorCode, NetworkError
from sdkwork.common.http import BaseHttpClient, create_http_client
from sdkwork.common.utils import Logger, create_logger, create_retry_config, noop_logger


def test_common_namespace_imports_resolve():
    config = SdkConfig(base_url="https://api.example.com")
    client = create_http_client(config)
    retry_config = create_retry_config(max_retries=2)
    logger = create_logger()
    silent_logger = noop_logger()

    assert isinstance(client, BaseHttpClient)
    assert retry_config.max_retries == 2
    assert isinstance(logger, Logger)
    assert isinstance(silent_logger, Logger)


def test_error_model_exports_consistent_codes_and_exception_message():
    error = NetworkError("temporary outage")

    assert ErrorCode.NETWORK_ERROR == "NETWORK_ERROR"
    assert error.is_network_error()
    assert str(error) == "temporary outage"
