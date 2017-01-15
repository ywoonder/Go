package config


type DBConfig struct {
	MasterConn   string
	SlaveConn    string
	MaxOpenConns int
}

// type PaymentDBConfig struct {
// 	MasterConn   string
// 	SlaveConn    string
// 	MaxOpenConns int
// }
