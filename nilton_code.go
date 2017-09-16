package programa2

import (
"strings"
)


func validateCode() string  {
	return "dsfa"
}

func readFile() (bool,[]string){
	array:=abrirArchivo2("archivo.txt")
	if(len(array)>0){
		return true,array
	}else {
		return false,array
	}
}

func encontrarMetodos() []int {
	var lineas []int
	array:=abrirArchivo2("archivo.txt")
	for i:=0;i<len(array) ;i++  {
		if strings.Index(array[i],"func")!=-1{
			lineas=append(lineas, i)
		}
	}
	return lineas
}

func conteoDeLineas() (int,[]int, int) {
	var numeroLineas []int
	minusculas:="abcdefghijklmnñopqrstuvwxyz"
	var errores []int
	logico, array:=readFile()
	arrayLineas:=encontrarMetodos()
	if logico {
		for i:=0;i<len(arrayLineas) ;i++  {
			indice:=strings.Index(array[arrayLineas[i]],"func")
			substring := array[arrayLineas[i]][strings.IndexAny(array[arrayLineas[i]], " ")+1:len(array[arrayLineas[i]])]
			primeraLetra := substring[0]
			substring2 := substring[strings.IndexAny(substring, "("):len(substring)]
			if indice!=0 || strings.IndexAny(minusculas,string(primeraLetra))==-1 || len(substring2)-1!=strings.IndexAny(substring2,"{") {
				errores=append(errores,arrayLineas[i])
			}else {
				if strings.IndexAny(substring2[1:len(substring2)],")")==-1{
					errores=append(errores,arrayLineas[i])
					break
				}else {
					substring3 := substring2[strings.IndexAny(substring2, ")")+1:len(substring2)]
					if(strings.IndexAny(string(substring3[0]),"(")!=-1){
						retorno:=substring3[1:strings.IndexAny(substring3, ")")]
						if len(retorno)<=0 {
							errores=append(errores,arrayLineas[i])
							break
						}else {
							numeroLineas=append(numeroLineas, contar(array, arrayLineas[i]))
						}
					}else {
						numeroLineas=append(numeroLineas, contar(array, arrayLineas[i]))
					}
				}
			}

		}
	}
	suma:=0
	for i:=0;i<len(numeroLineas);i++ {
		suma=suma+numeroLineas[i]
	}
	return suma, errores, len(arrayLineas)
}

func contar(array []string,arrayLineas int) int{
	resultado:=1
	contadorDeLlaves:=1
	i:=arrayLineas+1
	for contadorDeLlaves>0{
		resultado++
		if strings.IndexAny(array[i],"{")!=-1{
			contadorDeLlaves++
		}
		if strings.IndexAny(array[i],"}")!=-1{
			contadorDeLlaves--
		}
		i++
	}
	return resultado
}


