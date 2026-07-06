# 📡 API Documentation

Base URL: `http://localhost:3000/api`

## Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

### POST /api/auth/login

Login with email and password.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "your-password"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": "uuid",
      "first_name": "John",
      "last_name": "Doe",
      "email": "user@example.com",
      "role": "patient"
    }
  }
}
```

### POST /api/auth/register

Register a new user.

**Request:**
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "+1234567890",
  "password": "password123"
}
```

## Patients

### GET /api/patients

List all patients. Requires auth.

### POST /api/patients

Create a new patient.

**Request:**
```json
{
  "first_name": "Jane",
  "last_name": "Smith",
  "email": "jane@example.com",
  "phone": "+0987654321",
  "date_of_birth": "1990-05-15",
  "gender": "Female",
  "blood_group": "A+"
}
```

### GET /api/patients/:id

Get patient by ID.

### PUT /api/patients/:id

Update patient.

### DELETE /api/patients/:id

Delete patient.

## Doctors

### GET /api/doctors

List all doctors.

### POST /api/doctors

Create a new doctor.

### GET /api/doctors/:id

Get doctor by ID.

## Appointments

### GET /api/appointments

List all appointments.

### POST /api/appointments

Book an appointment.

**Request:**
```json
{
  "patient_id": "1",
  "doctor_id": "1",
  "appointment_date": "2024-02-15",
  "time_slot": "10:00 AM",
  "type": "consultation",
  "reason": "Annual checkup"
}
```

### PUT /api/appointments/:id/cancel

Cancel an appointment.

## Medical Records

### GET /api/records/:patientId

Get records for a patient.

### POST /api/records

Create a medical record.

## Error Responses

```json
{
  "success": false,
  "error": "Error message here"
}
```

| Status | Meaning |
|--------|---------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 404 | Not Found |
| 500 | Server Error |
