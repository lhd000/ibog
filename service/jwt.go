package service

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

const (
    ErrorServerBusy = "server is busy"
    ErrorReLogin = "relogin"
)

type JWTClaims struct {
    jwt.StandardClaims
    ID int `json:"id"`
    Pass string `json:"pass"`
    Name string `json:"name"`
}


var (
    Secret = "123#111"  //salt
    ExpireTime = 3600  //token expire time
)

//generate jwt token
func genToken(claims *JWTClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(Secret))
    if err != nil {
        return "", errors.New(ErrorServerBusy)
    }
    return signedToken, nil
}

func GenerateToken(id int,name string ,pass string)(string,error){
	claims := &JWTClaims{
        ID: id,
        Name: name,
        Pass: pass,
	}

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	singedToken, err := genToken(claims)

	return singedToken,err
}

func ParseToken(strToken string) (*JWTClaims,error){
  
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(Secret), nil
    })
    if err != nil {
        return nil, errors.New(ErrorServerBusy)
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok {
        return nil, errors.New(ErrorReLogin)
    }
    if err := token.Claims.Valid(); err != nil {
        return nil, errors.New(ErrorReLogin)
    }

   // fmt.Println("verify")
    return claims, nil
}










