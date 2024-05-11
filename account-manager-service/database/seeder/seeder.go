package seeder

import (
	"account-manager-service/internal/entity"
	"account-manager-service/internal/repository"
	"account-manager-service/internal/util"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, repositorySetup *repository.RepositorySetup) error {
	users, err := seedUsers(db, repositorySetup.UserRepository)
	if err != nil {
		return err
	}

	accountTypes, err := seedAccountType(db, repositorySetup.AccountTypeRepository)
	if err != nil {
		return err
	}

	currencies, err := seedCurrency(db, repositorySetup.CurrencyRepository)
	if err != nil {
		return err
	}

	accounts, err := seedAccount(db, repositorySetup.AccountRepository, users, accountTypes)
	if err != nil {
		return err
	}

	_, err = seedTransaction(db, repositorySetup.TransactionRepository, accounts, currencies)
	if err != nil {
		return err
	}

	return nil
}

func seedUsers(db *gorm.DB, userRepository *repository.UserRepository) ([]entity.User, error) {
	var numUsers = 9
	var users = make([]entity.User, numUsers)

	for i := 1; i <= numUsers; i++ {
		newPassword, _ := util.GenerateFromPassword("password")
		newUser := entity.User{
			ID:        uuid.NewString(),
			FirstName: "User",
			LastName:  strconv.Itoa(i),
			Password:  newPassword,
			Phone:     strconv.Itoa(i * 10000000),
			Email:     fmt.Sprintf("user%d@example.com", i),
		}
		tx := db.Begin()
		if err := userRepository.Repository.Create(tx, &newUser); err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
		users[i-1] = newUser
	}

	return users, nil
}

func seedAccountType(db *gorm.DB, accountTypeRepository *repository.AccountTypeRepository) ([]entity.AccountType, error) {
	var accountTypeNames []string = []string{
		"credit",
		"debit",
		"loan",
	}
	var accountTypes = make([]entity.AccountType, len(accountTypeNames))

	for i, aN := range accountTypeNames {
		newAccountType := entity.AccountType{
			ID:   uuid.NewString(),
			Name: aN,
		}
		tx := db.Begin()
		if err := accountTypeRepository.Repository.Create(tx, &newAccountType); err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
		accountTypes[i] = newAccountType
	}

	return accountTypes, nil
}

func seedCurrency(db *gorm.DB, currencyRepository *repository.CurrencyRepository) ([]entity.Currency, error) {
	var currencyToIDR map[string]float64 = map[string]float64{
		"IDR": 1.0,
		"USD": 0.000062,
		"GBP": 0.00005,
		"JPY": 0.0097,
	}
	var currencies = make([]entity.Currency, len(currencyToIDR))

	i := 0
	for code, nom := range currencyToIDR {
		newCurrency := entity.Currency{
			ID:             uuid.NewString(),
			Code:           code,
			CurrentPer1IDR: nom,
		}
		tx := db.Begin()
		if err := currencyRepository.Repository.Create(tx, &newCurrency); err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
		currencies[i] = newCurrency
		i += 1
	}

	return currencies, nil
}

func seedAccount(db *gorm.DB, accountRepository *repository.AccountRepository, users []entity.User, accountTypes []entity.AccountType) ([]entity.Account, error) {
	var accounts []entity.Account = make([]entity.Account, len(users))

	for i, usr := range users {
		randAccType := accountTypes[util.GetRandomNumberBetween(0, len(accountTypes)-1)]
		newAcc := entity.Account{
			ID:            uuid.NewString(),
			UserID:        usr.ID,
			AccountTypeID: randAccType.ID,
			Number: fmt.Sprintf("%s%d%s%d-%s%d%s%d",
				usr.LastName, 0, usr.LastName, 0,
				usr.LastName, 0, usr.LastName, 0),
			BalanceIDR: float64(util.GetRandomNumberBetween(50000, 10000000)),
		}
		tx := db.Begin()
		if err := accountRepository.Repository.Create(tx, &newAcc); err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
		accounts[i] = newAcc
	}

	return accounts, nil
}

func seedTransaction(db *gorm.DB, transactionRepository *repository.TransactionRepository, accounts []entity.Account, currencies []entity.Currency) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	for _, senderAcc := range accounts {
		numOfTransactions := util.GetRandomNumberBetween(0, 5)

		for i := 0; i < numOfTransactions; i++ {
			var recipientAcc entity.Account
			for {
				recipientAcc = accounts[util.GetRandomNumberBetween(0, len(accounts)-1)]
				if recipientAcc.ID != senderAcc.ID {
					break
				}
			}
			isTransactionFailed := util.GetRandomNumberBetween(0, 2) == 1
			currencyUsed := currencies[util.GetRandomNumberBetween(0, len(currencies)-1)]

			newTransaction := new(entity.Transaction)
			newTransaction.ID = uuid.NewString()
			newTransaction.SenderAccountID = senderAcc.ID
			newTransaction.RecipientAccountID = recipientAcc.ID
			newTransaction.CurrencyID = currencyUsed.ID
			newTransaction.IDRAmount = float64(util.GetRandomNumberBetween(10000, 40000))
			newTransaction.OriginalAmount = newTransaction.IDRAmount * currencyUsed.CurrentPer1IDR
			if isTransactionFailed {
				newTransaction.Status = 0
				failedAt := time.Now().Add(30 * time.Second)
				newTransaction.FailedAt = &failedAt
			} else {
				newTransaction.Status = 2
				succeedAt := time.Now().Add(30 * time.Second)
				newTransaction.SucceedAt = &succeedAt
			}
			newTransaction.CreatedAt = time.Now()

			tx := db.Begin()
			if err := transactionRepository.Repository.Create(tx, newTransaction); err != nil {
				tx.Rollback()
				return nil, err
			}
			tx.Commit()
			transactions = append(transactions, *newTransaction)
		}
	}

	return transactions, nil
}
