# Python modules cheat

for dev & exploitation

## Exp

binascii

```
>>> import binascii
>>> binascii.hexlify('zam')
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: a bytes-like object is required, not 'str'
>>> binascii.unhexlify('7a616d')
b'zam'
>>> binascii.hexlify(b'zam')
b'7a616d'
```

pwntool
numpy





## Dev

```
json
math
numpy
statistics
```



## stdlibs musts

```
//strings
re
string

//cryptos / encoding / numbers
Crypto
base64
struct
random
hmac
hashlib
md5

// os / files sys / net
os
zlib
gzip
itertools
sys
PIL
zipfile
socket
```



## Next

```
time
signal
marshal
gmpy2
z3
frida
gdb
openssl
subprocess
```

## Next march 2021

```
asyncio
task_bg
redis
```



## refs

- https://stackoverflow com/questions/46904355/aes-128-cbc-decryption-in-python
