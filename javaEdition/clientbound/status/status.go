package status

import common2 "github.com/camdenorrb/minecraftPackets/javaEdition/protocol/common"

type PingResponse struct {
	Payload int64
}

func (*PingResponse) PacketType() string {
	return string(common2.ClientBoundStatusPong)
}

type Response struct {
	Version           Version     `json:"version"`
	Players           Players     `json:"players"`
	Description       Description `json:"description"`
	Favicon           string      `json:"favicon"`
	EnforceSecureChat bool        `json:"enforceSecureChat"`
}

func (*Response) PacketType() string {
	return string(common2.ClientBoundStatusResponse)
}
