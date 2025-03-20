package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	k, t, n := 0, 2*numRows-2, len(s)
	out := make([]byte, n)

	for row := range numRows {
		for i := 0; ; i++ {
			index1 := i*t + row
			if index1 < n {
				out[k] = s[index1]
				k++
			} else {
				break // no more characters in this row
			}

			if row != 0 && row != numRows-1 {
				index2 := (i+1)*t - row
				if index2 < n {
					out[k] = s[index2]
					k++
				} else {
					break // no more characters in this row
				}
			}
		}
	}

	return string(out)
}
