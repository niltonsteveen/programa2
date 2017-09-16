package programa2

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func abrirArchivo(s_nombreArchivo string) string{
	dat, err := ioutil.ReadFile(s_nombreArchivo)
	if err != nil {
		return ""
	}
	return string(dat)
}

func abrirArchivo2(nombre string) []string{
	file, err := os.Open(nombre)
	var array []string
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		array = append(array, strings.Trim(scanner.Text(), " "))
	}
	return  array
}
