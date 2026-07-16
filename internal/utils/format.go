package utils

import "strconv"

// FormatRupiah mengubah 250000 menjadi "250.000"
func FormatRupiah(amount int64) string {
	str := strconv.FormatInt(amount, 10)

	n := len(str)
	if n <= 3 {
		return str
	}

	var result []byte
	for i, c := range []byte(str) {
		if i > 0 && (n-i)%3 == 0 {
			result = append(result, '.')
		}
		result = append(result, c)
	}

	return string(result)
}
