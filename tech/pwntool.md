#### pwntool mini cheatsheet

1.3.1 Making Connections
1.3.2 Packing Integers
1.3.3 Setting the Target Architecture and OS
1.3.4 Setting Logging Verbosity
1.3.5 Assembly and Disassembly
1.3.6 Misc Tools
1.3.7 ELF Manipulation

##### 1.3.1 Making Connections

```
>>> conn = remote('ftp.ubuntu.com',21)
>>> conn.recvline()
'220 ...'
>>> conn.send('USER anonymous\r\n')
>>> conn.recvuntil(' ', drop=True)
'331'
>>> conn.recvline()
'Please specify the password.\r\n'
>>> conn.close()
```

##### 1.3.2 Packing Integers

```
>>> import struct
>>> p32(0xdeadbeef) == struct.pack('I', 0xdeadbeef)
True
>>> leet = '37130000'.decode('hex')
>>> u32('abcd') == struct.unpack('I', 'abcd')[0]
True
```

##### 1.3.5 Assembly and Disassembly
