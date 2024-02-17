package gout

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/drewstinnett/gout/v2/formats"
	"github.com/drewstinnett/gout/v2/formats/yaml"
)

// Gout is a structure you can use that contains a formatter, and a target
// io.Writer
type Gout struct {
	formatter formats.Formatter
	writer    io.Writer
}

// Use this for doing things without explicitely creating a new gout, similar to
// viper.Viper
var gi *Gout

func init() {
	gi = New()
}

// Get gets the default Gout instance
func Get() *Gout {
	return gi
}

// SetWriter set the io.Writer that will be used for printing. By default, this
// will be os.Stdout
func SetWriter(i io.Writer) { gi.SetWriter(i) }

func (g *Gout) SetWriter(i io.Writer) {
	g.writer = i
}

// SetFormatter sets the Formatter to use for the text.
func SetFormatter(f formats.Formatter) { gi.SetFormatter(f) }

func (g *Gout) SetFormatter(f formats.Formatter) {
	g.formatter = f
}

// Print print an interface using the given Formatter and io.Writer
func Print(v interface{}) (err error) { return gi.Print(v) }

func (g *Gout) Print(v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic while attempting to format: %v", r)
		}
	}()
	var b []byte
	b, err = g.itemizedFormatter(v)
	if err != nil {
		return err
	}
	fmt.Fprint(g.writer, string(b))
	return err
}

// PrintMulti useful when wanting to print multiple interfaces to a single
// serialized item
func PrintMulti(v ...interface{}) (err error) { return gi.PrintMulti(v) }

func (g *Gout) PrintMulti(v ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic while attempting to format: %v", r)
		}
	}()
	var b []byte
	// b, err = c.condensedFormatter(v)
	b, err = g.itemizedFormatter(v)
	if err != nil {
		return err
	}
	fmt.Fprint(g.writer, string(b))
	return err
}

// MustPrint print an interface and panic if there is any sort of error
func MustPrint(v interface{}) { gi.MustPrint(v) }

func (g *Gout) MustPrint(v interface{}) {
	err := g.Print(v)
	if err != nil {
		panic(err)
	}
}

// MustPrintMulti print an multiple interfaces and panic if there is any sort of
// error
func MustPrintMulti(v ...interface{}) { gi.MustPrintMulti(v) }

func (g *Gout) MustPrintMulti(v ...interface{}) {
	err := g.PrintMulti(v)
	if err != nil {
		panic(err)
	}
}

// Option is an option that can be passed in to help configure a Gout instance
type Option func(*Gout)

// WithWriter can be passed to New(), specifying which writer should be used for
// output
func WithWriter(w io.Writer) Option {
	return func(g *Gout) {
		g.writer = w
	}
}

// WithFormatter can be passed to New(), specifying which Formatter should be
// used for output
func WithFormatter(f formats.Formatter) Option {
	return func(g *Gout) {
		g.formatter = f
	}
}

// New creates a pointer to a new Gout, with some sensible defaults
func New(opts ...Option) *Gout {
	g := &Gout{
		formatter: yaml.Formatter{},
		writer:    os.Stdout,
	}

	for _, opt := range opts {
		opt(g)
	}
	return g
}

func (g *Gout) itemizedFormatter(v ...interface{}) ([]byte, error) {
	var buf bytes.Buffer
	for _, item := range v {
		bi, err := g.formatter.Format(item)
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
