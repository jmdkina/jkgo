
set INSTALL_PATH=srctmp

go get code.google.com/p/goprotobuf/
go get github.com/sbinet/go-config/config

https://github.com/golangers/session.git
https://github.com/golangers/utils.git

https://github.com/golangers/webrouter.git
https://github.com/golangers/config.git
https://github.com/golangers/framework.git
https://github.com/golangers/log.git
https://github.com/golangers/validate.git
https://github.com/golangers/urlmanage.git
https://github.com/golangers/i18n.git
https://github.com/golangers/middleware.git
https://github.com/golangers/database.git

go get labix.org/v2/mgo
go get gopkg.in/mgo.v2

if [ ! -d src/github.com/andybalholm ] ; then
    mkdir -p src/github.com/andybalholm
fi
go get golang.org/x/net/html
go get github.com/sevlyar/go-daemon
go get github.com/tyranron/daemonigo

go get github.com/astaxie/beego
go get github.com/beego/bee


go get github.com/mattn/go-sqlite3
go get github.com/go-sql-dirver/mysql
go get github.com/lib/pq

go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git

git clone https://github.com/jeffallen/mqtt.git
git clone https://github.com/huin/mqtt.git
git clone github.com/surgemq/surgemq.git
git clone https://github.com/andybalholm/cascadia.git src/github.com/andybalholm/cascadia
git clone https://github.com/PuerkitoBio/goquery.git src/github.com/goquery
git clone https://github.com/surge/glog.git
git clone https://github.com/surgemq/message.git
hg clone https://code.google.com/p/go.net/  %INSTALL_PATH%/
