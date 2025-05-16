package parseNumberToTest

import (
	"NumberToText/src/utils"
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

func dismemberHundreds(number int64) []int64 {
	var numbers []int64
	var next int64

	next = number

	for next > 0 {
		numbers = append(numbers, next%1000)
		next = next / 1000
	}

	return utils.ReverseSlice(numbers)
}

func parseValue(number int64) (int, string) {
	thousands := []string{"", "mil", "milhão", "bilhão", "trilhão"}
	buffer := []string{}
	decimalsCounted := 0
	partsOfTheNumbers := dismemberHundreds(number)

	for i := range partsOfTheNumbers {
		part := partsOfTheNumbers[i]
		if part == 0 {
			decimalsCounted += 1
			continue
		}

		index := len(partsOfTheNumbers) - 1 - i

		if !(len(partsOfTheNumbers) == 2 && part == 1) {
			bufferHundred := parseHundreds(int(part))
			buffer = append(buffer, bufferHundred)
		}

		if thousands[index] != "" {
			if part == 1 && (index) > 1 {
				buffer = append(buffer, thousands[index])
			} else {
				if thousands[index] != "mil" {
					var sufaux = ""
					if part > 1 {
						sufaux = "ões"
					}
					// should be -2, but Go works with bytes with special characters taking into account 1 extra byte
					buffer = append(buffer, thousands[index][:len(thousands[index])-3]+sufaux)
				} else {
					buffer = append(buffer, thousands[index])
				}
			}
		}

		if (partsOfTheNumbers[index] > 100) && (len(partsOfTheNumbers) > 1) && ((i + 1) < len(partsOfTheNumbers)) {
			buffer[len(buffer)-1] += ","
		}
	}

	joined := strings.Join(buffer, " ")
	result := strings.TrimSpace(joined)

	return decimalsCounted, result
}

func parseMoneyValue(number int64) (string, error) {
	buffer := []string{}
	decimalsCounted, bufferValue := parseValue(number)
	buffer = append(buffer, bufferValue)

	if number == 1 {
		buffer = append(buffer, "real")
	} else {
		if decimalsCounted >= 2 {
			buffer = append(buffer, "de")
		}

		buffer = append(buffer, "reais")
	}

	joined := strings.Join(buffer, " ")
	result := strings.TrimSpace(joined)

	return result, nil
}

func parseCents(number int) (string, error) {
	buffer := []string{}
	_, bufferValue := parseValue(int64(number))
	buffer = append(buffer, bufferValue)

	if number == 1 {
		buffer = append(buffer, "centavo")
	} else {
		buffer = append(buffer, "centavos")
	}

	joined := strings.Join(buffer, " ")
	result := strings.TrimSpace(joined)

	return result, nil
}

func NumberToText(number int64) (string, error) {
	buffer := []string{}
	value := number / 100
	cents := int(number % 100)

	if value > 0 {
		bufferMoneyValue, _ := parseMoneyValue(value)
		buffer = append(buffer, bufferMoneyValue)
	}

	if cents > 0 {
		if value > 0 {
			buffer = append(buffer, "e")
		}

		bufferCents, _ := parseCents(cents)
		buffer = append(buffer, bufferCents)
	}

	joined := strings.Join(buffer, " ")
	result := strings.TrimSpace(joined)

	return result, nil
}
