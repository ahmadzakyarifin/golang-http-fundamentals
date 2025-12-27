# 🚀 02-basic: Dynamic HTTP Server (Query Parameters)

> **Level 2: Membaca Input User**
> Jika di Level 1 server kita seperti "Robot Kaku" yang hanya bisa menyapa hal yang sama, di Level 2 ini server sudah menjadi "Pelayan Pintar" yang bisa menyapa sesuai nama tamu.

---

## 🎯 Tujuan Belajar

Project ini mendemonstrasikan cara mengambil data dinamis dari URL menggunakan **Query Parameters**.
* **Static (Level 1):** `Hello World` (Selalu sama).
* **Dynamic (Level 2):** `Hello, Budi` / `Hello, Siti` (Berubah sesuai input).

---

## 🍔 Analogi Logika: "Membaca Stiker Nama"

Kita update analogi Restoran kita. Sekarang tamu yang datang tidak diam saja, tapi mereka menempelkan **Stiker Nama** di baju mereka.

| Komponen Kode | Peran di Restoran | Logika Baru (Level 2) |
| :--- | :--- | :--- |
| **`r.URL.Query().Get("nama")`** | **Mata Pelayan** | Pelayan melihat ke baju tamu, mencari stiker bertuliskan "nama", lalu membaca isinya. |
| **`if nama == ""`** | **Logika Pelayan** | *"Wah, tamu ini gak pakai stiker nama."* -> Sapa standar ("Hello"). |
| **`else`** | **Logika Pelayan** | *"Oh, stikernya tulisan Budi."* -> Sapa spesifik ("Hello, Budi"). |
| **`fmt.Sprintf`** | **Mulut Pelayan** | Teknik menyusun kalimat gabungan: *"Kata Sapaan"* + *"Nama Tamu"*. |

---

## 💡 Bedah Teknis (Deep Dive)

### 1. Apa itu Query Parameter?
Query parameter adalah cara user menitipkan data lewat URL browser.
Formatnya selalu dimulai dengan tanda tanya `?`.

Contoh URL: `http://localhost:8080/?nama=Dian`
* **Endpoint:** `http://localhost:8080/`
* **Key:** `nama`
* **Value:** `Dian`

### 2. Kode: `r.URL.Query().Get("nama")`
* `r` (Request): Kita bongkar kertas pesanan/data tamu.
* `URL`: Kita lihat alamat yang dia ketik.
* `Query()`: Kita cari bagian setelah tanda tanya `?`.
* `Get("nama")`: Kita ambil nilai dari kunci "nama".

### 3. Kode: `fmt.Sprintf`
```go
pesan := fmt.Sprintf("Hello, %s", nama)


## 📦 Aturan Penamaan: "Siapa yang Boleh Lihat?"

Di Golang, huruf awal sebuah function menentukan privasinya. Ini disebut **Exported (Public)** vs **Unexported (Private)**.

| Gaya Penulisan | Huruf Awal | Status | Analogi Restoran | Contoh |
| :--- | :--- | :--- | :--- | :--- |
| **PascalCase** | **Besar** (Kapital) | **Public / Exported** | **Buku Menu.** Bisa dilihat dan dipesan oleh pelanggan dari luar. | `fmt.Println`, `http.ListenAndServe` |
| **camelCase** | **Kecil** | **Private / Unexported** | **SOP Dapur.** Rahasia internal. Hanya karyawan dalam toko ini (package ini) yang tahu. | `func hello`, `func hitungGaji` |

### Kenapa `func hello` kita pakai huruf kecil?
Karena fungsi `hello` hanya dipakai **di dalam file ini saja** (di dalam `package main`). Tidak ada orang luar yang perlu memanggil fungsi ini. Jadi, kita membuatnya **Private** (huruf kecil) untuk menjaga kerapian dan keamanan kode internal.

### Kenapa `fmt.Println` pakai huruf besar?
Karena fungsi `Println` dibuat oleh orang lain (Tim Google) di dalam package `fmt`. Agar kita (orang luar) bisa memakainya, mereka **WAJIB** menamainya dengan huruf besar (`Println`). Kalau mereka menamainya `println` (huruf kecil), kita tidak akan bisa memakainya.