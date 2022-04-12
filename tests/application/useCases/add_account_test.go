package usecases_test

import (
	"errors"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

type AccountRepository interface {
	checkAccountByEmail(email string) bool
	checkAccountByUserName(userName string) bool
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

	emailTaken := useCase.accountRepository.checkAccountByEmail(email)
	if emailTaken {
		return nil, errors.New("email already taken")
	}

	userNameTaken := useCase.accountRepository.checkAccountByUserName(userName)
	if userNameTaken {
		return nil, errors.New("username already taken")
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
	checkAccountOutput  bool
	checkUserNameOutput bool
}

func (repo *FakeAccountRepository) checkAccountByEmail(email string) bool {
	return repo.checkAccountOutput
}

func (repo *FakeAccountRepository) checkAccountByUserName(userName string) bool {
	return repo.checkUserNameOutput
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

func TestAddAccountUseCase_WithEmailAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkAccountOutput = true

	createdAccount, err := sut.addAccount(fakeUserName, "already_taken_email@gmail.com", fakePassword, fakePassword)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "email already taken")
}

func TestAddAccountUseCase_WithUsernameAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkUserNameOutput = true

	createdAccount, err := sut.addAccount("already_taken_username", fakeEmail, fakePassword, fakePassword)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "username already taken")
}
