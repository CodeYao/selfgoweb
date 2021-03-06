// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/tjfoc/gmsm/sm2"
)

var Elliptic string = "p256"
var NumberOfKey int = 1

const (
	SM2 = iota
	ECDSA
)

func genKey(path string) {
	switch Cfg.Cert.KeyType {
	case "sm2":
		genSM2Key(path)
	case "ecdsa":
		genECDSAKey(path)
	default:
		fmt.Println("err key type!")
	}
}

func genSM2Key(path string) {
	for i := 0; i < NumberOfKey; i++ {
		priv, err := sm2.GenerateKey()

		if err != nil {
			fmt.Println(err)
		}
		//s := fmt.Sprintf("priv%d.pem", i)
		s := "key.pem"
		//设置私钥路径
		s = path + "/" + s
		ok, err := sm2.WritePrivateKeytoPem(s, priv, nil)
		if !ok {
			fmt.Println(err)
		}

		pubKey, _ := priv.Public().(*sm2.PublicKey)
		//s = fmt.Sprintf("pub%d.pem", i)
		s = "pubkey.pem"
		//设置公钥路径
		s = path + "/" + s
		ok, err = sm2.WritePublicKeytoPem(s, pubKey, nil)
		if !ok {
			fmt.Println(err)
		}

	}
}

func genECDSAKey(path string) {
	for i := 0; i < NumberOfKey; i++ {
		var c elliptic.Curve
		switch Elliptic {
		case "p256":
			c = elliptic.P256()
		case "p384":
			c = elliptic.P384()
		case "p521":
			c = elliptic.P521()
		default:
			fmt.Println("err elliptic curve!")

		}
		priv, err := ecdsa.GenerateKey(c, rand.Reader)

		if err != nil {
			fmt.Println(err)
		}

		//s1 := fmt.Sprintf("priv%d.pem", i)
		s1 := "key.pem"
		//设置私钥路径
		s1 = path + "/" + s1
		//s2 := fmt.Sprintf("pub%d.pem", i)
		s2 := "pubkey.pem"
		//设置公钥路径
		s2 = path + "/" + s2
		err = ToPem(priv, s1, s2)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func ToPem(key *ecdsa.PrivateKey, privName, pubName string) (err error) {
	pubB, err := x509.MarshalPKIXPublicKey(&key.PublicKey)

	if err != nil {
		return err
	}

	privB, err := marshalEcdsaUnecryptedPrivateKey(key)
	if err != nil {
		fmt.Println("marshal ecdsa priv key err:", err)
		return err
	}

	_, err = x509.ParsePKCS8PrivateKey(privB)
	if err != nil {
		fmt.Println("pkcs8 parse err:", err)
	}
	fmt.Println("pkcs8 parse success!")

	pub := &pem.Block{
		Type:  "ECDSA PUBLIC KEY",
		Bytes: pubB,
	}

	file, err := os.Create(pubName)
	if err != nil {
		return err
	}

	err = pem.Encode(file, pub)
	if err != nil {
		return err
	}

	file.Close()

	priv := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privB,
	}

	file, err = os.Create(privName)
	if err != nil {
		return err
	}

	err = pem.Encode(file, priv)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}
