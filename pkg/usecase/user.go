package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"jakpat-test-2/entity"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *entity.Users `json:"users"`
}

const CtxUserKey = "user"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
)

func (u *Usecase) CreateUser(input entity.Users) (int, error) {
	pwd := sha1.New()
	pwd.Write([]byte(input.Password))
	pwd.Write([]byte(u.hashSalt))

	userID, err := u.services.CreateUser(entity.Users{
		Name:     input.Name,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
		Role:     input.Role,
		Status:   1,
	},
	)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *Usecase) GetUserByIdAndStatus(id int, status int) (entity.Users, error) {
	user, err := u.services.GetUserByIdAndStatus(id, status)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (u *Usecase) GetUserByNameAndPassword(name, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(u.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := u.services.GetUserByNameAndPassword(name, password)
	if err != nil {
		return "", ErrUserNotFound
	}

	claims := AuthClaims{
		User: &user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(u.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(u.signingKey)
}

func (u *Usecase) ParseToken(ctx context.Context, accessToken string) (*entity.Users, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, ErrInvalidAccessToken
}
