# 🏗️ Architecture

## System Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Mobile App    │     │  Public Website  │     │ Admin Dashboard │
│    (Flutter)    │     │    (Astro)       │     │  (React+Vite)   │
└────────┬────────┘     └────────┬─────────┘     └────────┬────────┘
         │                       │                         │
         └───────────────────────┼─────────────────────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │     Backend API        │
                    │   (Go + Fiber v2)      │
                    │   Port: 3000           │
                    └────────────┬───────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │     PostgreSQL         │
                    │   Port: 5432           │
                    └────────────────────────┘
```

## Component Details

### Mobile App (Flutter)
- **Pattern:** Provider + GoRouter
- **State:** Provider for global state
- **Navigation:** GoRouter declarative routing
- **Theme:** Material Design 3 with teal palette

### Backend API (Go)
- **Framework:** Fiber v2
- **Auth:** JWT with bcrypt passwords
- **Middleware:** CORS, Logger, Recovery, RequestID
- **Structure:** Clean architecture (handlers → models → routes)

### Public Website (Astro)
- **Rendering:** Static Site Generation (SSG)
- **Styling:** Tailwind CSS
- **Pages:** 5 static pages

### Admin Dashboard (React)
- **Build:** Vite 8
- **State:** Zustand
- **Routing:** React Router v7
- **Charts:** Recharts

## Data Models

### User
- id, first_name, last_name, email, phone, password, role

### Patient
- id, user_id, first_name, last_name, email, phone, date_of_birth, gender, address, blood_group, insurance_id, allergies

### Doctor
- id, user_id, first_name, last_name, email, phone, specialization, department, qualification, experience_years, consultation_fee, rating, is_available

### Appointment
- id, patient_id, doctor_id, appointment_date, time_slot, status, type, reason, notes, fee

### Medical Record
- id, patient_id, doctor_id, visit_date, diagnosis, symptoms, treatment, prescriptions, lab_results, notes

## API Flow

```
Client → JWT Token → Middleware → Handler → Response
  │
  ├── /api/auth/* (public)
  ├── /api/patients/* (protected)
  ├── /api/doctors/* (protected)
  ├── /api/appointments/* (protected)
  └── /api/records/* (protected)
```
