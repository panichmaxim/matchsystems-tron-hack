package neo4jstore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
	"time"
)

func CreateSession(driver neo4j.DriverWithContext) neo4j.SessionWithContext {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func CreateDriver(ctx context.Context, uri, username, password string) (neo4j.DriverWithContext, error) {
	auth := neo4j.BasicAuth(username, password, "")
	return neo4j.NewDriverWithContext(uri, auth, func(config *neo4j.Config) {
		config.MaxConnectionLifetime = 30 * time.Minute
		config.MaxConnectionPoolSize = 50
		config.ConnectionAcquisitionTimeout = 2 * time.Minute
		config.SocketConnectTimeout = 15 * time.Second
		config.MaxTransactionRetryTime = 15 * time.Second
		config.Log = NewNeoLogger(zerolog.Ctx(ctx))
	})
}

func BuildLimitOffset(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if pageSize < 0 || pageSize > 100 {
		pageSize = 100
	}

	var from int
	if page > 1 {
		from = (page - 1) * pageSize
	}

	return from, pageSize
}

func GetItems(ctx context.Context, result neo4j.ResultWithContext, key string) ([]*Node, error) {
	var items []*Node
	for result.Next(ctx) {
		node, err := ExtractNode(result.Record(), key)
		if err != nil {
			return nil, err
		}

		items = append(items, node)
	}

	return items, nil
}

func GetItem(ctx context.Context, result neo4j.ResultWithContext, key string) (*Node, error) {
	if result.Next(ctx) {
		return ExtractNode(result.Record(), key)
	}

	return nil, result.Err()
}

func ExtractNode(record *neo4j.Record, key string) (*Node, error) {
	item, ok := record.Get(key)
	if !ok {
		return nil, fmt.Errorf("skip missing node: %+v", record)
	}

	if item == nil {
		return nil, nil
	}

	switch item.(type) {
	case dbtype.Node:
		v, ok := item.(dbtype.Node)
		if !ok {
			return nil, fmt.Errorf("skip invalid node type: %+v", record)
		}

		return convert(v), nil

	case dbtype.Relationship:
		v, ok := item.(dbtype.Relationship)
		if !ok {
			return nil, fmt.Errorf("skip invalid node type: %+v", record)
		}

		return convertRelationship(v), nil
	}

	return nil, fmt.Errorf("unknown node type: %T", item)
}

func convert(v dbtype.Node) *Node {
	return &Node{
		ID:     v.Id,
		Labels: v.Labels,
		Props:  v.Props,
	}
}

func getProps(v any) (map[string]interface{}, error) {
	switch v.(type) {
	case dbtype.Node:
		v, ok := v.(dbtype.Node)
		if !ok {
			return nil, fmt.Errorf("skip invalid node type: %+v", v)
		}
		return v.Props, nil

	case dbtype.Relationship:
		v, ok := v.(dbtype.Relationship)
		if !ok {
			return nil, fmt.Errorf("skip invalid node type: %+v", v)
		}

		return v.Props, nil
	}

	return nil, fmt.Errorf("unknown type %T", v)
}

func ConvertType[T any](v any) (*T, error) {
	props, err := getProps(v)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(props)
	if err != nil {
		return nil, err
	}
	data := new(T)
	if err := json.Unmarshal(payload, data); err != nil {
		return nil, err
	}

	return data, nil
}

func convertRelationship(v dbtype.Relationship) *Node {
	return &Node{
		ID:    v.Id,
		Props: v.Props,
	}
}

func Count(ctx context.Context, session neo4j.SessionWithContext, query string, values map[string]interface{}, key string) (int, error) {
	result, err := session.Run(ctx, query, values)
	if err != nil {
		return 0, err
	}
	if result.Next(ctx) {
		record := result.Record()
		item, ok := record.Get(key)
		if !ok {
			return 0, fmt.Errorf("invalid node: %v", record.Values)
		}

		return cast.ToInt(item), nil
	}

	return 0, result.Err()
}
