import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/user.dart';

class ApiService {
  // GANTI dengan URL backend Go Anda
  // Kalau run di emulator/simulator, gunakan:
  // - Android Emulator: http://10.0.2.2:8080
  // - iOS Simulator: http://localhost:8080 atau http://127.0.0.1:8080
  // - Device fisik: http://YOUR_COMPUTER_IP:8080
  static const String baseUrl = 'http://localhost:8080/api';

  // Login
  Future<Map<String, dynamic>> login(String username, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'username': username, 'password': password}),
      );

      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');

      if (response.statusCode == 200) {
        // Login berhasil
        final data = jsonDecode(response.body);
        
        // Response structure baru: { success, message, token, data: { user: {...} } }
        return {
          'success': true,
          'message': data['message'] ?? 'Login berhasil',
          'user': User.fromJson(data['data']['user']), // ⚠️ Perhatikan 'data.user'
          'token': data['token'],
        };
      } else {
        // Login gagal
        final data = jsonDecode(response.body);
        return {'success': false, 'message': data['message'] ?? 'Login gagal'};
      }
    } catch (e) {
      print('Error: $e');
      return {'success': false, 'message': 'Gagal terhubung ke server: $e'};
    }
  }

  // Register (untuk nanti)
  Future<Map<String, dynamic>> register({
    required String username,
    required String email,
    required String password,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/register'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'username': username,
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 201 || response.statusCode == 200) {
        final data = jsonDecode(response.body);
        return {
          'success': true,
          'message': data['message'] ?? 'Register berhasil',
        };
      } else {
        final data = jsonDecode(response.body);
        return {
          'success': false,
          'message': data['message'] ?? 'Register gagal',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Gagal terhubung ke server: $e'};
    }
  }

  // Get user profile (dengan token)
  Future<Map<String, dynamic>> getUserProfile(String token) async {
    try {
      final response = await http.get(
        Uri.parse('$baseUrl/profile'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        return {'success': true, 'user': User.fromJson(data['user'])};
      } else {
        return {'success': false, 'message': 'Gagal mengambil data profile'};
      }
    } catch (e) {
      return {'success': false, 'message': 'Gagal terhubung ke server: $e'};
    }
  }
}
