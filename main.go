package main

import "git_project/utils"
import "git_project/wordz"
import "git_project/color"
import "fmt"

func main() {

a := utils.MadeUtils()
b := wordz.GetWords()
c := color.GetColor()
fmt.Printf("MakeUtils = %v, GetWords = %v, GetColor = %v\n", a, b, c)

}
