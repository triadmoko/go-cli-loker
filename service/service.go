package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/triadmoko/go-cli-loker/entity"
)

var loker []entity.Identitas

func Init(sum int) (res string, err error) {

	if len(loker) > 0 {
		err = errors.New("loker sudah ada gays")
		return
	}
	loker = make([]entity.Identitas, sum)

	fmt.Println("loker sukses di buat sebanyak : ", len(loker))
	return
}
func Status() (result string, err error) {
	if len(loker) < 1 {
		err := errors.New("loker sudah ada gays")
		return " ", err
	}
	tableRows := []string{
		"============================================",
		"No Loker\tTipe Identitas\tNo Identitas",
		"--------------------------------------------",
	}
	for i, kartu := range loker {
		if kartu != (entity.Identitas{}) {
			tableRows = append(tableRows, fmt.Sprintf("%d\t\t%s\t\t%d", i+1, kartu.Type, kartu.ID))
		}
	}
	res := strings.Join(tableRows, "\n")

	return res, nil
}

func AddCard(Type string, ID uint) (res string, err error) {

	if len(loker) < 1 {
		err := errors.New("Loker Belum Ada")
		return " ", err
	}
	for _, v := range loker {
		if v.ID == ID {
			err := errors.New("Gagal Menambahkan Identitas")
			return " ", err
		}
	}

	Type = strings.ToUpper(Type)

	lokerId := 0
	for i, v := range loker {
		if v == (entity.Identitas{}) {
			loker[i] = entity.Identitas{Type, ID}
			lokerId = i + 1
			break
		}
	}

	if lokerId == 0 {
		err = errors.New("loker sudah penuh")
		return " ", err
	}
	res = strconv.Itoa(lokerId)
	return res, nil
}

func Leave(ID int) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("loker belum di atur")
		return "", err
	}

	if ID > len(loker) {
		// nomor loker tidak tersedia
		err = errors.New(fmt.Sprintf("nomor loker tidak tersedia", ID))
		return
	}
	if loker[ID-1] == (entity.Identitas{}) {
		// nomor loker tersedia tapi kosong
		err = errors.New(fmt.Sprintf("nomor loker tersedia tapi kosong", ID))
		return
	}

	loker[ID-1] = entity.Identitas{}

	fmt.Println("Berhasil Mengosongkan loker", ID)
	return
}

func Find(ID int) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("loker belum di atur")
		return
	}

	for i, card := range loker {
		if int(card.ID) == ID {
			fmt.Println("Hasil : ", i+1)
			return
		}
	}
	res = "data kosong"
	return
}

func Search(Type string) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("loker belum di atur")
		return
	}

	Type = strings.ToUpper(Type)

	var searchResults []string
	for _, card := range loker {
		if card.Type == Type {
			searchResults = append(searchResults, fmt.Sprintf("%d", card.ID))
		}
	}

	if len(searchResults) == 0 {
		err = errors.New(fmt.Sprintf("Data tidak ditemukan", Type))
		return
	}

	res = fmt.Sprintf("No Identitas: %s", strings.Join(searchResults, ", "))
	return
}
