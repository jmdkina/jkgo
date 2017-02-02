package jkprotocol

import "errors"

const (
	JK_PROTOCOL_C_REGISTER = 1 << iota
	JK_PROTOCOL_C_LEAVE
	JK_PROTOCOL_C_KEEPALIVE
)

type JKProtocolWrap struct {
	Type     int   // protocol type
	CmdType  int
	Base     V5Base
}

func NewJKProtocolWrap(ptype int) (*JKProtocolWrap, error) {
	wrap := &JKProtocolWrap{
		Type: ptype,
	}

	wrap.Base = V5Base{}
	return wrap, nil
}

func (wrap *JKProtocolWrap) Register(data string) (string, error ) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		reg, err := wrap.Base.Register(data)
		return reg, err
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) RegisterResponse(data string) (string, error ) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		reg, err := wrap.Base.RegisterReponse(data)
		return reg, err
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) Keepalive(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		keep, err := wrap.Base.Keepalive(data)
		return keep, err
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) KeepaliveResponse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		keep, err := wrap.Base.KeepaliveResponse(data)
		return keep, err
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) Leave(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		leave, err := wrap.Base.Leave(data)
		return leave, err
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

func (wrap *JKProtocolWrap) LeaveResponse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		leave, err := wrap.Base.LeaveResponse(data)
		return leave, err
		break;
	default:
		break;
	}
	return "", errors.New("Unsupported protocol type")
}

// Set cmd type for which command
// Give common response of the sender
func (wrap *JKProtocolWrap) Parse(data string) (string, error) {
	switch wrap.Type {
	case JK_PROTOCOL_VERSION_5:
		v5base := V5Base{}
		v, t, e := v5base.Parse(data)
		if e != nil || v == nil {
			return "", e
		}
		switch t {
		case JK_PROTOCOL_C_REGISTER:
			v5regres, err := wrap.Base.RegisterReponse("RegisterResponse")
			wrap.CmdType = JK_PROTOCOL_C_REGISTER
			return v5regres, err
			break;
		case JK_PROTOCOL_C_KEEPALIVE:
			v5keep, err := wrap.Base.KeepaliveResponse("KeepaliveResponse")
			if err != nil {
				return "", err
			}
			wrap.CmdType = JK_PROTOCOL_C_KEEPALIVE
			return v5keep, err
			break;
		case JK_PROTOCOL_C_LEAVE:
			v5leave, err := wrap.Base.LeaveResponse("LeaveResponse")
			if err != nil {
				return "", err
			}
			wrap.CmdType = JK_PROTOCOL_C_LEAVE
			return v5leave, err
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