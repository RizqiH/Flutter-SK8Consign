import 'product.dart';

class Order {
  final String id;
  final String userId;
  final double totalAmount;
  final String status;
  final String paymentMethod;
  final String paymentStatus;
  final String shippingAddress;
  final String notes;
  final List<OrderItem> orderItems;
  final DateTime createdAt;

  Order({
    required this.id,
    required this.userId,
    required this.totalAmount,
    required this.status,
    required this.paymentMethod,
    required this.paymentStatus,
    required this.shippingAddress,
    required this.notes,
    required this.orderItems,
    required this.createdAt,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    var itemsList = json['order_items'] as List? ?? [];
    List<OrderItem> items =
        itemsList.map((item) => OrderItem.fromJson(item)).toList();

    return Order(
      id: json['id'] ?? '',
      userId: json['user_id'] ?? '',
      totalAmount: (json['total_amount'] ?? 0).toDouble(),
      status: json['status'] ?? '',
      paymentMethod: json['payment_method'] ?? '',
      paymentStatus: json['payment_status'] ?? '',
      shippingAddress: json['shipping_address'] ?? '',
      notes: json['notes'] ?? '',
      orderItems: items,
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'])
          : DateTime.now(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'total_amount': totalAmount,
      'status': status,
      'payment_method': paymentMethod,
      'payment_status': paymentStatus,
      'shipping_address': shippingAddress,
      'notes': notes,
      'order_items': orderItems.map((item) => item.toJson()).toList(),
      'created_at': createdAt.toIso8601String(),
    };
  }

  String get formattedTotal {
    return 'Rp ${totalAmount.toStringAsFixed(0).replaceAllMapped(
          RegExp(r'(\d{1,3})(?=(\d{3})+(?!\d))'),
          (Match m) => '${m[1]}.',
        )}';
  }

  String get statusDisplay {
    switch (status.toLowerCase()) {
      case 'pending':
        return 'Pending';
      case 'confirmed':
        return 'Confirmed';
      case 'shipped':
        return 'Shipped';
      case 'delivered':
        return 'Delivered';
      case 'cancelled':
        return 'Cancelled';
      default:
        return status;
    }
  }
}

class OrderItem {
  final String id;
  final String productId;
  final int quantity;
  final double price;
  final double subtotal;
  final Product product;

  OrderItem({
    required this.id,
    required this.productId,
    required this.quantity,
    required this.price,
    required this.subtotal,
    required this.product,
  });

  factory OrderItem.fromJson(Map<String, dynamic> json) {
    return OrderItem(
      id: json['id'] ?? '',
      productId: json['product_id'] ?? '',
      quantity: json['quantity'] ?? 1,
      price: (json['price'] ?? 0).toDouble(),
      subtotal: (json['subtotal'] ?? 0).toDouble(),
      product: Product.fromJson(json['product'] ?? {}),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'product_id': productId,
      'quantity': quantity,
      'price': price,
      'subtotal': subtotal,
      'product': product.toJson(),
    };
  }
}



