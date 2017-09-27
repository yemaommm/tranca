#coding=utf-8

import urllib2
import base64

import tornado.httpserver
import tornado.ioloop
import tornado.options
import tornado.web
from tornado.options import define, options
port = 8000
define("port", default=port, help="run on the given port", type=int)
settings = {'debug' : True}

def http_post(url, data):
    req = urllib2.Request(url, data)
    response = urllib2.urlopen(req)
    return response.read()


class IndexHandler(tornado.web.RequestHandler):
    def post(self, *args, **kwargs):
        body = self.request.body
        print body
        # jbody = base64.b64encode(body)
        url = 'https://sandbox.itunes.apple.com/verifyReceipt'
        print http_post(url, body)


if __name__ == '__main__':
    #启动tornado服务
    handlers = [(r"/", IndexHandler)]
    tornado.options.parse_command_line()
    app = tornado.web.Application(handlers = handlers, **settings)
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    a = tornado.ioloop.IOLoop.instance()
    tornado.ioloop.IOLoop.instance().start()