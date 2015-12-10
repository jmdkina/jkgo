#!/usr/bin/python

import socket
import sys
from pathlib import Path

lenargs = len(sys.argv)
if lenargs <=2 :
	print ("Tell me connect addr and port")
	exit(-1)

host = sys.argv[1]
host = socket.gethostbyname(host)
port = int(sys.argv[2])

if port <= 0 or port > 65535 :
	print("Error port ", port)
	exit(-2)


class BVFileOperation(object):
	"""docstring for BVFileOperation"""
	__fhandle = 0;
	def __init__(self, arg):
		super(BVFileOperation, self).__init__()
		self.arg = arg
		
	def create_file(self, name):
		__fHanlde = open(name, "w")

	def write_tofile(self, text):
		__fHanlde.write(text)

	def close_file(self):
		__fHanlde.close()

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

print ("Start to connect ... ", host, " :", port)
s.connect((host, port))

""" recever has not ok """
rv = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
rv.bind(("", 23433))
rv.listen(2)

print ("Connect OK, start to do something ... ")
while True:
    data = input('What do you want to do [pipe/file/exit] > ')
    if data == "exit":
        break

    if data == "pipe" :
    	text = input("Tell me the command value > ")
    	n = s.send(bytes("pipe["+text+"]", encoding='utf-8'))
    	print ("Command has send with len ", n)

    if data == "file" :
        text = input("Tell me what file do you want to get > ")
        n = s.send(bytes("file["+text+"]", encoding='utf-8'))
        print ("Command has send with len ", n, " Wait response ...")
        fileop = BVFileOperation(n)
        fileop.create_file(text)
        
        while True:
        	try:
        		getdata = s.recv(1024)
        		if not getdata:
        			break
        		fileop.write_tofile(getdata)
        	except socket.Error:
        	    break
        
        fileop.close_file

    if data == "cmd" :
        text = input("Tell me the command you want to send > ")
        n = s.send(bytes("cmd["+text+"]", encoding="utf-8"))
        print ("Command has send with len ", n, "Wait response ...")

        while True:
            try :
                get = s.recv(1024)
                if not get:
                    break
                print ("The out data : ", get)
                break
            except socket.Error:
                break
    
print ("everything down, exit ...")
s.close
