package main

import "fmt"

type Account struct {
  Id int
  Balance float64
  Name string
  Type string
}

func (account Account) String () string {
  return fmt.Sprintf("Account[%d]:\n\tName: %s\n\tBalance: %0.2f\n\tType: %s", account.Id, account.Name, account.Balance, account.Type)
}

func (account *Account) Deposit (amount float64) float64 {
  account.Balance += amount
  return account.Balance
}

func (account *Account) Withdraw (amount float64) float64 {
  account.Balance -= amount
  return account.Balance
}

func main () {
  var acct1 Account = Account{}
  fmt.Println(acct1.String())

  var acct2 Account = Account{Id: 1, Name: "funding", Balance: 100, Type: "savings"}
  fmt.Println(acct2.String())
}
