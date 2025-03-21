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

							fmt.Println(s[2:5])
							fmt.Println(s[:5])
							fmt.Println(s[2:])

							good := []string{"g", "o", "0", "d"}
							fmt.Println(good)



						//
						// 2.8 map
						// 在其他语言中 叫做 哈希 / 字典
						// 使用最频繁 (在深度学习项目里确实是这样！！！
						// 内部无序
						m := make(map[string]int)
						m["one"] = 1
						m["two"] = 2
						fmt.Println(m)
						fmt.Println(len(m))
						fmt.Println(m["one"])
						fmt.Println(m["unknow"])

						r, ok := m["unknow"] // ok 可以获取map里究竟有没有这个key存在
						fmt.Println(r, ok)   // 0 false

						delete(m, "one")

						m2 := map[string]int{"one": 1, "two": 2}
						var m3 = map[string]int{"one": 1, "two": 2}
						fmt.Println(m2, m3)



					//
					// 2.9 range
					nums := []int{2, 3, 4}
					sum := 0
					for i, num := range nums {
						sum += num
						if num == 2 {
							fmt.Println("index:", i, "num:", num)
						}
					}
					fmt.Println(sum)

					m := map[string]string{"a": "A", "b": "B"}
					for k, v := range m {
						fmt.Println(k, v)
					}
					for k := range m {
						fmt.Println("key", k)
					}



				//
				// 2.10 函数
				res := add(1, 2)
				fmt.Println(res)

				v, ok := exists(map[string]string{"a": "A"}, "a")
				fmt.Println(v, ok)


			//
			// 2.11 指针
			// 用途有限 仅对函数值进行修改
			n := 5

			add2(n)
			fmt.Println(n)
			add2ptr(&n)
			fmt.Println(n)




		a := user{name: "wang", password: "1024"}
		b := user{"wang", "1024"}
		c := user{name: "wang"}

		c.password = "1024"

		var d user
		d.name = "wang"
		d.password = "1024"

		fmt.Println(a, b, c, d)
		fmt.Println(checkPassword(a, "haha"))
		fmt.Println(checkPassword2(&a, "haha"))

	*/

	//
	// 2.13 结构体方法
	a := user{"zhangzheng", "123"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048"))
}

/*
//
// 2.10 函数
func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok // 返回两个值  第一个是值 第二个是错误信息
}
*/

// 2.11 指针
func add2(n int) {
	// 拷贝
	n += 2
}
func add2ptr(n *int) {
	*n += 2
}

// 2.12 结构体
type user struct {
	name     string
	password string
}

// 不用指针
func checkPassword(u user, password string) bool {
	return u.password == password
}

// 结构体也可以用指针 可以避免一些大结构体 拷贝的开销
func checkPassword2(u *user, password string) bool {
	return u.password == password
}

// 2.13 结构体方法
// 类比上面 就是函数写法 移动一下结构体位置
// 个人思考 ： 和类/对象 相比 没有了 self
func (u user) checkPassword(password string) bool {
	return u.password == password
}

// 带指针可以修改值
func (u *user) resetPassword(password string) {
	u.password = password
}
