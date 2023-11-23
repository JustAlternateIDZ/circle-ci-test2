package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    moi := Person{"Loïc", 21}
    fmt.Println(SayGoodBye(&moi))
}

func SayGoodBye(p *Person) string {
    return fmt.Sprintf("Good Bye %s", p.Name)
}
