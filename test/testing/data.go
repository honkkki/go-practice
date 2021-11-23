package testing

var cache = make(map[int]int)

func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib1(n-2) + fib1(n-1)
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}

	n1 := 0
	n2 := 1
	for i := 3; i <= n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
	}

	return n2
}

func fib3(n int) int {
	if n <= 1 {
		return n
	}

	if num, ok := cache[n]; ok {
		return num
	} else {
		res := fib3(n-2) + fib3(n-1)
		cache[n] = res
		return res
	}
}
