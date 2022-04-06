package helper

import "fmt"

func Result(res string, err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func UnpackInput(userInput []string) (command string, arg1 string, arg2 string) {
	if len(userInput) == 0 {
		return
	}

	command = userInput[0]

	if len(userInput) > 2 {
		arg1, arg2 = userInput[1], userInput[2]
	} else if len(userInput) > 1 {
		arg1 = userInput[1]
	}
	return
}

var Message = map[string]string{
	"welcome":          "Selamat datang, silahkan membuat loker terlebih dahulu",
	"init_success":     "Berhasil membuat loker dengan jumlah %d",
	"init_duplicate":   "! Loker sudah tersedia, silahkan keluar program untuk membuat baru",
	"loker_slot_empty": "! Loker belum dibuat",
	"input_success":    "Kartu identitas tersimpan di loker nomor %d",
	"input_duplicate":  "Kartu dengan no identitas tersebut sudah ada di loker nomor %d",
	"input_failed":     "! Maaf, loker sudah penuh",
	"leave_success":    "Loker nomor %d berhasil dikosongkan",
	"leave_empty":      "! Maaf, loker nomor %d belum ada isinya",
	"leave_failed":     "! Maaf, tidak ada loker dengan nomor %d",
	"search_empty":     "Tidak ada kartu identitas dengan tipe %s",
	"find_result":      "Kartu identitas tersebut berada di loker nomor %d",
	"find_empty":       "Nomor identitas tidak ditemukan",
	"exit":             "Keluar dari program",
	"empty_command":    "! Maaf, perintah tidak boleh kosong",
	"invalid_command":  "! Maaf, perintah yang anda masukkan salah",
	"empty_argument":   "! Maaf, perintah belum lengkap",
	"invalid_argument": "! Maaf, nilai yang anda masukkan salah",
}
