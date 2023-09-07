package jwt

import (
	"fmt"
	"github.com/lwzphper/go-mall/pkg/until"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lwzphper/go-mall/pkg/errors"
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
	ErrToken             = errors.Unauthorized(reason, "JWT token error")
	ErrMissingJwtToken   = errors.Unauthorized(reason, "JWT token is missing")
	ErrExpiredOrNotValid = errors.Unauthorized(reason, "JWT token expire or not valid")
)

// CustomClaims 自定义 claims
type CustomClaims struct {
	jwt.RegisteredClaims
}

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

// GenerateToken 生成 token
func (t *TokenGen) GenerateToken(id uint64, expireSec time.Duration) (string, error) {
	nowSec := t.nowFunc()
	claims := CustomClaims{}
	claims.Issuer = t.issuer
	claims.IssuedAt = jwt.NewNumericDate(nowSec)
	claims.ExpiresAt = jwt.NewNumericDate(nowSec.Add(expireSec))
	claims.Subject = until.Uint64ToString(id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
func (v *TokenValidator) Validator(token string, options ...jwt.ParserOption) (uint64, error) {
	var id uint64
	auths := strings.SplitN(token, " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return id, ErrMissingJwtToken
	}
	jwtToken := auths[1]
	var (
		tokenInfo *jwt.Token
		err       error
	)
	tokenInfo, err = jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.signKey, nil
	}, options...)
	if err != nil {
		// token 无效或过期
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return id, ErrExpiredOrNotValid
		}
		// 其他错误，如：格式有误 jwt.ErrTokenMalformed、签名错误 jwt.ErrTokenSignatureInvalid、其他错误
		return id, ErrToken
	}

	if tokenInfo == nil {
		return id, ErrToken
	}

	if claims, ok := tokenInfo.Claims.(*CustomClaims); ok && tokenInfo.Valid {
		intNum, _ := strconv.Atoi(claims.Subject)
		return uint64(intNum), nil
	}
	return id, ErrToken
}
