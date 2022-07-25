package scalars

import (
	"fmt"
	"io"
	"time"
)

type DateTime struct {
	t time.Time
}

func New(t time.Time) *DateTime {
	return &DateTime{
		t: t,
	}
}

func (d *DateTime) GetTime() time.Time {
	return d.t
}

func (d *DateTime) UnmarshalGQL(vi interface{}) (err error) {
	v, ok := vi.(string)
	if !ok {
		return fmt.Errorf("unknown type of DateTime: `%+v`", vi)
	}
	if d.t, err = time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	return nil
}

func (d DateTime) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(d.t.Format(time.RFC3339)))
}
