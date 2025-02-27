import "dart:convert";

import "package:booking_client/models/host.dart";

class Booking {
  Booking.fromMap(Map<String, dynamic> map)
    : id = map["id"],
      checkInDate = DateTime.parse(map["check_in_date"]),
      checkOutDate = DateTime.parse(map["check_out_date"]),
      totalPrice = map["total_price"],
      status = map["status"],
      host = Host.fromMap(map["host"]);

  factory Booking.fromJson(String json) => Booking.fromMap(jsonDecode(json));

  final int id;
  final DateTime checkInDate;
  final DateTime checkOutDate;
  final double totalPrice;
  final String status;
  final Host host;
}
