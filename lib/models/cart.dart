import 'product.dart';

class Cart {
  final String id;
  final String userId;
  final String productId;
  final int quantity;
  final Product product;
  final DateTime createdAt;

  Cart({
    required this.id,
    required this.userId,
    required this.productId,
    required this.quantity,
    required this.product,
    required this.createdAt,
  });

  factory Cart.fromJson(Map<String, dynamic> json) {
    return Cart(
      id: json['id'] ?? '',
      userId: json['user_id'] ?? '',
      productId: json['product_id'] ?? '',
      quantity: json['quantity'] ?? 1,
      product: Product.fromJson(json['product'] ?? {}),
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'])
          : DateTime.now(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'product_id': productId,
      'quantity': quantity,
      'product': product.toJson(),
      'created_at': createdAt.toIso8601String(),
    };
  }

  double get subtotal => product.price * quantity;

  String get formattedSubtotal {
    return 'Rp ${subtotal.toStringAsFixed(0).replaceAllMapped(
          RegExp(r'(\d{1,3})(?=(\d{3})+(?!\d))'),
          (Match m) => '${m[1]}.',
        )}';
  }
}



