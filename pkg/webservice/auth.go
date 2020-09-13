package webservice

import (
	"github.com/gin-gonic/gin"
)

type Auth interface {
	GetAuth() (interface{}, error)

	Validate() error

	OnError(c *gin.Context, err error)

	WriteAuth() (c *gin.Context)
}

//type JWTService struct {
//	Token string
//}
//
//func NewJWTService(token string) {
//
//}
//type authCustomClaims struct {
//	Name string `json:"name"`
//	User bool   `json:"user"`
//	jwt.StandardClaims
//}
//
//func (auth *JWTService) GenerateToken(email string, isUser bool) (string, error) {
//	claims := &authCustomClaims{
//		email,
//		isUser,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
//			Issuer:    service.issure,
//			IssuedAt:  time.Now().Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	//encoded string
//	t, err := token.SignedString([]byte(service.secretKey))
//	if err != nil {
//		panic(err)
//	}
//	return t
//}
//
//func (auth *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
//	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
//		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
//			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
//
//		}
//		return []byte(service.secretKey), nil
//	})
//}
