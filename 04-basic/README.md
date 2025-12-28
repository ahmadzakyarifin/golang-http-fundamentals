# 🚀 04-basic: http.NewServeMux & Modern Routing (Go 1.22)

Project ini adalah demonstrasi penerapan **Standard Library Go 1.22** yang baru. Kita meninggalkan cara lama yang "kotor" dan beralih ke struktur yang lebih **bersih, aman, dan modular** tanpa perlu library pihak ketiga (seperti Chi, Gin, atau Gorilla Mux).

---

## 🎯 Mengapa Tidak Pakai `http.HandleFunc` Biasa?

Ada 3 alasan fatal kenapa programmer profesional menghindari cara lama (Global State):

### 1. Masalah "Global State" (Analogi Papan Pengumuman Desa)
`http.DefaultServeMux` (yang dipakai saat kita mengetik `nil` pada parameter handler) adalah variabel **GLOBAL**.

> **Analogi:** Bayangkan Router Global itu seperti **Papan Pengumuman di Balai Desa**. Siapapun boleh menempel kertas di sana.

* **Bahaya (Collision):** Jika Anda menginstall library pihak ketiga (misalnya library pembayaran), dan pembuat library itu ceroboh ikut mendaftarkan rute `/admin`, maka aplikasi Anda bisa **CRASH (Panic)** karena rute `/admin` milik Anda tertimpa atau bentrok.
* **Solusi:** Dengan `NewServeMux`, kita membuat **Buku Catatan Pribadi**. Orang lain (library luar) tidak bisa mencoret-coret buku rute kita.

### 2. Tidak Bisa Menjalankan Multi-Server
Dengan cara lama (Global), Anda hanya punya 1 router. Anda tidak bisa membedakan rute untuk port yang berbeda. Dalam project ini, kita mendemonstrasikan kebutuhan level production:

* **Port 8080 (Public):** Untuk User (Home, Login, Product).
* **Port 9090 (Internal):** Untuk Admin/Metrics (Dashboard, Monitoring).

**Solusi:** Dengan `NewServeMux`, kita bisa membuat instance router berbeda (`muxPublic` & `muxAdmin`) dan menjalankannya bersamaan menggunakan **Goroutine**.

### 3. Kontrol Konfigurasi (Timeout)
Menggunakan `http.ListenAndServe` secara langsung dengan `nil` adalah cara yang "malas". Kita tidak bisa mengatur timeout.

* **Risiko:** Jika koneksi internet user lambat (atau ada serangan hacker *Slowloris*), server Anda bisa *hang* selamanya menunggu data yang tak kunjung sampai.
* **Solusi:** Kita membungkus `mux` ke dalam `http.Server` manual untuk mengatur `ReadTimeout` dan `WriteTimeout`.

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