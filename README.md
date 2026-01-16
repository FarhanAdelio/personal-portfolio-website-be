# Personal Website - Backend API

Backend API untuk personal website menggunakan Golang dan Gin framework.

## Tech Stack

- **Go** 1.21+
- **Gin** - Web framework
- **PostgreSQL** - Database (optional)
- **godotenv** - Environment variable management

## Struktur API

### Endpoints

#### Profile
- `GET /api/v1/profile` - Get profile information

#### Experience
- `GET /api/v1/experiences` - Get all work experiences
- `GET /api/v1/experiences/:id` - Get specific experience

#### Education
- `GET /api/v1/education` - Get education history
- `GET /api/v1/education/:id` - Get specific education detail

#### Skills
- `GET /api/v1/skills` - Get all skills (technical, soft skills, tools)

#### Projects
- `GET /api/v1/projects` - Get all projects
- `GET /api/v1/projects/:id` - Get specific project

#### Contact
- `POST /api/v1/contact` - Send contact message

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Update `.env` dengan konfigurasi Anda

4. Run the application:
```bash
go run .
```

Server akan berjalan di `http://localhost:8080`

## Development

### Build
```bash
go build -o personal-website-backend
```

### Run
```bash
./personal-website-backend
```

## Next Steps

- [ ] Implement database models dan migrations
- [ ] Add authentication (jika diperlukan)
- [ ] Implement email service untuk contact form
- [ ] Add file upload untuk images
- [ ] Add validation dan error handling yang lebih robust
- [ ] Add unit tests
- [ ] Setup Docker container
