/*
range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.
*/

package main

import "fmt"

func main() {
    // Here we use range to sum the numbers in a slice. Arrays work like this too.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

/*
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
*/
