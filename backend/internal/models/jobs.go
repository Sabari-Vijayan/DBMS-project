package models

import "time"

type Job struct {
    ID          int       `json:"id" db:"id"`
    EmployerID  int       `json:"employer_id" db:"employer_id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    CategoryID  *int      `json:"category_id" db:"category_id"`
    Location    string    `json:"location" db:"location"`
    SalaryMin   *float64  `json:"salary_min" db:"salary_min"`
    SalaryMax   *float64  `json:"salary_max" db:"salary_max"`
    Duration    *string   `json:"duration" db:"duration"`
    Requirements *string  `json:"requirements" db:"requirements"`
    ContactPhone *string  `json:"contact_phone" db:"contact_phone"`
    ContactEmail *string  `json:"contact_email" db:"contact_email"`
    ExpiresAt   time.Time `json:"expires_at" db:"expires_at"`
    Status      string    `json:"status" db:"status"`
    IsActive    bool      `json:"is_active" db:"is_active"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type JobWithDetails struct {
    Job
    EmployerName string  `json:"employer_name"`
    CategoryName *string `json:"category_name"`
}
