package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (acc *Account) Output() {
	fmt.Println("________________________")
	fmt.Println("Ваш логин:", acc.Login)
	fmt.Println("Ваш пароль:", acc.Password)
	fmt.Println("Ваш адрес:", acc.Url)
	fmt.Println("________________________")
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" || len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Url:      urlString,
		Login:    login,
		Password: password,
	}
	//field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	//fmt.Println(field.Tag)
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
