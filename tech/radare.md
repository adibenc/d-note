# radare2 cheat

## cheat
```
basic @https://zachgrace.com/

List functions
afl

Disassemble function:
aa
pdr@main

Print call graph:
agc > /tmp/foo.dot xdot /tmp/foo.dot

Print a detailed graph:
ag $$ > /tmp/c2.dot

Disassemble instruction:
pD 2

Seek to a specific memory location:
s 0x08048470

Write hex value:
wx eb

Set breakpoint
db 0x00401383

Remove breakpoint
db -0x00401383

List breakpoints
[0x00401383]> db 0x00401383 - 0x004013841 --x sw break enabled cmd="" name="0x00401383"

Continue
dc

Switch to Visual Mode
V

Cycle through Visual modes
p

Step through code
s

Switch to graph view (in Visual Mode)
V

```

## refs
- https://zachgrace.com/cheat_sheets/radare2/
- https://aboureada.com/security/radare2-cheat-sheet/
- https://ben.the-collective.net/my-notes/radare2-cheatsheet/
- https://gist.github.com/williballenthin/6857590dab3e2a6559d7
