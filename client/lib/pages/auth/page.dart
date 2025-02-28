import "package:booking_client/components/providers/api.dart";
import "package:booking_client/pages/auth/login_form.dart";
import "package:booking_client/pages/auth/sign_up_form.dart";
import "package:booking_client/state/auth.dart";
import "package:flutter/material.dart";
import "package:flutter_bloc/flutter_bloc.dart";

class AuthPage extends StatelessWidget {
  const AuthPage({super.key});

  @override
  Widget build(BuildContext context) => BlocProvider(
    create: (context) => AuthCubit(context.auth),
    child: Scaffold(
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          BlocBuilder<AuthCubit, AuthState>(
            builder: (context, state) => Text(state.type.text),
          ),
          const SizedBox(height: 16),
          BlocBuilder<AuthCubit, AuthState>(
            builder:
                (context, state) => switch (state.type) {
                  AuthType.signUp => const SignUpForm(),
                  AuthType.login => const LoginForm(),
                },
          ),
        ],
      ),
    ),
  );
}
