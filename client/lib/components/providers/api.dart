import "package:booking_client/api/auth.dart";
import "package:booking_client/api/bookings_api.dart";
import "package:booking_client/api/host_api.dart";
import "package:flutter/widgets.dart";

class ApiProvider extends InheritedWidget {
  const ApiProvider({
    required super.child,
    required this.hostApi,
    required this.auth,
    required this.bookingsApi,
    super.key,
  });

  final HostApi hostApi;
  final Auth auth;
  final BookingsApi bookingsApi;

  @override
  bool updateShouldNotify(covariant InheritedWidget oldWidget) => false;

  static ApiProvider of(BuildContext context) =>
      context.dependOnInheritedWidgetOfExactType<ApiProvider>()!;
}

extension ApiProviderX on BuildContext {
  ApiProvider get api => ApiProvider.of(this);
  BookingsApi get bookingsApi => api.bookingsApi;
  HostApi get hostApi => api.hostApi;
  Auth get auth => api.auth;
}
