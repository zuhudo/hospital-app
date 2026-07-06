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
# Edit .env and set JWT_SECRET to a random value
# Generate one with: openssl rand -hex 32
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
cp .env.example .env
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
# Required — generate with: openssl rand -hex 32
JWT_SECRET=your-random-secret-here

# Server
PORT=3000

# CORS — comma-separated origins
ALLOWED_ORIGINS=http://localhost:5173

# Database (optional — in-memory for demo)
DB_HOST=localhost
DB_PORT=5432
DB_USER=hospital
DB_PASSWORD=your-db-password
DB_NAME=hospital_db
```

### Dashboard (.env)

```env
# API URL — point to your backend
VITE_API_URL=http://localhost:3000/api
```

## First Steps

1. Start the backend
2. Register a new account via the dashboard or API:
   ```bash
   curl -X POST http://localhost:3000/api/auth/register \
     -H "Content-Type: application/json" \
     -d '{"first_name":"John","last_name":"Doe","email":"john@example.com","phone":"+1234567890","password":"your-password"}'
   ```
3. Login with your registered credentials

## Docker

```bash
# Set required environment variables
export JWT_SECRET=$(openssl rand -hex 32)

docker-compose up -d
```

## Next Steps

- Read the [[Architecture]] guide
- Explore the [[API Documentation]]
- Check out the [[Contributing]] guide
