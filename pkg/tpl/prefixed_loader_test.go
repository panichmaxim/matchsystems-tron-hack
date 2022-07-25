package tpl

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed views/*
var templatesdata embed.FS

func TestPrefixedFilesystem(t *testing.T) {
	s, err := NewPrefixedFilesystem(templatesdata, "views")
	require.NoError(t, err)
	_, err = s.Open("index.jet.html")
	require.NoError(t, err)
	require.True(t, s.Exists("index.jet.html"))
}
