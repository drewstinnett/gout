package gout

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/drewstinnett/go-output-format/v2/formats/yaml"
)

type Formatter interface {
	Format(interface{}) ([]byte, error)
}

// Client is a structure you can use that contains a formatter, and a target
// io.Writer
type Client struct {
	// The format!
	Formatter Formatter
	// Target io.Writer output
	Writer io.Writer
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
	var b []byte
	b, err = c.itemizedFormatter(v)
	if err != nil {
		return err
	}
	fmt.Fprint(c.Writer, string(b))
	return err
}

// PrintMulti useful when wanting to print multiple interfaces to a single
// serialized item
func (c *Client) PrintMulti(v ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic while attempting to format: %v", r)
		}
	}()
	var b []byte
	// b, err = c.condensedFormatter(v)
	b, err = c.itemizedFormatter(v)
	if err != nil {
		return err
	}
	fmt.Fprint(c.Writer, string(b))
	return err
}

// MustPrint print an interface and panic if there is any sort of error
func (c *Client) MustPrint(v interface{}) {
	err := c.Print(v)
	if err != nil {
		panic(err)
	}
}

// MustPrintMulti print an multiple interfaces and panic if there is any sort of
// error
func (c *Client) MustPrintMulti(v ...interface{}) {
	err := c.PrintMulti(v)
	if err != nil {
		panic(err)
	}
}

// NewClient creates a pointer to a new writer, with some sensible defaults
func New() (*Client, error) {
	c := &Client{
		Formatter: yaml.Formatter{},
		Writer:    os.Stdout,
	}
	return c, nil
}

func (c *Client) itemizedFormatter(v ...interface{}) ([]byte, error) {
	var buf bytes.Buffer
	for _, item := range v {
		bi, err := c.Formatter.Format(item)
		if err != nil {
			return nil, err
		}
		buf.Write(bi)
	}
	b := buf.Bytes()
	if !bytes.HasSuffix(b, []byte("\n")) {
		b = append(b, "\n"...)
	}
	return b, nil
}
