package account

type repository interface {
	Read(id string) (Account, error)
	Create(Account) error
	List() []Account
}
