# 🚀 03-basic: HTTP Methods & CRUD Logic (Modern Go 1.22+)

> **Level 2.5: Membedakan "Niat" & Validasi ID**
> Di level ini, server kita semakin cerdas. Tidak hanya bisa membaca nama tamu, tapi juga bisa membedakan **Metode (Niat)** tamu: apakah mau mengambil data, membuat baru, mengedit, atau menghapus.

---

## 🎯 Tujuan Belajar

Project ini mendemonstrasikan logika dasar **RESTful API** menggunakan fitur modern **Go 1.22**:
1.  **Method Switching:** Menggunakan `switch r.Method` untuk merespon beda-beda tergantung tombol yang ditekan client (GET, POST, PUT, dll).
2.  **Modern Routing (Path Value):** Menggunakan fitur baru Go (`r.PathValue`) untuk menangkap ID langsung dari URL (contoh: `/101`) tanpa perlu parsing manual atau query param.

---
## 📚 Teori: Filosofi HTTP Methods (Wajib Paham!)

Sebelum masuk ke kode, kita harus paham "Kenapa" hal ini ada. Bagian ini berisi konsep fundamental yang sering ditanyakan saat **Interview Backend Developer**.

### 1. Apa itu HTTP Method?
HTTP Method (disebut juga *HTTP Verbs*) adalah **Kata Kerja** yang dikirim Client (Browser/Postman) untuk memberitahu Server **tindakan apa** yang harus dilakukan terhadap data.

Bayangkan kalimat bahasa Indonesia:
* **URL (Kata Benda):** `/surat`
* **Method (Kata Kerja):** `BAKAR` (Delete), `TULIS` (Post), atau `BACA` (Get).

Tanpa Method, Server bingung: *"Kamu akses `/surat` ini mau diapain? Dibaca atau dibakar?"*

### 2. Kenapa Tidak Pakai `POST` Saja Untuk Semuanya?
Secara teknis, kita *bisa* membuat API yang isinya `POST` semua. Tapi ini adalah **Bad Practice** (Kebiasaan Buruk). Kenapa?

1.  **Semantik (Bahasa yang Jelas):** `GET` berarti aman dibaca berulang-ulang, `DELETE` berarti awas ada data yang hilang. Frontend Developer akan berterima kasih jika Anda pakai standar ini.
2.  **Caching (Kecepatan):** Browser hanya berani menyimpan (cache) data `GET` agar loading instan. Browser tidak akan menyimpan data `POST` karena dianggap berubah-ubah.

###  Kamus Lengkap Method (Bedah Perbedaan)

Berikut adalah 5 metode wajib bagi Backend Developer:

| Method | Fungsi Utama | Sifat (Idempotent?) | Keterangan Simpel |
| :--- | :--- | :--- | :--- |
| **GET** | Mengambil data. | ✅ Ya (Aman) | Cuma "lihat-lihat". Tidak mengubah database. |
| **POST** | Membuat data **BARU**. | ❌ Tidak | Tiap kali diklik, data baru tercipta (bisa duplikat). |
| **PUT** | Update **TOTAL**. | ✅ Ya | Ganti seluruh isi data dengan yang baru. |
| **PATCH** | Update **SEBAGIAN**. | ⚠️ Tergantung | Cuma tambal bagian yang rusak/diganti. |
| **DELETE** | Menghapus data. | ✅ Ya | Hilangkan data selamanya. |

### 💡 Topik Interview Pro: "Idempotency"
*"Idempotent"* artinya: **Mau tombol diklik 1 kali atau 1000 kali, hasil akhirnya tetap sama.**

* **DELETE itu Idempotent:** Hapus User ID 101.
    * Klik 1x: Data hilang.
    * Klik 100x lagi: Data tetap hilang (Server stabil, tidak ada kerusakan tambahan).
* **POST itu TIDAK Idempotent:** Buat Pesanan Pizza.
    * Klik 1x: Ada 1 pesanan.
    * Klik 100x (karena HP lag): Ada **100 pesanan** masuk & saldo terpotong 100x. (Bahaya!).

---
### 🛒 Studi Kasus: Logika Keranjang Belanja (Shopee/Tokopedia)

Sering ditanyakan: *"Apakah tombol 'Add to Cart' itu method POST?"*
Jawabannya unik. Fitur keranjang belanja biasanya menggunakan logika Hybrid yang disebut **Upsert (Update or Insert)**.

Mari kita bedah menggunakan cerita: **User Budi ingin membeli Sepatu (ID 101)**.

1.  **Cek Database (Query):**
    Saat Budi klik "Tambah", server mengecek: *"Barang ID 101 ini sudah ada belum di keranjang Budi?"*

2.  **Skenario A: Barang Belum Ada (Logika POST/INSERT)**
    * *Kondisi:* Budi baru pertama kali klik.
    * *Aksi Server:* Buat baris baru.
    * *Hasil Database:* `User: Budi | Item: 101 | Qty: 1`

3.  **Skenario B: Barang Sudah Ada (Logika PATCH/UPDATE)**
    * *Kondisi:* Budi klik lagi karena ingin beli 2 pasang.
    * *Aksi Server:* **Jangan buat baris baru!** Cukup edit jumlahnya.
    * *Logika:* `Qty_Baru = Qty_Lama + 1`
    * *Hasil Database:* `User: Budi | Item: 101 | Qty: 2` (Tetap 1 baris).

**Kesimpulan:**
* **Fase Keranjang:** Fokus pada kenyamanan user & kerapian database (Upsert Logic).
* **Fase Checkout (Bayar):** Fokus pada keamanan ketat. Di sinilah **Idempotency Key** wajib ada agar jika HP nge-lag, saldo tidak terpotong dua kali.

## 🍔 Analogi Logika: "Restoran & Nomor Meja"

Bayangkan pelayan restoran yang bekerja dengan mata tertutup, hanya mendengar **Nada Bicara (Method)** dan **Nomor Meja (ID)**.

| HTTP Method | Analogi Tindakan | Butuh ID? | Logika di Kode |
| :--- | :--- | :--- | :--- |
| **GET** | *"Mas, minta daftar menu."* | ❌ Tidak | `w.Write("Ambil data")` |
| **POST** | *"Mas, ini pesanan baru saya."* | ❌ Tidak | `w.Write("Kirim data")` |
| **PUT** | *"Mas, tolong **GANTI TOTAL** pesanan di **Meja 5**."* | ✅ **WAJIB** | Cek `if id == ""` -> Error. |
| **PATCH** | *"Mas, tolong **NAMBAH** sambal di **Meja 5**."* | ✅ **WAJIB** | Cek `if id == ""` -> Error. |
| **DELETE** | *"Mas, tolong **BATALKAN** pesanan di **Meja 5**."* | ✅ **WAJIB** | Cek `if id == ""` -> Error. |

---

## 🛠️ Cara Menjalankan

1.  **Inisialisasi Project:**
    ```bash
    go mod init 03-basic
    ```
2.  **Pastikan Versi Go:**
    Karena kode ini menggunakan fitur baru, pastikan versi Go Anda minimal **1.22**. Cek dengan `go version`.
3.  **Jalankan Server:**
    ```bash
    go run main.go
    ```

---

## 🧪 Cara Test (Wajib Postman)

Karena kita sudah pakai **Cara Baru (Path Value)**, cara test URL-nya menjadi lebih rapi (tidak pakai `?id=` lagi).

### 1. Test GET & POST (Tanpa ID)
* **Method:** `GET`
    * **URL:** `http://localhost:8080` (atau `http://localhost:8080/`)
    * **Respon:** "Hai saya sedang ambil data..."
* **Method:** `POST`
    * **URL:** `http://localhost:8080`
    * **Respon:** "Hai saya sedang kirim data..."

### 2. Test Error (PUT/PATCH/DELETE Tanpa ID)
Coba lakukan method berbahaya tapi akses ke root (tanpa angka ID).
* **Method:** `DELETE`
* **URL:** `http://localhost:8080/` (Tanpa angka ID)
* **Respon:** `Eror : Maaf id belum ada`
* *(Logika keamanan server berjalan!)*

### 3. Test Sukses (Dengan ID)
Masukkan ID langsung di belakang URL (Path Parameter).
* **Method:** `PUT` (Ganti Total)
    * **URL:** `http://localhost:8080/101`
    * **Respon:** "...edit data seluruhnya... id : 101"
* **Method:** `PATCH` (Edit Sebagian)
    * **URL:** `http://localhost:8080/99`
    * **Respon:** "...edit data sebagian... id : 99"
* **Method:** `DELETE` (Hapus)
    * **URL:** `http://localhost:8080/55`
    * **Respon:** "...hapus data... id : 55"

---

## 🧠 Bedah Kode (Deep Dive)

### 1. Evolusi Routing: Kenapa Kode Kita Berubah?

Di dunia Go, ada dua cara menangani URL dinamis. Kode ini menunjukkan transisi tersebut.

#### A. Cara Lama (The Old Way - Query Params)
Dulu (sebelum Go 1.22), Router bawaan Go (`DefaultServeMux`) sangat sederhana. Dia tidak mengerti bahwa `/users/101` itu artinya "User dengan ID 101". Dia menganggap itu alamat yang berbeda total dari `/users/102`.

Solusinya dulu adalah pakai **Query Parameter (`?id=`)**:
* **URL:** `http://localhost:8080?id=101`
* **Kode:** `id := r.URL.Query().Get("id")`
* **Kelemahan:** URL terlihat kurang rapi dan tidak standar RESTful API modern.

#### B. Cara Baru (The Modern Way - Path Value) 🚀
Sejak **Go 1.22 (Feb 2024)**, Go menjadi lebih pintar! Kita bisa menggunakan **Wildcard / Path Parameter** di routing.
Wildcard adalah simbol / placeholder di URL path yang bisa diganti dengan nilai apa pun, sehingga satu route bisa menangani banyak URL berbeda.
* **Setup Route:** `http.HandleFunc("/{id}", ...)`
* **URL:** `http://localhost:8080/101`
* **Kode:** `id := r.PathValue("id")`
* **Kelebihan:**
  1. URL bersih, rapi, dan sesuai standar RESTful API.
  2. Memisahkan resource utama dan info tambahan


###  Kenapa Harus Dua Route (`/` dan `/{id}`)? Apakah ini Kemunduran?

Mungkin Anda berpikir: *"Dulu di versi lama cukup satu baris `http.HandleFunc("/", ...)` semua beres. Kenapa sekarang harus dua? Bukankah ini jadi lebih ribet?"*

Jawabannya: **Ini bukan kemunduran, tapi "Ketegasan" (Strictness).**

* **Skenario Route Hanya `{id}`:**
    Jika kita hanya pasang `http.HandleFunc("/{id}")`, maka server hanya menerima URL yang ada buntutnya (misal `/101`). Akibatnya, halaman utama (`/`) untuk **GET/POST akan Error 404**.
    
* **Skenario Route Hanya `/`:**
    Jika kita hanya pasang `http.HandleFunc("/")`, fitur canggih `r.PathValue("id")` **tidak akan jalan**, karena router menganggap URL tersebut hanyalah string biasa, bukan variabel.

**Kesimpulan:**
Kita memisahkan "Pintu Masuk" agar penanganan lebih rapi:
1.  **Pintu `/`**: Khusus tamu umum (GET List / POST Create).
2.  **Pintu `/{id}`**: Khusus tamu yang membawa tiket nomor (PUT/PATCH/DELETE).
Meskipun setup awal nambah satu baris, kode di dalamnya jadi jauh lebih bersih daripada cara lama.