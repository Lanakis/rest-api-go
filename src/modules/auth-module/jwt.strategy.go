package auth_module

//import (
//	"authorization/src/modules/user-module/entity"
//	"crypto/hmac"
//	"crypto/sha256"
//	"encoding/base64"
//	"encoding/json"
//	"fmt"
//	"strings"
//	"time"
//)
//
//var secretKey = []byte("secretKey")
//
//func createToken(user entity.UserEntity) (string, error) {
//
//	payload := struct {
//		UserID int   `json:"userId"`
//		Exp    int64 `json:"exp"`
//	}{
//		UserID: user.Id,
//		Exp:    time.Now().Add(time.Hour * 24).Unix(),
//	}
//
//	payloadBytes, err := json.Marshal(payload)
//	if err != nil {
//		return "", err
//	}
//
//	token, err := createJWT(payloadBytes, secretKey)
//	if err != nil {
//		return "", err
//	}
//
//	return token, nil
//}
//
//func createJWT(payload []byte, secretKey []byte) (string, error) {
//	hash := hmac.New(sha256.New, secretKey)
//	hash.Write(payload)
//	signature := hash.Sum(nil)
//	signatureBase64 := base64.RawURLEncoding.EncodeToString(signature)
//	payloadBase64 := base64.RawURLEncoding.EncodeToString(payload)
//	jwtToken := payloadBase64 + "." + signatureBase64
//	return jwtToken, nil
//}
//func ParseToken(tokenString string) (map[string]interface{}, error) {
//	parts := strings.Split(tokenString, ".")
//	if len(parts) != 3 {
//		return nil, fmt.Errorf("invalid token format")
//	}
//	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
//	if err != nil {
//		return nil, fmt.Errorf("failed to decode token payload: %v", err)
//	}
//	var payload map[string]interface{}
//	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal token payload: %v", err)
//	}
//
//	return payload, nil
//}
