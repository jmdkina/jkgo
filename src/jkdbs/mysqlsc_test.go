package jkdbs

import (
	"testing"
	"time"
)

func TestMysqlSCDevStatusInsert(t *testing.T) {
	m, err := NewCMMysqlSC("v", "webfuture", "sctek_status")
	if err != nil {
		t.Fatal(err)
	}

	devstatus := SCDevStatus{
		Version:    "v1.0.9-20181023154622",
		Time:       time.Now().Unix(),
		DevType:    "rk",
		Mac:        "11:22:33:44:ff:ee",
		IP:         "192.168.5.215",
		RemoteIp:   "204.231.1.55",
		ActionType: "login",
		Online:     true,
	}
	err = m.AddDevStatus(devstatus)
	if err != nil {
		t.Error("error ", err)
	}

	devstatus.Mac = "aa:bb:cc:ee:dd:11"
	devstatus.ActionType = "logout"
	devstatus.Time = time.Now().Unix()
	err = m.AddDevStatus(devstatus)
	if err != nil {
		t.Error("error ", err)
	}

	m.Close()
}

func TestMysqlSCDevStatusQuery(t *testing.T) {
	m, err := NewCMMysqlSC("v", "webfuture", "sctek_status")
	if err != nil {
		t.Fatal(err)
	}

	// Query
	alldevstatus, err := m.QueryDevStatus()
	if err != nil {
		t.Fatal("query error ", err)
	} else {
		if 2 != len(alldevstatus) {
			t.Errorf("dev status should 2, but result is %d\n", len(alldevstatus))
		}
		if alldevstatus[0].Mac != "11:22:33:44:ff:ee" {
			t.Errorf("dev status 0 mac should be 11:22:33:44:ff:ee, but result is %s\n",
				alldevstatus[0].Mac)
		}
		if alldevstatus[1].Mac != "aa:bb:cc:ee:dd:11" {
			t.Errorf("dev status 1 mac should be aa:bb:cc:ee:dd:11, but result is %s\n",
				alldevstatus[1].Mac)
		}
	}

	m.Close()
}

func TestMysqlSCDevStatusQueryOne(t *testing.T) {
	m, err := NewCMMysqlSC("v", "webfuture", "sctek_status")
	if err != nil {
		t.Fatal(err)
	}

	// Query one
	qdevstatus, err := m.QueryDevStatusMac("11:22:33:44:ff:ee")
	if err != nil {
		t.Fatal("query one error ", err)
	} else {
		if qdevstatus.Mac != "11:22:33:44:ff:ee" {
			t.Errorf("dev status mac should be 11:22:33:44:ff:ee, but result is %s\n", qdevstatus.Mac)
		}
		if qdevstatus.IP != "192.168.5.215" {
			t.Errorf("dev status ip should be 192.168.5.215, but result is %s\n", qdevstatus.IP)
		}
	}

	m.Close()
}

func TestMysqlSCDevStatusUpdateOne(t *testing.T) {
	m, err := NewCMMysqlSC("v", "webfuture", "sctek_status")
	if err != nil {
		t.Fatal(err)
	}

	qdevstatus, err := m.QueryDevStatusMac("11:22:33:44:ff:ee")
	if err != nil {
		t.Fatal("update query error ", err)
	} else {
		if qdevstatus.Mac != "11:22:33:44:ff:ee" {
			t.Errorf("update query fail mac should be 11:22:33:44:ff:ee, but result is %s\n", qdevstatus.Mac)
		}
		qdevstatus.IP = "127.1.1.111"
		qdevstatus.DevType = "rk3308one"
		err = m.UpdateDevStatusMac(*qdevstatus)
		if err != nil {
			t.Errorf("update fail %v", err)
		} else {
			ndevstatus, err := m.QueryDevStatusMac("11:22:33:44:ff:ee")
			if err != nil {
				t.Fatal("update query again fail ", err)
			} else {
				if ndevstatus.IP != "127.1.1.111" {
					t.Fatalf("update ip should be 127.1.1.111, but result is %s", ndevstatus.IP)
				}
				if ndevstatus.DevType != "rk3308one" {
					t.Fatalf("update ip should be rk3308one, but result is %s", ndevstatus.DevType)
				}
			}
		}
	}

	m.Close()
}

func TestMysqlDevStatusRemove(t *testing.T) {
	m, err := NewCMMysqlSC("v", "webfuture", "sctek_status")
	if err != nil {
		t.Fatal(err)
	}

	err = m.RemoveDevStatus()
	if err != nil {
		t.Fatal("remove fail ", err)
	}

	m.Close()
}
