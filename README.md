# Job-seeker

## This was our project for our DBMS

## Overview

This is a platform designed similar to LinkedIn, but specifically for **handyman** and **local-gig work**. The backend is built with **Go (Golang)** and the frontend uses **React** (Vite). It provides a simple flow for employers to post jobs and for workers to apply and get hired.

---

# Backend

* Golang 1.21+
* Postgre SQL 14+
* Gin framework
* JWT Tokens

# Frontend

* React 18+ with Vite
* Axios

# Features

* Can register and login as a worker or employer
* Employer can post jobs
* Worker can apply for the jobs posted by Employers
* Worker can choose which applicant to accept
* Home Page for everyone to browse all available jobs

---

# How to run it on your device?

Good question...

First of all open your linux terminal or if you're on windows (Please switch) open your subshell

## 1. Install prerequisites

1. Install Go (Golang) and set its environment path.
2. Update your Node.js to a recent version.
3. Install PostgreSQL (v14+). Make sure `psql` CLI is available.

## 2. Clone this repository

Go to the directory where you want to work on this project and run:

```bash
git clone https://github.com/Sabari-Vijayan/DBMS-project.git
cd DBMS-project
```

## 3. Database setup

We will be using Postgres for this project. Make sure to install the latest version of Postgres on your device — check official docs for installation instructions.

On Linux, start and enable PostgreSQL:

```bash
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

### Connecting to Postgres

```bash
sudo -i -u postgres
```

**Do the next commands as it is itself (inside the postgres user session):**

In the PostgreSQL shell (psql), run:

```sql
CREATE DATABASE dbms_project;
CREATE USER dbms_user WITH PASSWORD 'your_password_here';
GRANT ALL PRIVILEGES ON DATABASE dbms_project TO dbms_user;
```

### Grant schema permissions

```sql
\c dbms_project
GRANT ALL ON SCHEMA public TO dbms_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dbms_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO dbms_user;
```

Exit the psql shell and postgres user session:

```bash
\q
exit
```

> Replace `your_password_here` with a strong password.

## 4. Running the migrations

From project root:

```bash
cd backend

# Run the first migration (users table)
psql -U dbms_user -d dbms_project -f migrations/001_create_users.sql

# Run the second migration (complete schema)
psql -U dbms_user -d dbms_project -f migrations/002_complete_schema.sql
```

## 5. Backend setup

```bash
cd backend
# create a .env file with the following content:
```

```
DATABASE_URL=postgres://dbms_user:your_password_here@localhost:5432/dbms_project?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

Then install Go dependencies:

```bash
go mod download
```

## 6. Start the server

```bash
go run cmd/server/main.go
```

## 7. Frontend

```bash
cd frontend
npm install
npm run dev
```

:)

ENJOY!

