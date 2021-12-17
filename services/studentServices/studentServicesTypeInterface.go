package studentServices

import "survey-report-api/repositories/studentRepositories"

type (
	studentInfomation struct {
		Std_code          string `json:"Std_code"`
		Prename_no        string `json:"Prename_no"`
		Prename_thai      string `json:"Prename_thai"`
		Prename_eng       string `json:"Prename_eng"`
		First_name_thai   string `json:"First_name_thai"`
		First_name_eng    string `json:"First_name_eng"`
		Last_name_thai    string `json:"Last_name_thai"`
		Last_name_eng     string `json:"Last_name_eng"`
		Birth_date        string `json:"Birth_date"`
		Faculty_no        string `json:"Faculty_no"`
		Faculty_name_thai string `json:"Faculty_name_thai"`
		Faculty_name_eng  string `json:"Faculty_name_eng"`
		Curr_no           string `json:"Curr_no"`
		Major_no          string `json:"Major_no"`
		Major_flag        string `json:"Major_flag"`
		Major_name_thai   string `json:"Major_name_thai"`
		Major_name_eng    string `json:"Major_name_eng"`
		Lev_id            string `json:"Lev_id"`
	}

	// claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	// -iss (issuer) : เว็บหรือบริษัทเจ้าของ token
	// -sub (subject) : subject ของ token
	// -aud (audience) : ผู้รับ token
	// -exp (expiration time) : เวลาหมดอายุของ token
	// -nbf (not before) : เป็นเวลาที่บอกว่า token จะเริ่มใช้งานได้เมื่อไหร่
	// -iat (issued at) : ใช้เก็บเวลาที่ token นี้เกิดปัญหา
	// -jti (JWT id) : เอาไว้เก็บไอดีของ JWT แต่ละตัวนะครับ
	// -name (Full name) : เอาไว้เก็บชื่อ
	ClaimsToken struct {
		Issuer              string `json:"issuer"`
		Subject             string `json:"subject"`
		Role                string `json:"role"`
		AccessTokenUUID     string `json:"access_token_uuid"`
		RefreshTokenUUID    string `json:"refresh_token_uuid"`
		ExpiresAccessToken  string `json:"expires_access_token"`
		ExpiresRefreshToken string `json:"expiration_refresh_token"`
	}

	// ข้อมูล Token ที่จะ Response ไปให้ front-end
	studentToken struct {
		AccessToken         string `json:"access_token"`
		RefreshToken        string `json:"refresh_token"`
		ExpiresAccessToken  int64  `json:"expires_access_token"`
		ExpiresRefreshToken int64  `json:"expires_refresh_token"`
		AccessTokenUUID     string `json:"access_token_uuid"`
		RefreshTokenUUID    string `json:"refresh_token_uuid"`
		Authorized          string `json:"authorized"`
	}

	StudentAuthenticationServicesResponse struct {
		StudentInfo  studentInfomation
		StudentToken studentToken
		StatusCode   string `json:"status_code"`
		Message      string `json:"message"`
		Role         string `json:"role"`
		RoleDetail   string `json:"role_detail"`
	}

	StudentAuthenticationServicesRequest struct {
		Std_code   string `json:"std_code"`
		Birth_date string `json:"birth_date"`
		Lev_id     string `json:"lev_id"`
	}

	studentServices struct {
		studentRepositories studentRepositories.StudentRepositories
	}

	StudentService interface {
		Authentication(std_code string, birth_date string, lev_id string) (*StudentAuthenticationServicesResponse, error)
	}
)

func NewStudentServices(studentRepositories studentRepositories.StudentRepositories) studentServices {
	return studentServices{studentRepositories: studentRepositories}
}
