# SK8 Consign Backend (Go)

Simple REST API untuk SK8 Consign mobile app.

## 🚀 Cara Menjalankan Backend

### 1. Install Dependencies
```bash
cd backend
go mod download
```

### 2. Run Server
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### 1. Health Check
```
GET /api/health
```

Response:
```json
{
  "status": "ok",
  "message": "SK8 Consign API is running",
  "version": "1.0.0"
}
```

### 2. Login
```
POST /api/login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

Response (Success):
```json
{
  "success": true,
  "message": "Login berhasil",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "1",
    "username": "admin",
    "email": "admin@sk8consign.com"
  }
}
```

Response (Failed):
```json
{
  "success": false,
  "message": "Username atau password salah"
}
```

### 3. Register
```
POST /api/register
Content-Type: application/json

{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

Response (Success):
```json
{
  "success": true,
  "message": "Register berhasil"
}
```

## 👤 Test Users

Sudah tersedia 2 user untuk testing:

1. **Admin**
   - Username: `admin`
   - Password: `admin123`

2. **User**
   - Username: `user`
   - Password: `user123`

## 🧪 Testing dengan cURL

### Test Health
```bash
curl http://localhost:8080/api/health
```

### Test Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

### Test Register
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"test123"}'
```

## 📱 Koneksi dari Flutter

### iOS Simulator
```dart
static const String baseUrl = 'http://localhost:8080/api';
```

### Android Emulator
```dart
static const String baseUrl = 'http://10.0.2.2:8080/api';
```

### Device Fisik
```dart
// Ganti dengan IP komputer Anda
static const String baseUrl = 'http://192.168.1.100:8080/api';
```

Cara cek IP komputer:
- Mac: `ifconfig | grep "inet " | grep -v 127.0.0.1`
- Windows: `ipconfig`

## 🔐 Security Notes

**PENTING untuk Production:**

1. ❌ **Jangan simpan password plain text** - gunakan bcrypt
2. ❌ **Jangan hardcode JWT secret** - gunakan environment variable
3. ❌ **Jangan gunakan CORS `*`** - sebutkan origin spesifik
4. ✅ **Gunakan HTTPS** di production
5. ✅ **Validasi semua input**
6. ✅ **Implement rate limiting**
7. ✅ **Gunakan database real** (PostgreSQL, MongoDB, dll)

## 📦 Dependencies

- `github.com/golang-jwt/jwt/v5` - JWT authentication
- `github.com/google/uuid` - Generate UUID
- `github.com/rs/cors` - CORS middleware

## 🔄 Next Steps

- [ ] Hash password dengan bcrypt
- [ ] Connect ke database real
- [ ] Implement refresh token
- [ ] Add middleware untuk auth
- [ ] Add validation library
- [ ] Add logging
- [ ] Add rate limiting
- [ ] Add unit tests

## 🐛 Troubleshooting

### Port sudah digunakan
```bash
# Mac/Linux
lsof -ti:8080 | xargs kill -9

# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

### CORS Error
Pastikan CORS middleware sudah disetup dengan benar di backend.

### Connection Refused
- Pastikan backend sudah running
- Cek URL yang benar sesuai device/emulator
- Pastikan tidak ada firewall yang block
