package main

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

//func main()  {
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
	//pkscript, err := hex.DecodeString("0020c600668682f0ec4cc5c718f12e44a082302098d640afe191f94aa58124947d89")
	//fmt.Println(pkscript)
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

//}