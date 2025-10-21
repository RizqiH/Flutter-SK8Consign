import 'package:flutter/material.dart';
import 'dart:async';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    // Timer untuk pindah ke halaman login setelah 3 detik
    Timer(const Duration(seconds: 3), () {
      // Navigate ke halaman login
      Navigator.pushReplacementNamed(context, '/login');
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        // Gradient background seperti di gambar
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
            colors: [
              Color(0xFF1a1a2e), // Dark navy purple (atas)
              Color(0xFF2d1b4e), // Medium purple (tengah)
              Color(0xFF4a2c6d), // Lighter purple (bawah)
            ],
          ),
        ),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              // Logo SK8 CONSIGN
              RichText(
                text: const TextSpan(
                  children: [
                    TextSpan(
                      text: 'SK8',
                      style: TextStyle(
                        fontSize: 48,
                        fontWeight: FontWeight.bold,
                        color: Colors.white,
                        letterSpacing: 2,
                      ),
                    ),
                    TextSpan(
                      text: ' CONSIGN',
                      style: TextStyle(
                        fontSize: 48,
                        fontWeight: FontWeight.w300,
                        color: Colors.white,
                        letterSpacing: 2,
                      ),
                    ),
                  ],
                ),
              ),
              // Garis underline di bawah SK8
              Container(
                margin: const EdgeInsets.only(top: 8, right: 195),
                width: 80,
                height: 3,
                color: Colors.white,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
