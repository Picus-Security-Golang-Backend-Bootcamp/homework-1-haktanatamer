package helper

import (
	"fmt"
	"log"
	"strings"
)

func StringLower(data string) string {
	return strings.ToLower(data)
}

func StringContains(line, word string) bool {
	return strings.Contains(StringLower(line), StringLower(word))
}

func AddLog(detail string, err error) {
	log.Printf("%s Hata : %s", detail, err)
}

func ConsoleBook(counter int, book string) {
	fmt.Printf("%d . %s \n", counter, book)
}

func AddSeparator() {
	fmt.Println("--------------------------")
}
