package main

import (
    "encoding/hex"
    "fmt"
    "github.com/btcsuite/btcd/btcec"
    "github.com/btcsuite/btcutil"
    "massnet.org/mass-wallet/config"
    "massnet.org/mass-wallet/massutil"
    "massnet.org/mass-wallet/masswallet/keystore"
)

//func GenerateBTC() (string, string, error) {
//	privKey, err := btcec.NewPrivateKey(btcec.S256())
//	if err != nil {
//		return "", "", err
//	}
//
//	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
//	if err != nil {
//		return "", "", err
//	}
//	pubKeySerial := privKey.PubKey().SerializeUncompressed()
//
//	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
//	if err != nil {
//		return "", "", err
//	}
//
//	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
//}
//
//func GenerateBTCTest() (string, string, error) {
//	privKey, err := btcec.NewPrivateKey(btcec.S256())
//	if err != nil {
//		return "", "", err
//	}
//
//	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.TestNet3Params, false)
//	if err != nil {
//		return "", "", err
//	}
//	pubKeySerial := privKey.PubKey().SerializeUncompressed()
//	fmt.Println(len(privKey.PubKey().SerializeCompressed()))
//
//	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.TestNet3Params)
//	if err != nil {
//		return "", "", err
//	}
//
//	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
//}

func mains()  {
	//massutil.DecodeAddress("ms1qqzgswmq459zqq0ehguckw6u5cn82xqtasr7z5c8yq5jaspsjswg6ql2g3qv", &config.ChainParams)
	//fmt.Println(address, wifKey)
	//pkscript, err := hex.DecodeString("00201220ed82b4288007e6e8e62ced729899d4602fb01f854c1c80a4bb00c2507234")

	//if (err!=nil) {
	//	fmt.Println(err)
	//}
	//
	//ps, err := txscript.ParsePkScript(pkscript)
	//if (err!=nil) {
	//	fmt.Println(err)
	//}
	//fmt.Println(ps)
	//wif, err := btcutil.DecodeWIF("L5ArwptUcASFPNMCW4sEzxYYCUG8WMjxGDpkt27LTTBjaNKojcoA")
	//if(err != nil) {
	//	fmt.Println(err)
	//}
	//var pubs []*btcec.PublicKey
	//pubs = append(pubs, wif.PrivKey.PubKey())
	////_, add, err  := keystore.NewNonPersistentWitSAddrForBtcec(pubs,1,0x0000,&config.ChainParams)
	//fmt.Println(wif.PrivKey.PubKey().SerializeCompressed())
	pkscript, err := hex.DecodeString("020249fec74623b44b8161ba1972d392ca14ebc312ee1baa6c51b2d0d2eaa85c8b")
    ////pkscript, err := hex.DecodeString("3a54d62f11f1589734de789deb85e2c1af846863")
    //
    //if(err != nil) {
    //	fmt.Println(err)
    //}
    pubKey, err := btcec.ParsePubKey(pkscript, btcec.S256())
    var pubss []*btcec.PublicKey
    pubss = append(pubss,pubKey)
    _, addr, err := keystore.NewNonPersistentWitSAddrForBtcec(pubss, 1, massutil.AddressClassWitnessV0, &config.ChainParams)

    //addr, err := massutil.NewAddressPubKeyHash(pkscript, &config.ChainParams)

    if(err != nil) {
       fmt.Println(err)
    }
    fmt.Println(addr.EncodeAddress())
    wif, err := btcutil.DecodeWIF("L5ArwptUcASFPNMCW4sEzxYYCUG8WMjxGDpkt27LTTBjaNKojcoA")

    if err != nil {
        fmt.Println(err)
    }
    var pubs []*btcec.PublicKey
    pubs = append(pubs, wif.PrivKey.PubKey())
    _, add, err := keystore.NewNonPersistentWitSAddrForBtcec(pubs, 1, massutil.AddressClassWitnessV0, &config.ChainParams)
    fmt.Println(add.EncodeAddress())
//sk1qqccqxdp5z7rkye3w8rrcju39qsgczpxxkgzh7ry0ef2jczfy50kysfhcjeh
	//val, err := StringToAmount(s)
	//if err != nil {
	//	logging.CPrint(logging.ERROR, "Invalid amount string", logging.LogFormat{
	//		"str": s,
	//		"err": err,
	//	})
	//	st := status.New(ErrAPIInvalidAmount, ErrCode[ErrAPIInvalidAmount])
	//	return massutil.ZeroAmount(), st.Err()
	//}
	//fmt.Println(val)

}