package db

type Querier interface {
	GetAccount(id string) Account
}

// これはコンパイル時にインターフェースの実装が正しいことを確認するための文法
// プログラムの実行時に影響を与えるものではない
var _ Querier = (*Queries)(nil)
