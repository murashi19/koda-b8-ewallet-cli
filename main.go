package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
)

type User struct {
	Username string
	Email    string
	Password string
}

var users []User

type Auth interface {
	Register()
	Login()
	ForgotPassword()
	ListUser()
	SearchUser()
	Logout()
}

type Authentication struct{}

func CreatePassword() string {

	password := Input("Enter a strong password: ")
	confirm := Input("Confirm your password: ")
	if password != confirm {
		fmt.Println("Passwords do not match !!, please enter to back")
		fmt.Scanf("&\n")
		CreatePassword()
	}
	return password
}

func (a *Authentication) Register() {
	ClearScreen()

	username := Input("What is your name: ")
	email := Input("What is your Email: ")
	password := CreatePassword()

	user := User{
		Username: username,
		Email:    email,
		Password: HashPassword(password),
	}

	users = append(users, user)

	fmt.Println("\nIs it true?")
	fmt.Println("Username: ", username)
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)
	close := Input("Continue (y/n): ")
	if close == "y" {
		fmt.Println("Register Succes")
	} else {
		a.Register()
	}
}
func (a *Authentication) Login() {
	ClearScreen()

	for {
		email := Input("Enter your Email: ")
		password := HashPassword(Input("Enter your password: "))

		for _, user := range users {
			if user.Email == email && user.Password == password {
				fmt.Println("Login Success, press enter to dashboard")
				fmt.Scanf("&\n")
				Dashboard()
			}
		}
		fmt.Println("Wrong Email or Password, please enter to back...")
		fmt.Scanf("&\n")
		ClearScreen()
	}

}

func (a *Authentication) ListUser() {
	ClearScreen()

	for i, user := range users {
		fmt.Println(i+1, user)
	}
	fmt.Println("\nEnter to Back Dashboard")
	fmt.Scanf("&\n")
	Dashboard()

}

func (a *Authentication) SearchUser() {
	ClearScreen()
	fmt.Println("Pencarian User")
	input := Input("\nSearch User : ")

	for _, user := range users {
		if input == user.Username {
			fmt.Println("\nUser telah ditemukan, ", user)
			fmt.Println("\nBack to dashboard, press enter")
			fmt.Scanf("&\n")
			Dashboard()
		}
	}
}

func (a *Authentication) Logout() {
	ClearScreen()
	main()
}
func Dashboard() {
	ClearScreen()

	var auth Auth = &Authentication{}

	fmt.Println("--- Welcome to System ---")
	fmt.Println("1. List Users")
	fmt.Println("2. Search User")
	fmt.Println("3. Logout")
	fmt.Println("")

	input := Input("Choose a menu: ")

	switch input {
	case "1":
		auth.ListUser()
	case "2":
		auth.SearchUser()
	case "3":
		auth.Logout()
	}
}

func (a *Authentication) ForgotPassword() {
	ClearScreen()

	email := Input("Enter your Email: ")
	newpassword := Input("Enter your new password: ")

	for i, user := range users {
		if user.Email == email {
			users[i].Password = newpassword
			fmt.Println("Update Password Success, press enter to back...")
			fmt.Scanf("&\n")
			return
		}
	}
}
func Exit() {
	os.Exit(0)
}

var scanner = bufio.NewScanner(os.Stdin)

func Input(promt string) string {
	fmt.Print(promt)
	if !scanner.Scan() {
		panic(scanner.Err())
	}
	return scanner.Text()
}

func ClearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func Recover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	var auth Auth = &Authentication{}
	defer Recover()
	ClearScreen()

	for {
		ClearScreen()
		fmt.Println("\n--- Welcome to System ---")
		fmt.Println("1. Register ")
		fmt.Println("2. Login ")
		fmt.Println("3. Forgot Password")
		fmt.Println("")
		fmt.Println("0. Exit")
		fmt.Println(" ")

		input := Input("Choose a menu: ")

		switch input {
		case "1":
			auth.Register()
		case "2":
			auth.Login()
		case "3":
			auth.ForgotPassword()
		case "0":
			Exit()
		default:
			ClearScreen()
			fmt.Println("\nInvalid Choose menu, enter back to menu")
			fmt.Scanf("&\n")
			main()

		}

	}
}
