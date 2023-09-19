package tencent

import (
	"context"
	"fmt"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	sms2 "go-book-server/internal/service/sms"
)

type Service struct {
	appId    *string
	signName *string
	client   *sms.Client
}

func NewService(client *sms.Client, appId string, signName string) *Service {
	return &Service{
		appId:    &appId,
		client:   client,
		signName: &signName,
	}
}

func (s *Service) Send(ctx context.Context, tplId string, args []sms2.NameArg, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SignName = s.signName
	req.SmsSdkAppId = s.appId
	req.TemplateId = &tplId
	req.PhoneNumberSet = s.toStringPrtSlice(numbers)
	req.TemplateParamSet = s.nameArgToPrtString(args)

	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}

	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败,%s,%s", *status.Code, *status.PhoneNumber)
		}
	}
	return nil
}

func (s *Service) toStringPrtSlice(val []string) []*string {
	pointerRes := make([]*string, len(val))
	for i, str := range val {
		pointerRes[i] = &str
	}
	return pointerRes
}

func (s *Service) nameArgToPrtString(val []sms2.NameArg) []*string {
	pointerRes := make([]*string, len(val))
	for i, cur := range val {
		pointerRes[i] = &cur.Val
	}
	return pointerRes
}
