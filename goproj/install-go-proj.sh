
PWD=`pwd`
GOPATH=$GOPATH:$PWD

FILES="github.com/astaxie/beego github.com/beego/bee "
FILES+=" code.google.com/p/graphics-go/graphics"
FILES+=" goconfig/config golanger.com/log golanger.com/utils"
FILES+=" jk/jkcommon jk/jkconfig jk/jkimage jk/jklog jk/jkmath jk/jkserver jk/jkprotocol jk/jkparsedoc jk/jknetwork"
FILES+=" labix.org/v2/mgo labix.org/v2/mgo/bson"
FILES+=" github.com/tyranron/daemonigo/"
FILES+=" github.com/jeffallen/mqtt github.com/surgemq/surgemq github.com/surgemq/surgemq/service"
FILES+=" github.com/deckarep/gosx-notifier"

for i in $FILES
do
    echo "go install $i ..."
    go install $i
done

echo "Install Success !"
