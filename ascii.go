package main
import (
"fmt"
)
const ascii = "\x22\x20\x80\x20AC\xF7\x20\xBE\x20\x64\x6F\x6C\x6C\x61\x72\x20\x22"


func main() {
iterateOverExtendedASCIIStringLiteral()
ExtendedASCIIText(ascii)
}

func iterateOverExtendedASCIIStringLiteral(){
for i := 128; i <= 255; i++ {
		fmt.Printf("%X %c %b \n", i, i, i)
	}
}

func ExtendedASCIIText(tekst string){
a := tekst
for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}
}