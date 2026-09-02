package account

import "github.com/Ramiro-Manoel/piggy/internal/external"

type Account struct {
	ID      string
	Ref     external.Reference
	Number  string
	Name    string
	Owner   string
	Balance int64
}
