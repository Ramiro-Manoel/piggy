package account

type Repository interface {
	Read(id string) (Account, error)
	Save(a Account) error
	List() []Account
}
