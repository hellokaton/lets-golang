package main

import "fmt"

func basic() {
    // == !=
    foo := map[int]string{1: "biezhi", 2: "rose"}
    fmt.Println(foo)
    emptyMap := map[int]int{}
    fmt.Println(emptyMap)
    var nilMap map[int]int
    fmt.Println(nilMap)
    bar := make(map[string]string, 10)
    bar["name"] = "biezhi"
    bar["title"] = "代码真香"
    bar["url"] = "https://youtube.com/biezhi"
    fmt.Println(bar["title"])
    fmt.Println(bar["title2"])
    if v, ok := bar["title2"]; ok {
        fmt.Println(v)
    }
    delete(bar, "name")
    fmt.Println(bar)
    fmt.Println(len(bar))
    for k, v := range bar {
        fmt.Println(k, "=", v)
    }
}

func valueNotAddressable() {
    foo := map[int]*[3]int{}
    foo[1] = &[3]int{1, 2, 3}
    fmt.Println(foo[1])
    foo[1][2] = 6
    fmt.Println(foo[1])
}

func change(m map[int]string)  {
    m[1] = "ddog"
}

func main() {
    foo := map[int]string{1:"biezhi"}
    fmt.Println(foo)
    change(foo)
    fmt.Println(foo)
}