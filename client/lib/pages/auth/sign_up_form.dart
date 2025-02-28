import "package:flutter/material.dart";

class SignUpForm extends StatefulWidget {
  const SignUpForm({super.key});

  @override
  State<SignUpForm> createState() => _SignUpFormState();
}

class _SignUpFormState extends State<SignUpForm> {
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) => Column(
    children: [
      Form(
        key: _formKey,
        child: Column(
          children: [TextFormField(decoration: const InputDecoration())],
        ),
      ),
    ],
  );
}
