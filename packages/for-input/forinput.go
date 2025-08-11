package forinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputNum(text string) int {
	var num int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("failed to read the number: %w", err)
	}

	number = strings.TrimSpace(number)
	num, err = strconv.Atoi(number)
	if err != nil {
		fmt.Errorf("failed to convert the number: %w", err)
	}

	return num
}
