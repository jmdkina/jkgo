
FILES=protocol.proto

for i in $FILES
do
    echo " protoc --go_out=. $i"
    protoc --go_out=. $i
done
