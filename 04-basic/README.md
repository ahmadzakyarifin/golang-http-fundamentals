# 🚀 04-basic: http.NewServeMux & Modern Routing (Go 1.22)

Project ini adalah demonstrasi penerapan **Standard Library Go 1.22** yang baru. Kita meninggalkan cara lama yang "kotor" dan beralih ke struktur yang lebih **bersih, aman, dan modular** tanpa perlu library pihak ketiga (seperti Chi, Gin, atau Gorilla Mux).

---

## 🎯 Mengapa Tidak Pakai `http.HandleFunc` Biasa?

Ada 3 alasan fatal kenapa programmer profesional menghindari cara lama (Global State):

## 🛑 Deep Dive: Kenapa Meninggalkan `http.HandleFunc` (Global State)?

Seringkali tutorial pemula mengajarkan:
```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```
Untuk proyek hobi, ini oke. Tapi untuk **Production**, cara ini memiliki 3 kelemahan fatal yang bisa menghancurkan aplikasi Anda.

---

### 1. Masalah "Global State" & Dependency Hell
**Masalah:** `http.DefaultServeMux` (yang digunakan saat parameter `nil`) adalah variabel **GLOBAL** yang bisa diakses oleh *siapa saja*, termasuk library pihak ketiga yang Anda import.

> **🏠 Analogi:** Bayangkan Router Global itu seperti **Papan Pengumuman Umum di Balai Desa**.
> * Anda menempel pengumuman: *"Rapat di Ruang 1"*.
> * Tiba-tiba, orang asing (library lain) datang dan menempel kertas di atas punya Anda: *"Ruang 1 dipakai Gudang"*.
> * **Hasilnya:** Kekacauan (Panic/Crash).

**Skenario Bahaya (Code Conflict):**
Bayangkan Anda membuat rute `/login`. Lalu Anda meng-import library analitik pihak ketiga. Tanpa sepengetahuan Anda, library itu memiliki fungsi `init()` yang *juga* mendaftarkan rute `/login`.

```go
// main.go (Kode Anda)
http.HandleFunc("/login", myLoginHandler)

// library_orang_lain.go (Library yang Anda import)
func init() {
    // 💣 Dhuar! Aplikasi Anda akan CRASH saat start-up
    // karena rute "/login" didaftarkan dua kali di Global Mux yang sama.
    http.HandleFunc("/login", theirAnalyticsHandler)
}
```

**✅ Solusi (NewServeMux):**
Kita membuat instance `mux` sendiri (`mux := http.NewServeMux()`). Ini ibarat **Buku Catatan Pribadi**. Library lain tidak punya akses ke buku ini, sehingga rute Anda dijamin aman dari tabrakan.

---

### 2. Keterbatasan Single Router (Tidak Bisa Multi-Server)
**Masalah:** `http.DefaultServeMux` hanya satu instance. Anda tidak bisa memilah rute mana yang boleh diakses publik, dan mana yang rahasia (internal), jika server dijalankan di port yang sama atau router yang sama.

**Kebutuhan Production:**
Aplikasi modern biasanya membutuhkan dua pintu:
1.  **Public (Port 8080):** Untuk User (Home, Login, API).
2.  **Internal (Port 9090):** Untuk Admin System, Metrics (Prometheus), Health Check.

**Bahaya Global Mux:**
Jika Anda memakai Global Mux, rute `/metrics` (data sensitif server) akan terekspos ke internet publik bersamaan dengan rute `/home`. Hacker bisa melihat beban server Anda.

**✅ Solusi (Isolated Instances):**
Dengan `NewServeMux`, kita membuat dua dunia berbeda:

```go
// 1. Router Publik (Aman untuk internet)
muxPub := http.NewServeMux()
muxPub.HandleFunc("GET /products", productHandler)

// 2. Router Internal (Hanya untuk tim IT/VPN)
muxInt := http.NewServeMux()
muxInt.HandleFunc("GET /metrics", prometheusHandler) // Data sensitif aman
```
Kedua router ini kemudian dijalankan di port terpisah (8080 & 9090) menggunakan **Goroutine**.

---

### 3. Tidak Adanya Kontrol Konfigurasi (Risiko Serangan)
**Masalah:** Menggunakan `http.ListenAndServe(":8080", nil)` adalah cara yang "malas" karena menggunakan konfigurasi default server yang **tidak memiliki batas waktu (timeout)**.

## 🛡️ Penjelasan Config Timeout (Mekanisme Anti-Gantung)

Dalam `&http.Server{}`, kita mengatur 3 batas waktu kritis. Tanpa ini, satu user dengan internet lemot bisa memblokir kinerja server untuk user lain.

### 1. `ReadTimeout` (Batas Waktu Mendengar)
Waktu maksimal server menunggu client/user mengirimkan **keseluruhan** data request (Header + Body).

* **⏱ Kapan Dihitung:** Mulai dari koneksi diterima sampai seluruh data request (misal: upload file JSON) selesai dibaca.
* **🗣 Analogi Restoran:** Pelayan mendatangi meja user.
    * *User:* "Saya mau pesan..." (diam 5 menit) "...nasi..." (diam 5 menit) "...goreng".
    * **Tanpa Timeout:** Pelayan terjebak menunggu selamanya.
    * **Dengan ReadTimeout 10s:** Jika dalam 10 detik user belum selesai ngomong pesanan, pelayan langsung pergi (putus koneksi).
* **🛡 Fungsi:** Mencegah serangan **Slowloris** (Hacker mengirim data super lambat untuk menghabiskan resource server).

### 2. `WriteTimeout` (Batas Waktu Melayani)
Waktu maksimal server untuk **memproses logic** (Handler) DAN **mengirim balik** respon ke user.

* **⏱ Kapan Dihitung:** Mulai dari request selesai dibaca -> Masuk ke Handler (Database, Hitung Gaji, dll) -> Sampai byte terakhir respon dikirim ke user.
* **🍳 Analogi Restoran:**
    * Dapur (Server) memasak pesanan + Pelayan mengantar makanan ke meja.
    * Jika koki memasak terlalu lama (Database lemot) ATAU user makan terlalu lambat (Internet user lemot saat download response), waktu habis.
* **⚠️ Penting:** Jika proses database Anda butuh 15 detik, tapi `WriteTimeout` diset 10 detik, koneksi akan diputus di tengah jalan (Cancel) sebelum selesai.

### 3. `IdleTimeout` (Batas Waktu Bengong/Keep-Alive)
Waktu maksimal server membiarkan koneksi tetap "nyala" (Keep-Alive) saat **tidak ada aktivitas** antar request.

* **⏱ Kapan Dihitung:** Setelah respon pertama selesai dikirim, sampai user mengirim request kedua.
* **🪑 Analogi Restoran:**
    * User sudah selesai makan (request 1 selesai). Tapi dia masih duduk di kursi (koneksi masih terbuka) sambil main HP.
    * Berapa lama kita biarkan dia duduk sebelum kita usir supaya kursi bisa dipakai pelanggan lain?
* **🛡 Fungsi:** Menghemat RAM. Web modern menggunakan fitur *Keep-Alive* agar tidak perlu *handshake* ulang tiap kali klik link. Tapi jika dibiarkan selamanya, server akan kehabisan memori menampung "penonton diam".

---

### 📊 Ringkasan Visual (Timeline)

```text
   [ KONEKSI MASUK ]
          │
          ▼
[ ReadTimeout Mulai ] ──────────────────────────┐
          │                                     │
   (Baca Request Header & Body)                 │
          │                                     │
[ ReadTimeout Selesai ] ────────────────────────┘
          │
          ▼
[ WriteTimeout Mulai ] ─────────────────────────┐
          │                                     │
    (Proses Logic / Database)                   │
          │                                     │
    (Kirim Respon Balik ke User)                │
          │                                     │
[ WriteTimeout Selesai ] ───────────────────────┘
          │
          ▼
[ IdleTimeout Mulai ] ──────────────────────────┐
          │                                     │
   (Menunggu Request Selanjutnya...)            │
          │                                     │
[ IdleTimeout Selesai (Tutup Koneksi) ] ────────┘

**✅ Solusi (`http.Server` Wrapper):**
Kita membungkus `mux` kita ke dalam struct `http.Server` untuk menetapkan batas tegas.

```go
server := &http.Server{
    Addr:    ":8080",
    Handler: mux,
    // "Jika user diam lebih dari 10 detik, putuskan koneksi!"
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
}
```
---

## ⏳ "Zaman Kegelapan" vs 🧭 Era Modern

Sebelum Go 1.22 (Februari 2024), `http.NewServeMux` bawaan Go itu sangat terbatas (bodoh dan kaku). Itulah alasan kenapa framework seperti Chi atau Gorilla Mux sangat populer dulu.

### Perbandingan Fitur

| Fitur | NewServeMux (Lama) ❌ | Framework (Chi/Gorilla) ✅ | **Go 1.22 (Sekarang)** 🚀 |
| :--- | :--- | :--- | :--- |
| **Cek Method** | Harus `if/else` manual | Otomatis (`r.Get`) | **Otomatis** (`GET /path`) |
| **Parameter URL** | Manual parsing string | Mudah (`chi.URLParam`) | **Native** (`r.PathValue`) |
| **Pola Rute** | Prefix matching saja | Regex & Wildcards | **Wildcards** (`{id}`) |

### Contoh Kode: Dulu vs Sekarang

**❌ Cara Lama (Menyiksa):**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 1. Cek Method Manual
    if r.Method != "GET" {
        http.Error(w, "Method salah", 405)
        return
    }
    // 2. Parsing ID Manual dari URL "/user/123"
    id := strings.TrimPrefix(r.URL.Path, "/user/")
    // ... logic ...
}
```

**🚀 Kode Modern (Go 1.22+):**
```go
// Definisi Rute langsung ada Method dan Wildcard
mux.HandleFunc("GET /barang/{id}", handler)

func handler(w http.ResponseWriter, r *http.Request) {
    // Tidak perlu if-else method lagi!
    // Langsung ambil ID secara native
    id := r.PathValue("id")
}
```

---

## 🛠️ Bedah Struktur Server & Concurrency

Di project ini, kita melakukan hal yang agak advanced: **Menjalankan Dua Server dalam satu aplikasi.**

### 1. Kenapa Multi-Port (8080 & 9090)?
Kita memisahkan pintu masuk User dan Admin.

> **🏢 Analogi: Pintu Depan vs Pintu Belakang Restoran**
> * **Port 8080 (User):** Pintu depan yang ramai. Terbuka untuk umum.
> * **Port 9090 (Admin):** Pintu belakang khusus staf. Di sini manajer melihat stok, grafik penjualan, dll.

**Kenapa dipisah?** Keamanan. Kita bisa memblokir akses ke port 9090 dari internet luar (firewall), tapi membiarkan 8080 terbuka. Hacker tidak bisa iseng coba login admin.

### 2. Kenapa Pakai `go func()`? (Goroutine)
Fungsi `ListenAndServe` itu sifatnya **Blocking** (Menahan/Looping).

**Jika tanpa Goroutine:**
```go
serverAdmin.ListenAndServe() // Program AKAN BERHENTI DI SINI selamanya (looping nunggu tamu)
serverUtama.ListenAndServe() // Baris ini TIDAK AKAN PERNAH DIJALANKAN
```

**Solusi dengan Goroutine:**
```go
go func() {
    serverAdmin.ListenAndServe() // "Asisten, tolong jaga pintu belakang ya!"
}()

serverUtama.ListenAndServe() // "Oke, saya (Bos) jaga pintu depan."
```

> **Catatan:** Kita tidak menggunakan `sync.WaitGroup` di sini karena triknya adalah membiarkan `serverUtama` berjalan di *Main Thread* (tanpa `go func`) agar program tidak langsung exit/selesai.

### 3. Kenapa Pakai `&http.Server{}`?
Kita membungkus konfigurasi server secara manual demi keamanan.

**❌ Cara Malas (Berbahaya):**
```go
http.ListenAndServe(":8080", mux)
```
*Tidak ada Timeout. Rentan serangan Slowloris.*

**✅ Cara Pro (`&http.Server`):**
```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  10 * time.Second, // "Kalau 10 detik diam, usir!"
    WriteTimeout: 10 * time.Second,
}
```
*Ini membuat server kita Stabil dan tahan banting.*