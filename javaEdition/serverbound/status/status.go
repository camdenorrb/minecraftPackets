package status

import common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"

type Request struct{}

func (*Request) PacketType() string {
	return string(common2.ServerBoundStatusRequest)
}

type PingRequest struct {
	Payload int64
}

func (*PingRequest) PacketType() string {
	return string(common2.ServerBoundStatusPing)
}
