import click
from flask.cli import with_appcontext
from app import db
from models import VirusData
from tools import scrape_data, filter_data, import_data

@click.command(name='create_table')
@with_appcontext
def create_table():
    db.create_all()

@click.command(name='import_data')
def import_data():
    import_data()
