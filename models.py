from app import db
from datetime import datetime

class VirusData(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String())
    case_total = db.Column(db.String())
    case_today = db.Column(db.String())
    case_active = db.Column(db.String())
    case_serious = db.Column(db.String())
    recovered_total = db.Column(db.String())
    death_today = db.Column(db.String())
    death_total = db.Column(db.String())
    timestamp = db.Column(db.Date, nullable=False, default=datetime.utcnow)

    def __repr__(self):
        return f'ID:{self.id}, Country: {self.name}, Total Cases: {self.case_total}, Today Cases: {self.case_today}, \
                Active Case: {self.case_active}, Serious: {self.case_serious}, Recovered: {self.recovered_total}, \
                Today Deaths: {self.death_today}, Total Deaths: {self.death_total},'

