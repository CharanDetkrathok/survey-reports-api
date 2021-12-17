// Bussiness Logic
package studentServices

import (
	"fmt"
	"net/http"
	"survey-report-api/middleware"
	"survey-report-api/repositories/studentRepositories"
)

func (s studentServices) Authentication(std_code string, birth_date string, lev_id string) (*StudentAuthenticationServicesResponse, error) {

	var student *studentRepositories.StudentAuthenticationRepositoriesFromDB
	var err error
	var roleDetail string

	switch lev_id {
	case "1":

		// ส่ง std_code และ birth_date ไปตรวจสอบใน Database
		student, err = s.studentRepositories.AuthenticateBachelor(std_code, birth_date, lev_id)
		if err != nil {

			// ถ้าไม่พบข้อมูล ทำการเตรียมโครงสร้างข้อมูลแบบว่างเปล่ากลับไป
			studentResponseInfomation := studentInfomation{}
			studentResponse := StudentAuthenticationServicesResponse{
				studentResponseInfomation,
				studentToken{AccessToken: "", RefreshToken: "", Authorized: ""},
				fmt.Sprint(http.StatusNoContent),
				MakeMessage(lev_id),
				"",
				"",
			}
			return &studentResponse, err

		}

		roleDetail = fmt.Sprint(lev_id, ": บัณฑิต ปริญญาตรี")

	case "2":
		// ส่ง std_code และ birth_date ไปตรวจสอบใน Database
		student, err = s.studentRepositories.AuthenticateMaster(std_code, birth_date, lev_id)
		if err != nil {

			// ถ้าไม่พบข้อมูล ทำการเตรียมโครงสร้างข้อมูลแบบว่างเปล่ากลับไป
			studentResponseInfomation := studentInfomation{}
			studentResponse := StudentAuthenticationServicesResponse{
				studentResponseInfomation,
				studentToken{AccessToken: "", RefreshToken: "", Authorized: ""},
				fmt.Sprint(http.StatusNoContent),
				MakeMessage(lev_id),
				"",
				"",
			}
			return &studentResponse, err

		}

		roleDetail = fmt.Sprint(lev_id, ": บัณฑิตศึกษา ปริญญาโท")

	case "3":
		// ส่ง std_code และ birth_date ไปตรวจสอบใน Database
		student, err = s.studentRepositories.AuthenticatePhd(std_code, birth_date, lev_id)
		if err != nil {

			// ถ้าไม่พบข้อมูล ทำการเตรียมโครงสร้างข้อมูลแบบว่างเปล่ากลับไป
			studentResponseInfomation := studentInfomation{}
			studentResponse := StudentAuthenticationServicesResponse{
				studentResponseInfomation,
				studentToken{AccessToken: "", RefreshToken: "", Authorized: ""},
				fmt.Sprint(http.StatusNoContent),
				MakeMessage(lev_id),
				"",
				"",
			}
			return &studentResponse, err

		}

		roleDetail = fmt.Sprint(lev_id, ": บัณฑิตศึกษา ปริญญาเอก")

	default:
		roleDetail = "unknown"
	}

	generateToken, err := middleware.GenerateToken(lev_id, student.Std_code, fmt.Sprint(" - "+student.First_name_thai+" - "+student.First_name_eng))
	if err != nil {
		return nil, err
	}

	studentResponseGenerateToken := studentToken{
		AccessToken:         generateToken.AccessToken,
		RefreshToken:        generateToken.RefreshToken,
		ExpiresAccessToken:  generateToken.ExpiresAccessToken,
		ExpiresRefreshToken: generateToken.ExpiresRefreshToken,
		AccessTokenUUID:     generateToken.AccessTokenUUID,
		RefreshTokenUUID:    generateToken.RefreshTokenUUID,
		Authorized:          generateToken.Authorized,
	}

	// เก็บข้อมูลที่ได้จากการ Query เพื่อเตรียม Respose
	studentResponseInfomation := studentInfomation{
		Std_code:          student.Std_code,
		Prename_no:        student.Prename_no,
		Prename_thai:      student.Prename_thai,
		Prename_eng:       student.Prename_eng,
		First_name_thai:   student.First_name_thai,
		First_name_eng:    student.First_name_eng,
		Last_name_thai:    student.Last_name_thai,
		Last_name_eng:     student.Last_name_eng,
		Birth_date:        student.Birth_date,
		Faculty_no:        student.Faculty_no,
		Faculty_name_thai: student.Faculty_name_thai,
		Faculty_name_eng:  student.Faculty_name_eng,
		Curr_no:           student.Curr_no,
		Major_no:          student.Major_no,
		Major_flag:        student.Major_flag,
		Major_name_thai:   student.Major_name_thai,
		Major_name_eng:    student.Major_name_eng,
		Lev_id:            lev_id,
	}

	// เตรียม Infomation และ Token สำหรับ Authorization ของนักศึกษา
	studentResponse := StudentAuthenticationServicesResponse{
		studentResponseInfomation,
		studentToken{AccessToken: studentResponseGenerateToken.AccessToken, RefreshToken: studentResponseGenerateToken.RefreshToken, Authorized: studentResponseGenerateToken.Authorized},
		fmt.Sprint(http.StatusCreated),
		"Created tokens successfully",
		lev_id,
		roleDetail,
	}

	return &studentResponse, nil
}

func MakeMessage(lev_id string) string {

	var MessageDetail string
	switch lev_id {
	case "1", "2", "3":
		MessageDetail = "รหัสนักษึกษาหรือ วัน/เดือน/ปีเกิด ไม่ถูกต้อง!"
	// case "2","3":
	// 	MessageDetail = "รหัสนักษึกษาหรือ วัน/เดือน/ปีเกิด ไม่ถูกต้อง! กรณีมหาบัณฑิตระดับปริญญาโท และปริญญาเอก กรุณาติดต่อฝ่ายทะเบียนบัณฑิตวิทยาลัยเพื่อ Update ข้อมูลของท่าน เนื่องจากไม่พบรหัสนักษึกษาหรือ วัน/เดือน/ปีเกิด หรือไม่พบข้อมูลคณะหรือหลักสูตรของท่าน"
	default:
		MessageDetail = "unknown"
	}

	return MessageDetail
}
