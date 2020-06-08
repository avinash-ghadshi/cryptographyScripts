#!/usr/bin/python

###=========================================================================###
### Author: Avinash Ghadshi                                                 ###
### Language Support: Python 2.7                                            ###
### Repository: https://github.com/avinash-ghadshi/cryptographyScripts      ###
### Details: This script demonstrates the working of Vigenere Cipher        ###
### Usage:                                                                  ###
### python vigenereCiphers.py  --help                                       ###
### python vigenereCiphers.py  -t 'Hello World' -k python -a 1              ###
### python vigenereCiphers.py  --text 'wcesc lmksr' --key python --action 2 ###
###=========================================================================###

import optparse
import re

def get_input(parser):
    parser.add_option("-t", "--text", dest="text", help="text string to convert")
    parser.add_option("-k", "--key", dest="key", help="key should contains only alphabets")
    parser.add_option("-a", "--action", dest="action", help="1: encryption, 2: Decryption")

    (options, aurgumets) = parser.parse_args()
    if not options.text or not options.key or not options.action:
        parser.error("[-] Please specify all options, use --help for more info.")

    if re.match(r'^[a-zA-Z\s]+$', options.text) == None:
        parser.error("[-] Text should not contains numbers or special characters other than space/tab, use --help for more info.")

    if re.match(r'^[a-zA-Z]+$', options.key) == None:
        parser.error("[-] Key should not contains numbers or special characters, use --help for more info.")

    options.text = options.text.lower()
    options.key = options.key.lower()

    if len(options.key) < len(options.text):
        options.key = options.key * (int(len(options.text)/len(options.key)) + 1)

    options.key = options.key[0:len(options.text)]

    return options


def initialize():
    global tabulaRecta
    tabulaRecta = dict()
    alphabets = [chr(ord('a')+x) for x in range(26)]
    
    for x in range(26):
        for y in range(26):
            tabulaRecta[alphabets[x]+''+alphabets[y]] = alphabets[(x+y)%26]


def encrypt(options):
    encstr = ''
    for x in range(len(options.text)):
        if options.text[x].strip() == '':
            encstr = encstr + options.text[x]
        else:
            encstr = encstr + tabulaRecta[options.text[x]+''+options.key[x]]
    
    print("[+] Encrypted Text = "+encstr)


def decrypt(options):
    decstr = ''
    for x in range(len(options.key)):
        if options.text[x].strip() == '':
            decstr = decstr + options.text[x]
            continue
        for trkey in tabulaRecta:
            if options.key[x] in trkey and tabulaRecta[trkey] is options.text[x]:
                dechar = trkey.replace(options.key[x],'',1)
                decstr = decstr + dechar
                break
            
    print("[+] Plain Text = "+decstr)


parser = optparse.OptionParser()

options = get_input(parser)

initialize()

if options.action == "1":
    encrypt(options)
elif options.action == "2":
    decrypt(options)
else:
    parser.error("[-] Please specify proper action, use --help for more info.")

