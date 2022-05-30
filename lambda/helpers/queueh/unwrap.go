package queueh

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// nolint: godox
// TODO: find original type.
type SNSWrap struct {
	Type              string            `json:"Type"`
	MessageID         string            `json:"MessageId"`
	TopicARN          string            `json:"TopicArn"`
	Subject           string            `json:"Subject"`
	Message           string            `json:"Message"`
	Timestamp         string            `json:"Timestamp"`
	SignatureVersion  string            `json:"SignatureVersion"`
	Signature         string            `json:"Signature"`
	SigningCertURL    string            `json:"SigningCertURL"`
	UnsubscribeURL    string            `json:"UnsubscribeURL"`
	MessageAttributes MessageAttributes `json:"MessageAttributes"`
}

type MessageAttributes struct {
	EventType MessageAttribute `json:"eventType"`
	Source    MessageAttribute `json:"source"`
}

type MessageAttribute struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

func Unwrap(message string, dest interface{}) error {
	var wrap SNSWrap

	err := json.Unmarshal([]byte(message), &wrap)
	if err != nil {
		return errors.Wrap(err, "can't extract SNS wrap from SQS message")
	}

	err = json.Unmarshal([]byte(wrap.Message), dest)
	if err != nil {
		return errors.Wrap(err, "can't extract original message from SNS notification wrap")
	}

	return nil
}
