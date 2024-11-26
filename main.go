// Задача 3.12 из учебника "Язык программирования Go - Алан А.А. Донован, Брайан У. Керниган"
// Напишите функцию, которая сообщает, являются ли две строки
// анаграммами одна другой, т.е. состоят ли они из одних и тех же букв в другом порядке

package main

import (
	"fmt"
	"sort"
)

func areAnagrams(s, ss string) bool {
	slice1 := []rune(s)
	slice2 := []rune(ss)

	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })

	return string(slice1) == string(slice2)
}

func main() {
	fmt.Println(areAnagrams("listen", "silent")) // true
	fmt.Println(areAnagrams("hello", "world"))   // false
}
