package helper

import "fmt"

func SplitArgument(input []string) (command string, arg1 string, arg2 string) {
	if len(input) == 0 { //cek total input, jika nol ulangi
		return
	}

	command = input[0]

	if len(input) > 2 {
		arg1, arg2 = input[1], input[2]
	} else if len(input) > 1 {
		arg1 = input[1]
	}
	return
}

func Result(res string, err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
