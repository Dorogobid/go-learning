package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	// "github.com/satori/go.uuid"
)

type Subject struct {
	ID uuid.UUID		`gorm:"primaryKey;type:uuid" json:"id"`
	SubjectName string	`json:"subjectName"`
}

type SchoolClass struct {
	ID uuid.UUID		`gorm:"primaryKey;type:uuid" json:"id"`
	ClassName string	`json:"className"`
	Students []Student
}

type Teacher struct {
	ID uuid.UUID		`gorm:"primaryKey;type:uuid" json:"id"`
	Name string			`json:"name"`
	Photo string		`json:"photo"`
	Login string		`json:"login"`
	Password string		`json:"password"`
	DateOfBirth string	`json:"dateOfBirth"`
}

type Student struct {
	ID uuid.UUID		`gorm:"primaryKey;type:uuid" json:"id"`
	Name string			`json:"name"`
	Photo string		`json:"photo"`
	Login string		`json:"login"`
	Password string		`json:"password"`
	DateOfBirth string	`json:"dateOfBirth"`
	SchoolClassId uuid.UUID 	`json:"schoolClassId"`
}

type Mark struct {
	ID uuid.UUID		`gorm:"primaryKey;type:uuid" json:"id"`
	MarkDate string		`json:"markDate"`
	Mark uint			`json:"mark"`
	TeacherID uuid.UUID	`json:"teacherId"`
	StudentID uuid.UUID	`json:"studentId"`
	SubjectID uuid.UUID	`json:"subjectId"`
}

type MarkFull struct {
	ID uuid.UUID		`gorm:"type:uuid" json:"id"`
	MarkDate string		`json:"markDate"`
	Mark uint			`json:"mark"`
	TeacherID uuid.UUID	`json:"teacherId"`
	StudentID uuid.UUID	`json:"studentId"`
	SubjectID uuid.UUID	`json:"subjectId"`
	SubjectName string	`json:"subjectName"`
	TeacherName string	`json:"teacherName"`	
	StudentName string	`json:"studentName"`
}

type LoginData struct {
	UserType string		`json:"userType" param:"userType" query:"userType" form:"userType" `
	UserName string		`json:"userName" param:"userName" query:"userName" form:"userName" `
	Pass string			`json:"pass" param:"pass" query:"pass" form:"pass" `
	Submit string		`json:"submit" param:"submit" query:"submit" form:"submit" `
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (s *SchoolClass) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (s *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (s *Mark) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}