package main

import "github.com/google/uuid"

func insertDefaults() {
	var subjects []Subject
	var schoolClasses []SchoolClass
	var teachers []Teacher
	var students []Student

	DB.Find(&subjects)
	DB.Find(&schoolClasses)
	DB.Find(&teachers)
	DB.Find(&students)

	if len(subjects) == 0 {
		DB.Create(&Subject{SubjectName: "Українська мова"})
		DB.Create(&Subject{SubjectName: "Українська література"})
		DB.Create(&Subject{SubjectName: "Математика"})
		DB.Create(&Subject{SubjectName: "Історія"})
		DB.Create(&Subject{SubjectName: "Географія"})
		DB.Create(&Subject{SubjectName: "Іноземна мова"})
	}

	if len(schoolClasses) == 0 {
		DB.Create(&SchoolClass{ClassName: "1-А"})
		DB.Create(&SchoolClass{ClassName: "2-А"})
		DB.Create(&SchoolClass{ClassName: "3-А"})
		DB.Create(&SchoolClass{ClassName: "4-А"})
		DB.Create(&SchoolClass{ClassName: "5-А"})
		DB.Create(&SchoolClass{ClassName: "6-А"})
		DB.Create(&SchoolClass{ClassName: "7-А"})
		DB.Create(&SchoolClass{ClassName: "8-А"})
		DB.Create(&SchoolClass{ClassName: "9-А"})
		DB.Create(&SchoolClass{ClassName: "10-А"})
		DB.Create(&SchoolClass{ClassName: "11-А"})
		DB.Create(&SchoolClass{ClassName: "11-Б"})
		DB.Create(&SchoolClass{ClassName: "11-В"})
	}
	if len(teachers) == 0 {
		DB.Create(&Teacher{
			Name: "Девід Гарбор", 
			Photo: "teacher1.jpeg", 
			Login: "teacher1", 
			Password: StrToHash("teacher1"), 
			DateOfBirth: "1988-09-10"})
		DB.Create(&Teacher{
			Name: "Вайнона Райдер", 
			Photo: "teacher2.jpeg", 
			Login: "teacher2", 
			Password: StrToHash("teacher2"), 
			DateOfBirth: "1979-05-11"})			
	}	
	if len(students) == 0 {
		schoolClassID, _ := uuid.Parse("216bf8fc-fca8-4f28-81b3-0b52b9b4c339")
		DB.Create(&Student{
			Name: "Міллі Боббі Браун",
			Photo: "student1.jpeg",
			Login: "student1",
			Password: StrToHash("student1"),
			DateOfBirth: "2004-09-11",
			SchoolClassId: schoolClassID,
		})
	}
}