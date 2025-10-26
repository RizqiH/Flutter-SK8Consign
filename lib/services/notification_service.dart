import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/notification.dart';
import 'storage_service.dart';

class NotificationService {
  static const String baseUrl = 'http://localhost:8080/api';

  Future<Map<String, dynamic>> getNotifications() async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.get(
        Uri.parse('$baseUrl/notifications'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        List<AppNotification> notifications = [];
        if (data['data'] != null && data['data']['notifications'] != null) {
          for (var item in data['data']['notifications']) {
            notifications.add(AppNotification.fromJson(item));
          }
        }
        return {'success': true, 'notifications': notifications};
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to get notifications',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> markAsRead(String notificationId) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.put(
        Uri.parse('$baseUrl/notifications/read?id=$notificationId'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);
      return {
        'success': response.statusCode == 200,
        'message': data['message'],
      };
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> markAllAsRead() async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.put(
        Uri.parse('$baseUrl/notifications/read-all'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);
      return {
        'success': response.statusCode == 200,
        'message': data['message'],
      };
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> getUnreadCount() async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'count': 0};
      }

      final response = await http.get(
        Uri.parse('$baseUrl/notifications/unread-count'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        return {
          'success': true,
          'count': data['data']['unread_count'] ?? 0,
        };
      } else {
        return {'success': false, 'count': 0};
      }
    } catch (e) {
      return {'success': false, 'count': 0};
    }
  }
}




