package history

import "github.com/mochi-co/mqtt/server/persistence"

type LogMessage struct {
	Header  persistence.FixedHeader
	Payload []byte
}
