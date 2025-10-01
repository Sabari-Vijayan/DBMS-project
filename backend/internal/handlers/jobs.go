package handlers

import (
    "context"
    "database/sql"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
		//"strconv"
)

type JobHandler struct {
    DB *pgxpool.Pool
}

type CreateJobRequest struct {
    EmployerID   int      `json:"employer_id" binding:"required"` // Add this line
	  Title        string   `json:"title" binding:"required"`
    Description  string   `json:"description" binding:"required"`
    CategoryID   *int     `json:"category_id"`
    Location     string   `json:"location" binding:"required"`
    SalaryMin    *float64 `json:"salary_min"`
    SalaryMax    *float64 `json:"salary_max"`
    Duration     string   `json:"duration"`
    Requirements string   `json:"requirements"`
    ContactPhone string   `json:"contact_phone"`
    ContactEmail string   `json:"contact_email"`
    ExpiryDays   int      `json:"expiry_days" binding:"required,min=1,max=7"` // 1-7 days
}

// Create new job posting
func (h *JobHandler) CreateJob(c *gin.Context) {
    var req CreateJobRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: Get employer ID from JWT token
    // For now, we'll accept it from request or use a default
    /* employerID := c.GetInt("user_id") // We'll implement this with JWT later
    if employerID == 0 {
        // Temporary: accept from query param for testing
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Employer ID required"})
        return
    }*/
		// Use employer_id from request
employerID := 9

// Verify the user is an employer
var userType string
checkQuery := `SELECT user_type FROM users WHERE id = $1`
err := h.DB.QueryRow(context.Background(), checkQuery, employerID).Scan(&userType)

if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employer ID"})
    return
}

if userType != "employer" {
    c.JSON(http.StatusForbidden, gin.H{"error": "Only employers can post jobs"})
    return
}

    // Calculate expiry date
    expiresAt := time.Now().AddDate(0, 0, req.ExpiryDays)

    query := `
        INSERT INTO jobs (
            employer_id, title, description, category_id, location,
            salary_min, salary_max, duration, requirements,
            contact_phone, contact_email, expires_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING id, employer_id, title, description, category_id, location,
                  salary_min, salary_max, duration, requirements,
                  contact_phone, contact_email, expires_at, status, is_active, created_at
    `

    var job struct {
        ID          int
        EmployerID  int
        Title       string
        Description string
        CategoryID  sql.NullInt64
        Location    string
        SalaryMin   sql.NullFloat64
        SalaryMax   sql.NullFloat64
        Duration    sql.NullString
        Requirements sql.NullString
        ContactPhone sql.NullString
        ContactEmail sql.NullString
        ExpiresAt   time.Time
        Status      string
        IsActive    bool
        CreatedAt   time.Time
    }

    err = h.DB.QueryRow(context.Background(), query,
        employerID, req.Title, req.Description, req.CategoryID, req.Location,
        req.SalaryMin, req.SalaryMax, req.Duration, req.Requirements,
        req.ContactPhone, req.ContactEmail, expiresAt,
    ).Scan(
        &job.ID, &job.EmployerID, &job.Title, &job.Description, &job.CategoryID,
        &job.Location, &job.SalaryMin, &job.SalaryMax, &job.Duration,
        &job.Requirements, &job.ContactPhone, &job.ContactEmail,
        &job.ExpiresAt, &job.Status, &job.IsActive, &job.CreatedAt,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Job created successfully",
        "job": job,
    })
}

// Get all active jobs
func (h *JobHandler) GetJobs(c *gin.Context) {
    query := `
        SELECT j.id, j.employer_id, j.title, j.description, j.category_id, j.location,
               j.salary_min, j.salary_max, j.duration, j.requirements,
               j.contact_phone, j.contact_email, j.expires_at, j.status, j.is_active, j.created_at,
               u.full_name as employer_name,
               c.name as category_name
        FROM jobs j
        JOIN users u ON j.employer_id = u.id
        LEFT JOIN categories c ON j.category_id = c.id
        WHERE j.is_active = true 
          AND j.status = 'open'
          AND j.expires_at > NOW()
        ORDER BY j.created_at DESC
    `

    rows, err := h.DB.Query(context.Background(), query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
        return
    }
    defer rows.Close()

    var jobs []map[string]interface{}

    for rows.Next() {
        var (
            id, employerID int
            title, description, location, status, employerName string
            categoryID sql.NullInt64
            salaryMin, salaryMax sql.NullFloat64
            duration, requirements, contactPhone, contactEmail, categoryName sql.NullString
            expiresAt, createdAt time.Time
            isActive bool
        )

        err := rows.Scan(
            &id, &employerID, &title, &description, &categoryID, &location,
            &salaryMin, &salaryMax, &duration, &requirements,
            &contactPhone, &contactEmail, &expiresAt, &status, &isActive, &createdAt,
            &employerName, &categoryName,
        )

        if err != nil {
            continue
        }

        job := map[string]interface{}{
            "id":            id,
            "employer_id":   employerID,
            "employer_name": employerName,
            "title":         title,
            "description":   description,
            "location":      location,
            "status":        status,
            "is_active":     isActive,
            "expires_at":    expiresAt,
            "created_at":    createdAt,
        }

        if categoryID.Valid {
            job["category_id"] = categoryID.Int64
        }
        if categoryName.Valid {
            job["category_name"] = categoryName.String
        }
        if salaryMin.Valid {
            job["salary_min"] = salaryMin.Float64
        }
        if salaryMax.Valid {
            job["salary_max"] = salaryMax.Float64
        }
        if duration.Valid {
            job["duration"] = duration.String
        }
        if requirements.Valid {
            job["requirements"] = requirements.String
        }
        if contactPhone.Valid {
            job["contact_phone"] = contactPhone.String
        }
        if contactEmail.Valid {
            job["contact_email"] = contactEmail.String
        }

        jobs = append(jobs, job)
    }

    c.JSON(http.StatusOK, gin.H{
        "jobs": jobs,
        "count": len(jobs),
    })
}

// Get single job by ID
func (h *JobHandler) GetJob(c *gin.Context) {
    jobID := c.Param("id")

    query := `
        SELECT j.id, j.employer_id, j.title, j.description, j.category_id, j.location,
               j.salary_min, j.salary_max, j.duration, j.requirements,
               j.contact_phone, j.contact_email, j.expires_at, j.status, j.is_active, j.created_at,
               u.full_name as employer_name,
               c.name as category_name
        FROM jobs j
        JOIN users u ON j.employer_id = u.id
        LEFT JOIN categories c ON j.category_id = c.id
        WHERE j.id = $1
    `

    var (
        id, employerID int
        title, description, location, status, employerName string
        categoryID sql.NullInt64
        salaryMin, salaryMax sql.NullFloat64
        duration, requirements, contactPhone, contactEmail, categoryName sql.NullString
        expiresAt, createdAt time.Time
        isActive bool
    )

    err := h.DB.QueryRow(context.Background(), query, jobID).Scan(
        &id, &employerID, &title, &description, &categoryID, &location,
        &salaryMin, &salaryMax, &duration, &requirements,
        &contactPhone, &contactEmail, &expiresAt, &status, &isActive, &createdAt,
        &employerName, &categoryName,
    )

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
        return
    }

    job := map[string]interface{}{
        "id":            id,
        "employer_id":   employerID,
        "employer_name": employerName,
        "title":         title,
        "description":   description,
        "location":      location,
        "status":        status,
        "is_active":     isActive,
        "expires_at":    expiresAt,
        "created_at":    createdAt,
    }

    if categoryID.Valid {
        job["category_id"] = categoryID.Int64
    }
    if categoryName.Valid {
        job["category_name"] = categoryName.String
    }
    if salaryMin.Valid {
        job["salary_min"] = salaryMin.Float64
    }
    if salaryMax.Valid {
        job["salary_max"] = salaryMax.Float64
    }
    if duration.Valid {
        job["duration"] = duration.String
    }
    if requirements.Valid {
        job["requirements"] = requirements.String
    }
    if contactPhone.Valid {
        job["contact_phone"] = contactPhone.String
    }
    if contactEmail.Valid {
        job["contact_email"] = contactEmail.String
    }

    c.JSON(http.StatusOK, job)
}
