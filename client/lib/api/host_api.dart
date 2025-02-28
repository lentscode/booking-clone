import "package:booking_client/api/api.dart";
import "package:booking_client/models/host.dart";
import "package:booking_client/utils/result.dart";

class HostApi extends Api {
  const HostApi(super.client, super.sessionManager);

  AsyncResult<List<Host>, ApiError> getHosts() async {
    final response = await client.get<List<dynamic>>("/hosts");

    if (response.statusCode != 200 || response.data == null) {
      return AsyncResultExtension.failure(
        ApiError("failed to get hosts", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(
      response.data!.map((e) => Host.fromMap(e)).toList(),
    );
  }

  AsyncResult<Host, ApiError> getHost(int id) async {
    final response = await client.get<Map<String, dynamic>>("/hosts/$id");

    if (response.statusCode != 200 || response.data == null) {
      return AsyncResultExtension.failure(
        ApiError("failed to get host", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(Host.fromMap(response.data!));
  }
}
