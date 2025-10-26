class Product {
  final String id;
  final String userId;
  final String name;
  final String description;
  final double price;
  final String category;
  final String condition;
  final String status;
  final String? imageUrl;
  final int viewCount;
  final bool isActive;
  final DateTime createdAt;
  final DateTime updatedAt;
  final String? sellerName;
  final String? sellerUsername;

  Product({
    required this.id,
    required this.userId,
    required this.name,
    required this.description,
    required this.price,
    required this.category,
    required this.condition,
    required this.status,
    this.imageUrl,
    required this.viewCount,
    required this.isActive,
    required this.createdAt,
    required this.updatedAt,
    this.sellerName,
    this.sellerUsername,
  });

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'] ?? '',
      userId: json['user_id'] ?? '',
      name: json['name'] ?? '',
      description: json['description'] ?? '',
      price: (json['price'] ?? 0).toDouble(),
      category: json['category'] ?? '',
      condition: json['condition'] ?? '',
      status: json['status'] ?? '',
      imageUrl: json['image_url'],
      viewCount: json['view_count'] ?? 0,
      isActive: json['is_active'] ?? true,
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'])
          : DateTime.now(),
      updatedAt: json['updated_at'] != null
          ? DateTime.parse(json['updated_at'])
          : DateTime.now(),
      sellerName: json['seller_name'],
      sellerUsername: json['seller_username'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'name': name,
      'description': description,
      'price': price,
      'category': category,
      'condition': condition,
      'status': status,
      'image_url': imageUrl,
      'view_count': viewCount,
      'is_active': isActive,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
      'seller_name': sellerName,
      'seller_username': sellerUsername,
    };
  }

  String get formattedPrice {
    return 'Rp ${price.toStringAsFixed(0).replaceAllMapped(
          RegExp(r'(\d{1,3})(?=(\d{3})+(?!\d))'),
          (Match m) => '${m[1]}.',
        )}';
  }

  String get categoryDisplay {
    switch (category.toLowerCase()) {
      case 'gaming':
        return 'Gaming';
      case 'laptop':
        return 'Laptop';
      case 'phone':
        return 'Phone';
      case 'audio':
        return 'Audio';
      case 'camera':
        return 'Camera';
      case 'watch':
        return 'Watch';
      case 'tablet':
        return 'Tablet';
      case 'accessories':
        return 'Accessories';
      default:
        return category;
    }
  }

  String get conditionDisplay {
    switch (condition.toLowerCase()) {
      case 'new':
        return 'New';
      case 'like_new':
        return 'Like New';
      case 'good':
        return 'Good';
      case 'fair':
        return 'Fair';
      default:
        return condition;
    }
  }

  String get statusDisplay {
    switch (status.toLowerCase()) {
      case 'available':
        return 'Available';
      case 'sold':
        return 'Sold';
      case 'reserved':
        return 'Reserved';
      default:
        return status;
    }
  }
}
