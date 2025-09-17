package util

import (
	"blog/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 自定义 Claims
type Claims struct {
	UserID   uint
	Username string
	Password string
	jwt.RegisteredClaims
}

// 生成jwt
func GenJWT(user *model.User) string {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return ""
	}
	//密钥
	secretKey := []byte(config.jwt.Secret)

	claims := &Claims{
		UserID:   user.ID,
		Username: user.UserName,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //表示令牌的过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //表示令牌的生效时间
			Issuer:    "hss",                                              //表示令牌的签发者
			Subject:   "go-blog",                                          //表示令牌的主题
			Audience:  jwt.ClaimStrings{"hss"},                            //表示令牌的受众
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //表示令牌的签发时间
			ID:        fmt.Sprintf("go-api-%d", user.ID),                  //表示令牌的唯一标识符
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return ""
	}
	return token
}

// 解析jwt
func ParseJWT(tokenString string) (*Claims, error) {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return nil, err
	}
	//密钥
	secretKey := []byte(config.jwt.Secret)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
