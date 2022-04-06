package handler

import (
	"errors"
	"strconv"

	"github.com/triadmoko/go-cli-loker/helper"
	"github.com/triadmoko/go-cli-loker/service"
)

func Init(input string) {
	var (
		res      string
		err      error
		lokerCap int
	)

	if input == "" {
		err = errors.New(helper.Message["empty_argument"])
		helper.Result(res, err)
		return
	}

	lokerCap, err = strconv.Atoi(input)
	if err != nil || lokerCap < 1 {
		err = errors.New(helper.Message["invalid_argument"])
		helper.Result(res, err)
		return
	}

	res, err = service.ServiceLoker(lokerCap)
	helper.Result(res, err)
}
