#!/usr/bin/python

import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

host = ""
port = 11133

s.bind((host, port))

s.listen(1)

exit = False;

while True:
    if exit == True:
        break;
    print ("Accept now, wait data in ...")
    lin, addr = s.accept()
    print ("Recevied connected from ", addr)
  
    while True:
        data = lin.recv(1024)
        print ("Get data", data)

        if data == "exit" :
            exit = True
            break

        lin.send(bytes("Yes I got", encoding='utf-8'))

print ("Everything down, exit ...")
s.close()
