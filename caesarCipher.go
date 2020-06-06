/*
###=====================================================================###
### Author: Avinash Ghadshi                                             ###
### Language: GO 1.14		                                        ###
### Repository: https://github.com/avinash-ghadshi/cryptographyScripts  ###
### Details: This script demonstrates the working of Caesor Cipher      ###
### Usage:                                                              ###
### go run caesarCipher.go 	                                        ###
###=====================================================================###
*/

package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
  _"reflect"
)


func encrypt(text string, key int) {
	var aEncryptedText = make([]string, len(text))
	aText := strings.Split(text,"")
	for i := 0; i < len(aText); i++ {
		tmp := []rune(aText[i])[0]
		if tmp == 9 || tmp == 32 {
			aEncryptedText = append(aEncryptedText, aText[i])
		} else if tmp >= rune('a') && tmp <= rune('z') {
			if tmp + rune(key) > rune('z') {
				aEncryptedText = append(aEncryptedText,string(tmp + rune(key) - 26))
			} else {
				aEncryptedText = append(aEncryptedText,string(tmp + rune(key)))
			}
		} else if tmp >= rune('A') && tmp <= rune('Z') {
			if tmp + rune(key) > rune('Z') {
				aEncryptedText = append(aEncryptedText,string(tmp + rune(key) - 26))
			} else {
				aEncryptedText = append(aEncryptedText,string(tmp + rune(key)))
			}
		}
	}
	sEncryptedText := strings.Join(aEncryptedText, "")
	fmt.Println("[+] Encrypted Text = "+sEncryptedText)
}

func decrypt(text string, key int) {
	var aEncryptedText = make([]string, len(text))
	aText := strings.Split(text,"")
	for i := 0; i < len(aText); i++ {
		tmp := []rune(aText[i])[0]
		if tmp == 9 || tmp == 32 {
			aEncryptedText = append(aEncryptedText, aText[i])
		} else if tmp >= rune('a') && tmp <= rune('z') {
			if tmp - rune(key) < rune('a') {
				aEncryptedText = append(aEncryptedText,string(tmp - rune(key) + 26))
			} else {
				aEncryptedText = append(aEncryptedText,string(tmp - rune(key)))
			}
		} else if tmp >= rune('A') && tmp <= rune('Z') {
			if tmp - rune(key) < rune('A') {
				aEncryptedText = append(aEncryptedText,string(tmp - rune(key) + 26))
			} else {
				aEncryptedText = append(aEncryptedText,string(tmp - rune(key)))
			}
		}
	}
	sEncryptedText := strings.Join(aEncryptedText, "")
	fmt.Println("[+] Plain Text = "+sEncryptedText)
}

func main() {

  fmt.Println("###------------------------------------NOTE----------------------------------------###")
  fmt.Println("### Text should not contains numbers and special characters other than Space / Tab.###")
  fmt.Println("### Key should be number and between 1 to 25.					   ###")
  fmt.Println("### Action should be 1 for encrytion and 2 for decryption.			   ###")
  fmt.Println("###--------------------------------------------------------------------------------###\n")

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text to encrypt / decrypt: ")
  text,_ := reader.ReadString('\n')
  text = strings.Replace(text, "\n", "", -1)
  fmt.Print("Enter key to encrypt / decrypt the text: ")
  skey,_ := reader.ReadString('\n')
  key,_ := strconv.Atoi(strings.Replace(skey, "\n", "", -1))
  if key < 1 || key > 25 {
	  fmt.Println("Invalid Key\nPlease read above NOTE")
	  os.Exit(1)
  }
  fmt.Print("Enter action (1: Encryption\t2: Decryption): ")
  action,_ := reader.ReadString('\n')
  action = strings.Replace(action, "\n", "", -1)

  switch action {
	case "1":
		encrypt(text,key)
	case "2":
		decrypt(text,key)
	default:
		fmt.Println("Invalid action.\nAction should be 1 or 2")
	}
}
