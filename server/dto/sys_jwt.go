package dto

import "github.com/dgrijalva/jwt-go"

//Claimes载荷,可以再加一些子集需要的信息
//Claims实现了jwt里的Clainms接口
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	RoleID   string `json:"role_id"`
	jwt.StandardClaims
}
