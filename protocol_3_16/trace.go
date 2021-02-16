package protocol

import (
	"fmt"
	"sync"

	"github.com/tliron/glsp"
)

var traceValue TraceValue = TraceValueOff
var traceValueLock sync.Mutex

func GetTraceValue() TraceValue {
	traceValueLock.Lock()
	defer traceValueLock.Unlock()
	return traceValue
}

func SetTraceValue(value TraceValue) {
	traceValueLock.Lock()
	defer traceValueLock.Unlock()

	// The spec clearly says "message", but some implementations use "messages" instead
	if value == "messages" {
		value = TraceValueMessage
	}

	traceValue = value
}

func HasTraceLevel(value TraceValue) bool {
	value_ := GetTraceValue()
	switch value_ {
	case TraceValueOff:
		return false

	case TraceValueMessage:
		return value == TraceValueMessage

	case TraceValueVerbose:
		return true

	default:
		panic(fmt.Sprintf("unsupported trace level: %s", value_))
	}
}

func HasTraceMessageType(type_ MessageType) bool {
	switch type_ {
	case MessageTypeError, MessageTypeWarning, MessageTypeInfo:
		return HasTraceLevel(TraceValueMessage)

	case MessageTypeLog:
		return HasTraceLevel(TraceValueVerbose)

	default:
		panic(fmt.Sprintf("unsupported message type: %d", type_))
	}
}

func Trace(context *glsp.Context, type_ MessageType, message string) error {
	if HasTraceMessageType(type_) {
		go context.Notify(ServerWindowLogMessage, &LogMessageParams{
			Type:    type_,
			Message: message,
		})
	}
	return nil
}
