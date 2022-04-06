package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/triadmoko/go-cli-loker/handler"
	"github.com/triadmoko/go-cli-loker/helper"
	"github.com/triadmoko/go-cli-loker/model"
)

func ServiceLoker(input int) (res string, err error) {
	var loker = []model.Identitas{}
	if len(loker) > 0 {
		err = errors.New(helper.Message["init_duplicate"])
		return
	}
	loker = make([]model.Identitas, input)

	res = fmt.Sprintf(helper.Message["init_success"], len(loker))
	return
}

func process(userInput []string) bool {
	command, arg1, _ := helper.UnpackInput(userInput)
	command = strings.ToLower(command)
	var err error

	switch command {

	case "init":
		handler.Init(arg1)

	case "exit":
		return false
	case "":
		err = errors.New(helper.Message["empty_command"])
		helper.Result("", err)
	default:
		err = errors.New(helper.Message["invalid_command"])
		helper.Result("", err)
	}

	return true
}
