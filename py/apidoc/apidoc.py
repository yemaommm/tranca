#!/usr/bin/python
# -*- coding: UTF-8 -*- 

import re
import os
import logging
import thread
import pyinotify
import time

global path, mapidoc, apidoc, apiindex
path = ['../../src/IginServer/web/handler', '../../src/IginServer/framework']  #源码地址
mapidoc = True

'''
api注释样例
/**
 * @api {post} /admin/banner/add adminbanneradd添加banner
 * @apiName adminbanneradd添加banner
 * @apiGroup admin
 *
 * @apiVersion 2.0.0
 *
 * @apiParam {String} imageurl 图片
 * @apiParam {String} adurl 链接url
 * @apiParam {Number} sortflg 排序
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"status":200,"data":[{"Fpname":"鲜花","Fid":1,"Fimageurl":"","Fcreatetime":0,"Fsort":0,"Fprice":0}]}
 *
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 * 
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://139.196.22.120/xt/admin/banner/add
 */
'''

# pattern = re.compile(r'/\*\*[ \t]*\r?\n(([ \t]*\*([ \t]*[^ \f\n\r\t\v]*){0,}\r?\n){0,})[ \t]*\*/', re.M)
pattern = re.compile(r'/\*\*[ \t]*\r?\n(([ \t]*\*.*\r?\n){0,})[ \t]*\*/', re.M)
reapi = re.compile(r'[ \t]*\*[ \t]*@api [ \t]*\{(.*)\} [ \t]*([^ \f\n\r\t\v]*) [ \t]*(.*)', re.M)
rename = re.compile(r'[ \t]*\*[ \t]*@apiName [ \t]*(.*)', re.M)
reversion = re.compile(r'[ \t]*\*[ \t]*@apiVersion [ \t]*(.*)', re.M)
regroup = re.compile(r'[ \t]*\*[ \t]*@apiGroup [ \t]*(.*)', re.M)
reparam = re.compile(r'[ \t]*\*[ \t]*@apiParam [ \t]*\{(.*)\} [ \t]*([^ \f\n\r\t\v]*) [ \t]*(.*)', re.M)
resuccess = re.compile(r'[ \t]*\*[ \t]*@apiSuccess [ \t]*\((.*)\) [ \t]*\{(.*)\} [ \t]*([^ \f\n\r\t\v]*) [ \t]*(.*)', re.M)
reerror = re.compile(r'[ \t]*\*[ \t]*@apiError [ \t]*\((.*)\) [ \t]*\{(.*)\} [ \t]*([^ \f\n\r\t\v]*) [ \t]*(.*)', re.M)
reSampleRequest = re.compile(r'[ \t]*\*[ \t]*@apiSampleRequest [ \t]*(.*)', re.M)
reSuccessExample = re.compile(r'[ \t]*\*[ \t]*@apiSuccessExample [ \t]*([^ \f\n\r\t\v]*):[ \t]*\r?\n(([ \t]*\*[ \t].*\r?\n){0,})[ \t]*\*[ \t]*', re.M)


apidoc = {}
apiindex = {}

def regulardoc(string):
    '''
    获取api注释
    '''
    global path, mapidoc, apidoc, apiindex
    ret = []
    # print string
    stmp = pattern.findall(string)
    for i in stmp:
        ret.append(i[0])
    return ret

def getparam(string):
    '''
    获取api注释变迁
    '''
    global path, mapidoc, apidoc, apiindex
    api = reapi.findall(string)
    apiname = rename.findall(string)
    apiversion = reversion.findall(string)
    apigroup = regroup.findall(string)
    apiparams = reparam.findall(string)
    apisuccess = resuccess.findall(string)
    apierror = reerror.findall(string)
    apiSampleRequest = reSampleRequest.findall(string)
    apiexapmple = reSuccessExample.findall(string)
    # print apiexapmple
    if len(api)<=0 and len(apiname)<=0 and len(apiparams)<=0 and len(apisuccess)<=0 and len(apierror)<=0 and len(apiSampleRequest)<=0:
        return None, None, None
    ret = {
        'api':api,
        'apiname':apiname,
        'apigroup':apigroup,
        'apiparams':apiparams,
        'apisuccess':apisuccess,
        'apierror':apierror,
        'apiSampleRequest':apiSampleRequest,
        'apiexapmple':apiexapmple,
    }
    if len(apiexapmple) > 0:
        ret['apiexapmple'] = apiexapmple[0]
    else:
        ret['apiexapmple'] = ['', '']
    stmp = {}
    for i in ret['apisuccess']:
        if i[0] not in stmp:
            stmp[i[0]] = []
        stmp[i[0]].append(i)
    for i in ret['apierror']:
        if i[0] not in stmp:
            stmp[i[0]] = []
        stmp[i[0]].append(i)
    ret['apireturn'] = stmp

    stmparam = {}
    for i in ret['apiparams']:
        stmparam[i[1]] = i
    ret['apiparams'] = stmparam
    if len(apiversion) <= 0:
        apiversion = 'debug'
    else:
        apiversion = apiversion[0]
    return ret, apigroup[0], apiversion
    # print apiexapmple

def walk(path):
    '''
    遍历文件夹
    '''
    ret = []
    for rt, dirs, files in os.walk(path):
        for i in files:
            ret.append('/'.join([rt, i]))
    return ret

def readfile(path):
    f = open(path, 'rb')
    body = f.read()
    f.close()
    return body


def createapi():
    global path, mapidoc, apidoc, apiindex
    if mapidoc:
        mapidoc = False
        try:
            apidoc = {}
            apiindex = {}
            apiindex['ALL'] = {}
            apidoc['ALL'] = []
            for p in path:
                flist = walk(p)
                for i in flist:
                    print i
                    api = regulardoc(readfile(i))
                    for j in api:
                        p, groupname, version = getparam(j)
                        if p is not None:
                            groupname = ''.join(groupname.split('\r'))
                            version = ''.join(version.split('\r'))
                            p['apigroup'][0] = ''.join(p['apigroup'][0].split('\r'))
                            p['filepath'] = i
                            if version not in apidoc:
                                apidoc[version] = []
                            if version not in apiindex:
                                apiindex[version] = {}
                            apiindex[version][groupname] = groupname
                            apidoc[version].append(p)
                            apiindex['ALL'][groupname] = groupname
                            apidoc['ALL'].append(p)
        except Exception, e:
            print e
        mapidoc = True

class MyEventHandler(pyinotify.ProcessEvent):

    def process_IN_ACCESS(self, event):
        pass

    def process_IN_ATTRIB(self, event):
        pass

    def process_IN_CLOSE_NOWRITE(self, event):
        pass

    def process_IN_CLOSE_WRITE(self, event):
        print 'process_IN_CLOSE_WRITE'
        thread.start_new_thread(createapi, ())

    def process_IN_CREATE(self, event):
        print 'process_IN_CREATE'
        thread.start_new_thread(createapi, ())

    def process_IN_DELETE(self, event):
        print 'process_IN_DELETE'
        thread.start_new_thread(createapi, ())

    def process_IN_MODIFY(self, event):
        print 'process_IN_MODIFY'
        thread.start_new_thread(createapi, ())

    def process_IN_OPEN(self, event):
        pass

def filechange():
    # watch manager
    wm = pyinotify.WatchManager()
    for p in path:
        wm.add_watch(p, pyinotify.ALL_EVENTS, rec=True)
    #/tmp是可以自己修改的监控的目录
    # event handler
    eh = MyEventHandler()

    # notifier
    notifier = pyinotify.Notifier(wm, eh)
    notifier.loop()

if __name__ == '__main__':
    # global path, mapidoc, apidoc, apiindex
    print mapidoc
    createapi()
    thread.start_new_thread(filechange, ())
    #启动tornado服务
    import tornado.httpserver
    import tornado.ioloop
    import tornado.options
    import tornado.web
    from tornado.options import define, options
    port = 8888
    define("port", default=port, help="run on the given port", type=int)
    settings = {'debug' : True}

    class IndexHandler(tornado.web.RequestHandler):
        def get(self):
            version = self.get_argument('version', 'debug')
            if version in apiindex:
                self.render("index.html", list=apiindex, version=version, mlist=sorted(apiindex[version]))
            else:
                self.render("index.html", list=apiindex, version=version, mlist=sorted({}))

    class InfoHandler(tornado.web.RequestHandler):
        def get(self, *args, **kwargs):
            version = self.get_argument('version', 'debug')
            group = args[0]
            group = group.encode('utf8')
            if group not in apiindex[version]:
                self.write('none')
                self.flush()
                return
            ret = []
            for i in apidoc[version]:
                if i['apigroup'][0] == group:
                    ret.append(i)
            self.render("info.html", list=ret, indexlist=apiindex, version=version)

    handlers = [(r"/apidoc/*", IndexHandler), (r"/apidoc/(.*)", InfoHandler)]
    tornado.options.parse_command_line()
    app = tornado.web.Application(handlers = handlers, **settings)
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    logging.info(' '.join(["server start", "port:", str(port)]))
    a = tornado.ioloop.IOLoop.instance()
    tornado.ioloop.IOLoop.instance().start()