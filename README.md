# 🌱 Golang HTTP Fundamentals: Catatan Belajar Saya

Halo! 👋

Repositori ini adalah tempat saya **"ngulik" dan belajar** dasar-dasar backend menggunakan Golang (`net/http`).

Jujur saja, saya masih dalam proses belajar dan belum begitu mahir. Jadi, repo ini sengaja saya buat sebagai **dokumentasi pribadi** agar saya tidak lupa konsep-konsep penting sebelum nanti lanjut menggunakan Framework (seperti Gin, Echo, Fiber).

---

## 💡 Apa yang Saya Pelajari? (Kesimpulan)

Daripada bingung dengan banyak folder, intinya project-project di sini adalah bukti pemahaman saya terhadap 4 hal dasar ini:

1.  **Fundamental Server:** Ternyata membuat web server di Go itu simpel, cukup pakai library bawaan `net/http` tanpa perlu install apa-apa.
2.  **Routing Go 1.22:** Saya belajar fitur baru Go (versi 1.22) yang ternyata routing-nya sudah canggih (bisa baca `GET /barang/{id}` langsung).
3.  **Bahaya Global Variable:** Saya baru paham kalau menyimpan data user di variabel global itu **berbahaya** (bisa tertukar datanya saat trafik tinggi). Solusinya wajib pakai **Context**.
4.  **Middleware & Context:** Ini bagian tersulit tapi paling penting. Saya belajar gimana caranya "menitipkan" data user dari middleware ke handler dengan aman.

---

## 🎯 Tujuan Akhir

Kode-kode di sini mungkin belum siap untuk *production*, tapi ini adalah **pondasi** saya. Target saya selanjutnya adalah membawa pemahaman manual ini untuk belajar Framework  dengan lebih lancar.

*Selamat belajar buat saya, dan semoga bermanfaat buat yang mampir!* 🚀