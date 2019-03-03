package main

// func send(r io.Reader, e int64, t string) (err error) {
// 	recipPubkeys, err := base64keyRing(config().PairedDevice.PublicKey)
// 	senderPrivkey, err := base64keyRing(config().DeviceIdentity.PrivateKey)
// 	data, sig, err := crypto.en(r, recipPubkeys, senderPrivkey[0])
// 	msgParams := message.NewSendMessageParams()
// 	// msgParams.Messageinput.Message.Content = &data
// 	// msgParams.Messageinput.Message.Signature = &sig
// 	// msgParams.Messageinput.Message.ExitCode = e
// 	// msgParams.Messageinput.Message.Title = t
// 	msgParams.SetMessageinput(&models.MessageInput{&models.MessageInputMessage{Content: &data, Signature: &sig, ExitCode: e, Title: t}})
// 	msgParams.SetTimeout(10 * time.Second)
// 	_, err = polyrhythmAPI.Message.SendMessage(msgParams)
// 	check(err)
// 	// fmt.Printf("Data: %v\nSig: %v", data, sig)
// 	return
// }
