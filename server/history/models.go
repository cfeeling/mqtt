package history

import "github.com/cfeeling/mqtt/server/persistence"

type LogMessage struct {
	Topic   string
	Header  persistence.FixedHeader
	Payload []byte
}
