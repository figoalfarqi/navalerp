# NavalERP Frontend

Frontend Next.js untuk aplikasi NavalERP. Project ini berada di dalam pnpm
monorepo dan dikonfigurasi untuk deployment di Vercel.

## Local development

Salin `.env.example` menjadi `.env.local`, sesuaikan nilainya, lalu jalankan
dari root repository:

```bash
pnpm dev:web
```

Frontend tersedia di `http://localhost:3000`.

## Vercel project settings

Saat mengimpor Git repository di Vercel, gunakan pengaturan berikut:

| Setting | Value |
| --- | --- |
| Framework Preset | Next.js |
| Root Directory | `apps/web` |
| Build Command | `pnpm run build` |
| Output Directory | Biarkan default Next.js |
| Install Command | Biarkan default/otomatis |
| Node.js Version | `22.x` |

Konfigurasi yang dapat disimpan di repository berada di `vercel.json`.
Install Command sengaja tidak di-override agar Vercel membaca
`pnpm-lock.yaml` dan `packageManager` dari root monorepo.

## Vercel environment variables

Tambahkan variabel berikut melalui **Project Settings → Environment
Variables** untuk Production dan Preview:

| Variable | Required | Production example |
| --- | --- | --- |
| `NEXT_PUBLIC_API_BASE_URL` | Ya | `https://pmlapi.virtualgate.id/api/v1` |
| `NEXT_PUBLIC_API_FILESERVICE_URL` | Ya | `https://pmlfileservice.virtualgate.id` |
| `NEXT_PUBLIC_SSE_BASE_URL` | Jika SSE digunakan | `https://pmlapi.virtualgate.id/sse/v1` |
| `NEXT_PUBLIC_BIGDATACLOUD_API_KEY` | Jika reverse geocoding digunakan | Isi API key |
| `NEXT_PUBLIC_VAPID_PUBLIC_KEY` | Jika push notification digunakan | Isi public key VAPID |

Jangan menggunakan URL `localhost` pada environment Production atau Preview.
Semua variabel dengan awalan `NEXT_PUBLIC_` akan tersedia di browser, sehingga
jangan memasukkan private key, password, atau token rahasia.

Setelah mengubah environment variable, lakukan deployment ulang karena nilai
`NEXT_PUBLIC_` dimasukkan ke bundle pada waktu build.

## Production build

Jalankan dari root repository:

```bash
pnpm build:web
```

Build production akan dihentikan dengan pesan yang jelas jika URL API atau
file service belum dikonfigurasi.
