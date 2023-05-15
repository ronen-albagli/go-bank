package ports

type Asset struct {
	assetType string
	amount    int64
	reason    string
	name      string
}

type LedgerGateway interface {
	save(asset Asset)
}
