<div align="center">

# 🏥 Hospital & Patient Management System

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)
[![GitHub Stars](https://img.shields.io/github/stars/zuhudo/hospital-app?style=social)](https://github.com/zuhudo/hospital-app/stargazers)
[![GitHub Issues](https://img.shields.io/github/issues/zuhudo/hospital-app)](https://github.com/zuhudo/hospital-app/issues)
[![GitHub Discussions](https://img.shields.io/github/discussions/zuhudo/hospital-app)](https://github.com/zuhudo/hospital-app/discussions)

A comprehensive, full-stack hospital and patient management system built with modern technologies.

[🚀 One-Click Setup](https://github.com/zuhudo/hospital-app-setup) • [📱 Mobile App](https://github.com/zuhudo/hospital-app-mobile) • [⚙️ Backend API](https://github.com/zuhudo/hospital-app-backend) • [🌐 Website](https://github.com/zuhudo/hospital-app-web) • [📊 Dashboard](https://github.com/zuhudo/hospital-app-dashboard) • [📖 Wiki](https://github.com/zuhudo/hospital-app/wiki)

</div>

---

## 📋 Table of Contents

- [Architecture](#-architecture)
- [Tech Stack](#-tech-stack)
- [Repositories](#-repositories)
- [Quick Start](#-quick-start)
- [Features](#-features)
- [API Reference](#-api-reference)
- [Documentation](#-documentation)
- [Contributing](#-contributing)
- [License](#-license)

## 🏗️ Architecture

```
hospital-app/
├── app/           → Flutter mobile app (iOS + Android)
├── backend/       → Go + Fiber REST API
└── web/
    ├── public/    → Astro public-facing website
    └── dashboard/ → Vite + React admin dashboard
```

## 🛠️ Tech Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **[📱 Mobile App](https://github.com/zuhudo/hospital-app-mobile)** | Flutter / Dart | Cross-platform iOS & Android app |
| **[⚙️ Backend API](https://github.com/zuhudo/hospital-app-backend)** | Go + Fiber v2 | High-performance REST API |
| **[🌐 Public Website](https://github.com/zuhudo/hospital-app-web)** | Astro + Tailwind CSS | Hospital's public marketing site |
| **[📊 Admin Dashboard](https://github.com/zuhudo/hospital-app-dashboard)** | Vite + React + TypeScript | Internal admin management panel |

## 📦 Repositories

| Repository | Description | Tech |
|------------|-------------|------|
| [hospital-app](https://github.com/zuhudo/hospital-app) | Main repo with docs & overview | — |
| [hospital-app-setup](https://github.com/zuhudo/hospital-app-setup) | One-click monorepo setup scripts | Bash |
| [hospital-app-mobile](https://github.com/zuhudo/hospital-app-mobile) | Mobile app for patients & doctors | Flutter |
| [hospital-app-backend](https://github.com/zuhudo/hospital-app-backend) | REST API server | Go + Fiber |
| [hospital-app-web](https://github.com/zuhudo/hospital-app-web) | Public hospital website | Astro |
| [hospital-app-dashboard](https://github.com/zuhudo/hospital-app-dashboard) | Admin dashboard | React + Vite |

## 🚀 Quick Start

### Prerequisites

- [Flutter SDK](https://docs.flutter.dev/get-started/install) 3.2+
- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+
- [pnpm](https://pnpm.io/) 8+
- [PostgreSQL](https://www.postgresql.org/download/) (optional, in-memory for demo)

### 0. One-Click Setup (Recommended)

```bash
git clone https://github.com/zuhudo/hospital-app-setup.git
cd hospital-app-setup
./setup.sh
```

This will clone all repos, install dependencies, and verify everything builds. See [hospital-app-setup](https://github.com/zuhudo/hospital-app-setup) for more options.

### 1. Manual Setup — Clone All Repos

```bash
# Clone the main repo
git clone https://github.com/zuhudo/hospital-app.git
cd hospital-app

# Clone all component repos
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
# API running at http://localhost:3000
```

### 3. Start Public Website

```bash
cd web/public
pnpm install
pnpm dev
# Website at http://localhost:4321
```

### 4. Start Dashboard

```bash
cd web/dashboard
pnpm install
pnpm dev
# Dashboard at http://localhost:5173
```

### 5. Run Mobile App

```bash
cd app
flutter pub get
flutter run
```

## ✨ Features

### 📱 Mobile App
- 🔐 Login & Registration
- 🏠 Home dashboard with stats
- 👥 Patient management
- 📅 Appointment booking with calendar
- 👨‍⚕️ Doctor directory
- 📋 Medical records
- 👤 User profile
- 🎨 Material Design 3 theme

### ⚙️ Backend API
- 🔑 JWT authentication
- 👥 Patient CRUD
- 👨‍⚕️ Doctor management
- 📅 Appointment scheduling
- 📋 Medical records
- 🛡️ CORS, logging, recovery middleware
- 🐳 Docker support

### 🌐 Public Website
- 📄 5 pages (Home, About, Services, Doctors, Contact)
- 📱 Fully responsive
- 🎨 Tailwind CSS
- ⚡ Astro static site generation

### 📊 Admin Dashboard
- 📊 Dashboard with charts
- 👥 Patient table
- 👨‍⚕️ Doctor cards
- 📅 Appointment tracking
- 📋 Medical records
- 💰 Billing overview
- ⚙️ Settings
- 🔒 Protected routes

## 📡 API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/login` | User login |
| POST | `/api/auth/register` | User registration |
| GET | `/api/patients` | List all patients |
| POST | `/api/patients` | Create patient |
| GET | `/api/patients/:id` | Get patient details |
| PUT | `/api/patients/:id` | Update patient |
| DELETE | `/api/patients/:id` | Delete patient |
| GET | `/api/doctors` | List all doctors |
| POST | `/api/doctors` | Create doctor |
| GET | `/api/doctors/:id` | Get doctor details |
| POST | `/api/appointments` | Book appointment |
| GET | `/api/appointments` | List appointments |
| PUT | `/api/appointments/:id/cancel` | Cancel appointment |
| GET | `/api/records/:patientId` | Get medical records |
| POST | `/api/records` | Create medical record |

## 📖 Documentation

- 📖 [Wiki](https://github.com/zuhudo/hospital-app/wiki) — Comprehensive guides
- 🏗️ [Architecture](https://github.com/zuhudo/hospital-app/wiki/Architecture) — System design
- 🚀 [Getting Started](https://github.com/zuhudo/hospital-app/wiki/Getting-Started) — Setup guide
- 📡 [API Docs](https://github.com/zuhudo/hospital-app/wiki/API-Documentation) — REST API reference
- 🎨 [Design System](https://github.com/zuhudo/hospital-app/wiki/Design-System) — UI components
- 🔧 [Deployment](https://github.com/zuhudo/hospital-app/wiki/Deployment) — Production setup

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**[⬆ Back to Top](#-hospital--patient-management-system)**

Made with ❤️ by [Zuhudo](https://github.com/zuhudo)

</div>
