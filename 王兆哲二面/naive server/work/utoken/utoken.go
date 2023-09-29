package utoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Utoken struct {
	Uname string
	jwt.StandardClaims
}

// 签名
var sign = "string11"

// GenerateToken() 生成token
// 调用库的NewWithClaims(加密方式,载荷).SignedString(签名) 生成token
func GetnerateToken(name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "server"
	//	 赋值给结构体
	claims := Utoken{
		Uname: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign))
	return toke, err
}

//	ParseToken 解析token
//
// 调用ParseWithClaims进行解析，使用签名解密出数据
func ParseToken(token string) (*Utoken, error) {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Utoken{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(sign), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*Utoken); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
