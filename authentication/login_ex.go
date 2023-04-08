package authentication

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginEx sets the LoginEx handler function
func (protocol *AuthenticationProtocol) LoginEx(handler func(err error, client *nex.Client, callID uint32, username string, authenticationInfo *AuthenticationInfo)) {
	protocol.LoginExHandler = handler
}

func (protocol *AuthenticationProtocol) HandleLoginEx(packet nex.PacketInterface) {
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

	username, err := parametersStream.ReadString()

	if err != nil {
		go protocol.LoginExHandler(err, client, callID, "", nil)
		return
	}

	dataHolderName, err := parametersStream.ReadString()

	if err != nil {
		go protocol.LoginExHandler(err, client, callID, "", nil)
		return
	}

	if dataHolderName != "AuthenticationInfo" {
		err := errors.New("[Authentication::LoginEx] Data holder name does not match")
		go protocol.LoginExHandler(err, client, callID, "", nil)
		return
	}

	_ = parametersStream.ReadUInt32LE() // length including this field

	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go protocol.LoginExHandler(err, client, callID, "", nil)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, protocol.Server)

	authenticationInfo, err := dataHolderContentStream.ReadStructure(NewAuthenticationInfo())

	if err != nil {
		go protocol.LoginExHandler(err, client, callID, "", nil)
		return
	}

	go protocol.LoginExHandler(nil, client, callID, username, authenticationInfo.(*AuthenticationInfo))
}
