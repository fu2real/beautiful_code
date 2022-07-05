import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	input = strings.ReplaceAll(input, " ", "")
	o1 := ""
	o2 := ""
	o3 := ""
	oi1 := 0
	oi2 := 0
	sym1 := ""
	sym2 := ""
	sep1 := false
	sep2 := false

	for i, v := range input {

		_, err1 := strconv.Atoi(string(v))
		if err1 != nil && o2 != "" {
			sep2 = true
		}

		x, err2 := strconv.Atoi(string(v))
		if err2 == nil {

			if sep1 == false {
				o1 = o1 + strconv.Itoa(x)
			} else if sep2 == false {
				o2 = o2 + strconv.Itoa(x)
			} else {
				o3 = o3 + strconv.Itoa(x)
			}

			if i == 0 {
				sym1 = "+"
			}
			if i > 0 && o2 == "" {
				if string(input[i-1]) == "-" {
					sym1 = "-"
				}
			}
			if i > 0 && o2 != "" {
				if string(input[i-1]) == "-" && o2 != "" {
					sym2 = "-"
				}
			}

		} else if o1 != "" {
			sep1 = true
		}
	}

	if o3 != "" {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	if o2 == "" {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	oi1, _ = strconv.Atoi(sym1 + o1)
	oi2, _ = strconv.Atoi(sym2 + o2)
	output = strconv.Itoa(oi1 + oi2)
	return output, nil

}