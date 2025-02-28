import "package:booking_client/api/api.dart";
import "package:booking_client/models/booking.dart";
import "package:booking_client/utils/result.dart";
import "package:dio/dio.dart";

class BookingsApi extends Api {
  const BookingsApi(super.client, super.sessionManager);

  AsyncResult<List<Booking>, ApiError> getBookingsOfUser() async {
    final response = await client.get<List<dynamic>>(
      "/bookings",
      options: Options(headers: await sessionId),
    );

    if (response.statusCode != 200 || response.data == null) {
      return AsyncResultExtension.failure(
        ApiError("failed to get bookings", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(
      response.data!.map((e) => Booking.fromMap(e)).toList(),
    );
  }

  AsyncResult<Booking, ApiError> getBooking(int id) async {
    final response = await client.get<Map<String, dynamic>>(
      "/bookings/$id",
      options: Options(headers: await sessionId),
    );

    if (response.statusCode != 200 || response.data == null) {
      return AsyncResultExtension.failure(
        ApiError("failed to get booking", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(Booking.fromMap(response.data!));
  }
}
