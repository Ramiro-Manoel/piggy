package account

type repository interface {
	Read(id string) (Account, error)
	Save(a Account) error
	List() []Account
}
