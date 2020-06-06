#!/usr/bin/python

###=====================================================================###
### Author: Avinash Ghadshi                                             ###
### Language Support: Python 2.7                                        ###
### Repository: https://github.com/avinash-ghadshi/cryptographyScripts  ###
### Details: This script demonstrates the working of Caesor Cipher      ###
### Usage:                                                              ###
### python caesarCipher.py  --help                                      ###
### python caesarCipher.py  -t 'Hello World' -k 3 -a 1                  ###
### python caesarCipher.py  --text 'Khoor Zruog' --key 3 --action 2     ###
###=====================================================================###

import optparse
import re, sys

def get_input(parser):
    parser.add_option("-t", "--text", dest="text", help="text string to convert")
    parser.add_option("-k", "--key", dest="key", help="key must be integer and between 1 to 25")
    parser.add_option("-a", "--action", dest="action", help="1: encryption, 2: Decryption")

    (options, aurgumets) = parser.parse_args()
    if not options.text or not options.key or not options.action:
        parser.error("[-] Please specify all options, use --help for more info.")
    return options

def initialize():
    global smallA, smallZ, capsA, capsZ
    smallA = ord('a')
    smallZ = ord('z')
    capsA = ord('A')
    capsZ = ord('Z')
    
    #print(str(smallA)+"\t"+str(smallZ)+"\t"+str(capsA)+"\t"+str(capsZ))


def encrypt(options):
    if re.match(r'^[a-zA-Z\s]+$', options.text) == None:
        print("[-] Text should not contains numbers or special characters other than space/tab.")
        sys.exit()

    aEncryptedText = list()
    key = int(options.key)

    for x in options.text:
        if ord(x) == 9 or ord(x) == 32:
            aEncryptedText.append(x)

        elif ord(x) >= smallA and ord(x) <= smallZ:
            if ord(x) + key > smallZ:
                aEncryptedText.append(chr(ord(x) + key - 26 ))
            else:
                aEncryptedText.append(chr(ord(x) + key))

        elif ord(x) >= capsA and ord(x) <= capsZ:
            if ord(x) + key > capsZ:
                aEncryptedText.append(chr(ord(x) + key - 26 ))
            else:
                aEncryptedText.append(chr(ord(x) + key))

    sEncryptedText = ''.join(aEncryptedText)
    print("[+] Plain Text = "+options.text)
    print("[+] Encrypted Text = "+sEncryptedText)



def decrypt(options):
    if re.match(r'^[a-zA-Z\s]+$', options.text) == None:
        print("[-] Text should not contains numbers or special characters other than space/tab.")
        sys.exit()

    aDecryptedText = list()
    key = int(options.key)

    for x in options.text:
        if ord(x) == 9 or ord(x) == 32:
            aDecryptedText.append(x)

        elif ord(x) >= smallA and ord(x) <= smallZ:
            if ord(x) - key < smallA:
                aDecryptedText.append(chr(ord(x) - key + 26 ))
            else:
                aDecryptedText.append(chr(ord(x) - key))

        elif ord(x) >= capsA and ord(x) <= capsZ:
            if ord(x) - key < capsA:
                aDecryptedText.append(chr(ord(x) - key + 26 ))
            else:
                aDecryptedText.append(chr(ord(x) - key))

    sDecryptedText = ''.join(aDecryptedText)
    print("[+] Encrypted Text = "+options.text)
    print("[+] Plain Text = "+sDecryptedText)


parser = optparse.OptionParser()

options = get_input(parser)

initialize()

if options.action == "1":
    encrypt(options)
elif options.action == "2":
    decrypt(options)
else:
    parser.error("[-] Please specify proper action, use --help for more info.")

