package security

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/hkdf"
	"crypto/sha256"
)

// WORK IN PROGRESS

func CreateKeyPair() (*ecdh.PublicKey, *ecdh.PrivateKey, error){
	Curve := ecdh.X25519()

	PrivateKey,err := Curve.GenerateKey(rand.Reader)

	if err != nil{
		return nil,nil, err
	}

	PublicKey := PrivateKey.PublicKey()

	return PublicKey,PrivateKey,nil
}


func CalculateSecret(PrivateKey *ecdh.PrivateKey,OtherPublicKey *ecdh.PublicKey) ([]byte,error){
	return PrivateKey.ECDH(OtherPublicKey)
}

// Secret, Secret I wanna secret
func DeriveKey(Secret []byte) ([]byte,error){

	return hkdf.Key(sha256.New,Secret,nil,"GoChat Encryption Key",32)

}


func GetSalt() ([]byte,error){
	salt := make([]byte,32) 
	_,err := rand.Read(salt)

	if err != nil{
		return  nil,err
	}

	return salt,nil

}