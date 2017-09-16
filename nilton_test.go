package programa2

import (
	"testing"
	"strconv"
	"fmt"
)

func TestVacio(t *testing.T ){
	expected := "dsfa"
	actual := validateCode()
	if actual!=expected {
		t.Errorf("valor no esperado")
	}
}

func TestLeer(t *testing.T){
	valor, array:=readFile()
	if valor!=true && len(array)<=0 {
		t.Errorf("El archivo no fue correctamente leido o esta vacio")
	}
}

func TestEncontrarMetodos(t *testing.T){
	lineas:=encontrarMetodos()
	if len(lineas)<=0 {
		t.Errorf("En el archivo no hay metodos")
	}
}

func TestValidarFuncPrimeraPalabra(t *testing.T)  {
	numeroLineas, errores, numeroMetodos :=conteoDeLineas()
	if(len(errores)>0){
		for i:=0;i<len(errores);i++  {
			t.Errorf("la función de la línea "+strconv.Itoa(errores[i])+" no esta acorde al estandar")
		}
	}else {
		fmt.Println("El numero de lineas es: "+strconv.Itoa(numeroLineas))
		texto2:="El numero de metodos son:"+strconv.Itoa(numeroMetodos)
		fmt.Println(texto2)
	}
}

func TestValidarMinusculaNombreMetodo(t *testing.T){
	numeroLineas, errores, numeroMetodos :=conteoDeLineas()
	if(len(errores)>0){
		for i:=0;i<len(errores);i++  {
			t.Errorf("la función de la línea "+strconv.Itoa(errores[i])+" no esta acorde al estandar")
		}
	}else {
		fmt.Println("El numero de lineas es: "+strconv.Itoa(numeroLineas))
		texto2:="El numero de metodos son:"+strconv.Itoa(numeroMetodos)
		fmt.Println(texto2)
	}
}

func TestValidarArgumentos(t *testing.T)  {
	numeroLineas, errores, numeroMetodos :=conteoDeLineas()
	texto:="El numero de lineas son: "+strconv.Itoa(numeroLineas)
	if(len(errores)>0){
		for i:=0;i<len(errores);i++  {
			t.Errorf("la función de la línea "+strconv.Itoa(errores[i])+" no esta acorde al estandar")
		}
	}else {
		fmt.Println(texto)
		texto2:="El numero de metodos son:"+strconv.Itoa(numeroMetodos)
		fmt.Println(texto2)
	}
}