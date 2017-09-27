#coding=utf8

import os
import pymysql
import urllib
import urllib2
from poster.encode import multipart_encode
from poster.streaminghttp import register_openers
import json

register_openers()


def Schedule(a,b,c):
    '''
    a:已经下载的数据块
    b:数据块的大小
    c:远程文件的大小
    '''
    per = 100.0 * a * b / c
    if per > 100 :
        per = 100
    # print '%.2f%%' % per

def http_post_file(url, name, path):
    datagen, headers = multipart_encode({name: open(path, "rb+")})
    request = urllib2.Request(url, datagen, headers)
    response = urllib2.urlopen(request)
    the_page = response.read()
    return the_page

def http_get(url):
    getf = urllib.urlopen(url)
    return getf.read()

def removefile(a, b):
    try:
        os.remove(a)
    except Exception, e:
        pass
    try:
        os.remove(b)
    except Exception, e:
        pass

def changepath():
    # db = 'tranca'
    # dbhost = 'rds77366gohgd760h0ko.mysql.rds.aliyuncs.com'
    # dbport = 3306
    # dbuser = 'dbtranca'
    # dbpass = 'd4520937132d4733a6262803371dc1e4'
    # imageurl = "http://127.0.0.1/xt/imageupload/upyun?token=test&type=topic"
    # mediaurl = "http://127.0.0.1/xt/mediaupload/upyun?token=test&type=topic"
    # settopicurl = "http://127.0.0.1:3333/tranca/setTopic?topicid=%d"

    db = 'tranca'
    dbhost = '139.196.22.120'
    dbport = 3306
    dbuser = 'root'
    dbpass = 'root123'
    imageurl = "http://127.0.0.1/xt/imageupload/upyun?token=test&type=topic"
    mediaurl = "http://127.0.0.1/xt/mediaupload/upyun?token=test&type=topic"
    settopicurl = "http://127.0.0.1:3333/tranca/setTopic?topicid=%d"

    savephoto = 'photo'
    savemadia = 'media'
    logfile = 'log.txt'
    file = open(logfile, 'wb')
    page = 0
    size = 10
    isok = True

    #清楚文件
    removefile(savephoto, savemadia)
    #链接数据库
    conn = pymysql.connect(host=dbhost, user=dbuser, passwd=dbpass, port=dbport, db=db, charset="utf8")
    conn.autocommit(True)
    cursor = conn.cursor()

    #查询所要处理的文件
    sql = '''SELECT * FROM tc_topicrelation 
    WHERE photopath like "http://dou.image.alimmdn.com/%" 
    OR photopath like "http://dou.file.alimmdn.com/%"
    OR videourl like "http://dou.file.alimmdn.com/%" '''
    LIMIT = 'ORDER BY id DESC LIMIT %d, %d'
    #更新处理好的信息
    updatesql = 'UPDATE tc_topicrelation SET %s WHERE id = %d'

    while isok:
        stmp = "".join([sql, LIMIT%(page, size)])  #拼接查询语句
        cursor.execute(stmp)                       #数据库查询
        n = cursor.fetchall()                      #获取所有数据
        for i in xrange(len(n)):
            id = n[i][0]
            topicid = n[i][1]
            out = ""
            changeimageurl = ''
            changemediaurl = ''
            #对符合条件的数据进行处理
            if n[i][2].find("http://dou.image.alimmdn.com/") == 0 or n[i][2].find("http://dou.file.alimmdn.com/") == 0:
                print n[i][2]
                out += "\r\nid:%d:topicid:%d:photo:%s" % (id, topicid, n[i][2])
                #下载文件
                urllib.urlretrieve(n[i][2], savephoto, Schedule)
                #上传文件获取地址
                ret = http_post_file(imageurl, "trancaimage", savephoto)
                print ret
                out += "\r\n"+ret
                jret = json.loads(ret)['data']
                if len(jret) > 0:
                    changeimageurl = jret[0]

            if n[i][8].find("http://dou.file.alimmdn.com/") == 0:
                print n[i][8]
                out += "\r\nid:%d:topicid:%d:media:%s" % (id, topicid, n[i][8])
                #下载文件
                urllib.urlretrieve(n[i][8], savemadia, Schedule)
                #上传文件获取地址
                ret = http_post_file(mediaurl, "trancamedia", savemadia)
                print ret
                out += "\r\n"+ret
                jret = json.loads(ret)['data']
                if len(jret) > 0:
                    changemediaurl = jret[0]

            updateset = ''
            if changeimageurl != '':
                updateset += 'photopath="%s", photoname="",' %(changeimageurl)
            if changemediaurl != '':
                updateset += 'videourl="%s",' %(changemediaurl)
            if updateset != '':
                upstmp = updatesql % (updateset[0:-1], id)
                out += "\r\n%s" % (upstmp)
                # print out
                file.write(out)
                cursor.execute(upstmp)                    #更新数据库
                http_get(settopicurl %(topicid))          #通知接口进行缓存更新
            removefile(savephoto, savemadia)              #删除临时文件
        page += 1
        # if len(n) == 0:
        #     isok = False
        isok = False

    file.close()
    cursor.close()
    conn.close()

if __name__ == '__main__':
    changepath()