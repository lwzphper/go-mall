package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TokenGen 生成
type TokenGen struct {
	issuer  string
	signKey []byte
	nowFunc func() time.Time
}

func NewJwtTokenGen(issuer string, signKey []byte) *TokenGen {
	return &TokenGen{
		issuer:  issuer,
		signKey: signKey,
		nowFunc: time.Now,
	}
}

// GenerateToken 生成 token ， subject 可以为会员 id
func (t *TokenGen) GenerateToken(subject string) (string, error) {
	nowSec := t.nowFunc()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(nowSec),
		ExpiresAt: jwt.NewNumericDate(nowSec),
		Subject:   subject,
	})
	return token.SignedString(t.signKey)
}

// TokenVerify token 校验
type TokenVerify struct {
	signKey []byte
}

func (v *TokenVerify) Verify(token string, options ...jwt.ParserOption) (string, error) {
	/*jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {

	})*/
	return "", nil
}
