import "dart:convert";

import "package:booking_client/api/api.dart";
import "package:booking_client/models/booking.dart";
import "package:booking_client/utils/result.dart";

class BookingsApi with Api {
  AsyncResult<List<Booking>, ApiError> getBookingsOfUser() async {
    final response = await client.get(
      getUrl("/bookings"),
      headers: await sessionId,
    );

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        ApiError("failed to get bookings", StackTrace.current),
      );
    }

    final json = jsonDecode(response.body) as List<dynamic>;

    return AsyncResultExtension.success(
      json.map((e) => Booking.fromMap(e)).toList(),
    );
  }

  AsyncResult<Booking, ApiError> getBooking(int id) async {
    final response = await client.get(
      getUrl("/bookings/$id"),
      headers: await sessionId,
    );

    if (response.statusCode != 200) {
      return AsyncResultExtension.failure(
        ApiError("failed to get booking", StackTrace.current),
      );
    }

    return AsyncResultExtension.success(Booking.fromJson(response.body));
  }
}
