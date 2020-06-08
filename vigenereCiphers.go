/*
###========================================================================###
### Author: Avinash Ghadshi                                                ###
### Language Support: go 1.14                                              ###
### Repository: https://github.com/avinash-ghadshi/cryptographyScripts     ###
### Details: This script demonstrates the working of Vigenere Cipher       ###
### Usage:                                                                 ###
### go run vigenereCiphers.go --help                                       ###
### go run vigenereCiphers.go -t 'Hello World' -k python -a 1              ###
### go run vigenereCiphers.go --text 'wcesc lmksr' --key python --action 2 ###
### Requirements:                                                          ###
### This code need to install go-getoptions to build and run successfully  ###
### Use below command to install above package				   ###
### go get github.com/DavidGamba/go-getoptions				   ### 
###========================================================================###
*/
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/DavidGamba/go-getoptions"
)

var alphabets [26]string
var tabulaRecta = make(map[string]string)
var text string
var action int
var key string

//func  initialize() (map[string]string) {
func  initialize() {
	var count int = 0
	for i:=rune('a'); i <= rune('z'); i++ {
		alphabets[count] = string(i)
		count++
	}
	for i:=0; i<26; i++ {
		for j:=0; j<26; j++ {
			tabulaRecta[alphabets[i]+""+alphabets[j]]  = alphabets[(i+j)%26]
		}
	}
	//fmt.Println(alphabets)
	//fmt.Println(tabulaRecta)
	//os.Exit(0)
	//return tabulaRecta
}

func getAndValidateInput() {
	opt := getoptions.New()

	opt.Bool("help", false, opt.Alias("h", "?"))
	opt.StringVar(&text, "text", "",
		opt.Alias("t"),
		opt.Required(),
		opt.Description("Plain / Cipher Text to Encrypt / Decrypt"))
	opt.StringVar(&key, "key", "",
		opt.Alias("k"),
		opt.Required(),
		opt.Description("Key should contains only alphabets"))
	opt.IntVar(&action, "action", 0,
		opt.Alias("a"),
		opt.Required(),
		opt.Description("1: encryption, 2: Decryption"))

	_, err := opt.Parse(os.Args[1:])
	if opt.Called("help") {
		fmt.Fprintf(os.Stderr, opt.Help())
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\nUse --help for more info.\n", err)
		//fmt.Fprintf(os.Stderr, opt.Help(getoptions.HelpSynopsis))
		os.Exit(0)
	}

	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	re1 := regexp.MustCompile(`^[a-zA-Z]+$`)

	if !re.MatchString(text) {
		fmt.Println("[-] Text should not contains numbers or special characters other than space/tab, use --help for more info.")
		os.Exit(0)
	}

	if !re1.MatchString(key) {
		fmt.Println("[-] Key should not contains numbers or special characters, use --help for more info.")
		os.Exit(0)
	}

	text = strings.ToLower(text)
	key  = strings.ToLower(key)

	if len(key) < len(text) {
		key = strings.Repeat(key, (len(text)/len(key)) + 1)
	}

	key = key[0:len(text)]

}

func encrypt(tabulaRecta map[string]string, text string, key string) {
	var encstr string
	for i:=0; i<len(text); i++ {
		if ok := strings.TrimSpace(string(text[i])); ok == "" {
			encstr = encstr +""+ string(text[i])
		} else {
			encstr = encstr +""+ tabulaRecta[string(text[i])+""+string(key[i])]
		}
	}
	fmt.Println("[+] Encrypted Text = "+encstr)
}

func decrypt(tabulaRecta map[string]string, text string, key string) {
	decstr := ""
	for i:=0; i<len(key); i++ {
		if strings.TrimSpace(string(text[i])) == "" {
			decstr = decstr + string(text[i])
			continue
		}
		for trkey, v := range tabulaRecta {
			if strings.Contains(trkey, string(key[i])) && v == string(text[i]) {
				dechar := strings.Replace(trkey,string(key[i]),"",1)
				decstr = decstr + dechar
				break
			}
		}
	}
	fmt.Println("[+] Plain Text = "+decstr)
}


func main() {

	//tabulaRecta = initialize()
	initialize()
	getAndValidateInput()

	switch action {
	case 1:
		encrypt(tabulaRecta, text, key)
	case 2:
		decrypt(tabulaRecta, text, key)
	default:
		fmt.Println("[-]Invalid action.\nAction should be 1 or 2")
	}
}

