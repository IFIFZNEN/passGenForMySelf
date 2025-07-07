package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"

	"passGenForMySelf/account"
	"passGenForMySelf/encrypter"
	"passGenForMySelf/files"
	"passGenForMySelf/output"
)

var menuVariants = []string{
	"Введите: 1 - Создать аккаунт",
	"Введите: 2 - Найти аккаунт по URL",
	"Введите: 3 - Найти аккаунт по логину",
	"Введите: 4 - Удалить аккаунт",
	"Введите: 5 или 0 (или любое иное число для выхода из приложения)",
	"Введите команду",
}

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println("Итерация:", i)
	}
}

func main() {
	fmt.Println("Программа для создания учётных записей для разных сайтов")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти .env файл")
	}
	vault := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
	counter := menuCounter()
Menu:
	for {
		counter()
		variant := promtData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			color.Red("Вы ввели неверную команду")
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите ссылку")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	myAccount.Output()
	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Ввести URL для поиска")
	accounts := vault.FindAccounts(url, func(aсс account.Account, str string) bool {
		return strings.Contains(aсс.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promtData("Ввести логин для поиска")
	accounts := vault.FindAccounts(login, func(aсс account.Account, str string) bool {
		return strings.Contains(aсс.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(account *[]account.Account) {
	if len(*account) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range *account {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData("Ввести URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		fmt.Println("Успешно удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promtData(promt ...string) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

// func twoSum(nums []int, target int) []int {
// 	result := make([]int, 0)

// 	for i := range nums {
// 		if nums[i] + nums[i+1] == target {
// 			result = append(result, i, i+1)
// 			break
// 		}
// 	}
// 	return result
// }
