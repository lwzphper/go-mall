package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lwzphper/go-mall/pkg/errors"
	"strings"
	"time"
)

const (
	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken = errors.Unauthorized(reason, "JWT token is missing")
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
func (t *TokenGen) GenerateToken(subject string, expireSec time.Duration) (string, error) {
	nowSec := t.nowFunc()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(nowSec),
		ExpiresAt: jwt.NewNumericDate(nowSec.Add(expireSec)),
		Subject:   subject,
	})
	jwtStr, err := token.SignedString(t.signKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(bearerFormat, jwtStr), nil
}

// TokenValidator token 校验
type TokenValidator struct {
	signKey []byte
}

func NewTokenValidator(signKey []byte) *TokenValidator {
	return &TokenValidator{
		signKey: signKey,
	}
}

// Validator 校验 token
func (v *TokenValidator) Validator(token string, options ...jwt.ParserOption) (string, error) {
	auths := strings.SplitN(token, " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return "", ErrMissingJwtToken
	}
	jwtToken := auths[1]
	var (
		tokenInfo *jwt.Token
		err       error
	)
	tokenInfo, err = jwt.ParseWithClaims(jwtToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.signKey, nil
	}, options...)
	if err != nil {
		return "", err
	}
	return tokenInfo.Claims.GetSubject()
}
