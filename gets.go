package gets

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetInt Some shit function
func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return value, nil
}
