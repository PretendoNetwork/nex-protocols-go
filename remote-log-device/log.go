package remote_log_device

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Log sets the Log handler function
func (protocol *RemoteLogDeviceProtocol) Log(handler func(err error, client *nex.Client, callID uint32, strMessage string)) {
	protocol.LogHandler = handler
}

func (protocol *RemoteLogDeviceProtocol) HandleLog(packet nex.PacketInterface) {
	if protocol.LogHandler == nil {
		globals.Logger.Warning("RemoteLogDevice::Log not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	var err error
	var strMessage string
	strMessage, err = parametersStream.ReadString()
	if err != nil {
		go protocol.LogHandler(err, client, callID, "")
	}

	go protocol.LogHandler(nil, client, callID, strMessage)
}
