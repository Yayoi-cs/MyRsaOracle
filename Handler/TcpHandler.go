package Handler

import (
	"RsaOracle/RsaCrypt"
	"encoding/hex"
	"fmt"
	"net"
)

func TcpHandle(conn net.Conn) {
	defer conn.Close()
	//estr := strconv.Itoa(RsaCrypt.PrivateKeys.E)
	//nstr := RsaCrypt.PublicKeys.N.String()
	//conn.Write([]byte("**** The Rsa Oracle ****\nHere is public key\nE : " + estr + "\nN : " + nstr + "\n"))
	conn.Write([]byte("**** The Rsa Oracle ****\n"))
	modeInput := make([]byte, 2)
	textInput := make([]byte, 1024)
	conn.Write([]byte("E :Encrypt  D :Decrypt Q :Quit\n"))
	for {
		_, err := conn.Read(modeInput)
		if err != nil {
			fmt.Println("Error while reading", err)
			return
		}
		mode := ""
		for _, c := range modeInput {
			if rune(c) == '\n' || c == 0 {
				break
			}
			mode += string(rune(c))
		}
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
			outputText := formatCipherText(cipherText)
			if err != nil {
				fmt.Println("Error while encrypt", err)
				continue
			}
			conn.Write([]byte("PlainText : " + string(plainText) + "\nCipherText : " + outputText + "\nE :Encrypt  D :Decrypt Q :Quit\n"))
		}
		if mode == "D" {
			conn.Write([]byte("Decrypt\nEnter the text your want to decrypt :"))
			_, err = conn.Read(textInput)
			if err != nil {
				fmt.Println("Error while text reading")
				return
			}
			cipherText := inputToStr(textInput)
			cipher := decodeCipher(cipherText)
			if err != nil {
				fmt.Println("Error while reading")
			}
			plainText, err := RsaCrypt.RsaDecrypt(cipher)
			if err != nil {
				fmt.Println("Error while decrypt", err)
				continue
			}
			conn.Write([]byte("CipherText : " + string(cipherText) + "\nPlainText : " + string(plainText) + "\nHex : " + fmt.Sprintf("%x", plainText) + "\nE :Encrypt  D :Decrypt Q :Quit\n"))
		}
		if mode == "Q" {
			conn.Write([]byte("Quit\n"))
			break
		}
	}
}

func decodeCipher(text string) []byte {
	var retByte []byte
	if len(text)%2 == 1 {
		text = "0" + text
	}
	retByte, err := hex.DecodeString(text)
	if err != nil {
		return nil
	}
	return retByte

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

func formatCipherText(cipherText []byte) string {
	retStr := fmt.Sprintf("%x", cipherText)
	return retStr
}
