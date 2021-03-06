package dtls

import "fmt"

type CipherSuite uint16

const (
	CipherSuite_TLS_PSK_WITH_AES_128_CCM_8      CipherSuite = 0xC0A8
	CipherSuite_TLS_PSK_WITH_AES_128_CBC_SHA256 CipherSuite = 0x00ae
)

type Cipher interface {
	GetPrfSize() int
	GenerateKeyBlock(masterSecret []byte, rawKeyBlock []byte) *keyBlock
	Encrypt(rec *record, key []byte, iv []byte, mac []byte) ([]byte, error)
	Decrypt(rec *record, key []byte, iv []byte, mac []byte) ([]byte, error)
}

func getCipher(peer *Peer, cipherSuite CipherSuite) Cipher {
	switch cipherSuite {
	case CipherSuite_TLS_PSK_WITH_AES_128_CCM_8:
		return CipherCcm{peer: peer}
	case CipherSuite_TLS_PSK_WITH_AES_128_CBC_SHA256:
		return CipherCBC{peer: peer}
	}
	return nil
}

func cipherSuiteToString(c CipherSuite) string {
	switch c {
	case CipherSuite_TLS_PSK_WITH_AES_128_CCM_8:
		return "TLS_PSK_WITH_AES_128_CCM_8(0xC0A8)"
	case CipherSuite_TLS_PSK_WITH_AES_128_CBC_SHA256:
		return "TLS_PSK_WITH_AES_128_CBC_SHA256(0xC0AE)"
	}
	return fmt.Sprintf("Unknown(0x%X)", c)
}
