package interceptors

import (
	"context"
	"fmt"

	"github.com/Diaku49/grpc-game-server/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var ProtectedMethods = []string{
	"",
}

type AuthApplyed func(method string) bool

func RegisterAuthInterceptor(secret string) grpc.UnaryServerInterceptor {
	auth := MakeAuthApplyer()
	return AuthInterceptor(auth, secret)
}

func AuthInterceptor(auth AuthApplyed, secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !auth(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "Metadata missing")
		}

		userId, err := AuthMiddleware(md, secret)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		nCtx := context.WithValue(ctx, "user_id", userId)

		return handler(nCtx, req)
	}
}

func MakeAuthApplyer() AuthApplyed {
	temp := make(map[string]struct{}, len(ProtectedMethods))
	for _, m := range ProtectedMethods {
		temp[m] = struct{}{}
	}

	return func(method string) bool {
		_, ok := temp[method]
		return ok
	}
}

func AuthMiddleware(m metadata.MD, secret string) (string, error) {
	authorization := m.Get("authorization")
	if len(authorization) == 0 {
		return "", fmt.Errorf("authorization missing")
	}
	token := authorization[0]

	claim, err := pkg.ValidateToken(token, secret)
	if err != nil {
		return "", fmt.Errorf("Invalid token err: %v", err)
	}

	return claim.UserId, nil
}
