import "package:booking_client/logic/session_manager.dart";
import "package:dio/dio.dart";

class Api {
  const Api(this.client, this.sessionManager);
  final Dio client;
  final SessionManager sessionManager;

  Future<Map<String, String>> get sessionId async {
    final sessionId = await sessionManager.getSession();
    if (sessionId == null) {
      throw Exception("user not authenticated");
    }

    return {"Cookie": "sessionId=$sessionId"};
  }

  Future<void> setSessionId(String? sessionId) async {
    if (sessionId == null) {
      throw Exception("sessionId is null");
    }

    await sessionManager.setSession(sessionId);
  }
}
