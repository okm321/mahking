package domain

import (
	"unicode/utf8"

	pkgerror "github.com/okm321/mahking/go/pkg/error"
)

// requireText 必須の文字列が空でなく、max文字以内であることを検証する
func requireText(label, value string, max int) error {
	if value == "" {
		return pkgerror.NewClientErrorf("%sは必須です", label)
	}
	return optionalText(label, value, max)
}

// optionalText 任意の文字列がmax文字以内であることを検証する
func optionalText(label, value string, max int) error {
	if utf8.RuneCountInString(value) > max {
		return pkgerror.NewClientErrorf("%sは%d文字以内で入力してください", label, max)
	}
	return nil
}
