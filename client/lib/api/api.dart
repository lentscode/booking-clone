import "package:booking_client/logic/session_manager.dart";
import "package:http/http.dart";

mixin Api {
  final client = Client();

  final _baseUrl = const String.fromEnvironment("BASE_URL");

  Uri getUrl(String path) => Uri.parse("$_baseUrl$path");

  Future<Map<String, String>> get sessionId async {
    final sessionId = await SessionManager().getSession();
    if (sessionId == null) {
      throw Exception("user not authenticated");
    }

    return {"Cookie": "sessionId=$sessionId"};
  }

  Future<void> setSessionId(String? sessionId) async {
    if (sessionId == null) {
      throw Exception("sessionId is null");
    }

    await SessionManager().setSession(sessionId);
  }
}
