package main

import (
	"fmt"
)

var letras = []rune("ABCDEÉGHIÍJKLMNÑOÓPQRSTUÚVWXYZabcdeéfghiíjklmnñoópqrstuúvwxyz1234567890 .:,;¡!¿?''+-*/-=_{}[]()&%$#\\\"<>|°") // 105 Caracteres

var texto = "A" //Í

var displaceGlobal = 10 // Desplazamiento

func main() {
	fmt.Println("Desplazamiento: ", displaceGlobal)
	fmt.Println("Texto: ", texto)
	fmt.Println(" ")
	cifrarTexto()
}

func cifrarTexto() string {
	fmt.Println("Cifrado: ")
	cifrado := ""
	x := string(letras[67:105])
	runeDos := []rune(x + string(letras[0:67]))
	for _, e := range texto {
		for k, v := range letras {
			displace := k + displaceGlobal
			if e == v && displace <= 104 {
				cifrado = cifrado + string(letras[displace])
			}
			if e == v && displace > 104 {
				for krD, vrD := range runeDos {
					if vrD == e {
						cifrado = cifrado + string(runeDos[krD+displaceGlobal])
					}
				}
			}
		}
	}
	fmt.Println(cifrado)
	return cifrado
}
func descifrar(cifrado string) {

}
