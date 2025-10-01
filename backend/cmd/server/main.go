package main

import(
	"log"
	"os"

  "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
	"github.com/Sabari-Vijayan/DBMS-project/internal/db"
	"github.com/Sabari-Vijayan/DBMS-project/internal/handlers"
)

func main() {

  //Loading the environment variables
	if err:=godotenv.Load();err != nil{
		log.Println("No .env file found")
	}

	//INITIALIZING THE DATABASE
	database,err:=db.Connect()
	if err!=nil{
		log.Fatal("Failed to connect to database",err)
	}
	defer database.Close()

  //CREATE handlers
	authHandler := &handlers.AuthHandler{DB: database}
  profileHandler := &handlers.ProfileHandler{DB: database}  // Add this

	//CREATING Gin router
	router:=gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // React dev server
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	//Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"Status": "ok"})
	})

  //Auth routes
	router.POST("api/register", authHandler.Register)
	router.POST("api/login", authHandler.Login)

  //Profile routes
	router.GET("api/profile/:id", profileHandler.GetProfile)
	router.PUT("api/profile/:id", profileHandler.UpdateProfile)

	//Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting at port %s", port)
	router.Run(":" + port)

}
