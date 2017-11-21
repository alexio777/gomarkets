package connector

// Connector - универсальный интерфейс к биржам
type Connector interface {
	IsConnected() bool
}
