package cfg

import (
	"github.com/stretchr/testify/require"
	"gitlab.com/falaleev-golang/config"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := DistributedConfig{}
	err := config.Load([]string{".env.test"}, &cfg)
	require.NoError(t, err)
	require.Equal(t, "https://example.com", cfg.App.SiteURL)
}
