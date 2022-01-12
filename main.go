package main

import (
	"fmt"
)

var letras = []rune("ABCDEÉGHIÍJKLMNÑOÓPQRSTUÚVWXYZaábcdeéfghiíjklmnñoópqrstuúvwxyz1234567890 .:,;¡!¿?'/+-*=_{}[]()&%$#\\\"<>|@°€¥£") // 108 Caracteres

var texto = "Encriptación Cesár - Ejercicio de programación en Go."

var displaceGlobal = 20 // Desplazamiento

func main() {
	fmt.Println("Desplazamiento: ", displaceGlobal)
	fmt.Print("Texto: ", texto)
	fmt.Println(" ")

	cifrarTexto()

}

func cifrarTexto() string {
	cifrado := ""
	runeDos := []rune("7890 .:,;¡!¿?'/+-*=_{}[]()&%$#\"<>|@°€¥£ABCDEÉGHIÍJKLMNÑOÓPQRSTUÚVWXYZaábcdeéfghiíjklmnñoópqrstuúvwxyz123456")
	for _, e := range texto {
		for k, v := range letras {
			displace := k + displaceGlobal
			if e == v && displace <= 107 {
				cifrado = cifrado + string(letras[displace])
			}
			if e == v && displace > 107 {
				for krD, vrD := range runeDos {
					if vrD == e {
						sumarPosiciones := krD + displaceGlobal
						cifrado = cifrado + string(runeDos[sumarPosiciones])
					}
				}
			}
		}
	}
	fmt.Println("Cifrado: ", cifrado)

	descifrar(cifrado)

	return cifrado
}

func descifrar(cifrado string) {
	descifrado := ""
	x := []rune("RSTUÚVWXYZaábcdeéfghiíjklmnñoópqrstuúvwxyz1234567890 .:,;¡!¿?'/+-*=_{}[]()&%$#\"<>|@°€¥£ABCDEÉGHIÍJKLMNÑOÓPQ")
	for _, v := range cifrado {
		for k, vl := range letras {
			displace := k - displaceGlobal
			if v == vl && k >= displaceGlobal {
				descifrado += string(letras[displace])
			}
			if vl == v && k < displaceGlobal {
				for kX, vX := range x {
					if vX == v && kX >= displaceGlobal {
						restarPosiciones := kX - displaceGlobal
						descifrado += string(x[restarPosiciones])
					}
					if vX == v && kX <= displaceGlobal {
						restarPosiciones := displaceGlobal - kX
						descifrado += string(x[restarPosiciones])
					}
				}
			}
		}
	}
	fmt.Println("Descifrado: ", descifrado)
}
