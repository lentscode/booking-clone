import "dart:convert";

import "package:booking_client/api/api.dart";
import "package:booking_client/utils/result.dart";

class Auth with Api {
  AsyncResult<void, AuthError> login(String email, String password) async {
    final body = {"email": email, "password": password};

    final response = await client.post(
      getUrl("/login"),
      body: jsonEncode(body),
    );

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        AuthError("failed to login", StackTrace.current),
      );
    }

    final sessionId =
        response.headers["set-cookie"]?.split(";")[0].split("=")[1];

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

    final response = await client.post(
      getUrl("/signup"),
      body: jsonEncode(body),
    );

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        AuthError("failed to sign up", StackTrace.current),
      );
    }

    final sessionId =
        response.headers["set-cookie"]?.split(";")[0].split("=")[1];

    await setSessionId(sessionId);

    return AsyncResultExtension.voidSuccess();
  }
}
