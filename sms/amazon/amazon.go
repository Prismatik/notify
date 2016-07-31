package amazonSms

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/prismatik/notify/types"
)

func Send(m types.SMS) error {
	fmt.Println("send called")

	svc := sns.New(
		session.New(),
		&aws.Config{
			Region: aws.String("us-west-2"),
		},
	)

	params := &sns.PublishInput{
		Message: aws.String(m.Body),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(m.From),
			},
		},
		PhoneNumber: aws.String(m.To),
	}

	_, err := svc.Publish(params)
	return err
}
