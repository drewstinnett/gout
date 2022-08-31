package writer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/drewstinnett/go-output-format/formats/yaml"
)

type Formatter interface {
	Format(interface{}) ([]byte, error)
}

// Client is a structure you can use that contains a formatter, and a target
// io.Writer
type Client struct {
	Formatter Formatter
	Writer    io.Writer
}

// SetWriter set the io.Writer that will be used for printing. By default, this
// will be os.Stdout
func (c *Client) SetWriter(i io.Writer) {
	c.Writer = i
}

// SetFormatter sets the Formatter to use for the text.
func (c *Client) SetFormatter(f Formatter) {
	c.Formatter = f
}

// Print print an interface using the given Formatter and io.Writer
func (c *Client) Print(v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic while attempting to format: %v", r)
		}
	}()
	b, err := c.Formatter.Format(v)
	if err != nil {
		return err
	}
	s := string(b)
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	// fmt.Fprintf(c.Writer, s)
	fmt.Fprint(c.Writer, s)
	return err
}

// MustPrint print an interface and panic if there is any sort of error
func (c *Client) MustPrint(v interface{}) {
	err := c.Print(v)
	if err != nil {
		panic(err)
	}
}

// NewClient creates a pointer to a new writer, with some sensible defaults
func NewClient() (*Client, error) {
	c := &Client{
		Formatter: yaml.Formatter{},
		Writer:    os.Stdout,
	}
	return c, nil
}
