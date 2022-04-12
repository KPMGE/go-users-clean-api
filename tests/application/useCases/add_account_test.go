package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type AccountRepository interface {
	checkAccountByEmail(email string) bool
	save(accout *entities.Account) error
}

type Hasher interface {
	hash(plainText string) string
}

type AddAccountUseCase struct {
	accountRepository AccountRepository
	hasher            Hasher
}

type AddAccountOutputDTO struct {
	ID       string
	UserName string
	Email    string
}

func (useCase *AddAccountUseCase) addAccount(userName string, email string, password string, confirmPassword string) (*AddAccountOutputDTO, error) {
	account, _ := entities.NewAccount(userName, email, password)
	useCase.hasher.hash(password)
	output := AddAccountOutputDTO{
		ID:       account.ID,
		UserName: account.UserName,
		Email:    account.Email,
	}

	if password != confirmPassword {
		return nil, errors.New("password and confirmPassword must match")
	}

	return &output, nil
}

func NewAddAccountUseCase(accountRepository AccountRepository, hasher Hasher) *AddAccountUseCase {
	return &AddAccountUseCase{
		accountRepository: accountRepository,
		hasher:            hasher,
	}
}

type hasherSpy struct {
	input string
}

func (hasher *hasherSpy) hash(plainText string) string {
	hasher.input = plainText
	return "hashed_plain_text"
}

type FakeAccountRepository struct {
}

func (repo *FakeAccountRepository) checkAccountByEmail(email string) bool {
	return true
}

func (repo *FakeAccountRepository) save(accout *entities.Account) error {
	return nil
}

func NewHasherSpy() *hasherSpy {
	return &hasherSpy{}
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{}
}

const fakeUserName string = "any_user_name"
const fakeEmail string = "any_valid_email@gmail.com"
const fakePassword string = "any_password"

func makeSut() (*AddAccountUseCase, *hasherSpy, *FakeAccountRepository) {
	repo := NewFakeAccountRepository()
	hasher := NewHasherSpy()
	sut := NewAddAccountUseCase(repo, hasher)
	return sut, hasher, repo
}

func TestAddAccountUseCase_WithRightData(t *testing.T) {
	sut, hasher, _ := makeSut()
	createdAccount, err := sut.addAccount(fakeUserName, fakeEmail, fakePassword, fakePassword)

	require.Nil(t, err)
	require.Equal(t, hasher.input, fakePassword)
	require.Equal(t, createdAccount.Email, fakeEmail)
	require.Equal(t, createdAccount.UserName, fakeUserName)
}

func TestAddAccountUseCase_WithDifferentPasswordAndConfirmPassword(t *testing.T) {
	sut, _, _ := makeSut()
	createdAccount, err := sut.addAccount(fakeUserName, fakeEmail, fakePassword, "any_other_value")

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "password and confirmPassword must match")
}
