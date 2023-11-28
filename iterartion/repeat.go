package iteration

import (
	"strings"
)

const repeatedCount = 5

func Repeat(character string) string {
    var repeated string
    for i := 0; i < repeatedCount ; i++ {
        repeated += character
    }
    return repeated
}



func ExampleRepeat(character string, repeated int) string {
    var resultado string
    for i := 0; i < repeated; i++ {
        // resultado += character
        resultado = strings.Repeat(character, repeated)
    }
    return resultado
}
