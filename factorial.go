package main

// 1W的阶乘， 结果是非常大非常大的数字

import "fmt"

func factorial(n int) *bigInt {
	result := new(bigInt).SetInt64(1)
	for i := 2; i <= n; i++ {
		result.Mul(i)
	}
	return result
}

type bigInt struct {
	digits []int
}

func (a *bigInt) SetInt64(n int64) *bigInt {
	a.digits = []int{int(n)}
	return a
}

func (a *bigInt) Mul(b int) *bigInt {
	carry := 0
	for i := 0; i < len(a.digits) || carry > 0; i++ {
		if i == len(a.digits) {
			a.digits = append(a.digits, 0)
		}
		product := a.digits[i]*b + carry
		a.digits[i] = product % 10
		carry = product / 10
	}
	return a
}

func (a *bigInt) String() string {
	result := ""
	for i := len(a.digits) - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d", a.digits[i])
	}
	return result
}
