package middleware

import (
	"fmt"
	"survey-report-api/databaseConnection"
	"survey-report-api/errorsHandlers"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {

	// แกะ Bearer ออก เอาแค่เฉพาะ token
	token, err := getToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, errorsHandlers.NewUnauthorizedError())
		c.Abort()
		return
	}

	// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
	isToken, err := verifyAccessToken(token)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusUnauthorized, errorsHandlers.NewUnauthorizedError())
		c.Abort()
		return
	}

	if isToken {
		c.Next()
	}

}

func getToken(c *gin.Context) (string, error) {

	const BEARER_SCHEMA = "Bearer "

	AUTH_HEADER := c.GetHeader("Authorization")
	if len(AUTH_HEADER) == 0 {
		return "", errorsHandlers.NewMessageAndStatusCode(http.StatusUnauthorized, "authorization key in header not found")
	}

	if strings.HasPrefix(AUTH_HEADER, BEARER_SCHEMA) {
		tokenString := AUTH_HEADER[len(BEARER_SCHEMA):]
		return tokenString, nil
	} else {
		return "", errorsHandlers.NewMessageAndStatusCode(http.StatusUnauthorized, "Bearer signature key was not found")
	}

}

func verifyAccessToken(token string) (bool, error) {

	rdb := databaseConnection.NewDatabaseConnection().RedisConnection()
	defer rdb.Close()

	claims, err := getClaims(token)
	if err != nil {
		return false, errorsHandlers.NewUnauthorizedError()
	}

	_, err = rdb.Get(ctx, claims.AccessTokenUUID).Result()
	if err != nil {
		return false, errorsHandlers.NewUnauthorizedError()
	}

	return true, nil
}

func getClaims(encodedToken string) (*ClaimsToken, error) {

	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("S!U@R#V$E%Y~S!U@R#V$E%Y~S!U@R#V$E%Y"), nil
	})
	if err != nil {
		return nil, errorsHandlers.NewUnauthorizedError()
	}

	claimsToken := &ClaimsToken{}
	parseClaims := parseToken.Claims.(jwt.MapClaims)

	if parseClaims["issuer"] != nil {
		claimsToken.Issuer = parseClaims["issuer"].(string)
	}

	if parseClaims["subject"] != nil {
		claimsToken.Subject = parseClaims["subject"].(string)
	}

	if parseClaims["role"] != nil {
		claimsToken.Role = parseClaims["role"].(string)
	}

	if parseClaims["access_token_uuid"] != nil {
		claimsToken.AccessTokenUUID = parseClaims["access_token_uuid"].(string)
	}

	if parseClaims["refresh_token_uuid"] != nil {
		claimsToken.RefreshTokenUUID = parseClaims["refresh_token_uuid"].(string)
	}

	return claimsToken, nil
}

func RefreshAuthorization(c *gin.Context) {

	mapRefreshToken := map[string]string{} 
	if err := c.ShouldBindJSON(&mapRefreshToken); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errorsHandlers.NewMessageAndStatusCode(http.StatusUnprocessableEntity, "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"))
		return
	}

	Std_code := mapRefreshToken["std_code"]
	First_name_thai := mapRefreshToken["first_name_thai"]
	First_name_eng := mapRefreshToken["first_name_eng"]
	Lev_id := mapRefreshToken["lev_id"]
	Refresh_token := mapRefreshToken["refresh_token"]

	_, err := verifyRefreshToken(Refresh_token)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errorsHandlers.NewMessageAndStatusCode(http.StatusUnprocessableEntity, "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"))
		return
	}

	claimsDetail, err := getClaims(Refresh_token)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errorsHandlers.NewMessageAndStatusCode(http.StatusUnprocessableEntity, "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"))
		return
	}

	newToken, err := GenerateToken(Lev_id, Std_code, fmt.Sprint(" - "+First_name_thai+" - "+First_name_eng))
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errorsHandlers.NewMessageAndStatusCode(http.StatusUnprocessableEntity, "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"))
		return
	}

	isRevokeTokenFromRedisCache, err := revokeToken(claimsDetail.AccessTokenUUID,claimsDetail.RefreshTokenUUID)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, errorsHandlers.NewMessageAndStatusCode(http.StatusUnprocessableEntity, "กรุณา Sign-in เพื่อเข้าสู่ระบบใหม่"))
		return
	}

	responseNewToken := TokenAuthMiddlewareResponse {
		AccessToken:        newToken.AccessToken,
		RefreshToken:        newToken.RefreshToken,
		ExpiresAccessToken:  0,
		ExpiresRefreshToken: 0,
		AccessTokenUUID:     "",
		RefreshTokenUUID:    "",
		Authorized:          "true",
	}

	if isRevokeTokenFromRedisCache {
		c.IndentedJSON(http.StatusCreated, responseNewToken)
	}	

}

func verifyRefreshToken(refreshToken string) (bool, error) {

	rdb := databaseConnection.NewDatabaseConnection().RedisConnection()
	defer rdb.Close()

	claimsRefresh, err := getClaims(refreshToken)
	if err != nil {
		return false, err
	}

	_, err = rdb.Get(ctx, claimsRefresh.RefreshTokenUUID).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func revokeToken(accessTokenUUID string, refreshTokenUUID string) (bool, error) {

	rdb := databaseConnection.NewDatabaseConnection().RedisConnection()
	defer rdb.Close()

	rdb.Del(ctx,fmt.Sprint(accessTokenUUID), fmt.Sprint(refreshTokenUUID))

	return true, nil
}