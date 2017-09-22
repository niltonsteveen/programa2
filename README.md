# Contador de líneas de código
### Pruebas de Software 2017-2
#### Universidad de Antioquia
#### Estado de las pruebas ![Estado](https://circleci.com/gh/camigomez35/programa2-pruebasSoftware.svg?style=shield&circle-token=8cc112e65a14c168467af4b8367ccad49bcce930)

Este proyecto está basado en el "Personal Software Process (PSP) for Engineers: Part I"

### Metodología de desarrollo 💻
Test-driven development (TDD)

### Objetivo
Diseñar programa que cuente la cantidad de líneas del código cumpliento estándar de codificación.

### Pasos
> Se lee el código como archivo .txt

### Estándar de conteo
- Se contarán las líneas de código por partes (Lineas, métodos, clases).
- Cada librería en el import cuenta como una linea de código.º
- Tabla discriminante

**_Programa 1_**

| | Clase A | Clase B |
| --- | --- | --- |
| Número de métodos | Número de métodos de A | Número de métodos de B |
| Número de líneas | Número de líneas de A | Número de líneas de B |

Estimar el esfuerzo (Me voy a demorar 3 horas en hacer x líneas de código)
Cuando acabe, toma apunte de los resultados (En realidad me demoré 6 horas)


### Estándar de codificación:
#### Para variables:
- Las variables se nombran en español
- Las variables iniciarán en minúscula y palabras adicionales inician con mayúscula.
- Las variables deben ser nombradas con un identificador inicial según su tipo.
```go
i_contador int := 14
```
El identificador según el tipo se define a continuación

| Tipo Variable | Identificador | Ejemplo  |
| ------------- | :-----------: | :-------: |
| String | s | s_nombre |
| Int | i | i_edad |
| Boolean | b | b_existe |
| []anyType | a[tipo] | ai_edades |
| Testing | t | t_test |
| File | f | f_code |

- Empiezan con la palabra reservada **var**, seguido del nombre de la variable, su tipo y por último su valor inicial.
```go
s_nombre string := "Sakura"
```
- No deben haber varias variables en una misma línea
```go
i_inicio, i_contador int := 1, 0
```
- No se terminan las lineas en **;**
- Las variables se deben declarar con **:=**

#### Para Métodos:
- Los métodos se nombrarán en estañol
- Los métodos iniciarán en minúscula y palabras adicionales inician con mayúscula.
```go
func calcularLineas() int {}
```
- Empiezan con la palabra reservada **func**, luego el nombre, los parámetros de entrada, a continuación el tipo de dato que retorna y por último las llaves.
```go
func calcularMedida(i_alto int, i_ancho int) int {}
```
- Si tiene más de un valor de retorno se pone entre paréntesis los tipos de datos que se va a retornar.
```go
func calcularMedidas(i_alto int, i_ancho int) (int,int) {}
```
- Se pueden tener un número indeterminado de parámetros o argumentos en la función func
```go
sum(nums ...int) {}
```
- Las funciones de testing deben iniciar por la palabra 'Test'.
```
 func TestAverage(t_test *testing.T)
```

#### Para estructuras:
_(En go no hay clases)_
- Los nombres de las estructuras comienzan en minúscula, en igual caso para nombres compuestos
- Se nombran las estructuras en español
- Empieza por la palabra reservada type, seguido del nombre, luego sigue la palabra reservada struct.
```go
type contador struct {}
```
- Dentro de la estructura cada atributo se define como una variables, es decir, nombre seguido de tipo de dato
- Debe haber uno y solo un atributo por línea en el contenido de la estructura

> ADICIONAL:
> - La llave que abre una porcion de codigo debe iniciar en la misma linea de la declaración de un método o estructura
> - La primera línea de código debe ser package, seguido del nombre
> - Los imports deben ser únicos
