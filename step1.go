package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var a, b, c int //默认值为0
	fmt.Printf("%d,%d,%d\n", a, b, c)
	a, b, c = 1, 2, 3 //多重赋值
	fmt.Printf("%d,%d,%d\n", a, b, c)

	a, b, c = c+10, b+10, a+10 //过程中不影响
	fmt.Printf("%d,%d,%d\n", a, b, c)

	s := "hello"
	fmt.Println(s)
	s = ""
	//s = ''//错误

	e := true
	e = false

	f := 1.2

	g := 1.2 + 3.4i

	h := [...]int{1, 2, 3} //数组
	hh := []int{1, 2, 3}   //切片
	hh = h[:]
	hh = make([]int, 3, 5)
	//hh[4] = 4 //超出len限制
	hhh := append(hh, 4, 5, 6) //自动扩充容量

	j := map[string]string{"a": "b"}
	for i, v := range hh {
		fmt.Printf("%d => %d\n", i, v)
	}
	for i, v := range hhh {
		fmt.Printf("%d => %d\n", i, v)
	}
	for k, v := range j {
		fmt.Printf("%s => %s\n", k, v)
	}

	const PI = 3.14 //不能用 const PI := 3.14
	const (
		Monday = iota
		Tuesday
	)

	fmt.Println(a, b, c, e, f, g, PI)
	fmt.Println(Monday, Tuesday) //0 ,1

	fmt.Println(a, s)
	fa(a, s)
	fmt.Println(a, s)

	t := 11
	var p *int
	p = &t

	*p = *p + 1
	fmt.Println(t, *p)
	t += 10
	t++
	fmt.Println(t, *p)
	fmt.Println(&t, p)

}

func fa(aa int, bb string) (int, string) {
	aa = aa + 1
	bb = bb + "22"
	fmt.Println(aa, bb)
	var f func(int) int
	f = func(a int) int { //匿名函数
		return a + 1
	}
	aa = f(aa)
	return aa, bb
}
