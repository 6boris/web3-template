package model

type DefiPrice struct {
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
	SourceToken string  `json:"source_token"`
	TargetToken string  `json:"target_token"`
}
