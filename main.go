package main

import "fmt"

func main() {
	/*
					//
					// 2.1 hello world
					fmt.Println("Hello World!")

					//
					// 2.2 变量
					var a = "initial"
					var b, c int = 1, 2
					var d = true

					var e float64
					f := float32(e)

					g := a + "foo"
					fmt.Println(a, b, c, d, e, f, g)
					fmt.Println(g)

					const s string = "constant"
					const h = 500000
					const i = 3e20 / h
					fmt.Println(s, h, i, math.Sin(h), math.Cos(h))

					//
					// 2.3 if else
					if 7%2 == 0 {
						fmt.Println("7 is even/偶数")
					} else {
						fmt.Println("7 is odd/奇数")
					}

					if 8%4 == 0 {
						fmt.Println("8 is even/偶数")
					}

				//
				// 2.4 循环
				i := 1
				for {
					fmt.Println("loop")
					break
				}
				for j := 8; j < 9; j++ {
					fmt.Println(j)
				}

				for n := 0; n < 5; n++ {
					if n%2 == 0 {
						continue
					}
					fmt.Println(n)
				}

				for i <= 3 {
					fmt.Println(i)
					i = i + 1
				}



			//
			// 2.5 switch
			// 不需要写break cpp需要
			// 更强大 可以case很多类型
			a := 2
			switch a {
			case 1:
				fmt.Println("one")
			case 2:
				fmt.Println("two")
			default:
				fmt.Println("other")
			}
			// 可以直接当if使用
			// 可以代替多个if else
			t := time.Now()
			switch {
			case t.Hour() < 12:
				fmt.Println("Good morning")
			default:
				fmt.Println("Good afternoon")
			}



		//
		// 2.6 数组
		// 真实业务很少用数组 因为他的长度是固定的
		var a [5]int
		a[4] = 100
		fmt.Println(a[4], len(a))

		b := [5]int{1, 2, 3, 4, 5}
		fmt.Println(b)

		var twoD [2][3]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)

	*/

	//
	// 2.7 切片
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	fmt.Println("\ntest empty 'c'")
	fmt.Println(c)
	fmt.Println(len(c))
	copy(c, s)
	fmt.Println("\ntest copy 'c'")
	fmt.Println(c)
}
