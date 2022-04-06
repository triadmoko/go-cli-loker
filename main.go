package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/triadmoko/go-cli-loker/handler"
	"github.com/triadmoko/go-cli-loker/helper"
	entity "github.com/triadmoko/go-cli-loker/input"
)

func main() {
	fmt.Println(`=============================================`)
	fmt.Println(`     Selamat Datang di Loker ID Card         `)
	fmt.Println(`=============================================`)
	i := 0
	for {
		cmdinput := entity.Input()
		if !argumentInput(cmdinput) {
			fmt.Println(errors.New("input gagal"))
			break
		}
		i++
	}
	fmt.Println(`=============================================`)
	fmt.Println(`     Sampai Jumpa Kembali di Loker ID Card   `)
	fmt.Println(`=============================================`)
}

func argumentInput(input []string) bool {
	cmd, arg1, arg2 := helper.SplitArgument(input)
	cmd = strings.ToLower(cmd)

	switch cmd {

	case "init":
		handler.Init(arg1)
	case "status":
		handler.Status()
	case "input":
		handler.Input(arg1, arg2)
	case "leave":
		handler.Leave(arg1)
	case "find":
		handler.Find(arg1)
	case "search":
		handler.Search(arg1)
	case "exit":
		return false
	case " ":
		err := errors.New("input kosong")
		helper.Result("", err)
	default:
		err := errors.New("input gagal")
		helper.Result("", err)
	}

	return true
}
