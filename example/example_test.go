package person

import (
	"fmt"
	"strings"
	"reflect"
)

func ExamplePerson1_SetName() {
	p := &Person1{
		Id:   1939,
		Name: "Batman",
	}
	fmt.Println(checkMethodNotExists(p, "SetId"))

	fmt.Println(strings.TrimSpace(p.String()))
	p.SetName("Bruce Wayne")
	fmt.Println(strings.TrimSpace(p.String()))
	// Output:
	// true
	// id:1939 name:"Batman"
	// id:1939 name:"Bruce Wayne"
}

func ExamplePerson2_SetName() {
	p := &Person2{
		Id:   1939,
		Name: "Batman",
	}
	fmt.Println(checkMethodNotExists(p, "SetId"))

	fmt.Println(strings.TrimSpace(p.String()))
	p.SetName("Bruce Wayne")
	fmt.Println(strings.TrimSpace(p.String()))
	// Output:
	// true
	// id:1939 name:"Batman"
	// id:1939 name:"Bruce Wayne"
}

func ExamplePerson3_SetId() {
	p := &Person3{
		Id:   1929,
		Name: "Batman",
	}
	fmt.Println(checkMethodNotExists(p, "SetName"))

	fmt.Println(strings.TrimSpace(p.String()))
	p.SetId(1910)
	fmt.Println(strings.TrimSpace(p.String()))
	// Output:
	// true
	// id:1929 name:"Batman"
	// id:1910 name:"Batman"
}

func ExamplePerson4_NoSetters() {
	p := &Person4{
		Id:   1929,
		Name: "Batman",
	}
	fmt.Println(checkMethodNotExists(p, "SetId"))
	fmt.Println(checkMethodNotExists(p, "SetName"))
	// Output:
	// true
	// true
}

func checkMethodNotExists(obj interface{}, name string) bool {
	return reflect.ValueOf(obj).Elem().MethodByName(name) == reflect.Value{}
}
