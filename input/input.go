package entity

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Input() []string {
	newReader := bufio.NewReader(os.Stdin)   // Mengambil input
	input, err := newReader.ReadString('\n') // mengambil 1 field dari input
	// cek error
	if err != nil {
		fmt.Println(errors.New("inputan gagal"))
	}
	inputFields := strings.Fields(input) // menjadikan input ke dalam slice
	return inputFields                   // mengembalikan data
}
