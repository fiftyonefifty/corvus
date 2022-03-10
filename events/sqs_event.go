package events

// SQSRecordAttributes is the attributes of a SQS message
type SQSRecordAttributes struct {
	ApproximateReceiveCount          string `json:"approximateReceiveCount"`
	SentTimestamp                    string `json:"sentTimestamp"`
	SenderId                         string `json:"senderId"`
	ApproximateFirstReceiveTimestamp string `json:"approximateFirstReceiveTimestamp"`
}

// SQSMessageAttribute is a single attribute of a SQS message.
type SQSMessageAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// SQSMessage is a single SQS message
type SQSMessage struct {
	Body                   string                         `json:"body"`
	MessageID              string                         `json:"messageId"`
	ReceiptHandle          string                         `json:"receiptHandle"`
	Attributes             SQSRecordAttributes            `json:"attributes"`
	MessageAttributes      map[string]SQSMessageAttribute `json:"messageAttributes"`
	MD5OfBody              string                         `json:"md5OfBody"`
	MD5OfMessageAttributes string                         `json:"md5OfMessageAttributes"`
	EventSource            string                         `json:"eventSource"`
	EventSourceARN         string                         `json:"eventSourceARN"`
	AWSRegion              string                         `json:"awsRegion"`
}

// SQSEvent is the event structure for SQS events
type SQSEvent struct {
	Records []SQSMessage `json:"records"`
}
