package model

type DefiPrice struct {
	Price       float64 `json:"float_64"`
	Timestamp   int64   `json:"timestamp"`
	SourceToken string  `json:"source_token"`
	TargetToken string  `json:"target_token"`
}
