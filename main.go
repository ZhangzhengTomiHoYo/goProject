package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

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



								//
								// 2.13 结构体方法
								a := user{"zhangzheng", "123"}
								a.resetPassword("2048")
								fmt.Println(a.checkPassword("2048"))



							//
							// 2.14 错误处理
							u, err := findUser([]user{{"wang", "1024"}}, "wang")
							if err != nil {
								fmt.Println(err)
								return
							}
							fmt.Println(u.name)
							// 还可以这样写 !?
							if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
								fmt.Println(err)
								return
							} else {
								fmt.Println(u.name)
							}



						//
						// 2.15 字符串操作
						a := "hello"
						fmt.Println(strings.Contains(a, "ll"))
						fmt.Println(strings.Count(a, "l"))
						fmt.Println(strings.HasPrefix(a, "he"))
						fmt.Println(strings.HasSuffix(a, "llo"))
						fmt.Println(strings.Index(a, "ll"))
						fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
						fmt.Println(strings.Repeat(a, 2))
						fmt.Println(strings.Split("a-b-c", "-"))
						fmt.Println(strings.ToLower(a))
						fmt.Println(strings.ToUpper(a))
						fmt.Println(len(a))
						// 一个中文对应多个字符
						b := "你好"
						fmt.Println(len(b))



					//
					// 2.16 字符串格式化
					s := "hello"
					n := 123
					p := point{1, 2}
					fmt.Println(s, n)
					fmt.Println(p)
					// go 语言中 可以使用 %v打印任何类型的变量
					// 不需要区分数字用%d 字符串%s
					fmt.Printf("s=%v\n", s)
					fmt.Printf("n=%v\n", n)
					fmt.Printf("p=%v\n", p)
					// %+v 更详细
					fmt.Printf("p=%+v\n", p)
					//
					fmt.Printf("p=%#v\n", p)

					f := 3.141592653
					fmt.Println(f)
					fmt.Printf("%.2f\n", f)



				//
				// 2.17 JSON处理
				a := userInfo{Name: "wang", Age: 18, Hobby: []string{"GO", "4dgs"}}
				buf, err := json.Marshal(a)
				if err != nil {
					panic(err)
				}
				fmt.Println(buf)
				fmt.Println(string(buf))
				// Marshal 打包
				// Indent 缩进
				buf, err = json.MarshalIndent(a, "", "\t")
				if err != nil {
					panic(err)
				}
				fmt.Println(string(buf))

				// 定义一个空变量
				var b userInfo
				// 反序列化到这个空变量
				err = json.Unmarshal(buf, &b)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%#v\n", b)



			//
			// 2.18 时间处理
			now := time.Now()
			fmt.Println(now)
			t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
			t2 := time.Date(2025, 3, 21, 15, 38, 22, 0, time.UTC)

			fmt.Println(t)
			fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
			// 格式化不是 YYDD什么什么的 而是用一个字符串
			// 这个时间是写在time包的官方文档上的
			// 另外，注意这个格式化t是上面初始化过的 很神奇吧
			fmt.Println(t.Format("2006-01-02 15:04:05"))
			diff := t2.Sub(t)
			fmt.Println(diff)
			fmt.Println(diff.Minutes(), diff.Seconds())
			// 同样的 用2006-01-02 15:04:05去解析时间
			t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:24:24")
			if err != nil {
				panic(err)
			}
			fmt.Println(t3 == t)
			// 获取一个时间戳， 不知道是干啥的
			fmt.Println(now.Unix())



		//
		// 2.19 数字解析
		f, _ := strconv.ParseFloat("1.234", 64)
		fmt.Println(f)

		n, _ := strconv.ParseInt("111", 10, 64)
		fmt.Println(n)

		n, _ = strconv.ParseInt("0x1000", 0, 64)
		fmt.Println(n)

		n2, _ := strconv.Atoi("123")
		fmt.Println(n2)

		n2, err := strconv.Atoi("AAA")
		fmt.Println(n2, err)

	*/

	fmt.Println(os.Args)
	fmt.Println()
	fmt.Println(os.Getenv("PATH"))
	fmt.Println()
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "127.0.01", "/etc/bosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

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

// 2.14 错误处理
// 单从这个例子看
// 要么 user是有 错误是空
// 要么 user是空 错误是有
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

// 2.16 字符串格式化
// x,y int 在结构体可以  go里定义变量不能这样
type point struct {
	x, y int
}

// 2.17 JSON处理
// 对于一个结构体 保证每个变量首字母大小 就可以使用json.Marshal序列化
// marshal 打包
type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}
