package encrypter

//go:generate mockery --name=Encrypter
type Encrypter interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
	EncryptDataToCode(data string, expire int64, expireUnit string) (string, error)
	DecryptCodeToData(code string) (string, error)
}

type implEncrypter struct {
	key string
}

func NewEncrypter(key string) Encrypter {
	return &implEncrypter{
		key: key,
	}
}
