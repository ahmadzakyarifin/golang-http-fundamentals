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

## 🧠 Filosofi: Kenapa Wajib `go mod init`? (Walau Project Kecil)

Mungkin Anda bertanya:
> *"Kode saya cuma satu file, tidak pakai library luar, dan jalan lancar tanpa `go mod init`. Kenapa saya harus tetap melakukannya?"*

Kami membiasakan hal ini sejak baris pertama kode ditulis karena 3 alasan fundamental:

### 1. Memerdekakan Folder Project (Kill GOPATH)
**Tanpa `go mod` (Cara Kuno):**
Dulu, Go memaksa semua kodingan Anda HARUS disimpan di folder spesifik: `C:\Users\Nama\go\src\...`. Jika Anda simpan di `Desktop` atau `Documents`, Go tidak bisa membacanya. Ini sangat kaku.

**Dengan `go mod` (Cara Modern):**
File `go.mod` mengubah folder project Anda menjadi **Workspace Mandiri**. Anda bebas menyimpan kodingan di mana saja (Desktop, D:, Flashdisk), dan Go akan tetap paham cara menjalankannya. `go mod init` memberikan "KTP" pada folder tersebut di manapun ia berada.

### 2. Membuka Pintu "Modular Architecture"
Saat ini kode Anda mungkin hanya `main.go`. Tapi besok, Anda mungkin ingin memindahkan logika "Rumus" ke file terpisah agar rapi.



* **Tanpa `go.mod`:** Go akan menolak membaca file di folder lain (sub-folder) karena tidak punya root path yang jelas. Anda terjebak dengan struktur 1 file yang berantakan.
* **Dengan `go.mod`:** Anda sudah punya "Root Name". File `main.go` bisa dengan mudah memanggil file lain menggunakan jalur resmi: `import "nama-project/folder-lain"`.

### 3. Membentuk "Muscle Memory" Standar Industri
Di dunia profesional, **100% project Go menggunakan Go Modules**.
Tidak ada satupun perusahaan yang membiarkan kodenya berjalan tanpa `go.mod`.

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