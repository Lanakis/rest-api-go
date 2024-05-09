package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin-gorm/modules/auth/dto"
	"gin-gorm/modules/user/entity"
	"gin-gorm/modules/user/models"
	"gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const (
	secretKey = "secretKey"
)

type User struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
}
type RefreshToken struct {
	UserID    int
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type Service struct {
	userService entity.IUserService
}

func NewAuthService(userService entity.IUserService) *Service {
	return &Service{
		userService: userService,
	}
}

func (s *Service) GenerateToken(ctx context.Context, entity models.User) (string, error) {
	// Создаем структуру для хранения данных, которые будут включены в токен.
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := User{
		Username:  entity.Username,
		Role:      entity.Role,
		ExpiresAt: expirationTime,
	}

	// Кодируем структуру в JSON.
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Создаем заголовок токена.
	header := `{"alg":"HS256","typ":"JWT"}`

	// Кодируем заголовок и данные в base64url.
	headerBase64 := encodeBase64URL([]byte(header))
	claimsBase64 := encodeBase64URL(claimsJSON)

	// Создаем строку для подписи.
	signatureInput := fmt.Sprintf("%s.%s", headerBase64, claimsBase64)

	// Создаем подпись токена с использованием HMAC-SHA256 и секретного ключа.
	signature := createHMAC(signatureInput)

	// Собираем токен из заголовка, данных и подписи.
	token := fmt.Sprintf("%s.%s.%s", headerBase64, claimsBase64, signature)

	return token, nil
}
func (s *Service) Check(ctx context.Context, username string) (models.User, error) {

	el, err := s.userService.FindByUsername(ctx, username)

	return el, err
}
func (s *Service) SignIn(ctx context.Context, dto dto.SignAuthDto) (string, error) {
	el, err := s.userService.FindByUsername(ctx, dto.Username)
	if err != nil {
		return "", err // Возвращаем ошибку, если произошла ошибка при поиске пользователя
	}

	if el.Id == 0 {
		return "", errors.New("user not found") // Возвращаем ошибку, если пользователь не найден
	}
	hashEnteredPassword := utils.HashPassword(dto.Password)

	if hashEnteredPassword != el.Password {
		return "", errors.New("invalid password") // Возвращаем ошибку, если пароль неверен
	}

	return s.Login(ctx, el)
}
func (s *Service) Login(ctx context.Context, entity models.User) (string, error) {
	generateToken, err := s.GenerateToken(ctx, entity)
	return generateToken, err
}

// VerifyToken проверяет подлинность и декодирует JWT токен.
func (s *Service) VerifyToken(tokenString string) (*User, error) {
	// Разбиваем токен на заголовок, данные и подпись.
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	// Декодируем данные из base64url.
	claimsJSON, err := decodeBase64URL(parts[1])
	if err != nil {
		return nil, err
	}

	// Декодируем JSON в структуру User.
	var user User
	if err := json.Unmarshal(claimsJSON, &user); err != nil {
		return nil, err
	}

	// Проверяем подпись токена.
	signatureInput := fmt.Sprintf("%s.%s", parts[0], parts[1])
	signature := createHMAC(signatureInput)
	expectedSignature := parts[2]
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return nil, errors.New("invalid token signature")
	}

	return &user, nil
}

// Middleware проверяет JWT токен в заголовке запроса и добавляет информацию о пользователе в контекст запроса.
func (s *Service) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := s.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// encodeBase64URL кодирует байтовый срез в base64url.
func encodeBase64URL(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

// decodeBase64URL декодирует base64url обратно в байтовый срез.
func decodeBase64URL(data string) ([]byte, error) {
	// Добавляем недостающие символы =, чтобы декодирование работало правильно.
	padding := strings.Repeat("=", (4-len(data)%4)%4)
	data += padding
	return base64.URLEncoding.DecodeString(data)
}

// createHMAC создает подпись токена с использованием HMAC-SHA256 и секретного ключа.
func createHMAC(data string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return encodeBase64URL(h.Sum(nil))
}

func (s *Service) SaveRefreshToken(ctx context.Context, userID int, refreshToken string, expiresAt time.Time) error {
	// Здесь вы создаете новую запись RefreshToken в базе данных
	// Обычно это включает в себя выполнение SQL-запроса INSERT
	return nil
}

func (s *Service) FindRefreshToken(ctx context.Context, refreshToken string) (*RefreshToken, error) {
	// Здесь вы выполняете запрос к базе данных для поиска токена по его значению
	// Обычно это включает в себя выполнение SQL-запроса SELECT
	return nil, nil
}
