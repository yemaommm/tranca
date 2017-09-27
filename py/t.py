#coding=utf-8

import md5

s = '''16f2ffeb7d3e410abe10be54a25ddcb920161221,SHTsms,'''
s = md5.md5(s).hexdigest()
print s

import urllib

f = urllib.urlopen("http://www.shanghaitong.biz/api/v1/sms/sendSMS?MOBILE_PHONE=15618308028&VALID_CODE=222222&SHTKEY="+s)

print f.read()