var numeralToNum map[string]int = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	num, i := 0, 0

	for i < len(s) {
		a, b := string(s[i]), ""
		if i+1 < len(s) {
			b = string(s[i+1])
		}

		if val, ok := numeralToNum[a+b]; ok {
			num += val
			i += 2
		} else {
			num += numeralToNum[a]
			i += 1
		}
	}

	return num
}