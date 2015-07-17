
PWD=`pwd`
GOPATH=$GOPATH:$PWD

FILES="golanger.com/config golanger.com/framework/web golanger.com/i18n golanger.com/log golanger.com/middleware golanger.com/utils golanger.com/session/filesession golanger.com/session/cookiesession golanger.com/session/memorysession golanger.com/database/activerecord goconfig/config code.google.com/p/go.net/websocket code.google.com/p/goprotobuf/proto code.google.com/p/goprotobuf/protoc-gen-go jk/jkcommon jk/jklog jk/jkconfig jk/jkprotobuf jk/jkserver helper bveth labix.org/v2/mgo labix.org/v2/mgo/bson github.com/mattn/go-sqlite3"
FILES+=" code.google.com/p/graphics-go/graphics"
FILES+=" jk/jkprotocol jk/jkparsedoc jk/jknetwork"
FILES+=" github.com/tyranron/daemonigo/"

for i in $FILES
do
    echo "go install $i ..."
    go install $i
done

echo "Install Success !"
