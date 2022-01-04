package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"

	"github.com/google/uuid"
)

// NewSHA256 ...
func NewSHA256(timestamp int64, lastHash string, data string, nonce int64, difficulty int64) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%s%s%d%d", timestamp, lastHash, data, nonce, difficulty)))
	// fmt.Println(hash)
	toR := string(hash[:])
	encodedString := hex.EncodeToString([]byte(toR))
	// fmt.Println(encodedString)
	return encodedString
}

//GenerateEllepticKeyPair generate private public Keypair
func GenerateEllepticKeyPair() *ecdsa.PrivateKey {
	pubkeyCurve := elliptic.P256()
	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return privatekey
}

func GenerateUUID() string {
	return uuid.New().String()

}

func SignOutput(pkey string, data []byte) (string, *big.Int, *big.Int) {
	// byteInput := StructToByte(data)

	r, s, serr := ecdsa.Sign(rand.Reader, DecodeECDSAPrivateKey(pkey), data)

	if serr != nil {
		fmt.Println(serr)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	return string(signature), r, s
}
func NewSHA256ForByteData(data []byte) []byte {
	hash := sha256.Sum256(data)
	// fmt.Println(hash)
	toR := hash[:]
	return toR
}

func EncodeECDSAPrivateKey(privateKey *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	return string(pemEncoded)
}
func EncodeECDSAPublicKey(publicKey *ecdsa.PublicKey) string {

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncodedPub)
}

func DecodeECDSAPublicKey(pemEncodedPub string) *ecdsa.PublicKey {

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return publicKey
}

func DecodeECDSAPrivateKey(pemEncoded string) *ecdsa.PrivateKey {

	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	return privateKey
}

func VerifySignature(publicKey string, signature string, dataHash []byte, r, s *big.Int) bool {
	verify := ecdsa.Verify(DecodeECDSAPublicKey(publicKey), dataHash, r, s)
	return verify
}
