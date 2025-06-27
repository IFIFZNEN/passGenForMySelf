package main

import (
	"fmt"

	"passGenForMySelf/account"
	"passGenForMySelf/files"
	"passGenForMySelf/output"
)

func main() {
	fmt.Println("Программа для создания учётных записей для разных сайтов")
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		variant := promtData([]string{
			"Введите: число 1 - Создать аккаунт",
			"Введите: число 2 - Найти аккаунт",
			"Введите: число 3 - Удалить аккаунт",
			"Введите: число 4 или 0 (или любое иное число для выхода из приложения)",
			"Введите команду",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите ссылку"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	myAccount.Output()
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Ввести URL для поиска"})
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Ввести URL для поиска"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		fmt.Println("Успешно удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promtData[T any](promt []T) string {
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
