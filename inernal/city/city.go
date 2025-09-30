package city

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChekCity() (string, error) {
	var city string
	if len(os.Args) > 1 {
		city = strings.Join(os.Args[1:], "")
	} else {
		// Запрашиваем  у пользователя название  города, если аргумент не  указан
		fmt.Print("Enter the city name: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("Error reading input: %v", err)
		}
		city = strings.TrimSpace(input)
	}

	return city, nil
}