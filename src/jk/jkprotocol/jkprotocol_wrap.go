package jkprotocol

import "errors"

type JKProtocolWrap struct {
	Type    int   // protocol type
}

func NewJKProtocolWrap(ptype int) (*JKProtocolWrap, error) {
	wrap := &JKProtocolWrap{
		Type: ptype,
	}

	return wrap, nil
}

func (wrap *JKProtocolWrap) Register(data string) (string, error ) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		reg, err := NewV5Register(data)
		if err != nil {
			return "", err
		}
		return reg.String()
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) Keepalive(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		keep, err := NewV5Keepalive(data)
		if err != nil {
			return "", err
		}
		return keep.String()
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) Leave(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		leave, err := NewV5Leave(data)
		if err != nil {
			return "", err
		}
		return leave.String()
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}
