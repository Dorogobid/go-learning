package main

import (
	// "fmt"
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
)

func getAllSubjects(c echo.Context) error {
	var subjects []Subject
	result := DB.Find(&subjects)
	if result.Error == nil {
		return c.JSON(http.StatusOK, subjects)
	} else {
		return c.String(http.StatusNoContent, result.Error.Error())
	}
}

func getAllSchoolClasses(c echo.Context) error {
	var schoolClasses []SchoolClass
	result := DB.Find(&schoolClasses)
	if result.Error == nil {
		return c.JSON(http.StatusOK, schoolClasses)
	} else {
		return c.String(http.StatusNoContent, result.Error.Error())
	}
}

func apiLogin(c echo.Context) error {
	var teachers []Teacher
	var students []Student

	l := new(LoginData)
	if err := c.Bind(l); err != nil {
		return err
	}
	if l.UserType == "teacher" {
		DB.Where("login = ? and password = ?", l.UserName, StrToHash(l.Pass)).Find(&teachers)
		if len(teachers) > 0 {
			return c.String(http.StatusOK, "")
		}
		return c.String(http.StatusForbidden, "")
	} else if l.UserType == "student" {
		DB.Where("login = ? and password = ?", l.UserName, StrToHash(l.Pass)).Find(&students)
		if len(students) > 0 {
			return c.String(http.StatusOK, "")
		}
		return c.String(http.StatusForbidden, "")
	}
	return c.String(http.StatusForbidden, "")
}

func getTeachersByLogin(c echo.Context) error {
	var teachers []Teacher
	login := c.Param("login")
	DB.Where("login = ?", login).Find(&teachers)
	return c.JSON(http.StatusOK, teachers)
}

func getSchoolClassesByID(c echo.Context) error {
	var schoolClasses []SchoolClass
	id := c.Param("id")
	DB.Where("lower(id) = ?", id).Find(&schoolClasses)
	return c.JSON(http.StatusOK, schoolClasses)
}

func getMarksByStudentID(c echo.Context) error {
	var marks []MarkFull
	id := c.Param("id")
	DB.Raw("SELECT m.*, s.subject_name as subject_name, t.name as teacher_name, st.name as student_name FROM marks as m, subjects as s, teachers t, students st where s.id = m.subject_id and t.id = m.teacher_id and m.student_id=st.id and lower(student_id) = ?", id).Find(&marks)
	return c.JSON(http.StatusOK, marks)
}

func getMarksByTeacherID(c echo.Context) error {
	var marks []MarkFull
	id := c.Param("id")
	DB.Raw("SELECT m.*, s.subject_name as subject_name, t.name as teacher_name, st.name as student_name FROM marks as m, subjects as s, teachers t, students st where s.id = m.subject_id and t.id = m.teacher_id and m.student_id=st.id and lower(teacher_id) = ?", id).Find(&marks)
	return c.JSON(http.StatusOK, marks)
}

func getStudentsByLogin(c echo.Context) error {
	var students []Student
	login := c.Param("login")
	DB.Where("login = ?", login).Find(&students)
	return c.JSON(http.StatusOK, students)
}

func getAllStudents(c echo.Context) error {
	var students []Student
	DB.Find(&students)
	return c.JSON(http.StatusOK, students)
}

func postMark(c echo.Context) error {
	mark := new(Mark)
	if err := c.Bind(mark); err != nil {
		return err
	}
	DB.Create(&mark)
	return c.JSON(http.StatusOK, mark.ID)
}

func deleteMark(c echo.Context) error {
	mark := new(Mark)
	if err := c.Bind(mark); err != nil {
		return err
	}
	DB.Delete(&mark)
	return c.NoContent(http.StatusNoContent)
}