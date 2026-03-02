import 'dart:convert';
import 'dart:typed_data';

import 'package:http/http.dart' as http;

import '../core/types.dart';

class BaseHttpClient {
  final String baseUrl;
  final int timeout;
  final Map<String, String> headers;
  final http.Client _client;

  String? _apiKey;
  final String _apiKeyHeader;
  final bool _apiKeyAsBearer;
  String? _authToken;
  String? _accessToken;

  BaseHttpClient(SdkConfig config)
      : baseUrl = config.baseUrl.replaceFirst(RegExp(r'/+$'), ''),
        timeout = config.timeout,
        headers = Map<String, String>.from(config.headers),
        _client = http.Client(),
        _apiKey = config.apiKey,
        _apiKeyHeader = config.apiKeyHeader,
        _apiKeyAsBearer = config.apiKeyAsBearer,
        _authToken = config.authToken,
        _accessToken = config.accessToken {
    _applyAuthHeaders();
  }

  void _applyAuthHeaders() {
    headers.remove('Authorization');
    headers.remove('Access-Token');
    headers.remove('X-API-Key');

    if (_apiKey != null && _apiKey!.isNotEmpty) {
      headers[_apiKeyHeader] = _apiKeyAsBearer ? 'Bearer $_apiKey' : _apiKey!;
    }
    if (_authToken != null && _authToken!.isNotEmpty) {
      headers['Authorization'] = 'Bearer $_authToken';
    }
    if (_accessToken != null && _accessToken!.isNotEmpty) {
      headers['Access-Token'] = _accessToken!;
    }
  }

  void setApiKey(String apiKey) {
    _apiKey = apiKey;
    _authToken = null;
    _accessToken = null;
    _applyAuthHeaders();
  }

  void setAuthToken(String token) {
    _authToken = token;
    if (_apiKeyHeader.toLowerCase() != 'authorization') {
      _apiKey = null;
    }
    _applyAuthHeaders();
  }

  void setAccessToken(String token) {
    _accessToken = token;
    if (_apiKeyHeader.toLowerCase() != 'access-token') {
      _apiKey = null;
    }
    _applyAuthHeaders();
  }

  void setHeader(String key, String value) {
    headers[key] = value;
  }

  Uri _buildUri(String path, [Map<String, dynamic>? params]) {
    final normalizedPath = path.startsWith('/') ? path : '/$path';
    var url = Uri.parse('$baseUrl$normalizedPath');
    if (params != null && params.isNotEmpty) {
      url = url.replace(
        queryParameters: params.map((k, v) => MapEntry(k, v?.toString())),
      );
    }
    return url;
  }

  dynamic _parseResponse(http.Response response) {
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw Exception('HTTP ${response.statusCode}: ${response.body}');
    }

    if (response.body.isEmpty) {
      return null;
    }

    final contentType = response.headers['content-type'] ?? '';
    if (contentType.contains('application/json')) {
      return jsonDecode(response.body);
    }
    return response.body;
  }

  Future<dynamic> request(
    String method,
    String path, {
    Map<String, dynamic>? params,
    dynamic body,
    Map<String, String>? requestHeaders,
    String? contentType,
  }) async {
    final uri = _buildUri(path, params);
    final mergedHeaders = <String, String>{
      ...headers,
      ...?requestHeaders,
    };

    Future<http.Response> call;
    switch (method.toUpperCase()) {
      case 'GET':
        call = _client.get(uri, headers: mergedHeaders);
        break;
      case 'POST':
        call = _client.post(
          uri,
          headers: _buildHeaders(mergedHeaders, contentType, body),
          body: _encodeBody(body, contentType),
        );
        break;
      case 'PUT':
        call = _client.put(
          uri,
          headers: _buildHeaders(mergedHeaders, contentType, body),
          body: _encodeBody(body, contentType),
        );
        break;
      case 'PATCH':
        call = _client.patch(
          uri,
          headers: _buildHeaders(mergedHeaders, contentType, body),
          body: _encodeBody(body, contentType),
        );
        break;
      case 'DELETE':
        call = _client.delete(uri, headers: mergedHeaders);
        break;
      default:
        throw ArgumentError('Unsupported HTTP method: $method');
    }

    final response = await call.timeout(Duration(milliseconds: timeout));
    return _parseResponse(response);
  }

  String _normalizeContentType(String? contentType) {
    if (contentType == null || contentType.trim().isEmpty) {
      return 'application/json';
    }
    return contentType.trim();
  }

  Map<String, String> _buildHeaders(
    Map<String, String> mergedHeaders,
    String? contentType,
    dynamic body,
  ) {
    if (body == null) {
      return mergedHeaders;
    }

    final normalized = _normalizeContentType(contentType).toLowerCase();
    if (normalized.startsWith('multipart/form-data')) {
      // Let http.MultipartRequest set its own boundary.
      final copied = Map<String, String>.from(mergedHeaders);
      copied.remove('Content-Type');
      return copied;
    }

    return {
      ...mergedHeaders,
      'Content-Type': _normalizeContentType(contentType),
    };
  }

  dynamic _encodeBody(dynamic body, String? contentType) {
    if (body == null) {
      return null;
    }

    final normalized = _normalizeContentType(contentType).toLowerCase();
    if (normalized.startsWith('application/x-www-form-urlencoded')) {
      final entries = _toFormEntries(body);
      return entries
          .map(
            (entry) =>
                '${Uri.encodeQueryComponent(entry.key)}=${Uri.encodeQueryComponent(entry.value)}',
          )
          .join('&');
    }

    if (normalized.contains('json')) {
      return jsonEncode(body);
    }

    if (body is String || body is List<int>) {
      return body;
    }
    return body.toString();
  }

  List<MapEntry<String, String>> _toFormEntries(dynamic body) {
    final result = <MapEntry<String, String>>[];
    void addValue(String key, dynamic value) {
      if (value == null) {
        result.add(MapEntry(key, ''));
        return;
      }
      if (value is Iterable && value is! String && value is! List<int>) {
        for (final item in value) {
          addValue(key, item);
        }
        return;
      }
      result.add(MapEntry(key, value.toString()));
    }

    if (body is Map) {
      body.forEach((key, value) {
        if (key == null) {
          return;
        }
        addValue(key.toString(), value);
      });
    } else {
      addValue('value', body);
    }

    return result;
  }

  Future<http.Response> _sendMultipart(
    String method,
    Uri uri,
    dynamic body,
    Map<String, String> mergedHeaders,
  ) async {
    final request = http.MultipartRequest(method.toUpperCase(), uri);
    request.headers.addAll(_buildHeaders(mergedHeaders, 'multipart/form-data', body));

    void addField(String key, dynamic value) {
      if (value == null) {
        request.fields[key] = '';
        return;
      }
      if (value is Iterable && value is! String && value is! List<int>) {
        for (final item in value) {
          addField(key, item);
        }
        return;
      }
      if (value is http.MultipartFile) {
        request.files.add(value);
        return;
      }
      if (value is Uint8List || value is List<int>) {
        request.files.add(http.MultipartFile.fromBytes(key, List<int>.from(value as List<int>), filename: key));
        return;
      }
      request.fields[key] = value.toString();
    }

    if (body is Map) {
      body.forEach((key, value) {
        if (key == null) {
          return;
        }
        addField(key.toString(), value);
      });
    } else if (body != null) {
      addField('value', body);
    }

    final streamed = await _client
        .send(request)
        .timeout(Duration(milliseconds: timeout));
    return http.Response.fromStream(streamed);
  }

  Future<dynamic> get(
    String path, {
    Map<String, dynamic>? params,
    Map<String, String>? headers,
  }) {
    return request('GET', path, params: params, requestHeaders: headers);
  }

  Future<dynamic> post(
    String path, {
    dynamic body,
    Map<String, dynamic>? params,
    Map<String, String>? headers,
    String? contentType,
  }) {
    final normalized = _normalizeContentType(contentType).toLowerCase();
    if (normalized.startsWith('multipart/form-data')) {
      final uri = _buildUri(path, params);
      return _sendMultipart('POST', uri, body, {...this.headers, ...?headers})
          .then(_parseResponse);
    }
    return request(
      'POST',
      path,
      params: params,
      body: body,
      requestHeaders: headers,
      contentType: contentType,
    );
  }

  Future<dynamic> put(
    String path, {
    dynamic body,
    Map<String, dynamic>? params,
    Map<String, String>? headers,
    String? contentType,
  }) {
    final normalized = _normalizeContentType(contentType).toLowerCase();
    if (normalized.startsWith('multipart/form-data')) {
      final uri = _buildUri(path, params);
      return _sendMultipart('PUT', uri, body, {...this.headers, ...?headers})
          .then(_parseResponse);
    }
    return request(
      'PUT',
      path,
      params: params,
      body: body,
      requestHeaders: headers,
      contentType: contentType,
    );
  }

  Future<dynamic> patch(
    String path, {
    dynamic body,
    Map<String, dynamic>? params,
    Map<String, String>? headers,
    String? contentType,
  }) {
    final normalized = _normalizeContentType(contentType).toLowerCase();
    if (normalized.startsWith('multipart/form-data')) {
      final uri = _buildUri(path, params);
      return _sendMultipart('PATCH', uri, body, {...this.headers, ...?headers})
          .then(_parseResponse);
    }
    return request(
      'PATCH',
      path,
      params: params,
      body: body,
      requestHeaders: headers,
      contentType: contentType,
    );
  }

  Future<dynamic> delete(
    String path, {
    Map<String, dynamic>? params,
    Map<String, String>? headers,
  }) {
    return request('DELETE', path, params: params, requestHeaders: headers);
  }
}
