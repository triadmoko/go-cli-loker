package handler

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/triadmoko/go-cli-loker/helper"
	"github.com/triadmoko/go-cli-loker/service"
)

func Init(cmd string) {
	// cek apa input kosong atau tidak
	if cmd == "" {
		err := errors.New("input gagal")
		helper.Result(cmd, err)
	}

	// convert string to int

	kapasitasLoker, err := strconv.Atoi(cmd)
	if err != nil || kapasitasLoker < 1 {
		err := errors.New("input gagal")
		helper.Result(cmd, err)
	}

	res, err := service.Init(kapasitasLoker)
	helper.Result(res, err)
}
func Status() {
	result, err := service.Status()
	helper.Result(result, err)
}

func Input(typeCard, ID string) {

	if typeCard == "" || ID == "" {
		err := errors.New("input gagal")
		helper.Result("", err)
		return
	}

	regex, _ := regexp.Compile(`^[a-zA-Z]+$`)
	cardTypeIsAlphabet := regex.MatchString(typeCard)

	numberID, err := strconv.Atoi(ID)
	uintID := uint(numberID)

	if err != nil || numberID < 1 || !cardTypeIsAlphabet {
		err := errors.New("input gagal")
		helper.Result("", err)
		return
	}

	res, err := service.AddCard(typeCard, uintID)
	helper.Result(res, err)
}

func Leave(No string) {

	var res string
	var err error
	var lokerId int

	if No == "" {
		err = errors.New("argument tidak ditemukan")
		helper.Result(res, err)
		return
	}

	lokerId, err = strconv.Atoi(No)
	if err != nil || lokerId < 1 {
		err = errors.New("argument tidak ditemukan")
		helper.Result(res, err)
		return
	}

	res, err = service.Leave(lokerId)
	helper.Result(res, err)
}
func Find(ID string) {

	var res string
	var err error
	var cardNumber int

	if ID == "" {
		err = errors.New("argument tidak ditemukan")
		helper.Result(res, err)
		return
	}

	cardNumber, err = strconv.Atoi(ID)
	if err != nil || cardNumber < 1 {
		err = errors.New("argument tidak ditemukan")
		helper.Result(res, err)
		return
	}

	res, err = service.Find(cardNumber)
	helper.Result(res, err)
}

func Search(cardType string) {
	// Fungsi untuk menangani perintah `search`

	var res string
	var err error

	if cardType == "" {
		err = errors.New("argument tidak ditemukan")
		helper.Result(res, err)
		return
	}

	res, err = service.Search(cardType)
	helper.Result(res, err)
}
