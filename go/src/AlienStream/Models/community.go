package models

import (
    "AlienStream/Services/db"
    "labix.org/v2/mgo/bson"
    "strings"
)

type Community struct {
    Name string
    Description string
    Thumbnail string
    Popularity int
    FavoriteCount int
    Moderators []string
    Sources []Source
    Children []Community
    TrackIds   CommunityTracks `json:"-"`
}

type CommunityTracks struct {
    Hot      []string
    TopToday []string
    TopWeek  []string
    TopMonth []string
    TopYear  []string
    TopAll   []string
}

type CompactCommunity struct {
    Name string
    Description string
    Thumbnail string
    Popularity int
    FavoriteCount int
    Moderators []string
    Sources []Source
    Children []Community
    TrackIds   CommunityTracks `json:"-"`
}

// CRUD --------------------------------------- 

func (community Community) Create() bool {

    session := db.Connection()
    defer session.Close()
    collection := session.DB("alien").C("community")

    err := collection.Find(bson.M{"name": community.Name}).One(&community)

    // make all urls uniform
    for index, source := range community.Sources {
        var stripped string = source.URL
        stripped = strings.TrimPrefix(stripped, "http://www.")
        stripped = strings.TrimPrefix(stripped, "http://")
        stripped = strings.TrimPrefix(stripped, "https://")

        community.Sources[index].URL = "http://" + stripped
    }
    if err == nil {
        return false 
    } 

    err = collection.Insert(community)
    return true
}


func (community Community) Update() bool {

    // Connect to our database
    session := db.Connection()
    defer session.Close()

    // Specify the Community Collection
    collection := session.DB("alien").C("community")

    // Update Our Database Object
    err, _ := collection.Upsert(bson.M{"name": community.Name},community)

    // return false on failure
    if(err != nil) {
        return false
    }
    return true;
}

func (community Community) Delete() bool {

    session := db.Connection()
    collection := session.DB("alien").C("community")

    err := collection.Remove(bson.M{"name":community.Name})

    if(err != nil) {
        return false
    }
    
    return true
}



// Utility --------------------------------------- 

// Community{}.Compact()
//
// converts a fully fledged community object into a smaller object
// for JSON serialization in cases where all subreddits are desired
func (community Community) Compact() CompactCommunity {
    return CompactCommunity(community)
}

func (community Community) Tracks(time string) []Track {
    var tracks []Track
    //for _,source := range community.Sources{
        tracks = append(tracks,Track{}.BySource(community.Name+"&t="+time)...)
    //}
    
    return tracks
}

// Selection --------------------------------------- 

// Community{}.All()
//
// Fetches all Community Objects
func (community Community) All() []Community {
    
    // Connect to our database
    session := db.Connection()
    defer session.Close()

    // Specify the Community Collection
    collection := session.DB("alien").C("community")
    
    // get all communities
    communities := []Community{}
    err := collection.Find(bson.M{}).All(&communities)
    if(err !=nil) {
        panic(err)
    }

    return communities
}

// Community{}.ByName()
//
// Fetches a Community by it's name
func (community Community) ByName(name string) Community {
    
    // Connect to our database
    session := db.Connection()
    defer session.Close()

    // Specify the Community Collection
    collection := session.DB("alien").C("community")

    // Fetch the Community
    collection.Find(bson.M{"name": name}).One(&community)

    return community
}

