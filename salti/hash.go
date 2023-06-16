package salti

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateSalt()string{
	//alocate memory of 16 byte
	salt := make([]byte,16)

	//store random byte in each byte of 16bytes
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	//convert it to asci value to represent in string
	saltString := base64.StdEncoding.EncodeToString(salt);
	return saltString  

}

func SaltingHash(generatedSalt string, pwd string)(string,string) {
	salt := generatedSalt + pwd
	hash := sha256.Sum256([]byte(salt))
	hashedString := fmt.Sprintf("%x", hash[:])
	return generatedSalt,hashedString;
}

