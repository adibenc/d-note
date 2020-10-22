# Python cheat

## binascii
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

## refs
- https://stackoverflow com/questions/46904355/aes-128-cbc-decryption-in-python
