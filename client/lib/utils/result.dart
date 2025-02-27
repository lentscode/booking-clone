/// Base class for all application errors.
/// Provides a standardized way to handle errors with both an error message and stack trace.
sealed class AppError {
  const AppError(this.error, this.stackTrace);

  final String error;
  final StackTrace stackTrace;
}

/// Represents network-related errors (connectivity, timeout, etc.)
class NetworkError extends AppError {
  NetworkError(super.error, super.stackTrace);
}

/// Represents API-specific errors (invalid responses, server errors, etc.)
class ApiError extends AppError {
  ApiError(super.error, super.stackTrace);
}

class AuthError extends AppError {
  AuthError(super.error, super.stackTrace);
}

/// Represents unexpected or unhandled errors
class UnknownError extends AppError {
  UnknownError(super.error, super.stackTrace);
}

/// A type representing either a success value of type [T] or an error of type [E].
/// Uses Dart's record syntax to provide a lightweight Result type.
typedef Result<T, E> = (T? value, E? error);

/// An asynchronous version of [Result] wrapped in a [Future].
typedef AsyncResult<T, E> = Future<Result<T, E>>;

/// Extension methods for working with [Result] types.
extension ResultExtension<T, E> on Result<T, E> {
  /// Returns true if this Result contains a value and no error.
  bool get isSuccess => $2 == null;

  /// Returns true if this Result contains an error.
  bool get isFailure => $2 != null;

  /// Returns the success value. Throws if this is a failure Result.
  T get value =>
      $1 ?? (throw StateError("Cannot get value from failure Result"));

  /// Returns the error value. Throws if this is a success Result.
  E get error =>
      $2 ?? (throw StateError("Cannot get error from success Result"));

  /// Creates a success Result with the given value.
  static Result<T, E> success<T, E>(T value) => (value, null);

  /// Creates a failure Result with the given error.
  static Result<T, E> failure<T, E>(E error) => (null, error);

  /// Creates a success Result with a void value.
  static Result<void, E> voidSuccess<E>() => (null, null);

  /// Transforms the success value using the given mapping function.
  Result<R, E> map<R>(R Function(T) mapper) =>
      isSuccess ? (mapper(value), null) : (null, error);

  /// Executes one of the given functions based on whether this is a success or failure.
  R fold<R>(R Function(T) onSuccess, R Function(E) onFailure) =>
      isSuccess ? onSuccess(value) : onFailure(error);
}

/// Extension methods for working with [AsyncResult] types.
extension AsyncResultExtension<T, E> on AsyncResult<T, E> {
  /// Creates a success AsyncResult with the given value.
  static AsyncResult<T, E> success<T, E>(T value) async => (value, null);

  /// Creates a failure AsyncResult with the given error.
  static AsyncResult<T, E> failure<T, E>(E error) async => (null, error);

  /// Creates a success AsyncResult with a void value.
  static AsyncResult<void, E> voidSuccess<E>() async => (null, null);

  /// Transforms the success value using the given asynchronous mapping function.
  AsyncResult<R, E> map<R>(Future<R> Function(T) mapper) async {
    final result = await this;
    if (result.isSuccess) {
      try {
        final mappedValue = await mapper(result.value);
        return (mappedValue, null);
      } catch (e, st) {
        return (null, UnknownError(e.toString(), st) as E);
      }
    }
    return (null, result.error);
  }

  /// Executes one of the given functions based on whether this is a success or failure.
  Future<R> fold<R>(
    Future<R> Function(T) onSuccess,
    Future<R> Function(E) onFailure,
  ) async {
    final result = await this;
    return result.isSuccess
        ? await onSuccess(result.value)
        : await onFailure(result.error);
  }
}
