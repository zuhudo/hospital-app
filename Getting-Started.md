# 🚀 Getting Started

## Prerequisites

| Tool | Version | Install |
|------|---------|---------|
| Flutter SDK | 3.2+ | [flutter.dev](https://docs.flutter.dev/get-started/install) |
| Go | 1.21+ | [go.dev](https://go.dev/dl/) |
| Node.js | 18+ | [nodejs.org](https://nodejs.org/) |
| pnpm | 8+ | `npm install -g pnpm` |
| Docker | optional | [docker.com](https://www.docker.com/) |

## Installation

### 1. Clone All Repositories

```bash
# Main monorepo
git clone https://github.com/zuhudo/hospital-app.git
cd hospital-app

# Clone component repos
git clone https://github.com/zuhudo/hospital-app-mobile.git app
git clone https://github.com/zuhudo/hospital-app-backend.git backend
git clone https://github.com/zuhudo/hospital-app-web.git web/public
git clone https://github.com/zuhudo/hospital-app-dashboard.git web/dashboard
```

### 2. Start Backend

```bash
cd backend
cp .env.example .env
go mod tidy
go run main.go
```

API runs at `http://localhost:3000`

### 3. Start Public Website

```bash
cd web/public
pnpm install
pnpm dev
```

Website at `http://localhost:4321`

### 4. Start Admin Dashboard

```bash
cd web/dashboard
pnpm install
pnpm dev
```

Dashboard at `http://localhost:5173`

### 5. Run Mobile App

```bash
cd app
flutter pub get
flutter run
```

## Environment Variables

### Backend (.env)

```env
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=hospital
DB_PASSWORD=hospital_secret
DB_NAME=hospital_db
JWT_SECRET=your-secret-key
ALLOWED_ORIGINS=*
```

## Default Credentials

- **Email:** admin@hospital.com
- **Password:** admin123

## Docker

```bash
docker-compose up -d
```

This starts PostgreSQL, Backend API, Public Website, and Dashboard.

## Next Steps

- Read the [[Architecture]] guide
- Explore the [[API Documentation]]
- Check out the [[Contributing]] guide
