package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueID sets the AcquireNexUniqueID handler function
func (protocol *UtilityProtocol) AcquireNexUniqueID(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.AcquireNexUniqueIDHandler = handler
}

func (protocol *UtilityProtocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	if protocol.AcquireNexUniqueIDHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.AcquireNexUniqueIDHandler(nil, client, callID)
}
