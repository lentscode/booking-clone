import "dart:convert";

import "package:booking_client/api/api.dart";
import "package:booking_client/utils/result.dart";

class Auth extends Api {
  const Auth(super.client, super.sessionManager);

  AsyncResult<void, AuthError> login(String email, String password) async {
    final body = {"email": email, "password": password};

    final response = await client.post("/login", data: jsonEncode(body));

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        AuthError("failed to login", StackTrace.current),
      );
    }

    final sessionId =
        response.headers["set-cookie"]?.first.split(";")[0].split("=")[1];

    await setSessionId(sessionId);

    return AsyncResultExtension.voidSuccess();
  }

  AsyncResult<void, AuthError> signUp(
    String email,
    String password,
    String confirmPassword,
  ) async {
    final body = {
      "email": email,
      "password": password,
      "confirm_password": confirmPassword,
    };

    final response = await client.post("/signup", data: jsonEncode(body));

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        AuthError("failed to sign up", StackTrace.current),
      );
    }

    final sessionId =
        response.headers["set-cookie"]?.first.split(";")[0].split("=")[1];

    await setSessionId(sessionId);

    return AsyncResultExtension.voidSuccess();
  }
}
