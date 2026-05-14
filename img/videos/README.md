# Video Demo untuk Projects

Folder ini untuk menyimpan video demo dari project-project yang kamu buat.

## Format Video yang Didukung

1. **MP4** (Recommended) - Kompatibilitas terbaik di semua browser
2. **WebM** - Ukuran lebih kecil, cocok untuk web

## Cara Merekam & Membuat Video Demo

### 1. Merekam Layar
**MacOS:**
- Tekan `Cmd + Shift + 5`
- Pilih "Record Selected Portion" atau "Record Entire Screen"
- Klik Options → pilih "Show Mouse Clicks"
- Mulai rekam dan tunjukkan fitur-fitur website

**Windows:**
- Tekan `Win + G` (Xbox Game Bar)
- Klik tombol rekam
- Atau gunakan software: OBS Studio (gratis)

**Tools Recommended:**
- **OBS Studio** (gratis, powerful)
- **Loom** (gratis, langsung upload)
- **Screen Studio** (Mac, premium tapi hasil sangat bagus)

### 2. Edit Video (Optional)
- **iMovie** (Mac) - gratis
- **DaVinci Resolve** (gratis, profesional)
- **CapCut** (gratis, mudah digunakan)

Tips editing:
- Tambahkan intro text (judul project)
- Highlight fitur penting dengan zoom/pan
- Tambahkan music background (volume rendah)
- Maksimal durasi: 1-2 menit per project

### 3. Optimize Video untuk Web

**Kompres video agar loading cepat:**

```bash
# Install ffmpeg (MacOS)
brew install ffmpeg

# Kompres video menjadi MP4 (kualitas bagus, ukuran kecil)
ffmpeg -i input.mov -c:v libx264 -crf 23 -preset medium -c:a aac -b:a 128k portfolio-demo.mp4

# Buat versi WebM (backup)
ffmpeg -i input.mov -c:v libvpx-vp9 -crf 30 -b:v 0 -b:a 128k -c:a libopus portfolio-demo.webm
```

**Target ukuran file:**
- 1 menit video: 5-10 MB
- 2 menit video: 10-20 MB

### 4. Cara Menggunakan Video

#### Opsi 1: Video Lokal (Simpan di folder ini)

1. Letakkan video di folder `backend/videos/`
2. Update `handlers.go`:

```go
{
    "id": 1,
    "title": "My Awesome Project",
    "video": "/videos/portfolio-demo.mp4",
    "videoType": "local",
    // ... fields lainnya
}
```

#### Opsi 2: YouTube (Recommended untuk file besar)

1. Upload video ke YouTube
2. Set visibility: Public atau Unlisted
3. Copy Video ID dari URL
   - URL: `https://www.youtube.com/watch?v=dQw4w9WgXcQ`
   - Video ID: `dQw4w9WgXcQ`
4. Update `handlers.go`:

```go
{
    "id": 1,
    "title": "My Awesome Project",
    "video": "dQw4w9WgXcQ",  // Video ID aja
    "videoType": "youtube",
    // ... fields lainnya
}
```

#### Opsi 3: Vimeo

1. Upload ke Vimeo
2. Copy Video ID
3. Update `handlers.go`:

```go
{
    "id": 1,
    "title": "My Awesome Project",
    "video": "123456789",  // Vimeo video ID
    "videoType": "vimeo",
    // ... fields lainnya
}
```

## Struktur File

```
backend/videos/
├── README.md (file ini)
├── portfolio-demo.mp4      (contoh)
├── portfolio-demo.webm     (versi WebM, optional)
├── api-service-demo.mp4
└── task-app-demo.mp4
```

## Tips Membuat Video Demo yang Menarik

1. **Siapkan Script** - Tulis poin-poin fitur yang mau ditunjukkan
2. **Clean Browser** - Tutup tab yang tidak perlu, hide bookmark bar
3. **Smooth Navigation** - Jangan terlalu cepat klik-klik
4. **Show Key Features** - 
   - Login/Authentication
   - CRUD operations
   - Real-time features
   - Responsive design (resize browser)
5. **Background Music** - Tambahkan music instrumental (volume 20-30%)
6. **Text Overlays** - Tambahkan label untuk fitur penting
7. **Cursor Highlight** - Aktifkan highlight untuk memudahkan viewer

## Contoh Struktur Video (1-2 menit)

1. **Intro (5 detik)**
   - Title card: "Project Name - Demo"
   
2. **Overview (10 detik)**
   - Tampilkan landing page
   - Quick scroll untuk lihat design
   
3. **Features (40-60 detik)**
   - Fitur 1: Login/Register (10s)
   - Fitur 2: Main functionality (20s)
   - Fitur 3: Special feature (15s)
   - Fitur 4: Responsive demo (10s)
   
4. **Outro (5 detik)**
   - End screen dengan GitHub link/Logo

## Troubleshooting

### Video tidak muncul?
- Cek path video di `handlers.go` benar
- Pastikan backend server running
- Cek browser console untuk error

### Video terlalu besar?
- Kompres dengan ffmpeg (lihat di atas)
- Atau upload ke YouTube instead

### Video lag saat play?
- Reduce bitrate saat kompres
- Atau gunakan YouTube embed
