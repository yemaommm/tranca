#! /bin/bash
 
Ubuntu="Ubuntu"
Centos="Centos"
var2=`cat /proc/version`
 
#方法2
echo "$var2" | grep -q "$Ubuntu"
if [ $? -eq 0 ]; then
    sudo apt-get install golang
else
    sudo yum install golang
fi 
# 使用此方法，运行编译需安装apt-get install gccgo
# go build -compiler gccgo server.go