package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com//rudyjcruz831/golangTest/strings/greeting" // Importing a nested package
	"github.com/rudyjcruz831/golangTest/numbers"
	"github.com/rudyjcruz831/golangTest/strings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
