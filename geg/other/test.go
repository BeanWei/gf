package main

import (
    "fmt"
    "github.com/gogf/gf/g/os/gtime"
    "strconv"
    "strings"
)

func main() {
    //t := gconv.GTime("2010-10-10 00:00:01")
    fmt.Println(strings.ToUpper(strconv.FormatInt(gtime.Nanosecond(), 36)))
}