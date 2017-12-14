package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(u)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointerq
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	q := &v
	(*q).X = 1e9
	q.X = 1

	fmt.Println(v)


	asm := []string{"hello", "world", "ww"}   //也可以make创建 make(int[], 3, 3) 类型 长度 容量
	//asm[3] = "111" 会出错， 虽然不限长度，但其实在定义的时候定义了容量的
	fmt.Println(asm[0], asm[1])
	fmt.Println(asm)

	var pp []int
	fmt.Println("pp", pp, len(pp), cap(pp))

	// append works on nil slices.
	pp = append(pp, 0)
	fmt.Println("pp", pp, len(pp), cap(pp))

	for i, v := range asm {
		fmt.Printf("2**%d = %s\n", i, v)
	}

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	o, ok := m["Answer"]
	fmt.Println("The value:", o, "Present?", ok)

	aa := &Vertex1{3, 4}
	fmt.Println(aa.Abs())
	fmt.Println((*aa).Abs())

	var aaa Abser
	vvv := Vertex1{3, 4}
	aaa = vvv // aaa Vertex1 实现了 Abser
	aaa = &vvv // aaa *Vertex1 实现了 Abser，说明定义在值上的方法，默认它也算定义在指针类型上，但是反之，定义在指针类型上，并不算定义在了值类型上，虽然值类型可以使用

	fmt.Println(aaa.Abs())

	ama := Person{"Arthur Dent", 42}
	zmz := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(ama, zmz)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//接口是方法的集合
type Abser interface {
	Abs() float64
}

type Vertex1 struct {
	X, Y float64
}
//定义v上的Abs方法
//定义在指针和值上的区别在于，定义在指针上v指向原始值，而定义在值上v指向获取值的拷贝，即内部v的区别，但是不管定义在哪里，指针和值都能用这个方法，但方法的效果都是基于定义时来的，即里面的v的意义是根据定义时来的
func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//会报错，说重复声明
//func (v *Vertex1) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}


type Vertex struct {
	X int
	Y int
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (m ,n string) {
	m = y
	n = x
	return
}