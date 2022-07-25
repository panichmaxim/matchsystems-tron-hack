package elastic

type Entity struct {
	Date     string `json:"date"`
	Address  string `json:"address"`
	Chain    string `json:"chain"`
	Contact  string `json:"contact"`
	Category string `json:"category"`
	Data     any    `json:"data"`
}

type ElasticResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index   string   `json:"_index"`
			Id      string   `json:"_id"`
			Score   float64  `json:"_score"`
			Ignored []string `json:"_ignored"`
			Source  Entity   `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
