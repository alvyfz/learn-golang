package main

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		println("You are Blocked", name)
	} else {
		println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Alvy", blacklist)
}
