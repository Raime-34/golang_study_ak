package main

import "fmt"

func main() {
	fmt.Println(getFilePermissions(630))
}

func getFilePermissions(flag int) string {
	result := []byte{}

	binaryRep := make([]uint, 9)

	i := 0
	for flag > 0 {
		f := flag % 10
		flag /= 10

		leftBorder := len(binaryRep) - 3 - i
		rightBorder := len(binaryRep) - i
		setBinaryRep(f, binaryRep[leftBorder:rightBorder])
		i += 3
	}

	result = append(result, getPartialPermission("Owner", binaryRep[:3])...)
	result = append(result, getPartialPermission("Group", binaryRep[3:6])...)
	result = append(result, getPartialPermission("Other", binaryRep[6:])...)

	return string(result)
}

// Преобразование числа в его двоичную версию
func setBinaryRep(flag int, binaryRep []uint) {
	for i := 0; i < len(binaryRep); i++ {
		var b uint = 1 << i
		res := uint(flag) & b
		res = res >> i
		binaryRep[len(binaryRep)-1-i] = res
	}
}

// Заполнение разрешений для каждого типа пользователя
func getPartialPermission(prefix string, flag []uint) []byte {
	var (
		read  = "-"
		write = "-"
		exec  = "-"
	)

	if flag[0] == 1 {
		read = "Read"
	}

	if flag[1] == 1 {
		write = "Write"
	}

	if flag[2] == 1 {
		exec = "Execute"
	}

	return []byte(fmt.Sprintf(
		"%s:%s,%s,%s ",
		prefix, read, write, exec,
	))
}
