# 🚀 01-basic: Go HTTP Server

> **Catatan Dokumentasi:**
> Project ini menyertakan penjelasan teknis dan "Mental Model" (Analogi Logika) untuk mendemonstrasikan pemahaman mendalam mengenai cara kerja `net/http` dan fundamental Web Server.

---

## 🧐 Apa itu `net/http`?

`net/http` adalah library standar (Standard Library) yang disediakan langsung oleh Google di dalam bahasa pemrograman Go.

**Filosofi: "Batteries Included"**
Berbeda dengan bahasa lain (seperti PHP atau Node.js) yang sering membutuhkan instalasi Framework tambahan (seperti Laravel/Express) hanya untuk membuat server dasar, Go sudah menyediakan "Kotak Perkakas" lengkap.

Ibarat menyewa Ruko:
* **Bahasa Lain:** Ruko kosong. Harus beli meja/kompor dulu (Install Framework) baru bisa buka toko.
* **Go (`net/http`):** Ruko *Full Furnished*. Sudah ada Server, Router, dan HTTP Client bawaan pabrik. Kita tinggal masuk dan mulai koding.

---

## 🍔 Analogi Logika: "Restoran Go"

Untuk memvisualisasikan alur data pada kode `main.go`, project ini menggunakan analogi operasional sebuah **Restoran**:

| Komponen Kode | Peran | Penjelasan Logika |
| :--- | :--- | :--- |
| **`func main()`** | **Manajer Toko** | Titik awal program. Manajer datang untuk menginisialisasi sistem dan memulai operasional. |
| **`http.HandleFunc`** | **Resepsionis** | Petugas routing. Mengatur jalur: *"Jika ada tamu ke Lobi (`/`), arahkan ke pelayan ini."* |
| **`http.ListenAndServe`** | **Pintu Utama** | Membuka koneksi di Port `8080` dan melakukan **Blocking (Looping)** untuk stand-by menunggu tamu selamanya. |
| **`r` (Request)** | **Kertas Pesanan** | **(Input)** Data yang dibawa masuk oleh tamu dari luar (Browser, IP Address, Data Form). |
| **`w` (ResponseWriter)** | **Nampan Kosong** | **(Output)** Wadah untuk server menaruh respon/jawaban yang akan diantar balik ke meja tamu. |

### 🧩 Kesimpulan Arsitektur
Berdasarkan tabel di atas, **Server** didefinisikan sebagai **Satu Kesatuan Sistem**.

Server adalah proses kerja sama antara **Pintu Masuk** (`Listen`), **Pengatur Rute** (`Handler`), dan **Pelayan Data** (`Response`) yang berjalan terus-menerus di bawah komando **Manajer** (`Main`). Jika satu komponen hilang, sistem berhenti berfungsi.

---

## 💡 Tanya Jawab Interview (HRD/User)

Siapkan jawaban ini jika ditanya tentang keputusan teknis dalam kode:

### Q1: "Kenapa di kode ada `w.Write([]byte...)`?"
**Jawab:**
"Karena internet itu seperti pipa kabel yang hanya mengerti data biner (listrik), Pak. Internet tidak mengerti tulisan (String). Jadi tulisan 'Hello World' harus saya **hancurkan (convert)** dulu jadi serbuk data (`byte`) supaya bisa masuk dan dikirim lewat kabel."

### Q2: "Apa fungsi `ListenAndServe`?"
**Jawab:**
"Itu perintah untuk menyalakan server di port tertentu (8080). Perintah ini akan membuat program **Looping (Berputar)** terus menerus untuk menunggu tamu. Kalau baris ini dihapus, program langsung selesai dan server mati."

---

# 🧠 Kenapa Wajib `go mod init`? (Walau Project Kecil)

Sering muncul pertanyaan: *"Kenapa harus ribet ketik `go mod init` kalau kodenya cuma sedikit?"*

Jawabannya bukan soal kode banyak atau sedikit, tapi soal **Pondasi**. Tanpa file `go.mod`, project Go Anda ibarat rumah tanpa sertifikat. Berikut adalah 4 alasan praktisnya:

## 1. Syarat Mutlak Install Library (Paling Krusial)
Ini alasan paling teknis. Anda **TIDAK BISA** menginstall library luar (seperti Fiber, GORM, MySQL) jika belum melakukan `go mod init`.

* ❌ **Tanpa `go.mod`:** Saat Anda ketik `go get ...`, Terminal akan error: `go: go.mod file not found`. Go bingung mau mencatat library tersebut di mana.
* ✅ **Dengan `go.mod`:** File ini berfungsi sebagai **Catatan Belanja**. Go akan mendownload library dan mencatat versinya di sini agar project bisa berjalan.

## 2. Syarat Pemisahan Folder (Rapi)
Saat kode Anda mulai panjang, Anda pasti ingin memisahnya ke folder lain (misal: folder `handlers`, `helpers`).

* ❌ **Tanpa `go.mod`:** File `main.go` tidak akan bisa memanggil file di folder lain. Go akan bingung mencari jalurnya.
* ✅ **Dengan `go.mod`:** File `go.mod` menjadi titik nol (Root). `main.go` bisa dengan mudah memanggil folder lain menggunakan nama modul: `import "nama-project/handlers"`.

## 3. Identitas Project (KTP)
File `go.mod` memberikan **Nama Resmi** pada folder project Anda.

* **Tanpa `go.mod`:** Go menganggap folder Anda hanya kumpulan file teks biasa tanpa tuan.
* **Dengan `go.mod`:** Project Anda punya identitas (misal: `module toko-online`). Ini penting agar Go tahu file mana saja yang termasuk "anggota keluarga" project ini.

## 4. Agar Folder Project "Portable" (Fleksibel)
Zaman dulu (sebelum ada `go mod`), kodingan Go **wajib** ditaruh di folder khusus sistem (`C:\Go\src`). Sangat kaku.

* **Dengan `go.mod`:** Folder project Anda jadi **Portable**.
* Anda bebas menyimpannya di **Desktop**, **Documents**, **Drive D:**, atau **Flashdisk** sekalipun. Selama ada file `go.mod`, Go tetap bisa menjalankannya dengan lancar.

---

### 💡 Contoh Kasus Sederhana

Bayangkan struktur folder seperti ini:

```text
my-project/
├── go.mod        <-- (1) "Sertifikat" & "Catatan Belanja"
├── main.go       <-- (2) Ruang Tamu
└── helpers/      <-- (3) Dapur (Ruangan terpisah)
    └── rumus.go

Membiasakan diri mengetik `go mod init` di awal project (meskipun kecil) adalah investasi kebiasaan.
* **Bad Habit:** Terbiasa coding "asal jalan" tanpa inisialisasi. Saat masuk project besar, Anda akan bingung kenapa import error.
* **Good Habit:** Selalu menyiapkan wadah (Module) sebelum mengisi isinya. Ini membuat Anda siap untuk skala project sebesar apapun.

> **Kesimpulan:**
> Kami tidak menggunakan `go mod init` untuk kebutuhan kode *saat ini*, tapi untuk mempersiapkan **Skalabilitas** dan **Standarisasi** kode di masa depan.

## 🛠️ Cara Menjalankan

1.  Pastikan Go sudah terinstall.
2.  Buka terminal di folder ini.
3.  Ketik perintah:
    ```bash
    go run main.go
    ```
4.  Buka browser di: `http://localhost:8080`

---
*Dibuat sebagai referensi belajar Go Fundamental.*