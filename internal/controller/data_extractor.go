package controller

import (
	"errors"
	"io"

	"github.com/xuri/excelize/v2"
)

type Extractor interface {
	Next() bool
	Columns() ([]string, error)
	Close() error
}

type excelExtractor struct {
	*excelize.Rows
}

// Columns implements Extractor.
// Subtle: this method shadows the method (*Rows).Columns of excelExtractor.Rows.
func (e excelExtractor) Columns() ([]string, error) {
	return e.Rows.Columns()
}

type csvExtractor struct{}

// Close implements Extractor.
func (c csvExtractor) Close() error {
	panic("unimplemented")
}

// Columns implements Extractor.
func (c csvExtractor) Columns() ([]string, error) {
	panic("unimplemented")
}

// Next implements Extractor.
func (c csvExtractor) Next() bool {
	panic("unimplemented")
}

func newExtractor(fileType string, r io.Reader) (Extractor, error) {
	switch fileType {
	case "excel":
		r, err := excelize.OpenReader(r)
		if err != nil {
			return nil, err
		}
		name := r.GetSheetName(0)
		rows, err := r.Rows(name)
		if err != nil {
			return nil, err
		}
		return excelExtractor{
			rows,
		}, nil
	case "csv":
		return csvExtractor{}, nil
	default:
		return nil, errors.New("not implement")
	}
}
