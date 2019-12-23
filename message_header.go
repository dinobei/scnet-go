package scnet

type messageHeader struct {
	dataSize     int32
	packetType   int32
	messageType  int32
	cryptType    int32
	connectionID int32
}

const (
	mtProtobuf = iota
	mtRawbyte
	mtRawbyteRelay
	mtProtobufRelay
)

const magicPacketLength int = 2

var magicPacket = [...]byte{'I', 'J'}

const headerElements int = 5
const maxPacketHeaderSize int = 7 * headerElements
