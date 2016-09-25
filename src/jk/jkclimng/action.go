package jkclimng

import "encoding/json"

// What does it do
// It receive command and make response of data it needs

type ActionInfo struct {
	Item JKClientItem
}

func AIInit() *ActionInfo {
	ai := &ActionInfo{}
	return ai
}

func (ai *ActionInfo) Action(data []byte) (error, []byte) {

	gen := AIGeneral{}
	err := gen.Parse(data)
	if err != nil {
		// TODO: give fail response
		return err, nil
	}

    switch gen.Header.Cmd {
	case "GetItem":
		// get data and response
	    req := AIGetItemRequest{}
		err := req.Parse(data)
		if err != nil {
			// TODO: Give fail response
			return err, nil
		}
		cnts := req.Body.Count
		data := GlobalDataControl().GetItem(cnts)
		//data := []byte("{ \"Header\":{ \"ID\":\"1234\", \"Cmd\":\"GetItem\"}, \"Body\":{ \"Count\": 1, \"Result\": \"success\", \"Data\":[{ \"Content\":\"doning\", \"Author\":\"jmd\" }]}  }")
		ret, err := json.Marshal(data)
		if err != nil {
			// TODO: give fail response
			return err, nil
		} else {
			return nil, ret
		}
	case "Keepalive":

	}

	return nil, nil
}

