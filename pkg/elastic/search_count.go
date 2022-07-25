package elastic

import (
	"bytes"
	"context"
	"encoding/json"
)

func (e *elasticImpl) SearchCount(ctx context.Context, query string) (int, error) {
	elkQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"*"},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(elkQuery); err != nil {
		return 0, err
	}

	payload, err := e.search(ctx, &buf, e.es.Search.WithFilterPath("hits.total.value"))
	if err != nil {
		return 0, err
	}

	r := &ElasticResponse{}
	if err := json.NewDecoder(payload).Decode(r); err != nil {
		return 0, err
	}

	return r.Hits.Total.Value, nil
}
