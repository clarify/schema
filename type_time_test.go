package schema_test

import (
	"testing"
	"time"

	"github.com/searis/schema"
)

func TestTimeParser(t *testing.T) {
	correctTime := func(out interface{}, expect time.Time) func(t *testing.T) {
		return func(t *testing.T) {
			if ts, ok := out.(time.Time); !ok || !expect.Equal(ts) {
				t.Errorf("\n got: %v\nwant: %s", out, expect)
			}
		}
	}

	t.Run("given a Time schema with default settings", func(t *testing.T) {
		sc := schema.Time{}
		parser := sc.Parser()

		t.Run("when parsing an RFC3339 string", func(t *testing.T) {
			in := "2019-01-02T13:37:00Z"
			out, err := parser.Parse(in)

			t.Run("then there should be no error", noErr(err))
			expect := time.Date(2019, 01, 02, 13, 37, 0, 0, time.UTC)
			t.Run("then the correct time should be returned", correctTime(out, expect))
		})

		t.Run("when parsing an RFC1123Z string", func(t *testing.T) {
			in := "Wed, 02 Jan 2019 13:37:00 +0000"
			out, err := parser.Parse(in)

			t.Run("then there should be no error", noErr(err))
			expect := time.Date(2019, 01, 02, 13, 37, 0, 0, time.UTC)
			t.Run("then the correct time should be returned", correctTime(out, expect))
		})

		t.Run("when parsing an unsupported format", func(t *testing.T) {
			in := "02.01.2019-13:37:00+0000"
			out, err := parser.Parse(in)

			t.Run("then an error should be returned", errMatch(err, "invalid format"))
			t.Run("then out should be nil", isNill(out))
		})
	})
}
