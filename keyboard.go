package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func GetFloatFromInput(str string) (float64, error) {
	fmt.Print(str)
	reader := bufio.NewReader(os.Stdin)
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

