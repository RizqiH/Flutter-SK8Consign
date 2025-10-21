# 📱 Screens - Penjelasan untuk Pemula

Folder ini berisi semua halaman (screens) dalam aplikasi.

---

## 🎬 1. Splash Screen (`splash_screen.dart`)

### Apa itu Splash Screen?
Halaman pembuka yang muncul saat aplikasi pertama kali dibuka.

### Fitur:
- ✅ Background gradient ungu
- ✅ Logo SK8 CONSIGN di tengah
- ✅ Auto redirect ke Login setelah 3 detik

### Konsep yang Dipelajari:
- **StatefulWidget**: Widget yang bisa berubah
- **Timer**: Menjalankan aksi setelah delay
- **Navigator**: Pindah ke halaman lain
- **LinearGradient**: Membuat warna gradient

### Code Penting:
```dart
// Timer untuk auto navigate
Timer(const Duration(seconds: 3), () {
  Navigator.pushReplacementNamed(context, '/login');
});
```

---

## 🔐 2. Login Screen (`login_screen.dart`)

### Apa itu Login Screen?
Halaman untuk user memasukkan username dan password.

### Fitur:
- ✅ Logo SK8 CONSIGN di atas
- ✅ Card transparan dengan border
- ✅ Greeting "Hello !"
- ✅ Input Username dengan icon
- ✅ Input Password dengan icon (hidden text)
- ✅ Tombol Login dengan gradient
- ✅ Link "Tidak Punya Akun?"
- ✅ Validasi input kosong
- ✅ SnackBar notification

### Konsep yang Dipelajari:
- **TextEditingController**: Mengontrol input text
- **TextField**: Input field untuk user
- **ElevatedButton**: Tombol dengan elevasi
- **GestureDetector/TextButton**: Tombol text
- **SnackBar**: Notifikasi pop-up dari bawah
- **Form Validation**: Validasi input user

### Widget-Widget Penting:

#### 1. Container untuk Card
```dart
Container(
  padding: const EdgeInsets.all(32),
  decoration: BoxDecoration(
    color: Colors.white.withOpacity(0.1), // Transparan
    borderRadius: BorderRadius.circular(24), // Rounded corner
    border: Border.all(...), // Border putih
  ),
)
```

#### 2. TextField untuk Input
```dart
TextField(
  controller: _usernameController, // Controller
  style: TextStyle(color: Colors.white),
  decoration: InputDecoration(
    hintText: 'Username',
    prefixIcon: Icon(Icons.person_outline),
  ),
)
```

#### 3. Button dengan Gradient
```dart
Container(
  decoration: BoxDecoration(
    gradient: LinearGradient(
      colors: [Color(0xFF7B5FFF), Color(0xFF5B8DEE)],
    ),
  ),
  child: ElevatedButton(...),
)
```

### Cara Kerja Login:

1. User isi Username & Password
2. Klik tombol "Log In"
3. Fungsi `_handleLogin()` dipanggil
4. Validasi: Cek apakah field kosong
5. Jika kosong → tampilkan SnackBar merah
6. Jika terisi → tampilkan SnackBar hijau (demo)
7. TODO: Nanti connect ke backend/database

### Flow Data:
```
User ketik → Controller simpan → Button diklik → 
Fungsi validasi → Tampilkan hasil
```

---

## 🎨 Design System

### Colors:
- Background: Gradient ungu (`#1a1a2e` → `#4a2c6d`)
- Card: White 10% opacity
- Button: Gradient biru-ungu (`#7B5FFF` → `#5B8DEE`)
- Text: White

### Typography:
- Logo: 36px bold (SK8) + 300 weight (CONSIGN)
- Hello: 32px bold
- Input: 16px
- Button: 18px bold

### Spacing:
- Card padding: 32px
- Input spacing: 20px
- Button height: 56px

---

## 📝 Tips Coding untuk Pemula

### 1. Dispose Controller
Selalu dispose controller di akhir:
```dart
@override
void dispose() {
  _usernameController.dispose();
  _passwordController.dispose();
  super.dispose();
}
```

### 2. SafeArea
Gunakan SafeArea agar tidak tertutup notch:
```dart
SafeArea(
  child: // Your content
)
```

### 3. SingleChildScrollView
Agar bisa scroll kalau keyboard muncul:
```dart
SingleChildScrollView(
  child: // Your form
)
```

### 4. Validation
Selalu validasi input user:
```dart
if (username.isEmpty) {
  // Show error
  return;
}
```

---

## 🚀 Next Steps

- [ ] Buat halaman Register
- [ ] Connect ke API backend
- [ ] Simpan token login
- [ ] Navigate ke Home setelah login
- [ ] Remember me feature
- [ ] Forgot password

---

## 💡 Debugging Tips

### Hot Reload vs Hot Restart
- **Hot Reload** (r): Update UI tanpa restart
- **Hot Restart** (R): Restart app dari awal

### Print untuk Debug
```dart
print('Username: $username');
print('Password: $password');
```

### Check Console
Selalu lihat console untuk error message!

---

Dibuat dengan ❤️ untuk pembelajaran Flutter
