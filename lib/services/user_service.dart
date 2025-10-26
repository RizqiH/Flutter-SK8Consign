import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/user.dart';

class UserService {
  static const String baseUrl = 'http://localhost:8080/api';

  // Get user profile
  static Future<User> getProfile(String userId) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/profile'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'user_id': userId}),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        return User.fromJson(data['data'] as Map<String, dynamic>);
      } else {
        throw Exception(data['message'] ?? 'Failed to get profile');
      }
    } catch (e) {
      throw Exception('Error getting profile: $e');
    }
  }

  // Update profile
  static Future<User> updateProfile({
    required String userId,
    required String fullName,
    required String phone,
    required String email,
  }) async {
    try {
      final response = await http.put(
        Uri.parse('$baseUrl/profile/update'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'user_id': userId,
          'full_name': fullName,
          'phone': phone,
          'email': email,
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        return User.fromJson(data['data'] as Map<String, dynamic>);
      } else {
        throw Exception(data['message'] ?? 'Failed to update profile');
      }
    } catch (e) {
      throw Exception('Error updating profile: $e');
    }
  }

  // Change password
  static Future<void> changePassword({
    required String userId,
    required String oldPassword,
    required String newPassword,
  }) async {
    try {
      final response = await http.put(
        Uri.parse('$baseUrl/profile/change-password'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'user_id': userId,
          'old_password': oldPassword,
          'new_password': newPassword,
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode != 200 || data['success'] != true) {
        throw Exception(data['message'] ?? 'Failed to change password');
      }
    } catch (e) {
      throw Exception('Error changing password: $e');
    }
  }
}
