package jkprotocol

// This file is the init call of protocol
// We define different protocol version for this program to call

const (
	JK_PROTOCOL_VERSION_1 = 0x01
	JK_PROTOCOL_VERSION_2 = 0x02
	JK_PROTOCOL_VERSION_3 = 0x04
	JK_PROTOCOL_VERSION_4 = 0x08
)

type JKProtoMain struct {
	t  int
	v4 *JKProtoV4
}

func JKProtoNew(version int) (*JKProtoMain, error) {
	pm := &JKProtoMain{}
	switch version {
	case JK_PROTOCOL_VERSION_1:
		break
	case JK_PROTOCOL_VERSION_4:
		pm.v4, _ = JKProtoV4New()
		break
	}
	return pm, nil
}
