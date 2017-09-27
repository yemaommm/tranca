source ./port.sh
cd ..
# appname='run.go'
mode='EXE'
build='go run build.go'
server='go build main.go'
nowpath=`pwd`
export GOPATH=$nowpath
#go run $appname -r $build -r $server &
$build
$server
nohup ./main -host=$port -mode=$mode &
