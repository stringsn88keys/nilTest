package testy

import "fmt"

type foo struct {
}

func fooMeOnceNil() []foo {
	return nil
}

func fooMeOnceArray() []foo {
	return []foo{}
}

func fooMeOncePointerWithNil() *foo {
	return nil
}
func fooMeOncePointerToArrayWithNil() *[]foo {
	return nil
}

func main() {
	fmt.Println(fooMeOnceNil())
}
