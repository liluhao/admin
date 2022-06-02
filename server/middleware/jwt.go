package middleware

import (
	"admin/common"
	"admin/dto"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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
	return &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
}

// CreateToken 生成Token,根据传来的Claims来生成Token字符
func (j *JWT) CreateToken(claims dto.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token ,根据传来的Token字符串来解析验证
func (j *JWT) ParseToken(tokenString string) (*dto.Claims, error) {
	//生成Token字符串时候需要传入完整dto.Claims ,解析Token字符串传入空dto.Claims即可,返回error接口类型
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) { return j.SigningKey, nil })
	//fmt.Println(tokenString)
	/*eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInVzZXJuYW1lIjoiZ3Vlc3QiLCJyb2xlX2lkIjoiNTYiLCJleHAiOjE2NTQwODE0NTAsImlzcyI6ImFkbWluIiwibmJmIjoxNjU0MDczMjUwfQ.nCRZqHL4Gdk
	H2yMsxnpYgCXM2w8A0mzynLCpFIKtWvY*/
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
func GenerateToken(user *dto.UserInfoOut) (token string, msg string, ok bool) {
	j := &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
	res, err := strconv.ParseInt(common.CONFIG.Jwt.Expire, 10, 64)
	if err != nil {
		common.LOG.Error("jwt expire config info err")
	}
	TokenExpireAt = res
	claims := dto.Claims{
		user.ID,
		user.Username,
		user.RoleId,

		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),          // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + TokenExpireAt), // 过期时间 一小时
			Issuer:    common.CONFIG.Jwt.SigningKey,             //签名的发行者
		},
	}
	token, err = j.CreateToken(claims)
	if err != nil {
		common.LOG.Error("创建Token失败", zap.Any("err", err))
		return token, "创建token失败", false
	} else {
		bearer := "Bearer "
		token = fmt.Sprintf("%s%s", bearer, token)
		return token, "登录成功", true
	}
}

//从Context里获得Claims
func GetClaims(c *gin.Context) *dto.Claims {
	claims, exist := c.Get("claims")
	if !exist {
		return &dto.Claims{}
	}
	//c.Get返回的是空接口类型，返回时需要空接口.(具体类型)类型断言下，否则编译失败
	return claims.(*dto.Claims)
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		//fmt.Println(token)//同一用户登录每次生成的Token字符串都不一样
		/*Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInVzZXJuYW1lIjoiZ3Vlc3QiLCJyb2xlX2lkIjoiNTYiLCJleHAiOjE2NTM3Mzg0ODcsImlzcyI6ImFkbWluIiwibmJmIjoxNjUzNzMwMjg3fQ.K4zk
		1z73XHNlJH4HBEf6kK1W4CUaUyFg-Q3ahvymDUs*/
		/*Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInVzZXJuYW1lIjoiZ3Vlc3QiLCJyb2xlX2lkIjoiNTYiLCJleHAiOjE2NTQwODExMDYsImlzcyI6ImFkbWluIiwibmJmIjoxNjU0MDcyOTA2fQ.2izS
		UUISLhlPsnQJoiLvXpsBo_bQ14O6tbYMiIVK7Ss*/
		if token == "" {
			ResponseFail(c, 250, "请求未携带token，无权限访问")
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token[len("Bearer "):])
		if err != nil {

			if err == TokenExpired {
				ResponseFail(c, 251, "token 已经过期，请重新登录！")
				c.Abort()
				return
			}
			ResponseFail(c, 252, err.Error())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}
