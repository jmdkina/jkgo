// Package jkcommon defines some common varients.
package jkcommon

const (
	JK_RESULT_SUCCESS                 = 0
	JK_RESULT_E_FAIL                  = -100
	JK_RESULT_E_PARAM_ERROR           = -101
	JK_RESULT_E_PARSE_ERROR           = -102
	JK_RESULT_E_DATABASE_QUERY_ERROR  = -103
	JK_RESULT_E_DATABASE_INSERT_ERROR = -104
	JK_RESULT_E_DATABASE_REMOVE_ERROR = -105
	JK_RESULT_E_NOT_EXIST             = -106
	JK_RESULT_E_HAS_EXIST             = -107
	JK_RESULT_E_DATA_NOT_EXIST        = -108
	JK_RESULT_E_CODE_ERROR            = -109
	JK_RESULT_E_NET_DIAL_ERROR        = -110
	JK_RESULT_E_TIME_FAST             = -111
	JK_RESULT_E_NOTSUPPORT            = -112
)

type ResultStatus struct {
	RS map[string]interface{}
}

func NewResultStatus(v int, desc string) *ResultStatus {
	rs := map[string]interface{}{
		"status": v,
	}
	if len(desc) > 0 {
		rs["desc"] = desc
	}
	rsr := ResultStatus{
		RS: rs,
	}
	return &rsr
}

func (rs *ResultStatus) setStatus(v int, desc string) {
	rs.RS["status"] = v
	rs.RS["desc"] = desc
}

func (rs *ResultStatus) SetCustom(v int, desc string) {
	rs.setStatus(v, desc)
}

func (rs *ResultStatus) SetNoPermission() {
	rs.setStatus(-1, "NoPermission")
}

func (rs *ResultStatus) SetInsertFail() {
	rs.setStatus(-11, "InsertFail")
}

func (rs *ResultStatus) SetItemExist() {
	rs.setStatus(-12, "ItemExist")
}

func (rs *ResultStatus) SetItemModFail() {
	rs.setStatus(-13, "ItemModFail")
}

func (rs *ResultStatus) SetItemDelFail() {
	rs.setStatus(-14, "ItemDelFail")
}

func (rs *ResultStatus) SetItemNotExist() {
	rs.setStatus(-15, "ItemNotExist")
}
