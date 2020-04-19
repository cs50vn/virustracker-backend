from flask import Flask
from markupsafe import escape

app = Flask(__name__)

@app.route('/v1/app')
def hello_world1():
    return 'Hello, World!'

@app.route('/v1/app/version')
def hello_world1():
    return 'Check version'