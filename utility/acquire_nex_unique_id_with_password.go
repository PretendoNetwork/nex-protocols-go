package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AcquireNexUniqueIDWithPassword sets the AcquireNexUniqueIDWithPassword handler function
func (protocol *UtilityProtocol) AcquireNexUniqueIDWithPassword(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.AcquireNexUniqueIDWithPasswordHandler = handler
}

func (protocol *UtilityProtocol) handleAcquireNexUniqueIDWithPassword(packet nex.PacketInterface) {
	if protocol.AcquireNexUniqueIDWithPasswordHandler == nil {
		globals.Logger.Warning("Utility::AcquireNexUniqueIDWithPassword not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.AcquireNexUniqueIDWithPasswordHandler(nil, client, callID)
}
