package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccount []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

// Delete implements AccountRepository.
func (obj accountRepository) Delete(id string) error {

	return obj.db.Where("id =?", id).Delete(&BankAccount{}).Error
}

// FindAll implements AccountRepository.
func (obj accountRepository) FindAll() (bankAccount []BankAccount, err error) {
	err = obj.db.Find(&BankAccount{}).Error

	return bankAccount, err
}

// FindByID implements AccountRepository.
func (obj accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	err = obj.db.Where("id=?", id).First(&bankAccount).Error

	return bankAccount, err
}

// Save implements AccountRepository.
func (obj accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Save(bankAccount).Error
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	//db.Table("bank_account").AutoMigrate(&BankAccount{})
	db.AutoMigrate(&BankAccount{})

	return accountRepository{db}
}
