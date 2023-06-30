// Package nat_traversal implements the NAT Traversal NEX protocol
package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResult sets the ReportNATTraversalResult handler function
func (protocol *NATTraversalProtocol) ReportNATTraversalResult(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)) {
	protocol.ReportNATTraversalResultHandler = handler
}

func (protocol *NATTraversalProtocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	if protocol.ReportNATTraversalResultHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResult not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	// TODO - The NEX server should add a NATTraversalProtocolVersion method
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.ReportNATTraversalResultHandler(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), client, callID, 0, false, 0)
		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.ReportNATTraversalResultHandler(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), client, callID, 0, false, 0)
		return
	}

	var rtt uint32 = 0

	// TODO - Is this the right version?
	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 0 {
		rtt, err = parametersStream.ReadUInt32LE()
		if err != nil {
			go protocol.ReportNATTraversalResultHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, false, 0)
			return
		}
	}

	go protocol.ReportNATTraversalResultHandler(nil, client, callID, cid, result, rtt)
}
