package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/hoag/go-social-feed/pkg/util"
)

const (
	NotExpire = 0
)

func (e implEncrypter) createByteKey() []byte {
	return []byte(e.key)
}
func (e implEncrypter) Encrypt(plaintext string) (string, error) {
	key := e.createByteKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e implEncrypter) Decrypt(ciphertextStr string) (string, error) {
	key := e.createByteKey()

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext is too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (e implEncrypter) EncryptDataToCode(data string, expire int64, expireUnit string) (string, error) {
	var expireDuration time.Duration

	if expire != 0 {
		switch expireUnit {
		case "day":
			expireDuration = time.Hour * 24 * time.Duration(expire)
		case "hour":
			expireDuration = time.Hour * time.Duration(expire)
		case "minute":
			expireDuration = time.Minute * time.Duration(expire)
		case "second":
			expireDuration = time.Second * time.Duration(expire)
		default:
			return "", errors.New("invalid expire unit")
		}
	}

	dataWithExpire := DataWithExpire{
		Data: data,
	}

	if expire != 0 {
		dataWithExpire.ExpireTime = util.Now().Add(expireDuration).Unix()
	}

	finalData, err := json.Marshal(dataWithExpire)
	if err != nil {
		return "", err
	}

	encryptedStr, err := e.Encrypt(string(finalData))
	if err != nil {
		return "", err
	}

	return encryptedStr, nil
}

func (e implEncrypter) DecryptCodeToData(code string) (string, error) {
	decryptedStr, err := e.Decrypt(code)
	if err != nil {
		return "", err
	}

	var dataWithExpire DataWithExpire
	err = json.Unmarshal([]byte(decryptedStr), &dataWithExpire)
	if err != nil {
		return "", err
	}

	if dataWithExpire.ExpireTime != 0 && dataWithExpire.ExpireTime < util.Now().Unix() {
		return "", ErrExpireCode
	}

	data := dataWithExpire.Data
	return data, nil
}
