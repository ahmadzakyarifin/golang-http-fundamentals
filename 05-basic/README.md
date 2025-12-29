# 🚀 05-basic: HTTP Middleware Fundamentals

Project ini mendemonstrasikan cara membuat **HTTP Middleware** secara manual di Golang tanpa library pihak ketiga (seperti Chi atau Gin). Kita akan membedah konsep dasar di balik "Interceptor" request ini.

---

## 🧐 Apa itu Middleware?

Middleware adalah blok kode yang **menengah-nengahi** (intervene) antara **Request** yang masuk dari user dan **Handler** (fungsi utama) yang memproses logic bisnis.

### Analogi: "Pos Pemeriksaan Satpam"
Bayangkan aplikasi Anda adalah sebuah Gedung Kantor.
1.  **Middleware** = Satpam di gerbang depan.
2.  **Handler Utama** = Karyawan di dalam ruangan.

Setiap tamu (Request) harus melewati Satpam dulu:
* **Before:** Satpam mengecek ID & mencatat jam masuk (**Logging**).
* **Process:** Tamu bertemu karyawan (**Handler**).
* **After:** Satpam mencatat jam keluar tamu tersebut.

---

## ❓ FAQ: Mengapa Codingnya Seperti Ini?

Berikut adalah jawaban atas pertanyaan mendasar mengenai arsitektur kode ini:

### 1. Kenapa Harus Dipisah ke Package `middleware`?
Kenapa tidak disatukan saja di `main.go`?
* **Separation of Concerns (Pemisahan Tugas):** `main.go` berfungsi sebagai tempat merakit aplikasi (wiring). Logic teknis seperti logging tidak boleh mengotori logic utama.
* **Reusability (Bisa Dipakai Ulang):** Dengan memisahkannya, middleware ini bisa ditempelkan ke handler lain (misal: `/login`, `/dashboard`, `/payment`) tanpa menulis ulang kode.
* **Clean Code:** Menghindari "Spaghetti Code" dimana logic bisnis bercampur aduk dengan logic teknis.

### 2. Apa itu Higher-Order Function (HOF) & Closure?
Middleware di Go sangat bergantung pada dua konsep ini:

#### A. Higher-Order Function (HOF)
Adalah fungsi yang menerima fungsi lain sebagai argumen ATAU mengembalikan fungsi.
> **Di kode kita:** `LoggingMiddleware` menerima `next http.Handler` dan mengembalikan `http.Handler` baru. Ini memungkinkan kita membuat rantai (chaining).

## 🔁 Kenapa Menggunakan Higher-Order Function?

### Alasan Utama (Ini Intinya)

Karena **data yang dikirim ke middleware adalah FUNCTION**, yaitu:
- handler function (`http.Handler`)

Dan:
- **yang dikembalikan oleh middleware juga HARUS function**

---

### Kenapa Harus Higher-Order Function?

Karena **HANYA Higher-Order Function** yang bisa:
- menerima parameter berupa function
- mengembalikan function lagi

#### B. Closure
Adalah fungsi tanpa nama (anonymous function) yang berada di dalam fungsi lain, tapi bisa "mengingat" variabel di sekitarnya.
> **Di kode kita:**
> ```go
> return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
>     // Fungsi ini adalah Closure.
>     // Dia bisa mengakses variabel 'next' meskipun variabel itu ada di luar scope-nya.
>     next.ServeHTTP(w, r) 
> })
> ```
🔄 Kenapa Closure Wajib?
Karena:
Kontrak net/http tidak bisa diubah : Server hanya memanggil (w, r)
Middleware harus tetap bisa memanggil next
Closure adalah SATU-SATUNYA cara aman


## 🔄 Alur Kerja (Execution Flow)

Saat user mengakses `http://localhost:8080/`:

1.  **Browser** mengirim Request.
2.  **Middleware (Bagian Atas)** berjalan:
    * Mencetak: `before => method : GET ...`
3.  **Next Handler** dipanggil (`next.ServeHTTP`):
    * Program masuk ke Handler Utama.
    * Browser menerima tulisan "Hello Word".
4.  **Middleware (Bagian Bawah)** berjalan kembali:
    * Mencetak: `after`
5.  **Selesai.**

---

## 🏢 Best Practice di Industri

Apakah cara manual ini dipakai di perusahaan?
* **Konsepnya: YA.** Semua middleware (Auth, CORS, Logging, Metrics) bekerja dengan prinsip ini.
* **Implementasinya:** Di aplikasi besar, menumpuk middleware secara manual (`Logging(Auth(Cors(Handler)))`) akan membuat kode sulit dibaca (seperti kulit bawang).

**Solusi Modern:**
Developer biasanya menggunakan pola **Chaining** atau library router (seperti `Chi` atau `Gorilla Mux`) yang memudahkan penumpukan middleware:

```go
// Contoh konsep di production (supaya lebih rapi)
r.Use(Middleware1)
r.Use(Middleware2)
r.Use(LoggingMiddleware)


# Perbandingan HTTP Handler: Tanpa Middleware vs Dengan Middleware (Go)

Dokumen ini menjelaskan **bentuk dan alur eksekusi HTTP server di Go** dalam dua kondisi:
1. Tanpa middleware
2. Dengan middleware

Fokus utama:
- HTTP server **selalu** bekerja dengan `URL → handler`
- Middleware **tidak dikenal** oleh HTTP server
- Middleware bekerja dengan **membuat handler baru**

---


## Tanpa Middleware (Alur Paling Dasar)
URL → handler function
```go
http.HandleFunc("/", helloHandler)
```
## Dengan Middleware
HTTP Server → Handler (yang di dalamnya ada middleware + handler lama/ before sebagai penyaring + handler lama (logic) + after)


## Ringkasan Alur
Client
  ↓
HTTP Server
  ↓
Handler Baru (hasil middleware)
        ├─ BEFORE (validasi / logging)
        ├─ Handler Lama (logic bisnis)
        └─ AFTER (logging / cleanup)
