package history

import "github.com/cfeeling/mqtt/server/persistence"

type LogMessage struct {
	Header  persistence.FixedHeader
	Payload []byte
}
