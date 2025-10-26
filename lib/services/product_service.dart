import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/product.dart';

class ProductService {
  static const String baseUrl = 'http://localhost:8080/api';

  // Search products
  static Future<Map<String, dynamic>> searchProducts({
    String query = '',
    String category = 'all',
    double minPrice = 0,
    double maxPrice = 0,
    String status = 'available',
    int page = 1,
    int limit = 20,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/products/search'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'query': query,
          'category': category == 'all' ? '' : category,
          'min_price': minPrice,
          'max_price': maxPrice,
          'status': status,
          'page': page,
          'limit': limit,
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        final productsData = data['data']['products'] as List? ?? [];
        final products = productsData
            .map((json) => Product.fromJson(json as Map<String, dynamic>))
            .toList();

        return {
          'success': true,
          'products': products,
          'total': data['data']['total'] ?? 0,
          'page': data['data']['page'] ?? 1,
          'limit': data['data']['limit'] ?? 20,
        };
      } else {
        throw Exception(data['message'] ?? 'Failed to search products');
      }
    } catch (e) {
      throw Exception('Error searching products: $e');
    }
  }

  // Get product detail
  static Future<Product> getProductDetail(String productId) async {
    try {
      final response = await http.get(
        Uri.parse('$baseUrl/products/detail?id=$productId'),
        headers: {'Content-Type': 'application/json'},
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        return Product.fromJson(data['data'] as Map<String, dynamic>);
      } else {
        throw Exception(data['message'] ?? 'Failed to get product detail');
      }
    } catch (e) {
      throw Exception('Error getting product detail: $e');
    }
  }

  // Get user products
  static Future<Map<String, dynamic>> getUserProducts({
    required String userId,
    String status = 'all',
    int page = 1,
    int limit = 20,
  }) async {
    try {
      final response = await http.get(
        Uri.parse(
            '$baseUrl/products/user?user_id=$userId&status=$status&page=$page&limit=$limit'),
        headers: {'Content-Type': 'application/json'},
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        final productsData = data['data']['products'] as List? ?? [];
        final products = productsData
            .map((json) => Product.fromJson(json as Map<String, dynamic>))
            .toList();

        return {
          'success': true,
          'products': products,
          'total': data['data']['total'] ?? 0,
          'page': data['data']['page'] ?? 1,
          'limit': data['data']['limit'] ?? 20,
        };
      } else {
        throw Exception(data['message'] ?? 'Failed to get user products');
      }
    } catch (e) {
      throw Exception('Error getting user products: $e');
    }
  }

  // Create product
  static Future<Product> createProduct({
    required String name,
    required String description,
    required double price,
    required String category,
    required String condition,
    String? imageUrl,
    required String userId,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/products/create'),
        headers: {
          'Content-Type': 'application/json',
          'X-User-ID': userId,
        },
        body: jsonEncode({
          'name': name,
          'description': description,
          'price': price,
          'category': category,
          'condition': condition,
          'image_url': imageUrl ?? '',
        }),
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 201 && data['success'] == true) {
        return Product.fromJson(data['data'] as Map<String, dynamic>);
      } else {
        throw Exception(data['message'] ?? 'Failed to create product');
      }
    } catch (e) {
      throw Exception('Error creating product: $e');
    }
  }

  // Get categories
  static Future<List<String>> getCategories() async {
    try {
      final response = await http.get(
        Uri.parse('$baseUrl/products/categories'),
        headers: {'Content-Type': 'application/json'},
      );

      final data = jsonDecode(response.body);

      if (response.statusCode == 200 && data['success'] == true) {
        return List<String>.from(data['data'] ?? []);
      } else {
        throw Exception(data['message'] ?? 'Failed to get categories');
      }
    } catch (e) {
      throw Exception('Error getting categories: $e');
    }
  }
}
