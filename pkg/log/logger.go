package log

import (
	"fmt"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

func Log() logger {
	return logger{
		Level:            "DEFAULT",
		MaxMessageLentgh: 30,
	}
}

type logger struct {
	Level            string
	MaxMessageLentgh int
	Engine           interface{}
}

func (l *logger) Info(message string, meta interface{}) {
	msg := fmt.Sprintf("[%s]  ", blue("INFO"))
	spaceCharacterNumber := l.MaxMessageLentgh - len(message)
	for i := 0; i < len(message); i++ {
		if i <= l.MaxMessageLentgh {
			msg += string(message[i])
		}
	}
	for i := 0; i < spaceCharacterNumber+4; i++ {
		msg += " "
	}
	if len(message) > l.MaxMessageLentgh {
		msg += "..."
	}
	if meta != nil {
		for key, value := range meta.(map[string]string) {
			msg += fmt.Sprintf(" %s=%s", yellow(key), value)
		}
	}
	fmt.Println(msg)
}

func (l *logger) Debug(message string, meta interface{}) {
	msg := fmt.Sprintf("[%s] ", yellow("DEBUG"))
	spaceCharacterNumber := l.MaxMessageLentgh - len(message)
	for i := 0; i < len(message); i++ {
		if i <= l.MaxMessageLentgh {
			msg += string(message[i])
		}
	}
	for i := 0; i < spaceCharacterNumber+4; i++ {
		msg += " "
	}
	if len(message) > l.MaxMessageLentgh {
		msg += "..."
	}
	if meta != nil {
		for key, value := range meta.(map[string]string) {
			msg += fmt.Sprintf(" %s=%s", yellow(key), value)
		}
	}
	fmt.Println(msg)
}

func (l *logger) Error(message string, meta interface{}) {
	msg := fmt.Sprintf("[%s] ", red("ERROR"))
	spaceCharacterNumber := l.MaxMessageLentgh - len(message)
	for i := 0; i < len(message); i++ {
		if i <= l.MaxMessageLentgh {
			msg += string(message[i])
		}
	}
	for i := 0; i < spaceCharacterNumber+4; i++ {
		msg += " "
	}
	if len(message) > l.MaxMessageLentgh {
		msg += "..."
	}
	if meta != nil {
		for key, value := range meta.(map[string]string) {
			msg += fmt.Sprintf(" %s=%s", yellow(key), value)
		}
	}
	fmt.Println(msg)
}
