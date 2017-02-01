package jkprotocol

import "errors"

type JKProtocolWrap struct {
	Type    int   // protocol type
	CmdType  int
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

func (wrap *JKProtocolWrap) RegisterResponse(data string) (string, error ) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		reg, err := NewV5RegisterResponse(data)
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

func (wrap *JKProtocolWrap) KeepaliveResponse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		keep, err := NewV5KeepaliveResponse(data)
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

func (wrap *JKProtocolWrap) LeaveResponse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		leave, err := NewV5LeaveResponse(data)
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

func (wrap *JKProtocolWrap) Parse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		v5base := V5Base{}
		v, t, e := v5base.Parse(data)
		if e != nil || v == nil {
			return "", e
		}
		switch t {
		case JK_PROTOCOL_V5_REGESTER:
			v5regres, err := NewV5RegisterResponse("OK")
			if err != nil {
				return "", err
			}
			wrap.CmdType = JK_PROTOCOL_V5_REGESTER
			return v5regres.String()
			break;
		case JK_PROTOCOL_V5_KEEPALIVE:
			v5keep, err := NewV5KeepaliveResponse("OK")
			if err != nil {
				return "", err
			}
			wrap.CmdType = JK_PROTOCOL_V5_KEEPALIVE
			return v5keep.String()
			break;
		case JK_PROTOCOL_V5_LEAVE:
			v5leave, err := NewV5LeaveResponse("OK")
			if err != nil {
				return "", err
			}
			wrap.CmdType = JK_PROTOCOL_V5_LEAVE
			return v5leave.String()
			break;
		default:
			return "", errors.New("Unsupported command type")
			break;
		}

		break;
	default:
		break;
	}

	return "", errors.New("Parse error")
}