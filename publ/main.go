package main

import (
	"fmt"
)

func title() {
	fmt.Println(`             _     _ _      _     _   `)
	fmt.Println(`            | |   | (_)    (_)   | |  `)
	fmt.Println(` _ __  _   _| |__ | |_  ___ _ ___| |_ `)
	fmt.Println(`| '_ \| | | | '_ \| | |/ __| / __| __|`)
	fmt.Println(`| |_) | |_| | |_) | | | (__| \__ \ |_ `)
	fmt.Println(`| .__/ \__,_|_.__/|_|_|\___|_|___/\__|`)
	fmt.Println(`| |                                   `)
	fmt.Println(`|_|                                   `)
}
func version() {
}

func main() {
	title()
	fmt.Printf("publicist v0.0 build <DATE>\n")
}
