package webhook

// AssetClient : defines a client that can send a Asset anywhere
// could be any protocol: http, gRPC, apache kafka
type AssetClient interface {
	Send(asset *AssetWithStatus)
}
