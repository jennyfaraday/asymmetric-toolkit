package logger

/*
	Logger::PrintEvent() consumes a log event and determines if the log event is to be written (evaluating
	log level) and the destination (writer) if configured.  We then marshall the event struct into JSON.
*/
import (
	"encoding/json"
)

func (o *Logger) PrintEvent(event *LogEventStruct) {
	var out string = "{}"
	if o.PrintThisLine(event.Level.Get()) {
		event.Tags = o.tagMerge(&event.Tags)
		msg, err := json.Marshal(*event)
		if err != nil {
			panic(err)
		}
		out = string(msg)
	}
	if o.Writer != nil {
		o.Writer(&out)
	}
}
