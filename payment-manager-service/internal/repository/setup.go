package repository

type RepositorySetup struct {
	UserRepository        *UserRepository
	AccountTypeRepository *AccountTypeRepository
	CurrencyRepository    *CurrencyRepository
	AccountRepository     *AccountRepository
	TransactionRepository *TransactionRepository
}

func Setup() *RepositorySetup {
	return &RepositorySetup{
		UserRepository:        NewUserRepository(),
		AccountTypeRepository: NewAccounTypeRepository(),
		CurrencyRepository:    NewCurrencyRepository(),
		AccountRepository:     NewAccountRepository(),
		TransactionRepository: NewTransactionRepository(),
	}
}
