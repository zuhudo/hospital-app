<div align="center">

# 🌐 Hospital Website — Astro

[![Astro](https://img.shields.io/badge/Astro-5-blue?logo=astro)](https://astro.build)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind-3-blue?logo=tailwindcss)](https://tailwindcss.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/zuhudo/hospital-app-web/actions/workflows/astro-ci.yml/badge.svg)](https://github.com/zuhudo/hospital-app-web/actions/workflows/astro-ci.yml)

Public-facing hospital website built with Astro and Tailwind CSS.

[📱 Mobile App](https://github.com/zuhudo/hospital-app-mobile) • [⚙️ Backend](https://github.com/zuhudo/hospital-app-backend) • [📊 Dashboard](https://github.com/zuhudo/hospital-app-dashboard) • [📖 Wiki](https://github.com/zuhudo/hospital-app-web/wiki)

</div>

## 📋 Features

- 📄 **5 Pages** — Home, About, Services, Doctors, Contact
- 📱 **Fully Responsive** — Mobile-first design
- ⚡ **Static Site** — Blazing fast with Astro SSG
- 🎨 **Tailwind CSS** — Modern utility-first styling
- 🏥 **Medical Theme** — Teal/blue professional design
- ♿ **Accessible** — Semantic HTML & ARIA labels

## 🚀 Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) 18+
- [pnpm](https://pnpm.io/) 8+

### Installation

```bash
# Clone the repo
git clone https://github.com/zuhudo/hospital-app-web.git
cd hospital-app-web

# Install dependencies
pnpm install

# Start dev server
pnpm dev
```

The site will be available at `http://localhost:4321`

### Build

```bash
# Build for production
pnpm build

# Preview production build
pnpm preview
```

## 📁 Project Structure

```
web/public/
├── src/
│   ├── layouts/       # Page layouts
│   ├── pages/         # Route pages
│   │   ├── index.astro
│   │   ├── about.astro
│   │   ├── services.astro
│   │   ├── doctors.astro
│   │   └── contact.astro
│   ├── components/    # UI components
│   │   ├── Header.astro
│   │   ├── Footer.astro
│   │   ├── Hero.astro
│   │   ├── ServiceCard.astro
│   │   ├── DoctorCard.astro
│   │   └── TestimonialCard.astro
│   └── styles/        # Global CSS
├── public/            # Static assets
├── astro.config.mjs   # Astro config
└── tailwind.config.mjs # Tailwind config
```

## 🛠️ Tech Stack

- **Astro** 5 — Static site generator
- **Tailwind CSS** 3 — Utility-first CSS
- **TypeScript** — Type safety

## 📖 Documentation

- [Getting Started](https://github.com/zuhudo/hospital-app-web/wiki/Getting-Started)
- [Components](https://github.com/zuhudo/hospital-app-web/wiki/Components)
- [Deployment](https://github.com/zuhudo/hospital-app-web/wiki/Deployment)

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📄 License

MIT License — see [LICENSE](LICENSE) for details.
