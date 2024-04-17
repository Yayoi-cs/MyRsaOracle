package Handler

import (
	"RsaOracle/RsaCrypt"
	"fmt"
	"net"
)

func TcpHandle(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("**** The Rsa Oracle ****\n"))
	modeInput := make([]byte, 2)
	textInput := make([]byte, 64)
	for {
		conn.Write([]byte("E :Encrypt  D :Decrypt Q :Quit\n"))
		_, err := conn.Read(modeInput)
		if err != nil {
			fmt.Println("Error while reading")
			return
		}
		mode := ""
		for _, c := range modeInput {
			if rune(c) == '\n' || c == 0 {
				break
			}
			mode += string(rune(c))
		}
		fmt.Println(modeInput, mode)
		if mode != "E" && mode != "D" && mode != "Q" {
			continue
		}
		if mode == "E" {
			conn.Write([]byte("Encrypt\nEnter the text you want to encrypt :"))
			_, err = conn.Read(textInput)
			if err != nil {
				fmt.Println("Error while text reading")
				return
			}
			plainText := inputToStr(textInput)
			cipherText, err := RsaCrypt.RsaEncrypt(plainText)
			if err != nil {
				fmt.Println("Error while encrypt", err)
				continue
			}
			fmt.Println(RsaCrypt.RsaDecrypt(string(cipherText)))
			conn.Write([]byte("PlainText : " + string(plainText) + "\nCipherText : " + string(cipherText) + "\n"))
		}
		if mode == "D" {
			conn.Write([]byte("Decrypt\nEnter the text your want to decrypt :"))
			_, err = conn.Read(textInput)
			if err != nil {
				fmt.Println("Error while text reading")
				return
			}
			cipherText := inputToStr(textInput)
			fmt.Println(cipherText)
			plainText, err := RsaCrypt.RsaDecrypt(cipherText)
			if err != nil {
				fmt.Println("Error while decrypt", err)
				continue
			}
			conn.Write([]byte("CipherText : " + string(cipherText) + "\nPlainText : " + string(plainText) + "\n"))
		}
		if mode == "Q" {
			conn.Write([]byte("Quit\n"))
			break
		}
	}
}

func inputToStr(text []byte) string {
	retStr := ""
	for _, c := range text {
		if rune(c) == '\n' || c == 0 {
			break
		}
		retStr += string(rune(c))
	}
	return retStr
}
