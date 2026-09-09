# NavalERP

Monorepo aplikasi NavalERP dengan frontend Next.js di `apps/web` dan backend
Go di `apps/api`.

## Deploy frontend ke Vercel

1. Import repository ini di Vercel dan pilih **Root Directory: `apps/web`**.
2. Gunakan preset **Next.js** dan **Node.js 22.x**. Install dan build command
   sudah diatur di `apps/web/vercel.json`; biarkan Output Directory default.
3. Tambahkan environment variable untuk Production dan Preview:

   ```dotenv
   NEXT_PUBLIC_API_BASE_URL=https://navalerpapi.virtualgate.id/api/v1
   NEXT_PUBLIC_API_FILESERVICE_URL=https://navalerpfileservice.virtualgate.id
   NEXT_PUBLIC_SSE_BASE_URL=https://navalerpapi.virtualgate.id/sse/v1
   ```

   Alamat di atas adalah contoh dari `.env.example`. Sesuaikan dengan endpoint
   backend dan file service yang sudah aktif dan dapat diakses melalui HTTPS.
4. Klik **Deploy**. Setelah mengganti environment variable, lakukan redeploy.

Deployment ini membangun frontend saja. Backend Go, database PostgreSQL, dan
file service perlu dijalankan terpisah. Pastikan backend dan file service
mengizinkan request dari domain frontend Vercel melalui CORS.

Jangan memilih root repository sebagai Root Directory Vercel: script `build`
di root juga membangun backend Go.

Panduan lengkap: [apps/web/README.md](apps/web/README.md).
Referensi: [Vercel monorepos](https://vercel.com/docs/monorepos).

## Verifikasi build lokal

Salin `apps/web/.env.example` ke `apps/web/.env.local` jika belum ada,
sesuaikan nilainya, lalu jalankan dari root repository:

```bash
pnpm install --frozen-lockfile
pnpm build:web
```
