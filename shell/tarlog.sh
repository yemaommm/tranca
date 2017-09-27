bill=`date +%Y-%m-%d --date="-1 day"`
month=`date +%Y-%m`
path='/root/xt/go/martini-go'
if [ -f $path/log/tranca/info/$month/$bill.tar.gz ];then
	echo '文件已存在';
else
	tar czvPf $path/log/tranca/info/$month/$bill.tar.gz $path/log/tranca/info/$month/$bill.log*
	rm -rf $path/log/tranca/info/$month/$bill.log*
fi
if [ -f $path/log/tranca/error/$month/$bill.tar.gz ];then
        echo '文件已存在';
else
	tar czvPf $path/log/tranca/error/$month/$bill.tar.gz $path/log/tranca/error/$month/$bill.log*
	rm -rf $path/log/tranca/error/$month/$bill.log*
fi
if [ -f $path/log/$month/$bill.tar.gz ];then
        echo '文件已存在';
else
	tar czvPf $path/log/$month/$bill.tar.gz $path/log/$month/$bill.log*
	rm -rf $path/log/$month/$bill.log*
fi
