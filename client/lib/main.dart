import "package:booking_client/api/auth.dart";
import "package:booking_client/api/bookings_api.dart";
import "package:booking_client/api/host_api.dart";
import "package:booking_client/components/providers/api.dart";
import "package:booking_client/logic/session_manager.dart";
import "package:booking_client/pages/auth/page.dart";
import "package:booking_client/pages/home/page.dart";
import "package:dio/dio.dart";
import "package:flutter/material.dart";
import "package:go_router/go_router.dart";

void main() {
  runApp(MainApp());
}

final _router = GoRouter(
  routes: [
    GoRoute(path: "/", builder: (context, state) => const HomePage()),
    GoRoute(path: "/auth", builder: (context, state) => const AuthPage()),
  ],
  initialLocation: "/",
);

class MainApp extends StatelessWidget {
  MainApp({super.key});

  final client = Dio(
    BaseOptions(baseUrl: const String.fromEnvironment("BASE_URL")),
  );
  final sessionManager = SessionManager();

  @override
  Widget build(BuildContext context) => ApiProvider(
    hostApi: HostApi(client, sessionManager),
    auth: Auth(client, sessionManager),
    bookingsApi: BookingsApi(client, sessionManager),
    child: WidgetsApp.router(
      routerConfig: _router,
      color: const Color(0xFFFFFFFF),
    ),
  );
}
