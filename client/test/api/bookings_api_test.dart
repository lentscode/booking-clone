import "package:booking_client/api/bookings_api.dart";
import "package:dio/dio.dart";
import "package:flutter_test/flutter_test.dart";
import "package:http_mock_adapter/http_mock_adapter.dart";
import "package:mocktail/mocktail.dart";

import "../mocks.dart";

void main() {
  group("BookingsApi", () {
    late Dio dio;
    late DioAdapter dioAdapter;
    late MockSessionManager sessionManager;

    setUp(() {
      dio = Dio();
      dioAdapter = DioAdapter(dio: dio);
      sessionManager = MockSessionManager();

      when(
        () => sessionManager.getSession(),
      ).thenAnswer((_) => Future.value("token"));
    });

    test("getBookingsOfUser", () async {
      dioAdapter.onGet(
        "/bookings",
        (request) => request.reply(200, [
          {
            "id": 1,
            "check_in_date": "2024-01-01",
            "check_out_date": "2024-01-02",
            "total_price": 100.0,
            "status": "pending",
            "host": {
              "id": 1,
              "name": "John Doe",
              "location": "123 Main St, Anytown, USA",
              "rating": 4.5,
              "description": "John is a great host",
              "capacity": 2,
              "price": 100.0,
            },
          },
          {
            "id": 2,
            "check_in_date": "2024-01-03",
            "check_out_date": "2024-01-04",
            "total_price": 200.0,
            "status": "confirmed",
            "host": {
              "id": 2,
              "name": "Jane Doe",
              "location": "123 Main St, Anytown, USA",
              "rating": 4.5,
              "description": "Jane is a great host",
              "capacity": 2,
              "price": 100.0,
            },
          },
        ]),
      );

      final bookingsApi = BookingsApi(dio, sessionManager);
      final bookings = await bookingsApi.getBookingsOfUser();

      expect(bookings.$1, hasLength(2));
    });
  });
}
