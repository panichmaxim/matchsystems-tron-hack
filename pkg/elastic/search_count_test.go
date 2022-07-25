package elastic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestElasticImpl_InformationCount(t *testing.T) {
	es, err := NewElastic(testElasticHosts)
	require.NoError(t, err)

	r, err := es.SearchCount(ctx, "5039424408")
	require.NoError(t, err)
	require.Greater(t, r, 0)
}
