from bs4 import BeautifulSoup
from models import VirusData
import requests
import re
from app import db


def scrape_data():
    base_url = 'https://www.worldometers.info/coronavirus/'
    response = requests.get(base_url)
    if response:
        print('Successful')
    else:
        print('There is a problem')
    soup = BeautifulSoup(response.text, 'html.parser')
    table = soup.find('table', id=["main_table_countries_today"])
    table_rows = table.find_all('tr')
    data = []
    for tr in table_rows:
        td = tr.find_all('td')
        row = [i.text.strip() for i in td]
        data.append(row)
    return data

def filter_data(data):
    data = data[1:-8]
    data.pop(6)
    return data

# import data into database objects
def import_data():
    virus_data = VirusData()

    # crawling data
    data_scraped = scrape_data()

    # process data
    data = filter_data(data_scraped)

    for row in data:
        virus_data = VirusData(
            name = row[0], 
            case_total = row[1],
            case_today = row[2],
            case_active = row[6],
            case_serious = row[7],
            recovered_total = row[5],
            death_today = row[4],
            death_total = row[3],
        )
        # save each object into database
        db.session.add(virus_data)
    db.session.commit()

