package jkdbs

import (
	"database/sql"
	"errors"
	"fmt"
	"jk/jklog"

	_ "github.com/go-sql-driver/mysql"
)

type CMMysqlSC struct {
	db  *sql.DB
	dsn string
}

func NewCMMysqlSC(username, password string, dbname string) (*CMMysqlSC, error) {
	ms := &CMMysqlSC{
		dsn: username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname,
	}
	var err error
	ms.db, err = sql.Open("mysql", ms.dsn)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (ms *CMMysqlSC) Close() {
	if ms.db != nil {
		ms.db.Close()
	}
}

func (msc *CMMysqlSC) AddDevStatus(st SCDevStatus) error {
	insertstr := fmt.Sprintf(`INSERT INTO devstatus VALUES ("%s", "%s", "%s", "%s", %d, "%s", "%s", %t)`,
		st.Mac, st.IP, st.DevType, st.Version, st.Time, st.RemoteIp, st.ActionType, st.Online)

	ret, err := msc.db.Exec(insertstr)
	if err != nil {
		return err
	}
	if LastInsertId, err := ret.LastInsertId(); err != nil {
		jklog.L().Debugln("last insert id ", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); err != nil {
		jklog.L().Debugln("rows affected ", RowsAffected)
	}

	return nil
}

func (msc *CMMysqlSC) queryDevStatus(querystr string) ([]*SCDevStatus, error) {
	rows, err := msc.db.Query(querystr)
	if err != nil {
		errstr := err.Error()
		return nil, errors.New(errstr)
	}
	defer rows.Close()
	dev_status := []*SCDevStatus{}
	indx := 0
	for rows.Next() {
		item := &SCDevStatus{}
		rows.Scan(&item.Mac, &item.IP, &item.DevType, &item.Version,
			&item.Time, &item.RemoteIp, &item.ActionType, &item.Online)

		jklog.L().Debugf("scan data indx %d: mac %s, ip %s, devtype %s, version %s, time %d, remoteip %s, actiontype %s, online %t\n",
			indx, item.Mac, item.IP, item.DevType, item.Version, item.Time,
			item.RemoteIp, item.ActionType, item.Online)
		dev_status = append(dev_status, item)
		indx = indx + 1
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return dev_status, nil
}

func (msc *CMMysqlSC) QueryDevStatus() ([]*SCDevStatus, error) {
	querystr := fmt.Sprint(`select * from devstatus`)
	return msc.queryDevStatus(querystr)
}

func (msc *CMMysqlSC) QueryDevStatusMac(mac string) (*SCDevStatus, error) {
	querystr := fmt.Sprintf(`select * from devstatus where Mac = "%s"`, mac)
	devstatus, err := msc.queryDevStatus(querystr)
	if err != nil {
		return nil, err
	}

	if len(devstatus) > 0 {
		return devstatus[0], nil
	} else {
		return nil, errors.New("Not found")
	}
}

func (msc *CMMysqlSC) UpdateDevStatusMac(st SCDevStatus) error {

	stmt, err := msc.db.Prepare(`update devstatus set ip = ?, devtype = ?, version = ?, time = ?, remoteip = ?, actiontype = ?, online = ? where mac = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(st.IP, st.DevType, st.Version, st.Time, st.RemoteIp, st.ActionType, st.Online, st.Mac)
	if err != nil {
		return err
	}
	return nil
}

func (msc *CMMysqlSC) RemoveDevStatus() error {
	stmt, err := msc.db.Prepare(`delete from devstatus`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
