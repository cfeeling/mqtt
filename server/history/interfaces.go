package history

type LogWriter interface {
	Write(v LogMessage)
}
