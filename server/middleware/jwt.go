package middleware

import (
	"admin/dto"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

//JWT 结构
type JWT struct {
	SigningKey []byte
}

//以下根据jwt里面的错误而定制的一些error错误
var (
	//过期的Token
	TokenExpired = errors.New("Token is expired")
	//不是一个Token
	TokenMalformed = errors.New("That's not even a token")
	TokenExpireAt  int64
	//无效的Token
	TokenInvalid = errors.New("Couldn't  handle  this  token")
	//仍然无效的Token
	TokenNotValidYet = errors.New(("Token not  active  yet"))
)

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{}
}

//生成Token,根据传来的Claims来生成Token字符
func (j *JWT) CreateToken(claims dto.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//解析Token ,根据传来的自定义Token字符串来生成
func (j *JWT) ParseToken(tokenString string) (*dto.Claims, error) {
	//生成Token字符串时候需要传入完整dto.Claims ,解析Token字符串传入空dto.Claims即可,返回error接口类型
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) { return j.SigningKey, nil })
	if err != nil {
		//jwt.ValidationError结构体实现了error接口类型，需要断言，ve.Errors是个uint32字段
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				//不是一个Token
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//Token是过期的
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				//无效的Token
				return nil, TokenNotValidYet
			} else {
				//无效的Token
				return nil, TokenInvalid
			}
		}
	}
	//通过token结构体里的Vaild字段可以判断Token是否有效
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid

}

//更新Token
func (j *JWT) RefershToke(tokenString string) (string, error) {
	//var TimeFunc = time.Now
	//创建本地时间
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	//解析旧token
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//解析失败即更新失败
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix() // 默认token有效期为1个小时
		//本质上是通过个更新claims的过期时间字段，更新Claims字段 , 来重新创建token字符串
		return j.CreateToken(*claims)
	}
	//更新失败
	return "", TokenInvalid
}

//生成令牌
func GenrateToken() {}

//从Context里获得Claims
func GetClaims(c *gin.Context) *dto.Claims {
	claims, exist := c.Get("claims")
	if !exist {
		return &dto.Claims{}
	}
	//c.Get返回的是空接口类型，返回时需要空接口.(具体类型)类型断言下，否则编译失败
	return claims.(*dto.Claims)
}
