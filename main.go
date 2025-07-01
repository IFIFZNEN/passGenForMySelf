package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"passGenForMySelf/account"
	"passGenForMySelf/files"
	"passGenForMySelf/output"
)

var menuVariants = []string{
	"Введите: число 1 - Создать аккаунт",
	"Введите: число 2 - Найти аккаунт по URL",
	"Введите: число 3 - Найти аккаунт по логину",
	"Введите: число 4 - Удалить аккаунт",
	"Введите: число 5 или 0 (или любое иное число для выхода из приложения)",
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
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Программа для создания учётных записей для разных сайтов")
	vault := account.NewVault(files.NewJsonDB("data.json"))
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

		//
		//switch variant {
		//case "1":
		//	createAccount(vault)
		//case "2":
		//	findAccount(vault)
		//case "3":
		//	deleteAccount(vault)
		//default:
		//	break Menu
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
