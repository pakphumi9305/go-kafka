package repositories

//import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepositoriy interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccount []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

// type accountRepository struct {
// 	db *gorm.DB
// }

// func NewAccountRepository(db *gorm.DB) AccountRepository {
// 	db.Table("test_banks").AutoMigrate(&BankAccount{})

// 	return accountRepository{db}
// }
