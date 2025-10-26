// Model untuk User
class User {
  final String id;
  final String username;
  final String email;
  final String? fullName;
  final String? phone;
  final String role;
  final bool? isActive;
  final String? token;
  final DateTime? createdAt;

  User({
    required this.id,
    required this.username,
    required this.email,
    this.fullName,
    this.phone,
    required this.role,
    this.isActive,
    this.token,
    this.createdAt,
  });

  // Konversi dari JSON ke User object
  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'] ?? '',
      username: json['username'] ?? '',
      email: json['email'] ?? '',
      fullName: json['full_name'],
      phone: json['phone'],
      role: json['role'] ?? 'user',
      isActive: json['is_active'] ?? true,
      token: json['token'],
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'])
          : null,
    );
  }

  // Konversi dari User object ke JSON
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'email': email,
      'full_name': fullName,
      'phone': phone,
      'role': role,
      'is_active': isActive,
      'token': token,
      if (createdAt != null) 'created_at': createdAt!.toIso8601String(),
    };
  }
}
