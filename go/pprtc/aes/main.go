package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"github.com/xdg-go/pbkdf2"
)

// deriveKey ...
func deriveKey(salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(os.Getenv("CIPHER_TEXT")), salt, 1000, 32, sha256.New), salt
}

// Encrypt to encrypt product config
func Encrypt(plaintext string) string {
	key, salt := deriveKey(nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

// Decrypt to Decrypt product config
func Decrypt(ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}

func main() {
	encryptedDSN := Encrypt("host=127.0.0.1 port=5432 dbname=sphgpr_vc2 user=postgres password=root sslmode=disable")
	fmt.Println(encryptedDSN)

	encryptedDSNMysql := Encrypt("root:root@tcp(127.0.0.1:3307)/aubdem_vc_uat?parseTime=true")
	fmt.Println(encryptedDSNMysql)

	decrypted := Decrypt("e591be99db8284d4-37e963bd11d07c5fb569ec79-f7b631433295884061942e4ff0a8d492fecf5877c1f125cf3b1e499fcffef465e3f97bf0108e549c12ba0e2f265a17c533de4f99de623e247befecd42fef6ab13a3078177ca4d090127c1d2a8b4bf312a31e0577b950937b4ed87855c4d0249acdc5179c3955")
	fmt.Println(decrypted)

}
