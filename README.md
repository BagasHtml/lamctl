# lamctl

CLI helper untuk menyalakan XAMPP/LAMPP dan mengelola database MySQL dengan mudah — tanpa perlu ketik command panjang setiap kali.

> **Kenapa dibuat?** Karena males bolak-balik ke folder lampp (`/opt/lampp`) dan mengetik command panjang. Di Linux, XAMPP tidak punya GUI seperti Windows untuk menyalakan/mengelola database harus ketik `sudo /opt/lampp/lampp start`, `mysql -u root -e "SHOW DATABASES"`, dsb. yang panjang dan gampang lupa. `lamctl` membungkus semua itu jadi command yang singkat, konsisten, dan mudah diingat, dari mana saja.

## Fitur

- Control layanan XAMPP/LAMPP (start, stop, restart, status)
- Cek status per-service (misal cuma MySQL)
- Operasi database MySQL (list, query, create, drop)
- Buka mysql client interaktif (`lamctl mysql`)
- Setup kredensial interaktif (`lamctl init`)
- Support override via flag CLI, environment variable `.env`, dan nilai default
- Terintegrasi ke sistem — bisa dipanggil dari mana saja

## Instalasi

### Prasyarat

- [Go](https://golang.org/dl/) >= 1.20
- XAMPP/LAMPP terinstall di `/opt/lampp` (default, bisa diubah via env `LAMCTL_XAMPP_PATH`)

### Build dari source

```bash
make install
```

Ini akan build binary lalu menyalinnya ke `/usr/local/bin/lamctl`, sehingga bisa dipanggil dari direktori mana pun.

### Uninstall

```bash
make uninstall
```

## Setup Awal Kredensial

Jalankan `lamctl init` untuk setup kredensial database secara interaktif:

```bash
lamctl init
```

```text
? Database host (localhost):
? Database port (3306):
? Database username (root):
? Database password (kosongkan jika tidak ada):
? Database name (kosongkan untuk koneksi umum):
? Database Engine (mysql):

Kredensial tersimpan di .env
```

Kredensial disimpan ke file `.env` di direktori saat ini. File ini sudah di-ignore oleh git (lihat `.gitignore`).

Contoh template ini juga tersedia di `.env.example`:

```env
LAMCTL_DB_HOST=localhost
LAMCTL_DB_PORT=3306
LAMCTL_DB_USER=root
LAMCTL_DB_PASS=
LAMCTL_DB_NAME=mysql
```

## Penggunaan

### Kontrol XAMPP/LAMPP

```bash
lamctl start              # Start semua layanan XAMPP
lamctl stop               # Stop semua layanan XAMPP
lamctl restart            # Restart semua layanan XAMPP
lamctl status             # Status semua layanan
lamctl status mysql       # Status MySQL saja
```

### Operasi Database

```bash
lamctl db list                                    # List semua database
lamctl db query "SELECT * FROM users"             # Jalankan query SQL
lamctl db create nama_database                    # Buat database baru
lamctl db drop nama_database                      # Hapus database
lamctl mysql                                      # Buka mysql client interaktif
lamctl mysql --db nama_database                   # Buka client langsung ke database tertentu
```

> `lamctl mysql` membuka shell mysql interaktif (setara `mysql -u <user> -p`)
> tanpa harus `cd` ke `/opt/lampp/bin` dan tanpa `sudo`. Pakai credential dari `.env`
> atau override flag. Password dikirim lewat env `MYSQL_PWD`, tidak tampil di proses.

### Bantuan

```bash
lamctl                    # Tampilkan logo + daftar command & flag
lamctl --help
lamctl db --help
```

## Prioritas Konfigurasi

Nilai kredensial diambil dengan urutan prioritas berikut (tertinggi dulu):

1. **CLI flag** — `--host`, `--port`, `--user`, `--password`, `--db`, `--db_engine`
2. **Environment variable** — dari file `.env` (auto-load) atau variabel lingkungan
3. **Nilai default** — `localhost`, `3306`, `root`

Contoh override via flag:

```bash
lamctl db list --host 192.168.1.10 --port 3307 --user admin --password secret
```

### Variabel Environment

| Variabel | Default | Deskripsi |
|---|---|---|
| `LAMCTL_DB_HOST` | `localhost` | Host MySQL |
| `LAMCTL_DB_PORT` | `3306` | Port MySQL |
| `LAMCTL_DB_USER` | `root` | Username MySQL |
| `LAMCTL_DB_PASS` | empty | Password MySQL |
| `LAMCTL_DB_NAME` | empty | Nama database default |
| `LAMCTL_DB_ENGINE` | empty | Database engine yang digunakan (misal `mysql`) |
| `LAMCTL_XAMPP_PATH` | `/opt/lampp/lampp` | Path binary lampp |

## Command Reference

| Command | Deskripsi |
|---|---|
| `lamctl init` | Setup kredensial database interaktif |
| `lamctl start` | Start semua layanan XAMPP |
| `lamctl stop` | Stop semua layanan XAMPP |
| `lamctl restart` | Restart semua layanan XAMPP |
| `lamctl status [service]` | Cek status layanan (dengan/tanpa nama service) |
| `lamctl db list` | List semua database |
| `lamctl db query "<SQL>"` | Jalankan query SQL |
| `lamctl db create <name>` | Buat database |
| `lamctl db drop <name>` | Hapus database |
| `lamctl mysql` | Buka mysql client interaktif |

### Global Flags

| Flag | Deskripsi |
|---|---|
| `--host` | Host database |
| `--port` | Port database |
| `--user` | Username database |
| `--password` | Password database |
| `--db` | Nama database |
| `--db_engine` | Engine database yang digunakan (misal `mysql`) |
| `-h, --help` | Tampilkan bantuan |

## Arsitektur

Mengikuti prinsip Clean Architecture, dengan pemisahan tanggung jawab antar layer:

```
cmd/                     Delivery — handler perintah CLI (cobra)
internal/
  entity/                Entity — model domain (Database, Credential)
  repository/            Repository — akses data (MySQL, XAMPP, setting)
  usecase/               Use Case — logika bisnis (dbmanager, lamppctrl)
```

Alur permintaan: **CLI command** → **Use Case** (logika bisnis) → **Repository** (akses data/IO).

## Pengembangan

```bash
go build -o lamctl .     # Build binary
go vet ./...             # Static analysis
make clean               # Hapus binary hasil build
```

## Roadmap

- [ ] Dukungan database lain (PostgreSQL, SQLite) via engine abstraction
- [ ] Backup & restore database
- [ ] Export/import database
- [ ] Dukungan file `.env` global di home directory

## Lisensi

MIT
