package jkcenter

import (
	l4g "github.com/alecthomas/log4go"
	"time"
)

// Do with CommandLists

// Do cycle with thread
func (cc *CenterControl) DoCycle() error {
	go func() {
		for ;; {
			for k, item := range cc.lists {
				l4g.Debug("Will give response %s\n", item.resp)
				item.conn.Write([]byte(item.resp))
				cc.lists = append(cc.lists[:k], cc.lists[k+1:]...)
				break
			}
			time.Sleep(time.Millisecond*500)
		}
	}()

	return nil
}