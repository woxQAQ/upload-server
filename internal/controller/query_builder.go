package controller

import (
	"bytes"
	"strings"
	"text/template"
)

type queryBulder struct {
	sb        strings.Builder
	db, table string
	templ     string
	prebuilt  bool
}

func (b *queryBulder) prebuildInsert() error {
	if !b.prebuilt {
		templ, err := template.New("").Parse(
			`INSERT INTO
{{ .db }}.{{ .table }}
VALUES`,
		)
		if err != nil {
			return err
		}
		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]string{
			"db":    b.db,
			"table": b.table,
		})
		if err != nil {
			return err
		}
		b.templ = buf.String()
	}
	return nil
}

const sep = ","

func (b *queryBulder) PushRow(row []string) error {
	if !b.prebuilt {
		err := b.prebuildInsert()
		if err != nil {
			return err
		}
	}
	if b.sb.Len() == 0 {
		b.sb.WriteString(b.templ)
		b.sb.WriteString("(")
	} else {
		b.sb.WriteString(",(")
	}
	_, err := b.sb.WriteString(strings.Join(row, sep))
	if err != nil {
		return err
	}
	_, err = b.sb.WriteString(")")
	return err
}

func (b *queryBulder) String() string {
	if !b.prebuilt {
		return ""
	}
	return b.sb.String()
}

func (b *queryBulder) Clear() {
	b.sb.Reset()
}
