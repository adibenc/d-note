# Python


[TOC]

## Research
## Basic
### Fundamental

[wiki] **Python** is an [interpreted](https://en.wikipedia.org/wiki/Interpreted_language), [high-level](https://en.wikipedia.org/wiki/High-level_programming_language) and [general-purpose programming language](https://en.wikipedia.org/wiki/General-purpose_programming_language). Python's design philosophy emphasizes [code readability](https://en.wikipedia.org/wiki/Code_readability) with its notable use of [significant indentation](https://en.wikipedia.org/wiki/Off-side_rule). Its [language constructs](https://en.wikipedia.org/wiki/Language_construct) and [object-oriented](https://en.wikipedia.org/wiki/Object-oriented_programming) approach aim to help [programmers](https://en.wikipedia.org/wiki/Programmers) write clear, logical code for small and large-scale projects.[[29\]](https://en.wikipedia.org/wiki/Python_(programming_language)#cite_note-AutoNT-7-29)

## Concepts

1. Werkzeug abstractions

### Why

1. Simple, not so steep learning curve
2. Beautiful

### Why don't

1. Mobile app
2. Government project ?
3. I just know PHP to make website ?

### Libraries

1. basic utilities : stdlib - check at [doc](https://docs.python.org/3/)

2. web:

   1. f/w : django, flask

3. db : psycopg, pymysql, etc

4. mail

### 10 common cases

1. a

### 30 Common operations

1. IO
   1. cli stdin / stdout
   2. http request

### 10 products implementation

1. https://en.wikipedia.org/wiki/List_of_Python_software

### Official Doc & 3 refs
## Production
1. Basic api project skeleton & boilerplate

2. Basic api & frontend pluggable project skeleton & boilerplate

### Deploying
## FE
### Web HTML
### Desktop
### CLI
## Common cases
1. object to JSON, JSON to object

   ```python
   import json
   # dict to json string
   jsonString = json.dumps({"key":"value"})
   
   # json string to dict
   dicData = json.load('{"key":"value"}')
   ```

2. file system operations
   ```python
   mode = ['w','r','b']
   open('filename', mode[0]).write('text out') # write
text = open('filename', mode[1]).read() # read & save to text
   ```
   
3. db operations
   ```go
   
   ```


## Next

## refs

- py doc
