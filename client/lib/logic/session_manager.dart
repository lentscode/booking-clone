import "package:flutter_secure_storage/flutter_secure_storage.dart";

class SessionManager {
  SessionManager._(this._storage);

  factory SessionManager() {
    _instance ??= SessionManager._(const FlutterSecureStorage());
    return _instance!;
  }
  static SessionManager? _instance;

  final FlutterSecureStorage _storage;

  Future<void> setSession(String session) async {
    await _storage.write(key: "sessionId", value: session);
  }

  Future<String?> getSession() async => await _storage.read(key: "sessionId");

  Future<void> deleteSession() async {
    await _storage.delete(key: "sessionId");
  }
}
