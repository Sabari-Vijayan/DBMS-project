package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "github.com/Sabari-Vijayan/DBMS-project/internal/db"
    "github.com/Sabari-Vijayan/DBMS-project/internal/handlers"
    "github.com/Sabari-Vijayan/DBMS-project/internal/middleware"  // Add this
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    database, err := db.Connect()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer database.Close()

    // Create handlers
    authHandler := &handlers.AuthHandler{DB: database}
    profileHandler := &handlers.ProfileHandler{DB: database}
    jobHandler := &handlers.JobHandler{DB: database}
    applicationHandler := &handlers.ApplicationHandler{DB: database}

    // Create Gin router
    router := gin.Default()

    // CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Public routes (no authentication required)
    router.POST("/api/register", authHandler.Register)
    router.POST("/api/login", authHandler.Login)
    router.GET("/api/jobs", jobHandler.GetJobs)              // Anyone can view jobs
    router.GET("/api/jobs/:id", jobHandler.GetJob)           // Anyone can view job details

    // Protected routes (authentication required)
    protected := router.Group("/api")
    protected.Use(middleware.AuthRequired())
    {
        // Profile routes
        protected.GET("/profile/:id", profileHandler.GetProfile)
        protected.PUT("/profile/:id", profileHandler.UpdateProfile)

        // Job routes (employers only)
        protected.POST("/jobs", middleware.EmployerOnly(), jobHandler.CreateJob)

        // Application routes (workers only)
        protected.POST("/applications", middleware.WorkerOnly(), applicationHandler.ApplyToJob)
        protected.GET("/applications/worker/:workerId", middleware.WorkerOnly(), applicationHandler.GetWorkerApplications)

        // Application routes (employers only)
        protected.GET("/applications/job/:jobId", middleware.EmployerOnly(), applicationHandler.GetJobApplications)
        protected.PUT("/applications/:id", middleware.EmployerOnly(), applicationHandler.UpdateApplicationStatus)
    }

    // Get port from environment or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    router.Run(":" + port)
}
