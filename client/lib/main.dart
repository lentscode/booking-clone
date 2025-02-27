import "package:booking_client/pages/auth.dart";
import "package:booking_client/pages/home.dart";
import "package:flutter/material.dart";
import "package:go_router/go_router.dart";

void main() {
  runApp(const MainApp());
}

final _router = GoRouter(
  routes: [
    GoRoute(path: "/", builder: (context, state) => const HomePage()),
    GoRoute(path: "/auth", builder: (context, state) => const AuthPage()),
  ],
  initialLocation: "/",
);

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) =>
      WidgetsApp.router(routerConfig: _router, color: const Color(0xFFFFFFFF));
}
