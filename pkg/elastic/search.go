package elastic

import (
	"bytes"
	"context"
	"encoding/json"
)

func (e *elasticImpl) Search(ctx context.Context, query string, page, limit int) ([]*Entity, int, error) {
	if page <= 0 {
		page = 1
	}

	if limit < 0 || limit > 100 {
		limit = 100
	}

	var from int
	if page > 1 {
		from = (page - 1) * limit
	}

	elkQuery := map[string]interface{}{
		"from": from,
		"size": limit,
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"*"},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(elkQuery); err != nil {
		return nil, 0, err
	}

	payload, err := e.search(ctx, &buf, e.es.Search.WithFilterPath("hits.hits._source,hits.total.value"))
	if err != nil {
		return nil, 0, err
	}

	r := &ElasticResponse{}
	if err := json.NewDecoder(payload).Decode(r); err != nil {
		return nil, 0, err
	}

	entities := make([]*Entity, len(r.Hits.Hits))
	for i, c := range r.Hits.Hits {
		entities[i] = &Entity{
			Date:     c.Source.Date,
			Address:  c.Source.Address,
			Chain:    c.Source.Chain,
			Contact:  c.Source.Contact,
			Category: c.Source.Category,
			Data:     c.Source.Data,
		}
	}

	return entities, r.Hits.Total.Value, nil
}
