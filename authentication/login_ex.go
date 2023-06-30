// Package authentication implements the Authentication NEX protocol
package authentication

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginEx sets the LoginEx handler function
func (protocol *AuthenticationProtocol) LoginEx(handler func(err error, client *nex.Client, callID uint32, strUserName string, oExtraData *nex.DataHolder)) {
	protocol.LoginExHandler = handler
}

func (protocol *AuthenticationProtocol) handleLoginEx(packet nex.PacketInterface) {
	if protocol.LoginExHandler == nil {
		globals.Logger.Warning("Authentication::LoginEx not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.LoginExHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	oExtraData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.LoginExHandler(fmt.Errorf("Failed to read oExtraData from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	go protocol.LoginExHandler(nil, client, callID, strUserName, oExtraData)
}
