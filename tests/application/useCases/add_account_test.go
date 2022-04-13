package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type AddAccountUseCase struct {
	accountRepository protocols.AccountRepository
	hasher            protocols.Hasher
}

type AddAccountOutputDTO struct {
	ID       string
	UserName string
	Email    string
}

type AddAccountInputDTO struct {
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (useCase *AddAccountUseCase) addAccount(input *AddAccountInputDTO) (*AddAccountOutputDTO, error) {
	emailTaken := useCase.accountRepository.CheckAccountByEmail(input.Email)
	if emailTaken {
		return nil, errors.New("email already taken")
	}

	userNameTaken := useCase.accountRepository.CheckAccountByUserName(input.UserName)
	if userNameTaken {
		return nil, errors.New("username already taken")
	}

	if input.Password != input.ConfirmPassword {
		return nil, errors.New("password and confirmPassword must match")
	}

	hashedPassword := useCase.hasher.Hash(input.Password)
	account, err := entities.NewAccount(input.UserName, input.Email, hashedPassword)

	if err != nil {
		return nil, err
	}

	err = useCase.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	output := AddAccountOutputDTO{
		ID:       account.ID,
		UserName: account.UserName,
		Email:    account.Email,
	}

	return &output, nil
}

func NewAddAccountUseCase(accountRepository protocols.AccountRepository, hasher protocols.Hasher) *AddAccountUseCase {
	return &AddAccountUseCase{
		accountRepository: accountRepository,
		hasher:            hasher,
	}
}

type hasherSpy struct {
	input string
}

func (hasher *hasherSpy) Hash(plainText string) string {
	hasher.input = plainText
	return "hashed_text"
}

type FakeAccountRepository struct {
	input               *entities.Account
	checkAccountOutput  bool
	checkUserNameOutput bool
}

func (repo *FakeAccountRepository) CheckAccountByEmail(email string) bool {
	return repo.checkAccountOutput
}

func (repo *FakeAccountRepository) CheckAccountByUserName(userName string) bool {
	return repo.checkUserNameOutput
}

func (repo *FakeAccountRepository) Save(account *entities.Account) error {
	repo.input = account
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

func makeFakeInput() *AddAccountInputDTO {
	return &AddAccountInputDTO{
		UserName:        fakeUserName,
		Email:           fakeEmail,
		Password:        fakePassword,
		ConfirmPassword: fakePassword,
	}
}

func makeSut() (*AddAccountUseCase, *hasherSpy, *FakeAccountRepository) {
	repo := NewFakeAccountRepository()
	hasher := NewHasherSpy()
	sut := NewAddAccountUseCase(repo, hasher)
	return sut, hasher, repo
}

func TestAddAccountUseCase_WithRightData(t *testing.T) {
	sut, hasher, repo := makeSut()
	fakeInput := makeFakeInput()
	createdAccount, err := sut.addAccount(fakeInput)

	require.Nil(t, err)
	require.Equal(t, hasher.input, fakePassword)
	require.Equal(t, createdAccount.Email, fakeEmail)
	require.Equal(t, createdAccount.UserName, fakeUserName)
	require.Equal(t, repo.input.Password, "hashed_text")
}

func TestAddAccountUseCase_WithDifferentPasswordAndConfirmPassword(t *testing.T) {
	sut, _, _ := makeSut()
	fakeInput := makeFakeInput()
	fakeInput.ConfirmPassword = "any_different_password"

	createdAccount, err := sut.addAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "password and confirmPassword must match")
}

func TestAddAccountUseCase_WithEmailAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkAccountOutput = true
	fakeInput := makeFakeInput()

	createdAccount, err := sut.addAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "email already taken")
}

func TestAddAccountUseCase_WithUsernameAlreadyTaken(t *testing.T) {
	sut, _, repo := makeSut()
	fakeInput := makeFakeInput()
	repo.checkUserNameOutput = true

	createdAccount, err := sut.addAccount(fakeInput)

	require.Error(t, err)
	require.Nil(t, createdAccount)
	require.Equal(t, err.Error(), "username already taken")
}

func TestAddAccountUseCase_WithBlankFields(t *testing.T) {
	sut, _, repo := makeSut()
	repo.checkUserNameOutput = false
	repo.checkUserNameOutput = false

	fakeInput := makeFakeInput()
	fakeInput.UserName = ""
	createdAccount, err := sut.addAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.Password = ""
	createdAccount, err = sut.addAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.ConfirmPassword = ""
	createdAccount, err = sut.addAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)

	fakeInput = makeFakeInput()
	fakeInput.Email = ""
	createdAccount, err = sut.addAccount(fakeInput)
	require.Error(t, err)
	require.Nil(t, createdAccount)
}
