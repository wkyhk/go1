package goweek

import (
	"context"
	"fmt"
	"net/url"
)

/*
	func Go153() {
		a := []int{7, 8, 9}
		fmt.Println(a[real(2)])
	}
*/
func Go154() {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	fmt.Println("time out")
}

type Query struct {
	url.Values
}

func Go156() {
	q := Query{}
	q.Values["name"] = []string{"polarisxu"}
	fmt.Println(q.Get("name"))
}

// 157
func Go157() {
	const X = 7.0
	var x interface{} = X
	if y, ok := x.(int); ok {
		fmt.Println(y)
	} else {
		fmt.Println(int(y))
	}
}

/* //158
func Go158() {
	var s fmt.Stringer
	//	s = "s"   //异常
	fmt.Println(s)

} */
//159
func Go159() {
	m := make(map[int]int, 3)
	x := len(m)
	m[1] = m[1]
	y := len(m)
	fmt.Println(x, y)
}

// GO160
type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Print(elem)
	return &temp{}
}
func Go160() {
	tt := &temp{}
	defer tt.Add(1).Add(2)
	tt.Add(3)
}

func Go161() {
	s := []string{"A", "B", "C"}
	t := s[:1]
	fmt.Println(&s[0] == &t[0])
	u := append(s[:1], s[2:]...)
	fmt.Println(&s[0] == &u[0])
}
