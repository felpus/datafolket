package main

import (
	"fmt"
	"log"
	"os"
)

//Main-funksjon for å kjøre programmet, definerer arguments som tas inn fra kommandolinje
func main() {
	filnavn := os.Args[1]
	fileInfo(filnavn)
}


//Funksjon som legger inn filnavn som en stringvariabel, og finner informasjon om filen
func fileInfo(filnavn string) {
	fileInfo, err := os.Lstat(filnavn)
	if err != nil {
		log.Fatal(err)
	}

	//Definering av filstørrelse i B, KB, MB og GB
	bytes := float64(fileInfo.Size())
	kilo := bytes / 1024
	mega := kilo / 1024
	giga := mega / 1024

	//Informasjon om filen printes ut via Println
	fmt.Println("Infomation about file:", filenavn)
	fmt.Println("Bytes: ", bytes)
	fmt.Println("Kilobytes: ", kilo)
	fmt.Println("Megabytes: ", mega)
	fmt.Println("Gigabytes: ", giga)


	//Logikk for å identifisere flere av filens egenskaper

	//Hvis filen er en mappe, vil if-setningen vise true, hvis ikke vil den vise false
	if fileInfo.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if fileInfo.Mode().IsDir() == false {
		fmt.Println("Is not a directory")


		//Hvis filen er en "vanlig" fil, viser if-setningen true, hvis ikke vil den vise false
		if fileInfo.Mode().IsRegular() {
			fmt.Println("Is a regular file")
		} else {
			fmt.Println("Is not a regular file")
		}

		//Viser hvilke Unix Permission bits filen har
		fmt.Println("Has Unix permission bits:", fileInfo.Mode().Perm())

		//Sjekker om filen er "Append only" eller ikke
		if fileInfo.Mode() & os.ModeAppend == os.ModeAppend {
			fmt.Println("Is append only")
		} else {
			fmt.Println("Is not append only")
		}

		//Sjekker om filen er en device file eller ikke
		if fileInfo.Mode() & os.ModeDevice == os.ModeDevice {
			fmt.Println("Is a device file")
		} else {
			fmt.Println("Is not a device file")
		}

		//Viser om filen er en symbolsk link eller ikke
			if fileInfo.Mode() & os.ModeSymlink == os.ModeSymlink {
				fmt.Println("Is a symbolic link")
			} else {
				fmt.Println("Is not a Symbolic link")
			}
		}
	}
