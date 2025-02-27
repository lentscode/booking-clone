import "dart:convert";

import "package:booking_client/api/api.dart";
import "package:booking_client/models/host.dart";
import "package:booking_client/utils/result.dart";

class HostApi with Api {
  AsyncResult<List<Host>, ApiError> getHosts() async {
    final response = await client.get(getUrl("/hosts"));

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        ApiError("failed to get hosts", StackTrace.current),
      );
    }

    final json = jsonDecode(response.body) as List<dynamic>;

    return AsyncResultExtension.success(
      json.map((e) => Host.fromMap(e)).toList(),
    );
  }

  AsyncResult<Host, ApiError> getHost(int id) async {
    final response = await client.get(getUrl("/hosts/$id"));

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        ApiError("failed to get host", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(Host.fromJson(response.body));
  }
}
