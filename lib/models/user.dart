// Model untuk User
class User {
  final String id;
  final String username;
  final String email;
  final String? fullName;
  final String? phone;
  final String? role;
  final bool? isActive;
  final String? token;

  User({
    required this.id,
    required this.username,
    required this.email,
    this.fullName,
    this.phone,
    this.role,
    this.isActive,
    this.token,
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
    };
  }
}
