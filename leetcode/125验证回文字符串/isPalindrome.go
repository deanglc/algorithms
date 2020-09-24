package main

import "unicode"

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:
输入: "race a car"
输出: false
imp-yes:根据示例1，只需要考虑字母和数字正反顺序是否一致。 空字符直接忽略*/
func isPalindrome(s string) bool {
	le := len(s)
	if le <= 1 {
		return true
	}
	// 判定是否为 字母或数字
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}

	//最左最右开始
	l, r := 0, le-1
	for l < r {
		vLeft, vRight := rune(s[l]), rune(s[r])
		//两边都不是(字母或数字) 过
		if !isValid(vLeft) && !isValid(vRight) {
			l++
			r--
		} else if !isValid(vLeft) { //左边不是(字母或数字) 左边右移
			l++
		} else if !isValid(vRight) { //右边不是(字母或数字) 右边左移
			r--
		} else if unicode.ToUpper(vLeft) != unicode.ToUpper(vRight) { //两边值不等-->不是回文串 返回false
			return false
		} else { // 继续下一对比较
			l++
			r--
		}

	}
	return true

}
func main() {
	a := isPalindrome("A man, a plan, a canal: Panama")
	println(a)
}
