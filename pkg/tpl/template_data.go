package tpl

type Data map[string]interface{}

func (t Data) Set(key string, value interface{}) {
	t[key] = value
}
