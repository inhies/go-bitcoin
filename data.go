package bitcoin

// getinfo response
type ApiInfo struct {
	Balance         float64
	Blocks          int64
	Connections     int64
	Difficulty      float64
	Errors          string
	KeyPoolOldest   float64
	KeyPoolSize     int64
	PayTxFee        float64
	ProtocolVersion int64
	Proxy           string
	Testnet         bool
	Version         int64
	WalletVersion   int64
}

// This is here to enable direct unmarshalling of the JSON. If you know a better
// way I'm all eyes :)
type apiInfo struct {
	Error  *string
	ID     int64
	Result ApiInfo
}

// getblockcount response
type apiBlockCount struct {
	Error  *string
	ID     int64
	Result int64
}
