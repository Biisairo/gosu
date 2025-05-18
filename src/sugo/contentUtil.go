package sugo

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseIndexedName(s string) (int, string, error) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("'.' 구분자가 없거나 형식이 올바르지 않음")
	}

	i, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("%v 의 앞부분이 정수가 아님: %v", s, err)
	}

	return i, parts[1], nil
}
