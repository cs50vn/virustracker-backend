package model

type RecentItem struct {
    Timestamp int64
    Value int64
}

func MakeRecentItem(timestamp int64, value int64) *RecentItem {
    return &RecentItem{Timestamp: timestamp, Value: value}
}


//1 -> cases, 2 -> deaths