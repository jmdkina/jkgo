package helper

import (
	"encoding/gob"
	"github.com/golangers/log"
	//"github.com/golangers/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type TM map[string]interface{}

func init() {
	gob.Register(bson.M{})
}

type Mongo struct {
	session  *mgo.Session
	dbName   string
	colNames map[string]*mgo.Collection
}

func NewMongo(mgoDns string) *Mongo {
	DbPos := strings.LastIndex(mgoDns, "/")
	mgoAddr := mgoDns[:DbPos]
	dbName := mgoDns[DbPos+1:]
	mgoSession, err := mgo.Dial(mgoAddr)
	if err != nil {
		log.Fatal(err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return &Mongo{
		session:  mgoSession,
		dbName:   dbName,
		colNames: map[string]*mgo.Collection{},
	}
}

func (m *Mongo) C(col TM) *mgo.Collection {
	colName := col["name"].(string)

	if _, ok := m.colNames[colName]; !ok {
		m.colNames[colName] = m.session.DB(m.dbName).C(colName)
		if _, okIn := col["index"]; okIn {
			if colIndexs, okType := col["index"].([]string); okType {
				for _, colIndex := range colIndexs {
					colIndexArr := strings.Split(colIndex, ",")
					err := m.colNames[colName].EnsureIndex(mgo.Index{Key: colIndexArr, Unique: false})
					if err != nil {
						log.Fatal(colName+".Index:", err)
						return nil
					}
				}
			}
		}
	}

	return m.colNames[colName]
}

/* get databse */
func (m *Mongo) DB(dbname string) *mgo.Database {
	return m.session.DB(dbname)
}

/* get collections */
func (m *Mongo) Cnew(dbname string, cname string) *mgo.Collection {
	m.dbName = dbname
	return m.C(TM{"name": cname})
}

func (m *Mongo) DBSession() *mgo.Session {
	return m.session
}

/* get database name */
func (m *Mongo) DbName() string {
	return m.dbName
}

func (m *Mongo) Close() {
	m.session.Close()
}
