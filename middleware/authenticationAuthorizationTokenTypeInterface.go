package middleware

type (
	TokenAuthMiddlewareResponse struct {
		AccessToken         string `json:"access_token"`
		RefreshToken        string `json:"refresh_token"`
		ExpiresAccessToken  int64  `json:"expires_access_token"`
		ExpiresRefreshToken int64  `json:"expires_refresh_token"`
		AccessTokenUUID     string `json:"access_token_uuid"`
		RefreshTokenUUID    string `json:"refresh_token_uuid"`
		Authorized          string `json:"authorized"`
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
)