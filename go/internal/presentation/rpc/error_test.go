package rpc

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
)

func TestToConnectError(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		wantCode    connect.Code
		wantClient  bool
		wantMessage string
	}{
		{
			name:        "not found",
			err:         pkgerror.Errorf("get group: %w", pkgerror.ErrNotFound),
			wantCode:    connect.CodeNotFound,
			wantClient:  true,
			wantMessage: "存在しないデータです",
		},
		{
			name:        "client error",
			err:         pkgerror.Wrap(pkgerror.NewClientError("名前は必須です"), "create member"),
			wantCode:    connect.CodeInvalidArgument,
			wantClient:  true,
			wantMessage: "名前は必須です",
		},
		{
			name:        "connect error passes through",
			err:         connect.NewError(connect.CodePermissionDenied, errors.New("denied")),
			wantCode:    connect.CodePermissionDenied,
			wantClient:  true,
			wantMessage: "denied",
		},
		{
			name:        "internal error hides message",
			err:         pkgerror.New("db connection refused"),
			wantCode:    connect.CodeInternal,
			wantClient:  false,
			wantMessage: "エラーが発生しました。再試行しても発生し続ける場合ページをリロードしてください",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, isClient := toConnectError(tt.err)
			if got.Code() != tt.wantCode {
				t.Errorf("code = %v, want %v", got.Code(), tt.wantCode)
			}
			if isClient != tt.wantClient {
				t.Errorf("isClient = %v, want %v", isClient, tt.wantClient)
			}
			if got.Message() != tt.wantMessage {
				t.Errorf("message = %q, want %q", got.Message(), tt.wantMessage)
			}
		})
	}
}
