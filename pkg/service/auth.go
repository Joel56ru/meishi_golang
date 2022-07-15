package service

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"meishi_golang/pkg/repository"
	"meishi_golang/senti"
	"os"
)

//AuthService структура сервиса авторизации
type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	Det senti.UserJWT
}

//MyAuthService принимает репозиторий для работы с базой
func MyAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

//ParseToken парсинг токена
func (s *AuthService) ParseToken(accessToken, ip string) (senti.UserJWT, error) {
	var claimses senti.UserJWT
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return claimses, errors.New("invalid signing method")
	}
	if err != nil {
		return claimses, errors.New("token not valid")
	}
	if !token.Valid {
		return claimses, errors.New("token not valid")
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return claimses, errors.New("token claims are not of type")
	}
	if claims.Det.SecureIp == true && claims.Det.IP != generateIPHash(ip) {
		return claimses, errors.New("ip adress was changed")
	}
	claimses = claims.Det
	return claimses, nil
}

//generatePasswordHash генерация хэша ip
func generateIPHash(ip string) string {
	hash := sha1.New()
	hash.Write([]byte(ip + os.Getenv("SOLT_IP")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//CreateUser имплементируем CreateUser для передачи юзера в репозиторий
func (s *AuthService) CreateUser(user senti.UserRegister) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

//generatePasswordHash генерация хэша пароля
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + os.Getenv("SOLT_PASSWORD")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
