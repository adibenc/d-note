# Python


[TOC]

## Research
## Basic
### Fundamental

[wiki] **Flask** is a micro [web framework](https://en.wikipedia.org/wiki/Web_framework) written in [Python](https://en.wikipedia.org/wiki/Python_(programming_language)). It is classified as a [microframework](https://en.wikipedia.org/wiki/Microframework) because it does not require particular tools or libraries.[[2\]](https://en.wikipedia.org/wiki/Flask_(web_framework)#cite_note-2) It has no database abstraction layer, form validation, or any other components where pre-existing third-party libraries provide common functions.

### Concepts

1. Microfw
2. [WSGI](https://en.wikipedia.org/wiki/WSGI)
3. Venv

### Why

1. Simple, not so steep learning curve
2. Beautiful
3. Scalable, small to large
4. 100% [WSGI] 1.0 compliant

### Why don't

1. Mobile app
2. Government project ?
3. I just know PHP to make website ?

### Libraries

#### Flask-*

1. Flask
2. Flask-Assets
3. Flask-Compress
4. Flask-Login
5. Flask-Mail
6. Flask-Migrate
7. Flask-RQ
8. Flask-Script
9. Flask-SQLAlchemy
10. Flask-SSLify
11. Flask-WTF

#### Standard py app

1. Jinja2
2. psycopg2
3. WTForms
4. SQLAlchemy

### 10 common cases

1. Information System
2. Enterprise

### 30 Common operations

1. #todo

### 10 products implementation

1. [Linkedin](https://www.youtube.com/watch?v=OXN3wuHUBP0#t=46) - 
2. [Pinterest](https://www.quora.com/What-challenges-has-Pinterest-encountered-with-Flask/answer/Steve-Cohen?srid=hXZd&share=1)
3. other ?

### Official Doc / refs

1. https://www.python.org/doc/
2. https://docs.djangoproject.com/
3. https://flask.palletsprojects.com/en/1.1.x/

## Production
1. Basic api project skeleton & boilerplate : [github](https://github.com/adib-enc/flask-base-dib)

2. Basic api & frontend pluggable project skeleton & boilerplate : [github](https://github.com/adib-enc/flask-base-dib)

### Deploying

#### Heroku - python web app

1. use heroku cli, or CI from github / gitlab
2. https://devcenter.heroku.com/articles/python-support
3. https://devcenter.heroku.com/articles/kafka-on-heroku#kafka-concepts

#### Docker, VPS - cli app, challenge, etc

1. use Docker or use py std lib to listen at certain port.

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

2. file upload
   ```python
   #todo
   ```

3. db operations
   ```go
   #todo
   ```
