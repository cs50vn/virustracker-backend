from flask import Flask, render_template, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///test.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)
ma = Marshmallow(app)

from models import TodayData, YesterdayData

class TodaySchema(ma.SQLAlchemySchema):
    class Meta:
        model = TodayData
    
    id = ma.auto_field()
    name = ma.auto_field()
    case_total = ma.auto_field()
    case_today = ma.auto_field()
    case_active = ma.auto_field()
    case_serious = ma.auto_field()
    recovered_total = ma.auto_field()
    death_today = ma.auto_field()
    death_total = ma.auto_field()

class YesterdaySchema(ma.SQLAlchemySchema):
    class Meta:
        model = YesterdayData

@app.route('/')
def index():

    # init marshmallow schema
    countries_today = TodayData.query.all()
    today_schema = TodaySchema(many=True)
    countries = today_schema.dump(countries_today)

    return jsonify(countries)
    # return render_template('index.html', countries_today=countries_today)
