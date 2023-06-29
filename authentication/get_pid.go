package authentication

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPID sets the GetPID handler function
func (protocol *AuthenticationProtocol) GetPID(handler func(err error, client *nex.Client, callID uint32, strUserName string)) {
	protocol.GetPIDHandler = handler
}

func (protocol *AuthenticationProtocol) handleGetPID(packet nex.PacketInterface) {
	if protocol.GetPIDHandler == nil {
		globals.Logger.Warning("Authentication::GetPID not implemented")
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
		go protocol.GetPIDHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.GetPIDHandler(nil, client, callID, strUserName)
}
