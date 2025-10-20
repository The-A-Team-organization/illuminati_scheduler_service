## How to run this docker

# E-mail Service with scheduler and manual trigger

This project sends emails using Gmail SMTP and includes a MockServer setup for local testing.

---

## How to Run

### 1. Prerequisites

Before starting, make sure you have:

- Docker and Docker Compose installed
- A Google account with 2-Step Verification enabled
- Access to App Passwords (required for Gmail SMTP)

---

### 2. Generate a Gmail App Password

1. Go to [Google Account → Security → App passwords](https://myaccount.google.com/apppasswords)
2. Log in and verify your account
3. Choose "Mail" as the app and "Other" (for example, `DockerMailer`)
4. Copy the generated App Password — you will need it later

> Note: Regular Gmail passwords will not work. You must use the App Password.

---

### 3. Configure Environment Variables

Copy the example environment file and fill in your own secrets:

```bash
cp .env.example .env
```

---

### 4. Create MockServer Expectations

Change directory to `mock` in your project.  
Copy the example file and fill in your own data:

```bash
cd mock
cp participants.json.example participants.json
```

---

### 5. Run Docker Compose Configuration

This configuration builds and runs:
The Go email service container.
The MockServer container for testing.

```bash
docker-compose up --build
```

### 6. Verify Everything is Running

Open (http://127.0.0.1:1080/mockserver/dashboard) to view the MockServer dashboard

Trigger an email manually using the service endpoint (http://localhost:8080/trigger), or wait for the scheduler to send automatically

Check your terminal logs to see the email sending process
