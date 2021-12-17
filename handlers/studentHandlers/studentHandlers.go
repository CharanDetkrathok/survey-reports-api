package studentHandlers

import (
	"net/http"
	"survey-report-api/errorsHandlers"
	"survey-report-api/services/studentServices"

	"github.com/gin-gonic/gin"
)

// REST API CALL FOLLOW METHOD
func (s studentHandlers) Authentication(c *gin.Context) {

	// ประกาศตัวแปรเพื่อทำการรับข้อมูลที่ส่งมา จะต้องมีตัวแปร (--อย่างน้อย 1 ตัวที่--)ตรงกับ StudentServiceRequest struct{} ต้องเป็น JSON เท่านั้น
	var requestBoby studentServices.StudentAuthenticationServicesRequest

	// ทำการตรวจสอบข้อมูลว่าเป็น JSON formate หรือไม่
	err := c.ShouldBindJSON(&requestBoby)
	if err != nil {
		// (กรณีที่ไม่ได้ส่งมาในรูปแบบ JSON) ทำ ErrorHandler ตรงนี้เพื่อส่ง Message Error and Status Code Error ไปให้ Front-end
		c.IndentedJSON(http.StatusBadRequest, errorsHandlers.NewUnauthorizedError())
		return
	}

	// ส่งไป Query เพื่อค้นหาข้อมูลใน Database
	token, err := s.studentService.Authentication(requestBoby.Std_code, requestBoby.Birth_date, requestBoby.Lev_id)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, token)
		return
	}
	c.IndentedJSON(http.StatusCreated, token)

	// ทำการ Authenticate ตรงนี้โดยเรียกใช้งานที่ Service และไปสร้าง TOKEN ที่ฝั่ง Service (ส่วนของ Bussines Logic)
	// fmt.Println("ข้อมูลที่ส่งมา ==> ",token)

}

func (s studentHandlers) FetchData(c *gin.Context) {

	var requestBoby studentServices.StudentAuthenticationServicesRequest

	err := c.ShouldBindJSON(&requestBoby)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorsHandlers.NewUnauthorizedError())
		return
	}

	switch requestBoby.Lev_id {

	case "1":
		c.IndentedJSON(http.StatusOK, errorsHandlers.NewMessageAndStatusCode(http.StatusOK, "ป.ตรี"))
		return
	case "2":
		c.IndentedJSON(http.StatusOK, errorsHandlers.NewMessageAndStatusCode(http.StatusOK, "ป.โท"))
		return
	case "3":
		c.IndentedJSON(http.StatusOK, errorsHandlers.NewMessageAndStatusCode(http.StatusOK, "ป.เอก") )
		return

	default:
		c.IndentedJSON(http.StatusOK, errorsHandlers.NewMessageAndStatusCode(http.StatusOK, "User นี้ไม่มีสิทธิ์เข้าถึงข้อมูล!"))
		return
	}	
	

}
