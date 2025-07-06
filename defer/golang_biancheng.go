package main

import (
	"fmt"
	"reflect"
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24}, {Name: "li", Age: 23},
		{Name: "wang", Age: 22}}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}

func localVariable() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			time.Sleep(time.Second)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
}

//type People struct {
//}
//
//func (p *People) ShowA() {
//	fmt.Println("showA")
//	p.ShowB()
//}
//func (p *People) ShowB() {
//	fmt.Println("showB")
//}
//
//type Teacher struct {
//	People
//}
//
//func (t *Teacher) ShowB() {
//	fmt.Println("teachershowB")
//}

func select_eval() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//type People interface {
//	Speak(string) string
//}
//type Stduent struct {
//}
//
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "bitch" {
//		talk = "Youare a good boy"
//	} else {
//		talk = "hi"
//	}
//	return
//}

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {
}
func live() People {
	var stu *Student
	return stu
}

//func GetValue() int {
//	return 1
//}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

//func GetValue(m map[int]string, id int) (string, bool) {
//	if _, exist := m[id]; exist {
//		return "存在数据", true
//	}
//	return nil, false
//}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1

type MyStruct struct {
	T1
	T2
}

//func test() []func() {
//	var funs []func()
//	for i := 0; i < 2; i++ {
//		funs = append(funs, func() {
//			println(&i, i)
//		})
//	}
//	return funs
//}

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func panicFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic("deferpanic")
	}()
	panic("panic")
}

func panicFunc1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func printNum() {
	var counter int32
	go func() {
		for {
			curCounter := atomic.LoadInt32(&counter)
			if curCounter > 100 {
				return
			}
			if curCounter%2 == 0 {
				fmt.Println("g1 =", curCounter)
				atomic.CompareAndSwapInt32(&counter, curCounter, curCounter+1)
			}
		}
	}()

	go func() {
		for {
			curCounter := atomic.LoadInt32(&counter)
			if curCounter > 100 {
				return
			}
			if curCounter%2 == 1 {
				fmt.Println("g2 =", curCounter)
				atomic.CompareAndSwapInt32(&counter, curCounter, curCounter+1)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func main() {

	a := []int{1, 3, 2, 5, 3}
	slices.SortFunc(a, func(a, b int) int {
		cmpResult := a < b
		if cmpResult {
			return -1
		} else {
			return 1
		}
	})
	fmt.Println(a)

	//printNum()

	//fmt.Println("return:", *(c()))
	//panicFunc1()

	//a, b := test(100)
	//a()
	//b()

	//funs := test()
	//for _, f := range funs {
	//	f()
	//}

	//my := MyStruct{}
	//my.m1()

	//var i1 MyUser1
	//var i2 MyUser2
	//i1.m1()
	//i2.m2()

	//intmap := map[int]string{
	//	1: "a",
	//	2: "bb",
	//	3: "ccc",
	//}
	//v, err := GetValue(intmap, 3)
	//fmt.Println(v, err)

	//list := new([]int)
	//var a int = 1
	//list = append(list, &a)
	//fmt.Println(list)

	//println(DeferFunc1(1))
	//println(DeferFunc2(1))
	//println(DeferFunc3(1))

	//deferCall()
	//localVariable()
	//time.Sleep(1 * time.Minute)

	//t := Teacher{}
	//t.ShowA()

	//select_eval()

	//a := 1
	//b := 2
	//defer calc("1", a, calc("10", a, b))
	//a = 0
	//defer calc("2", a, calc("20", a, b))
	//b = 1

	//s := make([]int, 5)
	//s = append(s, 1, 2, 3)
	//fmt.Println(s)

	//var peo People = &Stduent{}
	//think := "bitch"
	//fmt.Println(peo.Speak(think))

	//if live() == nil {
	//	fmt.Println("AAAAAAA")
	//} else {
	//	fmt.Println("BBBBBBB")
	//}

	//i := GetValue()
	//switch i.(type)
	//{
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//	default:
	//		println("unknown")
	//}

	//sn1 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//sn2 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//if sn1 == sn2 {
	//	fmt.Println("sn1 == sn2")
	//}
	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//if sm1 == sm2 {
	//	fmt.Println("sm1== sm2")
	//}

}
