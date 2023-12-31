package util

import (
	"errors"
	"fmt"
	"log"
	"team_todo/global"
	"team_todo/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 设置jwt，jwt鉴权
type jwt_secret struct {
	Signedkey []byte
}

// 从配置中获取密钥
func NewJwtsecret() (*jwt_secret, error) {
	myconfig := global.GVA_CONFIG
	return &jwt_secret{Signedkey: []byte(myconfig.JWT_secret)}, nil
}

type BaseClaims struct {
	Email    string
	Password string
}

type RegisteredClaims struct {
	BaseClaims BaseClaims
	jwt.RegisteredClaims
}

// 创建jwt
func (j *jwt_secret) CreateClaims(baseClaims BaseClaims) (RegisteredClaims, error) {
	return RegisteredClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    baseClaims.Email,                                  // 发行人
			Subject:   "",                                                // 主题
			Audience:  nil,                                               // 用户
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 到期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 在此之前不可用
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 发布时间
			ID:        "",                                                // jwt的id
		},
	}, nil
}

// 生成token
func GenerateToken(req model.User) (string, int64, error) {
	//返回token,过期时间戳，错误
	j, err := NewJwtsecret()
	if err != nil {
		return "", 0, err
	}
	claims, err := j.CreateClaims(BaseClaims{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", 0, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //签名生成token
	tokenStr, err := token.SignedString(j.Signedkey)
	if err != nil {
		log.Printf("生成jwt的token失败，err: [%v]", err)
		return "", 0, err
	}
	expireAt := claims.RegisteredClaims.ExpiresAt
	expireTimestap := expireAt.Unix()
	return tokenStr, expireTimestap, nil

}

// 检查token
func CheckToken(token string) (*RegisteredClaims, error) {
	parse, err := jwt.ParseWithClaims(token, &RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名方式有误: [%v]", token.Header["alg"])
		}
		SigningKey := global.GVA_CONFIG.JWT_secret
		return []byte(SigningKey), nil
	})
	if parse == nil {
		return nil, errors.New("token为空/token有误")
	}
	if parse.Valid {
		if claims, ok := parse.Claims.(*RegisteredClaims); ok {
			return claims, nil
		} else {
			return nil, errors.New("token解析不正确")
		}
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, errors.New("令牌格式不正确")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("令牌已过期")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, errors.New("令牌签名无效")
	} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, errors.New("令牌尚未生效")
	} else {
		return nil, err
	}
}

// 从claims中获取邮箱密码
func Extract_From_Claims(Claims RegisteredClaims) *model.LoginReq {
	var User_Info model.LoginReq
	User_Info.Email = Claims.BaseClaims.Email
	User_Info.Password = Claims.BaseClaims.Password
	return &User_Info
}
