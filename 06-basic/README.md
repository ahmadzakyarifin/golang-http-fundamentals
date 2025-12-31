# 🛡️ HTTP Middleware Chaining & Context Injection

Project ini adalah demonstrasi tingkat lanjut penggunaan **Standard Library Go (net/http)** untuk membangun server yang aman dan modular menggunakan konsep **Middleware Chaining** dan **Context Passing**.

---

## 🧠 Konsep Utama

Ada dua teknik utama yang diterapkan di sini:

1.  **Middleware Chaining (Rantai Middleware)**
    Teknik membungkus *Handler utama* dengan lapisan-lapisan logika tambahan.
    * **Layer 1 (Luar):** `Timeout` (Membatasi durasi request).
    * **Layer 2 (Tengah):** `Authentication` (Identifikasi user).
    * **Core (Dalam):** `UserHandler` (Logic utama aplikasi).

2.  **Context Injection**
    Teknik mengirim data antar-layer (dari Middleware ke Handler) tanpa merubah parameter fungsi, menggunakan `context.Context`.

---

## 🛫 Analogi Sederhana: "Penerbangan Pesawat"

Bayangkan `UserHandler` adalah **Kabin Pesawat** (Tujuan Akhir). Sebelum penumpang bisa duduk di sana, mereka harus melewati pos pemeriksaan:

1.  **Pos 1 (Timeout Middleware):** *Petugas Timer*
    * Petugas menempelkan "Jam Waktu" di tas penumpang.
    * Jika penumpang lama di jalan (> 2 detik), mereka dipaksa keluar.

2.  **Pos 2 (Auth Middleware):** *Petugas Tiket*
    * Petugas mengecek identitas, lalu menempelkan stiker **"Nama: Zaky"** di tas penumpang (Context).

3.  **Kabin (UserHandler):** *Pramugari*
    * Pramugari melihat stiker di tas, lalu menyapa: *"Halo, Zaky!"*.

---

## 🔍 Bedah Kode & Sintaks (Deep Dive)

Berikut adalah penjelasan rinci mengenai sintaks-sintaks penting yang digunakan:

### 1. Private Context Key
Terletak di `middleware/auth.go`.

```go
type key int 
const userKey key = 0
```

* **Fungsi:** Mencegah Tabrakan Kunci (*Key Collision*).
* **Penjelasan:** Context bersifat global. Jika kita menggunakan string biasa `"user"` sebagai kunci, bisa jadi library lain (seperti Google Analytics atau Log) juga menggunakan kunci `"user"`. Data kita bisa tertimpa. Dengan membuat tipe data sendiri (`type key`), kunci kita dijamin unik dan aman.

### 2. Context Timeout
Menggunakan `context.WithTimeout`.

* **Fungsi:** Menyuntikkan Sinyal Deadline ke dalam request.
* **Logika:** Fungsi ini tidak mengisi data string/angka, melainkan memasang "Bom Waktu". Jika handler utama bekerja lebih dari 2 detik, context akan "meledak" (`Cancelled`), dan proses dihentikan paksa.

### 3. Context Value
Menggunakan `context.WithValue`.

* **Fungsi:** Menyuntikkan Data ke dalam request.
* **Logika:** Context di Go bersifat *Immutable* (tidak bisa diedit). Fungsi ini bekerja dengan cara:
    1.  Mengkopi Context lama (`r.Context()`).
    2.  Menambahkan data `currentUser`.
    3.  Mengembalikan Context baru (`ctx`) yang sudah berisi data.

### 4. Melanjutkan Request
Menggunakan `next.ServeHTTP`.

* **Fungsi:** Mengoper bola ke pemain selanjutnya.
* **Penting:** Perhatikan `r.WithContext(ctx)`. Kita tidak mengirim request lama, tapi mengirim **Request Baru** yang sudah menggendong context hasil modifikasi (berisi data user/timeout).

### 5. Error 401
Menggunakan `http.StatusUnauthorized`.

* **Arti:** "Maaf, saya tidak kenal kamu."
* **Fungsi:** Jika Handler merogoh saku context tapi tidak menemukan data user (variabel `ok` bernilai `false`), maka server membalas dengan kode HTTP 401. Ini standar internasional untuk "Gagal Login".

---

## 📂 Struktur File

```text
.
├── main.go                 # Titik masuk, merakit urutan middleware
├── handler/
│   └── user.go             # Handler akhir (Tujuan request)
└── middleware/
    ├── auth.go             # Logic menyisipkan data user
    └── timeout.go          # Logic membatasi waktu
```

---

## 🔄 Alur Data (Flow)

Saat user mengakses `http://localhost:8080`:

1.  **Request Masuk ➡️ TimeoutMiddleware**
    * Set Timer: 2 Detik.
    * *Context sekarang punya Deadline.*

2.  **Lanjut ➡️ AuthMiddleware**
    * Set Data: `currentUser = "zaky"` (Hardcoded untuk simulasi).
    * *Context sekarang punya Deadline + Data User.*

3.  **Sampai ➡️ UserHandler**
    * Handler membuka Context.
    * Ambil data user.
    * Print: "Halo, zaky 👋".

---

## 🚀 Cara Menjalankan & Testing

### 1. Jalankan Server
Pastikan Go sudah terinstall.

```bash
go run main.go
```

### 2. Test Berhasil (Happy Path)
Buka terminal baru dan jalankan curl:

```bash
curl -i http://localhost:8080
```

**Output:**
```http
HTTP/1.1 200 OK
Date: ...
Content-Length: ...
Content-Type: text/plain; charset=utf-8

Halo, zaky 👋
```