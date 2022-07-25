package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
)

var _ Client = (*elasticImpl)(nil)

func NewElastic(addresses []string) (Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addresses,
	})
	if err != nil {
		return nil, err
	}

	return &elasticImpl{
		es: es,
	}, nil
}

type elasticImpl struct {
	es *elasticsearch.Client
}

func (e *elasticImpl) search(ctx context.Context, buf *bytes.Buffer, o ...func(*esapi.SearchRequest)) (io.ReadCloser, error) {
	var opts []func(*esapi.SearchRequest)
	opts = append(opts, e.es.Search.WithContext(ctx), e.es.Search.WithBody(buf))
	opts = append(opts, o...)

	res, err := e.es.Search(opts...)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(
			"[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	return res.Body, nil
}
