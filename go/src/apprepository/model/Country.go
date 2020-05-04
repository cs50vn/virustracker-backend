package model

type Country struct {
    Id            string
    Name          string
    CapitalName   string
    Area          int64
    Population    int64
    FlagId        string
    FlagUrl       string
    FlagData      string
    FlagTimestamp int64
    Timestamp     int64
    Continent     *Continent
    Items         []*Item
}

func MakeCountry(id string, name string, capitalName string, area int64, population int64, flagId string, flagUrl string, flagData string, flagTimestamp int64, timestamp int64, continent *Continent, items []*Item) *Country {
    return &Country{Id: id, Name: name, CapitalName: capitalName, Area: area, Population: population, FlagId: flagId, FlagUrl: flagUrl, FlagData: flagData, FlagTimestamp: flagTimestamp, Timestamp: timestamp, Continent: continent, Items: items}
}

