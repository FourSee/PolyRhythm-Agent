package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"time"

	qrcode "github.com/Baozisoftware/qrcode-terminal-go"
	apiclient "github.com/foursee/swagger-go/client"
	"github.com/foursee/swagger-go/client/pairing_request"

	models "github.com/foursee/swagger-go/models"
)

func newPairingRequest() {
	_, pubKey, _ := myCryptoKey()
	user, _ := user.Current()
	name := user.Username
	hostname, _ := os.Hostname()
	deviceName := fmt.Sprintf("%s@%s", name, hostname)
	priParams := pairing_request.NewCreatePairingRequestParams()
	deviceType := "non_push"
	deviceInfo := models.PairingRequestInputPairingRequestDeviceInfo{DeviceType: &deviceType}
	priParams.SetPairingrequestinput(&models.PairingRequestInput{&models.PairingRequestInputPairingRequest{PublicKey: pubKey, DeviceName: deviceName, DeviceInfo: &deviceInfo}})
	priParams.SetTimeout(10 * time.Second)
	result, err := apiclient.Default.PairingRequest.CreatePairingRequest(priParams)

	if err != nil {
		fmt.Println(err)
	}
	obj := qrcode.New2(qrcode.ConsoleColors.NormalBlack, qrcode.ConsoleColors.BrightWhite, qrcode.QRCodeRecoveryLevels.Low)
	obj.Get([]byte(result.Payload.ShowURL)).Print()
	pairedDeviceKey, err := waitForAcceptance(result.Payload.ID)
	check(err)
	configInstance.PairedDevice.PublicKey = pairedDeviceKey
	check(configInstance.save())
	return
}

func waitForAcceptance(requestID string) (pubKey string, err error) {
	fmt.Println("Scan this QR code with the PolyRhythm app on your phone")
	fmt.Println("Waiting for pairing confirmation...")
	params := pairing_request.NewGetPairingRequestParams()
	params.SetID(requestID)
	params.SetTimeout(10 * time.Second)
	startedAt := time.Now()
	for {
		timeDiff := time.Now().Sub(startedAt)
		if timeDiff > 6000*time.Second {
			return "", errors.New("Timed out waiting for pairing acceptance")
		}
		accepted, pubKey, _ := getAcceptance(params)
		if accepted {
			return pubKey, nil
		}
		time.Sleep(5 * time.Second)
	}
}

func getAcceptance(params *pairing_request.GetPairingRequestParams) (accepted bool, pubKey string, err error) {
	prs, err := apiclient.Default.PairingRequest.GetPairingRequest(params)
	if err != nil {
		return
	}
	fmt.Println(prs.Payload.Status)
	if prs.Payload.Status == "accepted" {
		accepted = true
		pubKey = prs.Payload.AcceptedCryptoKey
	}
	return
}
