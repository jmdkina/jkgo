
set INSTALL_PATH=srctmp

hg clone https://code.google.com/p/go.net/  %INSTALL_PATH%/
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

git clone https://github.com/PuerkitoBio/goquery.git src/github.com/goquery
if [ ! -d src/github.com/andybalholm ] ; then
    mkdir -p src/github.com/andybalholm
fi
git clone https://github.com/andybalholm/cascadia.git src/github.com/andybalholm/cascadia
go get golang.org/x/net/html
go get github.com/sevlyar/go-daemon
go get github.com/tyranron/daemonigo

go get github.com/astaxie/beego
go get github.com/beego/bee


go get github.com/mattn/go-sqlite3
go get github.com/go-sql-dirver/mysql
go get github.com/lib/pq

