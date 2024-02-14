package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Komut satırı argümanlarını kontrol et
	if len(os.Args) != 2 {
		fmt.Println("Kullanim: go run dosya.go <kelime>")
		return
	}
	// Aranan kelimeyi al
	word := os.Args[1]
	if word == "" {
		return
	} else if word == "\\n" {
		fmt.Println()
		return
	}
	// Dosyayı aç
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Dosya açilamadi!")
		return
	}
	defer file.Close()
	// grupları bir harita olarak sakla
	harita := make(map[int]string)
	groups := readFile(file)

	// ASCII haritası
	for i, group := range groups {
		harita[32+i] = group
	}

	words := strings.Split(word, "\\n")
	// Her bir ASCII karakterin ve satırın ekrana yazdırılması
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for line := 0; line <= 8; line++ {
				for i := 0; i < len(word); i++ {
					if ascii, ok := harita[int(word[i])]; ok {
						asciiLines := strings.Split(ascii, "\n")
						if line < len(asciiLines) {
							fmt.Print(asciiLines[line])
						}
					}
				}
				if line <= 8 && line > 0 {
					fmt.Println()
				}
			}
		}
	}
}

// dosya okumma
func readFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	var group []string

	for scanner.Scan() {
		group = append(group, scanner.Text())
		if len(group) == 9 {
			lines = append(lines, strings.Join(group, "\n"))
			group = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Dosya okuma hatasi:", err)
	}

	return lines
}
