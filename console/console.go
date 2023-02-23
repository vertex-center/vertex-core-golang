package console

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Console struct {
	serviceName string
}

func New(serviceName string) *Console {
	return &Console{
		serviceName: serviceName,
	}
}

func (c *Console) formattedMessage(message string) string {
	return fmt.Sprintf("[%s] %s\n", c.serviceName, message)
}

func (c *Console) Log(msg string) {
	_, _ = color.New(color.FgHiBlue).Fprintf(os.Stdout, c.formattedMessage(msg))
}

func (c *Console) Warn(msg string) {
	_, _ = color.New(color.FgHiYellow).Fprintf(os.Stdout, c.formattedMessage(msg))
}

func (c *Console) Error(err error) {
	_, _ = color.New(color.FgHiRed).Fprintf(os.Stderr, c.formattedMessage(err.Error()))
}
