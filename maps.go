package main

import "fmt"

func main() {
    // To create an empty map, use the builtin make: make(map[key-type]val-type)
    m := make(map[string]int)

    m["k1"] = 7 // Set key/value pairs using typical name[key] = val syntax.
    m["k2"] = 13

    fmt.Println("map:", m) // Printing a map with e.g. fmt.Println will show all of its key/value pairs.

    v1 := m["k1"] // Get a value for a key with name[key].
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m)) // The builtin len returns the number of key/value pairs when called on a map.

    delete(m, "k2") // The builtin delete removes key/value pairs from a map.
    fmt.Println("map:", m)
    // The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
    _, prs := m["k2"] 
    fmt.Println("prs:", prs)
    
    // You can also declare and initialize a new map in the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

/*
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
*/
