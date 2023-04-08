package authentication

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestTicket sets the RequestTicket handler function
func (protocol *AuthenticationProtocol) RequestTicket(handler func(err error, client *nex.Client, callID uint32, userPID uint32, serverPID uint32)) {
	protocol.RequestTicketHandler = handler
}

func (protocol *AuthenticationProtocol) HandleRequestTicket(packet nex.PacketInterface) {
	if protocol.RequestTicketHandler == nil {
		globals.Logger.Warning("Authentication::RequestTicket not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	if len(parameters) != 8 {
		err := errors.New("[Authentication::RequestTicket] Parameters length not 8")
		go protocol.RequestTicketHandler(err, client, callID, 0, 0)
	}

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	userPID := parametersStream.ReadUInt32LE()
	serverPID := parametersStream.ReadUInt32LE()

	go protocol.RequestTicketHandler(nil, client, callID, userPID, serverPID)
}
