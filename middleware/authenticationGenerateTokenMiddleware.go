package middleware

import (
	"context"
	"fmt"
	"survey-report-api/databaseConnection"

	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

var ctx = context.Background()

// Role => สิทธิ์การเข้าถึงข้อมูล (นักศึกษาใช้ level_id : 1=>ป.ตรี, 2=>ป.โท, 3=>ป.เอก )
// Role => สิทธิ์การเข้าถึงข้อมูล (พนักงาน level_id : ใช้เลข 4 ขึ้นไป )
func GenerateToken(role string, user string, detail string) (*TokenAuthMiddlewareResponse, error) {
 
	// กำหนด Expiration time และ Universally Unique Identifier
	generateToken := &TokenAuthMiddlewareResponse{}
	generateToken.ExpiresAccessToken = time.Now().Add(time.Minute * 10).Unix()
	generateToken.AccessTokenUUID = uuid.New().String()

	generateToken.ExpiresRefreshToken = time.Now().Add(time.Minute * 30).Unix()
	generateToken.RefreshTokenUUID = uuid.New().String()

	generateToken.Authorized = "true"

	// ---------------------  Create Access Token  ----------------------------------------- //
	// กำหนด claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["issuer"] = "survey"
	accessTokenClaims["subject"] = user + detail
	accessTokenClaims["role"] = role
	accessTokenClaims["access_token_uuid"] = generateToken.AccessTokenUUID
	accessTokenClaims["refresh_token_uuid"] = generateToken.RefreshTokenUUID
	accessTokenClaims["expiration_time"] = generateToken.ExpiresAccessToken

	// map claims(payload) has algorithm (header)
	accessTokenHeaderAndPayload := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	// map claims(payload) และ has algorithm (header) เข้ากับ Signature key
	NEW_ACCESS_TOKEN, err := accessTokenHeaderAndPayload.SignedString([]byte("ใส่ Signature ตรงนี้"))
	if err != nil {
		return nil, err
	}
	// กำหนด token ที่เพิ่งสร้างมาให้กับ struct ที่เราจะ return
	generateToken.AccessToken = NEW_ACCESS_TOKEN

	// ---------------------  Create Refresh Token  ----------------------------------------- //
	// กำหนด claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["issuer"] = "survey"
	refreshTokenClaims["subject"] = user + detail
	refreshTokenClaims["role"] = role
	refreshTokenClaims["access_token_uuid"] = generateToken.AccessTokenUUID
	refreshTokenClaims["refresh_token_uuid"] = generateToken.RefreshTokenUUID
	refreshTokenClaims["expiration_time"] = generateToken.ExpiresRefreshToken

	// map claims(payload) has algorithm (header)
	refreshTokenHeaderAndPayload := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	// map claims(payload) และ has algorithm (header) เข้ากับ Signature key
	NEW_REFRESH_TOKEN, err := refreshTokenHeaderAndPayload.SignedString([]byte("ใส่ Signature ตรงนี้"))
	if err != nil {
		return nil, err
	}
	// กำหนด token ที่เพิ่งสร้างมาให้กับ struct ที่เราจะ return
	generateToken.RefreshToken = NEW_REFRESH_TOKEN

	// ---------------------------  redis cache database  ------------------------------------ //
	// save access token and refresh token in redis cache database
	redisCache := databaseConnection.NewDatabaseConnection().RedisConnection()
	defer redisCache.Close()

	// เริ่มนับเวลา ณ ตอนนี้
	timeNow := time.Now()

	// convertion Unix to UTC(to time object)
	redisCacheExpiresAccessToken := time.Unix(generateToken.ExpiresAccessToken, 0)
	// เก็บ uuid ลง redis cache database โดยใช้ uuid เป็น key และให้ username เป็น value
	err = redisCache.Set(ctx, fmt.Sprint(generateToken.AccessTokenUUID), user+detail, redisCacheExpiresAccessToken.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	// convertion Unix to UTC(to time object)
	redisCacheExpiresRefreshToken := time.Unix(generateToken.ExpiresRefreshToken, 0)
	// เก็บ uuid ลง redis cache database โดยใช้ uuid เป็น key และให้ username เป็น value
	err = redisCache.Set(ctx, fmt.Sprint(generateToken.RefreshTokenUUID), user+detail, redisCacheExpiresRefreshToken.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	// ลบ cache ทั้งหมด
	// redisCache.FlushAll(ctx)

	return generateToken, nil
}
