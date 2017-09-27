source ./port.sh
rm -rf /tmp/go-build*
sn=`ps -ef | grep $port | grep -v grep |awk '{print $2}'`
kill $sn
./clean.sh
echo "stop air:$sn"
