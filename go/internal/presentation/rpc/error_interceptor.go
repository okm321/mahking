package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/okm321/mahking/go/pkg/logger"
)

func errorInterceptor() connect.Interceptor {
	return connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				res, err := next(ctx, req)
				if err == nil {
					return res, nil
				}

				connectErr, isClientErr := toConnectError(err)
				if isClientErr {
					logger.WarnContext(ctx, err.Error(), "error", err)
				} else {
					logger.ErrorContext(ctx, err.Error(), "error", err)
				}
				return nil, connectErr
			}
		},
	)
}
