// REST API CALL FOLLOW METHOD
package studentHandlers

import "survey-report-api/services/studentServices"

// call from inside package studentHandler only.
type studentHandlers struct {
	studentService studentServices.StudentService
}

/*
	- return type studentHandler struct{} ที่เป็นข้อมูลของ studentService
	- ในส่วนของการทำ Bussiness logic ไปให้ผู้ที่ Implementation func NewStudentHandler(...) studentService
	- โดย NewStudentHandler( โดยจะต้องส่ง Parameter ที่เป็น studentService เข้ามาเพื่อ Assignment   และใช้ในการกำหนดค่าเริ่มต้นด้วย) studentHandler {...}
*/
func NewStudentHandlers(studentServices studentServices.StudentService) studentHandlers {
	return studentHandlers{studentService: studentServices}
}
