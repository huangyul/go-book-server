package sms

import "context"

type Service interface {
	Send(ctx context.Context, tpl string, args []NameArg, numbers ...string) error
}

type NameArg struct {
	Val  string
	Name string
}
