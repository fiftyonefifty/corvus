// Types

package corvus

type SQSMessage struct {
	MessageID     string            `json:"messageId"`
	ReceiptHandle string            `json:"receiptHandle"`
	Body          string            `json:"body"`
	Attributes    map[string]string `json:"attributes"`
}

type SQSEvent struct {
	Records []SQSMessage `json:"records"`
}
