package corvus

import (
	"encoding/json"
	"testing"
)

func TestSQSEvent(t *testing.T) {
	var event SQSEvent
	data := `{
		"Records": [
			{
				"messageId": "16d1526c-3a53-4cb0-9a69-f8039f2c1b06",
				"receiptHandle": "MessageReceiptHandle",
				"body": "Hello from SQS!",
				"attributes": {
					"ApproximateReceiveCount": "1",
					"SentTimestamp": "1523232000000",
					"SenderId": "123456789012",
					"ApproximateFirstReceiveTimestamp": "1523232000001"
				},
				"messageAttributes": {
					"Name": {
						"stringValue": "MyName",
						"dataType": "String"
					},
					"Age": {
						"stringValue": "31",
						"dataType": "Number.Integer"
					}
				},
				"md5OfBody": "7b270e59b47ff90a553787216d55d91d",
				"eventSource": "aws:sqs",
				"eventSourceARN": "arn:aws:sqs:us-east-1:123456789012:MyQueue",
				"awsRegion": "us-east-1"
			}
		]
	}`

	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Errorf("Error unmarshalling SQS event: %s", err)
	}

	if event.Records[0].MessageID != "16d1526c-3a53-4cb0-9a69-f8039f2c1b06" {
		t.Errorf("Expected MessageID to be 16d1526c-3a53-4cb0-9a69-f8039f2c1b06, but was %s", event.Records[0].MessageID)
	}

	if event.Records[0].ReceiptHandle != "MessageReceiptHandle" {
		t.Errorf("Expected ReceiptHandle to be MessageReceiptHandle, but was %s", event.Records[0].ReceiptHandle)
	}

}
