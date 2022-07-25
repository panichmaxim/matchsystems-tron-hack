package neoutils

type Node struct {
	ID     int64
	Labels []string
	Props  map[string]interface{}
}
