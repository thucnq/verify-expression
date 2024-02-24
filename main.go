package main

import (
	"fmt"
	"strings"
)

func main() {
	expressions := []string{
		"a/bx",
		"a/(b-c)",
		"a/(b-(c-d))",
		"a/(b-(b-c)-e/(b-c))",
		"a/(b-(b-c)-e/bxa-c))",
		"a/2 * (b-(b-c)-e/bxa-c))",
		"a/20 * (b-(b-c)-e/bxa-c))",
		"$1/20 * ($2-($3-$4)-$5/$123-$12))",
		"a1/20 * (a2-(a3-a4)-a5/a123-a12))",
	}

	for _, exp := range expressions {
		process(exp)
	}
}

func process(exp string) {
	exp = strings.ReplaceAll(exp, " ", "")
	for i := 0; i < len(exp); {
		if exp[i] == '/' && i < len(exp)-1 {
			if exp[i+1] != '(' {
				j := i + 1
				for ; j < len(exp); j++ {
					if !(exp[j] >= '0' && exp[j] <= '9') && !(exp[j] >= 'a' && exp[j] <= 'z') && exp[j] != '$' {
						break
					}
				}
				d := j - i
				exp = exp[:i+1] + "checkZero(" + exp[i+1:i+d] + ")" + exp[i+d:]
				i = i + len("checkZero(") + d
				continue
			} else {
				exp = exp[:i+1] + "checkZero" + exp[i+1:]
				i = i + len("checkZero")
				continue
			}
		}
		i++
	}
	fmt.Println(exp)
}
