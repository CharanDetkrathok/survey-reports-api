package main

import (
	"survey-report-api/databaseConnection"
	"survey-report-api/handlers/studentHandlers"
	"survey-report-api/middleware"
	"survey-report-api/repositories/studentRepositories"
	"survey-report-api/services/studentServices"
	"survey-report-api/timeZone"

	"github.com/gin-gonic/gin"
)

func main() {

	timeZone.Init()    

	db, err := databaseConnection.NewDatabaseConnection().OracleConnection() 
	if err != nil {
		panic(err)
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode) // set release mode using env:   export GIN_MODE=release
	router := gin.Default()  // เรียก function พื้นฐานของ gin-gonic
	router.Use(middleware.NewCorsMiddlewrerAccessControll().CorsMiddlewrerAccessControll()) 

	newStudnetRepo := studentRepositories.NewStudentRepositories(db)
	newStudnetService := studentServices.NewStudentServices(newStudnetRepo)
	newStudentHandler := studentHandlers.NewStudentHandlers(newStudnetService)

	// Authentication and Authorization
	studentAuthGeneratorRefreshToken := router.Group("/student/auth")
	{		
		// Genreate token
		studentAuthGeneratorRefreshToken.POST("/authentication", newStudentHandler.Authentication)
		// Refresh token
		studentAuthGeneratorRefreshToken.POST("/refresh-authentication", middleware.RefreshAuthorization)
	}

	router.Use(middleware.Authorization)

	student := router.Group("/student")
	{
		// เอา 1.ข้อคำถาม 2.ข้อคำตอบ-ข้อมูลที่อยู่ หมู่,ตำบล,อำเภอ,จังหวัด ในประเทศ 3.คณะ 4.สาขา หรือหลักสูตร
		student.GET("/fetch-data", newStudentHandler.FetchData)
	}

	router.Run(":8881")

}