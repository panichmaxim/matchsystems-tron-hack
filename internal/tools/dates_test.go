package tools

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseFromToDatesRaw(t *testing.T) {
	layout := "2006-01-02"
	_, err := time.Parse(layout, "2022-08-01")
	require.NoError(t, err)
	_, err = time.Parse(layout, "2022-08-09")
	require.NoError(t, err)
}

func TestParseFromToDates(t *testing.T) {
	from, to := ParseFromToDates(Ptr[string]("2022-08-01"), Ptr[string]("2022-08-09"))
	require.NotNil(t, from)
	require.NotNil(t, to)
	require.False(t, from.IsZero())
	require.False(t, to.IsZero())
}
