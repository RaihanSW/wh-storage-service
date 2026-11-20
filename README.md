# wh-storage-service
- install swagger on local machine : make install-swag
- install mockery on local machine : make install-mockery
- download go dependencies : go mod tidy
- generate docs : make docs or swag init

## Setup Database (Migration & Seeding)

Service memakai 1 database Postgres: `warehouse_db` untuk semua
operasi (read & write, tanpa replica).
Migrasi/seeding dijalankan lewat subcommand di `main.go` — tidak perlu
eksekusi SQL manual ke DB.

### 1. Siapkan config

```bash
cp template_application_local.json application_local.json
```

Isi di `application_local.json`:

- `database.warehouse_db` →
  `username`, `password`, `hostname` (`host:port`), `dbname`
- `secretkey` → base64 string

`ReadConfig` membaca file `application_<env>.json` sesuai flag `-env`
(mis. `-env local` → `application_local.json`). `InitDBConnection`
membuka koneksi `pgx` ke tiap key di map `database` dengan DSN
`postgres://<username>:<password>@<hostname>/<dbname>`.

Service logs use Go's standard `log/slog` package and are written to stdout.

### 2. Jalankan migration

```bash
# lihat status migrasi (applied/pending)
go run main.go -env local migrate status

# apply semua migrasi yang pending
go run main.go -env local migrate up

# rollback — default 1 step, atau n step
go run main.go -env local migrate down
go run main.go -env local migrate down 2
```

Yang dibuat migrasi saat ini (folder `migration/`):

| Version | Name | Isi |
|---|---|---|
| 202611130001 | `create_storage_schema_tables` | schema `"Storage"` + 4 tabel (lihat daftar di bawah) |

Semua tabel yang dicakup migrasi project ini:

| Tabel | Keterangan |
|---|---|
| `"Storage".storage` | lokasi penyimpanan — `uid` VARCHAR(16) PK (DB-generated), `name`, `goods_uid` (tanpa FK), `warehouse_id` → `warehouse`, `block_id` → `block_layout`, `status` → `storage_status`, `x_row`, `y_row` + audit cols |
| `"Storage".storage_status` | lookup status storage — `id` INT PK identity, `name` VARCHAR(32) |
| `"Storage".warehouse` | gudang — `id` INT PK identity, `name` VARCHAR(64), `code` VARCHAR(32), `address` TEXT, `ttl_area` INT |
| `"Storage".block_layout` | layout blok per warehouse — `id` INT PK identity, `warehouse_id` → `warehouse`, `block_name`, `floor`, `x_sum`, `y_sum` |
| `storage_schema_migration` | tracking migrasi (auto-create, bukan bagian schema `"Storage"`) |

Audit cols di semua tabel: `created_at`, `created_by` VARCHAR(64),
`updated_at`, `updated_by` VARCHAR(64).

Tiap migrasi tercatat di tabel `storage_schema_migration`
(auto-create) di DB target. `migrate up` aman dijalankan berulang —
yang sudah applied di-skip. Target DB ditentukan field `DBKey` di
tiap file migrasi (selalu writer, bukan replica).

Menambah tabel/kolom baru = buat file `migration/<version>_<name>.go`,
register `Migration{Version, Name, DBKey, Up, Down}` via `init()`.
Pakai timestamp-style version, mis. `202611100001`.

### 3. Jalankan seeding

```bash
# semua seeder
go run main.go -env local seed

# satu seeder tertentu
go run main.go -env local seed storage_status
```

Seeder yang tersedia (folder `seed/`, idempotent — aman diulang):

| Name | Isi |
|---|---|
| `storage_status` | `Empty`(1), `Occupied`(2), `Maintenance`(3) → `"Storage".storage_status` |

Catatan: nilai lookup masih placeholder — sesuaikan di tiap file
`seed/*_seeder.go`. Seeder juga menyelaraskan identity sequence
(`setval`) setelah insert id eksplisit, agar create lewat API tidak
bentrok PK.

### 4. Start service

```bash
go run main.go -env local
```

Shortcut Makefile: `make migrate-up`, `make migrate-down STEPS=n`,
`make migrate-status`, `make seed NAME=storage_status` (`ENV=local` default).

## Testing

Unit test memakai stub hand-written (tanpa DB/network), contoh:
`app/service/storage_service_test.go` dan `app/service/warehouse_service_test.go`.

```bash
# semua test di package service
go test ./app/service/

# verbose — tampilkan tiap test & subtest
go test ./app/service/ -v

# semua test di project
go test ./...

# hanya test tertentu (regex)
go test ./app/service/ -run TestStorage -v

# hanya satu subtest
go test ./app/service/ -run 'TestStorageDelete/sukses' -v

# coverage
go test ./app/service/ -coverprofile=cover.out && go tool cover -func=cover.out
```
