package rpc

import (
	"errors"

	"connectrpc.com/connect"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
)

// toConnectError はエラーをConnectのエラーに変換する。
// 戻り値のisClientErrはユーザー起因かどうか（ログレベルの判定に使う）。
func toConnectError(err error) (connectErr *connect.Error, isClientErr bool) {
	if errors.As(err, &connectErr) {
		return connectErr, connectErr.Code() != connect.CodeInternal
	}

	if notFoundErr, ok := errors.AsType[*pkgerror.ErrorNotFound](err); ok {
		return connect.NewError(connect.CodeNotFound, notFoundErr), true
	}

	if clientErr, ok := errors.AsType[*pkgerror.ClientError](err); ok {
		return connect.NewError(connect.CodeInvalidArgument, clientErr), true
	}

	return connect.NewError(
		connect.CodeInternal,
		errors.New("エラーが発生しました。再試行しても発生し続ける場合ページをリロードしてください"),
	), false
}
