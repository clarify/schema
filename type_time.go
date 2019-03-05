package schema

import (
	"time"
)

var (
	_ Type       = Time{}
	_ DocType    = Time{}
	_ LesserType = Time{}
)

var defaultTimeLayouts = []string{
	time.RFC3339Nano,
	time.RFC1123Z,
	time.RFC822Z,
}

// Time parses strings with a date-time format to time.Time.
type Time struct {
	// ParseLayouts lets you specify which layouts to use for parsing times. By
	// default, a predefined set of layouts from the time package will be tried.
	ParseLayouts []string

	// SerializeLayout lets you specify which layout to use for serialization.
	// Default is RFC3339Nano.
	SerializeLayout string

	// Truncate, when set, truncates time-stamps to the given precision relative
	// from the Go zero-time.
	Truncate time.Duration
}

func (Time) Doc(s Schema) Doc {
	return StringDoc{
		Schema: s,
		Type:   TypeString,
		Format: FormatDateTime,
	}
}

func (t Time) Serializer() Serializer {
	layout := t.SerializeLayout
	if layout == "" {
		layout = defaultTimeLayouts[0]
	}

	return SerializerFunc(func(in interface{}) (out interface{}, err error) {
		ts, ok := in.(time.Time)
		if !ok {
			return nil, ErrNotGoTime
		}
		return ts.Format(layout), nil
	})
}

func (t Time) Parser() Parser {
	layouts := t.ParseLayouts
	if len(layouts) == 0 {
		layouts = defaultTimeLayouts
	}

	return ParserFunc(func(in interface{}) (out interface{}, err error) {
		sIn, ok := in.(string)
		if !ok {
			return nil, ErrNotString
		}

		var ts time.Time
		for _, layout := range layouts {
			ts, err = time.Parse(layout, sIn)
			if err == nil {
				return ts.Truncate(t.Truncate), nil
			}
		}

		return nil, ErrInvalidFormat
	})
}

func (t Time) Validator() Validator {
	return ValidatorFunc(func(in, original interface{}) (out interface{}, err error) {
		ts, ok := in.(time.Time)
		if !ok {
			return nil, ErrNotGoTime
		}
		return ts, nil
	})
}

func (t Time) Lesser() Lesser {
	return LesserFunc(func(a, b interface{}) bool {
		at, ok := a.(time.Time)
		if !ok {
			return false
		}
		bt, ok := b.(time.Time)
		if !ok {
			return false
		}

		return at.Before(bt)
	})
}
