package main

import (
	"flag"
	"fmt"
)

// Command-line flags.
var (
	a = flag.Int("a", 13, "значення a у рівняннях")
	b = flag.Int("b", 2, "значення b у рівняннях")
	m = flag.Int("m", 53, "модуль за яким будуть вестись розрахунки")
)

func amodmx(a, m int) int {
	// x = a mod m
	return a % m
}

func apowbmodmx(a, b, m int) int {
	// x = a^b mod m
	x := 1
	for i := 1; i <= b; i++ {
		x = (x * a) % m
	}
	return x
}

func axbmodm(a, b, m int) int {
	// ax = b (mod n)
	// x = b * a^f(n)-1 (mod n)
	if gcd(a, m) != 1 {
		panic("can't solve")
	}
	x := b * apowbmodmx(a, phi(m)-1, m) % m
	return x
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	for b != 0 {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func phi(n int) int {
	r := 1
	for i := 2; i < n; i++ {
		if gcd(i, n) == 1 {
			r++
		}
	}
	return r
}

// genPrimes генерує прості числа у діапазоні [А, Б].
func genPrimes(A, B int) []int {
	primes := make([]int, 0)
	// Ігноруємо 0, 1 тому що вони не прості числа.
	if A <= 2 {
		A = 2
		// Оброблюємо 2 вручну.
		if B >= 2 {
			primes = append(primes, A)
			A++
		}
	}
	// Проходимо лише через непарні числа
	if A%2 == 0 {
		A++
	}
	for i := A; i <= B; i += 2 {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	flag.Parse()

	fmt.Println("1. Задаємо модуль за допомогою флагу -m.")
	m := *m

	a := *a
	b := *b

	fmt.Printf("Значення: m=%d, a=%d, b=%d\n", m, a, b)
	fmt.Println()

	fmt.Println("2. Розв'язуємо рівняння виду a mod m = x.")
	x := amodmx(a, m)

	fmt.Printf("%d mod %d = x\n", a, m)
	fmt.Printf("x = %d\n", x)
	fmt.Println()

	fmt.Println("3. Розв'язуємо рівняння виду a^b mod m = x.")
	x2 := apowbmodmx(a, b, m)

	fmt.Printf("%d^%d mod %d = x\n", a, b, m)
	fmt.Printf("x = %d\n", x2)
	fmt.Println()

	fmt.Println("4. Розв'язуємо рівняння виду a*x ≡ b mod m.")
	x3 := axbmodm(a, b, m)

	fmt.Printf("%d*x = %d mod %d\n", a, b, m)
	fmt.Printf("x = %d\n", x3)
	fmt.Println()

	fmt.Println("5. Генеруємо просте число у діапазоні від A до B.")
	A, B := 24, 35
	primes := genPrimes(A, B)
	fmt.Printf("Просте число у діапазоні [%d, %d] = %d\n", A, B, primes[0])
}
