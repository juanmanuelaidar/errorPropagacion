package main

import (
	"bufio"   // permite leer texto ingresado por consola.
	"fmt"     // imprime mensajes y resultados.
	"math"    // se usa para calcular valores absolutos con math.Abs.
	"os"      // permite leer desde la entrada estandar.
	"strconv" // convierte texto a numeros decimales.
	"strings" // limpia espacios y saltos de linea del texto ingresado.
)

func leerFloat(nombre string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(nombre + ": ")
		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)

		valor, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Valor inválido. Ingresá un número válido.")
			continue
		}

		return valor
	}
}

func calcularEnergia(v, i, fp, t float64) float64 {
	return (v * i * fp * t) / 1000
}

func calcularErrorPropagado(v, i, fp, t float64, ev, ei, efp, et float64) float64 {
	// Derivadas parciales calculadas manualmente
	dEdV := (i * fp * t) / 1000
	dEdI := (v * fp * t) / 1000
	dEdfp := (v * i * t) / 1000
	dEdt := (v * i * fp) / 1000

	errorPropagado := math.Abs(dEdV)*ev + math.Abs(dEdI)*ei + math.Abs(dEdfp)*efp + math.Abs(dEdt)*et
	return errorPropagado
}

func main() {
	fmt.Println("Propagación del error en el consumo eléctrico")
	fmt.Println("Fórmula: E = (V * I * fp * t) / 1000")
	fmt.Println()

	v := leerFloat("Ingrese la tensión V en volts")
	ev := leerFloat("Ingrese el error absoluto de V")

	i := leerFloat("Ingrese la corriente I en amperes")
	ei := leerFloat("Ingrese el error absoluto de I")

	fp := leerFloat("Ingrese el factor de potencia fp")
	efp := leerFloat("Ingrese el error absoluto de fp")

	t := leerFloat("Ingrese el tiempo t en horas")
	et := leerFloat("Ingrese el error absoluto de t")

	energia := calcularEnergia(v, i, fp, t)
	errorAbsoluto := calcularErrorPropagado(v, i, fp, t, ev, ei, efp, et)

	var errorRelativo float64
	var errorPorcentual float64

	if energia != 0 {
		errorRelativo = errorAbsoluto / math.Abs(energia)
		errorPorcentual = errorRelativo * 100
	}

	fmt.Println()
	fmt.Println("Resultados")
	fmt.Println("-----------------------------")
	fmt.Printf("Energía consumida: %.2f kWh\n", energia)
	fmt.Printf("Error absoluto propagado: %.2f kWh\n", errorAbsoluto)
	fmt.Printf("Error relativo: %.2f\n", errorRelativo)
	fmt.Printf("Error porcentual: %.2f %%\n", errorPorcentual)
}
