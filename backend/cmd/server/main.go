package main

import{
	"log"
	"net/http"
	"os"

  "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Sabari-Vijayan/DBMS-project/internal/db"
}

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

	//CREATING Gin router
	router:=gin.Default()

  //TESTING THE ENDPOINT
	router.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"status":"ok",
			"message":"Server is running!",
		})
	})

	//GET PORT FROM environment ot default to 8080
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}

	log.Printf("Server starting on port %s",port)
	router.Run(";"+port)

}
