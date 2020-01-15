package generativerecursion

// 25.1 bundle function implementation
// Bundle takes []string(1-character strings) and int(characters in chunk)
// and produces []string(for each s from []string len(s) <= 3)
// examples:
// 		Bundle([]string{"a","b","c","d"}, 3) => []{"abc", "d"}
// 		Bundle([]string{"a","b","c","d"}, 5) => []{"abcd"}
func Bundle(chars []string, n int) []string {
	return bundleV2(chars, n)
}

func bundleV1(chars []string, n int) []string {
	next := func() string {
		lim := n
		if len(chars) < lim {
			lim = len(chars)
		}

		o := ""
		for i := 0; i < lim; i++ {
			o += chars[i]
		}

		if len(chars) <= n {
			chars = nil
		} else {
			chars = chars[n:]
		}
		return o
	}

	var acc []string

	var s string
	for chars != nil {
		s = next()
		acc = append(acc, s)
	}

	return acc
}

func bundleV2(chars []string, n int) []string {
	l := n
	lc := len(chars)
	if lc < l {
		l = lc
	}

	o := ""
	for i := 0; i < l; i++ {
		o += chars[i]
	}

	result := []string{o}
	if lc <= n {
		return result
	}

	return append(result, bundleV2(chars[n:], n)...)
}

func bundleV3(chars []string, n int) []string {
	if len(chars) == 0 {
		return nil
	}

	return append([]string{take(chars, n)}, bundleV3(drop(chars, n), n)...)
}

func take(chars []string, n int) string {
	switch {
	case n == 0:
		return ""
	case len(chars) == 0:
		return ""
	default:
		return chars[0] + take(chars[1:], n-1)
	}
}

func drop(chars []string, n int) []string {
	switch {
	case n == 0:
		return chars
	case len(chars) == 0:
		return nil
	default:
		return drop(chars[1:], n-1)
	}
}
