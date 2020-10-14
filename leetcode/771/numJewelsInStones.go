package _71

func numJewelsInStones1(J string, S string) int {
	m := make(map[byte]int8)
	for _, v := range J {
		m[byte(v)] = 0
	}
	var sum int8

	for _, v := range S {
		if _, ok := m[byte(v)]; ok {
			sum++
		}
	}
	return int(sum)
}
func numJewelsInStones(J string, S string) (res int) {
	m := make(map[byte]struct{}, len(J))
	for i := range J {
		m[J[i]] = struct{}{}
	}
	for i := range S {
		if _, ok := m[S[i]]; ok {
			res++
		}
	}
	return res
}
