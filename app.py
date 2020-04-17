import os
from flask import Flask, render_template, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow
from sqlalchemy import func
from datetime import datetime

app = Flask(__name__)
app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY') or 'super-secret'
app.config['SQLALCHEMY_DATABASE_URI'] = os.environ.get('SQLALCHEMY_DATABASE_URI') or 'sqlite:///test.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)
ma = Marshmallow(app)

from models import VirusData

class DataSchema(ma.SQLAlchemySchema):
    class Meta:
        model = VirusData
    
    id = ma.auto_field()
    name = ma.auto_field()
    case_total = ma.auto_field()
    case_today = ma.auto_field()
    case_active = ma.auto_field()
    case_serious = ma.auto_field()
    recovered_total = ma.auto_field()
    death_today = ma.auto_field()
    death_total = ma.auto_field()
    timestamp = ma.auto_field()

    

@app.route('/')
def index():

    data = VirusData.query.filter(VirusData.timestamp == datetime.today().date()).all()

    # init marshmallow schema
    data_schema = DataSchema(many=True)
    countries = data_schema.dump(data)

    return jsonify(countries)
    # return render_template('index.html', data=data)
