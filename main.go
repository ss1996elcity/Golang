package main

import "git_project/utils"
import "git_project/wordz"
import "git_project/color"
import "fmt"

func main() {

a_ := utils.MadeUtils()
b_ := wordz.GetWords()
c_ := color.GetColor()
fmt.Printf("MakeUtils = %v, GetWords = %v, GetColor = %v\n", a, b, c)

}
