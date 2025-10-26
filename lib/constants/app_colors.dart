import 'package:flutter/material.dart';

class AppColors {
  // Primary Colors
  static const Color primary = Color(0xFF6B5FED);
  static const Color secondary = Color(0xFFE94DA3);
  
  // Background Colors
  static const Color background = Color(0xFF1A1A2E);
  static const Color surface = Color(0xFF2A2A4E);
  static const Color surfaceLight = Color(0xFF3A3A5E);
  
  // Text Colors
  static const Color textPrimary = Colors.white;
  static final Color textSecondary = Colors.white.withOpacity(0.7);
  static final Color textTertiary = Colors.white.withOpacity(0.5);
  
  // Accent Colors
  static const Color success = Color(0xFF4CAF50);
  static const Color error = Color(0xFFE53935);
  static const Color warning = Color(0xFFFFA726);
  static const Color info = Color(0xFF29B6F6);
  
  // Gradients
  static const LinearGradient primaryGradient = LinearGradient(
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
    colors: [primary, secondary],
  );
  
  static final LinearGradient bannerGradient = LinearGradient(
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
    colors: [
      Colors.purple.withOpacity(0.6),
      Colors.blue.withOpacity(0.6),
      Colors.pink.withOpacity(0.6),
    ],
  );
  
  // Overlay Colors
  static final Color overlay = Colors.black.withOpacity(0.5);
  static final Color overlayLight = Colors.black.withOpacity(0.3);
  
  // Border Colors
  static final Color border = Colors.white.withOpacity(0.1);
  static final Color borderLight = Colors.white.withOpacity(0.05);
}

