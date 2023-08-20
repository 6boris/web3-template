package redis

type Cmd struct {
	Key string `json:"key"`
	TTL int64  `json:"ttl"`
}
