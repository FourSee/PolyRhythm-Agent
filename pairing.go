package main

import (
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
	return
}
