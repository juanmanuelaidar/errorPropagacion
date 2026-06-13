# Propagacion del error en consumo electrico

Este proyecto es un programa de consola escrito en Go que calcula el consumo electrico en kWh y estima el error propagado a partir de los errores absolutos de cada medicion.

La formula principal utilizada es:

```text
E = (V * I * fp * t) / 1000
```

Donde:

- `E`: energia consumida, en kWh.
- `V`: tension electrica, en volts.
- `I`: corriente electrica, en amperes.
- `fp`: factor de potencia.
- `t`: tiempo de uso, en horas.

## Requisitos

Tener Go instalado.

Para verificarlo:

```bash
go version
```

## Como ejecutar el proyecto

Desde la carpeta del proyecto:

```bash
go run consumoServidor.go
```

Tambien se puede compilar y ejecutar:

```bash
go build -o consumoServidor consumoServidor.go
./consumoServidor
```

## Como se usa

Al ejecutar el programa, la consola solicita los siguientes datos:

1. Tension `V` en volts.
2. Error absoluto de `V`.
3. Corriente `I` en amperes.
4. Error absoluto de `I`.
5. Factor de potencia `fp`.
6. Error absoluto de `fp`.
7. Tiempo `t` en horas.
8. Error absoluto de `t`.

Ejemplo de carga:

```text
Ingrese la tension V en volts: 220
Ingrese el error absoluto de V: 1
Ingrese la corriente I en amperes: 5
Ingrese el error absoluto de I: 0.1
Ingrese el factor de potencia fp: 0.9
Ingrese el error absoluto de fp: 0.02
Ingrese el tiempo t en horas: 3
Ingrese el error absoluto de t: 0.05
```

El programa muestra:

- Energia consumida.
- Error absoluto propagado.
- Error relativo.
- Error porcentual.

## Que hace cada parte del codigo

### Importaciones

```go
import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)
```

Cada paquete cumple una funcion:

- `bufio`: permite leer texto ingresado por consola.
- `fmt`: imprime mensajes y resultados.
- `math`: se usa para calcular valores absolutos con `math.Abs`.
- `os`: permite leer desde la entrada estandar.
- `strconv`: convierte texto a numeros decimales.
- `strings`: limpia espacios y saltos de linea del texto ingresado.

### Funcion `leerFloat`

```go
func leerFloat(nombre string) float64
```

Esta funcion pide un dato numerico por consola.

Recibe como parametro el texto que se va a mostrar al usuario. Luego:

1. Lee lo que el usuario escribe.
2. Elimina espacios y saltos de linea.
3. Intenta convertir el texto a `float64`.
4. Si el valor no es valido, muestra un mensaje de error y vuelve a pedirlo.
5. Si el valor es valido, lo devuelve.

Sirve para evitar que el programa falle cuando el usuario ingresa texto no numerico.

### Funcion `calcularEnergia`

```go
func calcularEnergia(v, i, fp, t float64) float64
```

Calcula la energia consumida usando la formula:

```text
E = (V * I * fp * t) / 1000
```

El resultado se expresa en kWh porque se divide por `1000`.

### Funcion `calcularErrorPropagado`

```go
func calcularErrorPropagado(v, i, fp, t float64, ev, ei, efp, et float64) float64
```

Calcula el error absoluto propagado de la energia.

Para eso usa las derivadas parciales de la formula de energia respecto de cada variable:

```text
dE/dV  = (I * fp * t) / 1000
dE/dI  = (V * fp * t) / 1000
dE/dfp = (V * I * t) / 1000
dE/dt  = (V * I * fp) / 1000
```

Luego multiplica cada derivada por el error absoluto de su variable:

```text
Error propagado =
|dE/dV|  * error de V +
|dE/dI|  * error de I +
|dE/dfp| * error de fp +
|dE/dt|  * error de t
```

El resultado queda expresado en kWh.

### Funcion `main`

```go
func main()
```

Es el punto de entrada del programa. Ejecuta todo el flujo principal:

1. Muestra el titulo y la formula usada.
2. Pide los valores de `V`, `I`, `fp` y `t`.
3. Pide los errores absolutos de cada variable.
4. Calcula la energia consumida.
5. Calcula el error absoluto propagado.
6. Calcula el error relativo:

```text
error relativo = error absoluto / |energia|
```

7. Calcula el error porcentual:

```text
error porcentual = error relativo * 100
```

8. Imprime todos los resultados por consola.

## Archivos del proyecto

- `consumoServidor.go`: contiene todo el codigo fuente del programa.
- `consumoServidor`: ejecutable compilado generado con `go build`.
- `README.md`: documentacion del proyecto.

## Notas

El programa trabaja con errores absolutos. Por ejemplo, si la tension medida es `220 V` y el instrumento tiene una incertidumbre de `1 V`, se ingresa:

```text
V = 220
error absoluto de V = 1
```

Si la energia calculada da `0`, el programa evita dividir por cero y deja el error relativo y porcentual en `0`.
