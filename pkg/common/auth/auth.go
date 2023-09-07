package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lwzphper/go-mall/pkg/common/id"
	"github.com/lwzphper/go-mall/pkg/until"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	// ImpersonateAccountHeader defines the header for account
	// id impersonation.
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer "
)

// Interceptor 创建一个GRPC认证拦截器
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	//f, err := os.Open(publicKeyFile)
	//if err != nil {
	//	return nil, fmt.Errorf("cannnot open public key file: %v", err)
	//}

	//b, err := io.ReadAll(f)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot read public key: %v", err)
	//}
	//
	//pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	//if err != nil {
	//	return nil, fmt.Errorf("canot parse public key: %v", err)
	//}
	i := &interceptor{
		//verifier: &token.JWTTokenVerifier{
		//	PublicKey: pubKey,
		//},
	}
	return i.HandleReq, nil
}

type tokenVerifier interface {
	Verify(token string, options ...jwt.ParserOption) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	aid := impersonationFromContext(ctx)
	if aid != "" {
		return handler(ContextWithMemberID(ctx, id.MemberID(until.StringToUint64(aid))), req)
	}

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	aid, err = i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	return handler(ContextWithMemberID(ctx, id.MemberID(until.StringToUint64(aid))), req)
}

func impersonationFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}

// 从上下文中，获取token （token 不需要给外部知道，因此小写）
func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

type memberIDKey struct{}

// ContextWithMemberID 使用给定的帐户ID创建上下文
func ContextWithMemberID(c context.Context, aid id.MemberID) context.Context {
	return context.WithValue(c, memberIDKey{}, aid)
}

// MemberIDFromContext gets member id from context.
// Returns unauthenticated error if no member id is available.
func MemberIDFromContext(c context.Context) (id.MemberID, error) {
	v := c.Value(memberIDKey{})
	aid, ok := v.(id.MemberID)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
