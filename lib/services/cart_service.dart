import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/cart.dart';
import 'storage_service.dart';

class CartService {
  static const String baseUrl = 'http://localhost:8080/api';

  Future<Map<String, dynamic>> addToCart(String productId, int quantity) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.post(
        Uri.parse('$baseUrl/cart/add'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({
          'product_id': productId,
          'quantity': quantity,
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        return {
          'success': true,
          'message': data['message'],
        };
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to add to cart',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> getCart() async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.get(
        Uri.parse('$baseUrl/cart'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200) {
        List<Cart> carts = [];
        if (data['data'] != null) {
          for (var item in data['data']) {
            carts.add(Cart.fromJson(item));
          }
        }
        return {'success': true, 'carts': carts};
      } else {
        return {
          'success': false,
          'message': data['message'] ?? 'Failed to get cart',
        };
      }
    } catch (e) {
      return {'success': false, 'message': 'Connection error: $e'};
    }
  }

  Future<Map<String, dynamic>> updateCart(String cartId, int quantity) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.put(
        Uri.parse('$baseUrl/cart/update?id=$cartId'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({'quantity': quantity}),
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

  Future<Map<String, dynamic>> removeFromCart(String cartId) async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.delete(
        Uri.parse('$baseUrl/cart/remove?id=$cartId'),
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

  Future<Map<String, dynamic>> clearCart() async {
    try {
      final token = await StorageService.getTokenStatic();
      if (token == null) {
        return {'success': false, 'message': 'Not authenticated'};
      }

      final response = await http.delete(
        Uri.parse('$baseUrl/cart/clear'),
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
}




