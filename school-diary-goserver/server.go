package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB


func main() {
	DB = connectToDb()
	e := echo.New()

	insertDefaults()

	e.Static("/img", "static/img")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "host=${host},method=${method}, uri=${uri}, status=${status}\n"}))

	api := e.Group("/api")
	
	api.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ98Ng09a8", nil
	}))

	api.GET("/subjects", getAllSubjects)
	api.GET("/school_classes", getAllSchoolClasses)
	api.GET("/school_classes/:id", getSchoolClassesByID)
	api.POST("/login", apiLogin)
	api.GET("/teachers/:login", getTeachersByLogin)
	api.GET("/marks/1/:id", getMarksByStudentID)
	api.GET("/marks/2/:id", getMarksByTeacherID)
	api.POST("/marks", postMark)
	api.DELETE("/marks", deleteMark)
	api.GET("/students", getAllStudents)
	api.GET("/students/:login", getStudentsByLogin)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectToDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}
	db.AutoMigrate(&Subject{}, &SchoolClass{}, &Teacher{}, &Student{}, &Mark{})
	return db
}