import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/order.dart';
import 'storage_service.dart';

class OrderService {
  static const String baseUrl = 'http://localhost:8080/api';

  Future<Map<String, dynamic>> createOrder({
    required String paymentMethod,
    required String shippingAddress,
    String? notes,
  }) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.post(
        Uri.parse('$baseUrl/orders/create'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({
          'payment_method': paymentMethod,
          'shipping_address': shippingAddress,
          'notes': notes ?? '',
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 201) {
        return {
          'success': true,
          'message': data['message'],
          'order': Order.fromJson(data['data']),
        };
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to create order',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> getOrders({String? status}) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      String url = '$baseUrl/orders';
      if (status != null && status.isNotEmpty) {
        url += '?status=$status';
      }

      final response = await http.get(
        Uri.parse(url),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        List<Order> orders = [];
        if (data['data'] != null && data['data']['orders'] != null) {
          for (var item in data['data']['orders']) {
            orders.add(Order.fromJson(item));
          }
        }
        return {'success': true, 'orders': orders};
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to get orders',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> getOrderDetail(String orderId) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.get(
        Uri.parse('$baseUrl/orders/detail?id=$orderId'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        return {
          'success': true,
          'order': Order.fromJson(data['data']),
        };
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to get order',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> updatePaymentStatus(
      String orderId, String paymentStatus) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.put(
        Uri.parse('$baseUrl/orders/update-payment?id=$orderId'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({'payment_status': paymentStatus}),
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
}




