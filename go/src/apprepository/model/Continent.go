package model

type Continent struct {
    Id string
    Name string
    Value int64
}

func MakeContinent(id string, name string, value int64) *Continent {
    return &Continent{Id: id, Name: name, Value: value}
}