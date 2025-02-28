import "package:bloc/bloc.dart";
import "package:booking_client/api/auth.dart";
import "package:booking_client/utils/result.dart";
import "package:flutter/foundation.dart";

enum AuthType {
  login("Login"),
  signUp("Sign Up");

  const AuthType(this.text);

  final String text;
}

@immutable
sealed class AuthState {
  const AuthState({required this.type});

  final AuthType type;

  AuthState copyWith({AuthType? type});
}

class AuthInitial extends AuthState {
  const AuthInitial({required super.type});

  @override
  AuthState copyWith({AuthType? type}) => AuthInitial(type: type ?? this.type);
}

class AuthLoading extends AuthState {
  const AuthLoading({required super.type});

  @override
  AuthState copyWith({AuthType? type}) => AuthLoading(type: type ?? this.type);
}

class AuthSuccess extends AuthState {
  const AuthSuccess({required super.type});

  @override
  AuthState copyWith({AuthType? type}) => AuthSuccess(type: type ?? this.type);
}

class AuthFailure extends AuthState {
  const AuthFailure({required super.type, required this.error});
  final AuthError error;

  @override
  AuthState copyWith({AuthType? type, AuthError? error}) =>
      AuthFailure(type: type ?? this.type, error: error ?? this.error);
}

class AuthCubit extends Cubit<AuthState> {
  AuthCubit(this._auth) : super(const AuthInitial(type: AuthType.login));

  final Auth _auth;

  Future<void> login(String email, String password) async {
    emit(AuthLoading(type: state.type));

    final result = await _auth.login(email, password);
    if (result.isFailure) {
      emit(AuthFailure(type: state.type, error: result.error));
    } else {
      emit(AuthSuccess(type: state.type));
    }
  }

  Future<void> signUp(
    String email,
    String password,
    String confirmPassword,
  ) async {
    emit(AuthLoading(type: state.type));

    final result = await _auth.signUp(email, password, confirmPassword);
    if (result.isFailure) {
      emit(AuthFailure(type: state.type, error: result.error));
    } else {
      emit(AuthSuccess(type: state.type));
    }
  }

  void changeType() {
    emit(
      state.copyWith(
        type: state.type == AuthType.login ? AuthType.signUp : AuthType.login,
      ),
    );
  }
}
