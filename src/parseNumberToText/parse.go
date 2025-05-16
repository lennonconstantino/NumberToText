package parseNumberToTest

import (
	"fmt"
	"strings"
)

func parseHundreds(number int) string {
	if number == 0 {
		return ""
	}

	units := []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
	tens_lt_20 := []string{"dez", "onze", "doze", "treze", "quatorze", "quinze", "dezesseis", "dezessete", "dezoito",
		"dezenove"}
	tens_gt_20 := []string{"", "", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
	hundreds := []string{"cem", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos"}
	buffer := []string{}

	if number == 100 {
		return hundreds[0]
	}

	if number > 100 {
		buffer = append(buffer, hundreds[number/100])
		if number%100 > 0 {
			buffer = append(buffer, "e")
		}
	}

	var ten = number % 100
	var unit = number % 10

	if ten >= 10 && ten < 20 {
		buffer = append(buffer, tens_lt_20[ten-10])
	} else {
		if ten >= 20 {
			buffer = append(buffer, tens_gt_20[ten/10])
			if unit > 0 {
				buffer = append(buffer, "e")
			}
		}
		if unit > 0 {
			buffer = append(buffer, units[unit])
		}
	}

	joined := strings.Join(buffer, " ")
	result := strings.TrimSpace(joined)
	return result
}

func NumberToText(number int64) string {
	buffer := ""
	if number < 1000 {
		buffer = parseHundreds(int(number))
	}

	result := fmt.Sprintf("%s", buffer)
	return result
}
