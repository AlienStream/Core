package models

import (
    "AlienStream/Services/db"
    "labix.org/v2/mgo/bson"
)

type Track struct {
    Artist string
    Description string
    LastAccess string
    LastUpdate string
    Providers []Provider
    Rank int
    Title string
    Thumbnail string
}

// CRUD --------------------------------------- 

func (track Track) Upsert() bool {

    // Connect to our database
    session := db.Connection()
    defer session.Close()
    collection := session.DB("alien").C("track")

    // Check for Conflicting Sources

    // Update Our Database Object
    err, _ := collection.Upsert(bson.M{"title": track.Title},track)
    //err := collection.Insert(track)

    // return false on failure
    if(err != nil) {
        return false
        panic(err)
    }
    return true;
}

func (track Track) Delete() bool {

    // Connect to our database
    session := db.Connection()
    defer session.Close()
    collection := session.DB("alien").C("track")

    // Remove the Item
    err := collection.Remove(bson.M{"title":track.Title})

    if(err != nil) {
        return false
    }
    
    return true

}

func (track Track) ByIds(ids []string) []Track {
    
    // Connect to our database
    //session := db.Connection()
    //defer session.Close()

    // Specify the Track Collection
    //collection := session.DB("alien").C("track")

    var tracks []Track

    for _,track_id := range ids{
        var track Track
        track.Title = track_id
        tracks = append(tracks,track)
    }
    // Fetch the Tracks
    //collection.Find(bson.M{"title": bson.M{"$in" : ids}}).All(&tracks)

    //return tracks

    return tracks
}

func (track Track) BySource(source string) []Track {
    
    // Connect to our database
    session := db.Connection()
    defer session.Close()

    // Specify the Track Collection
    collection := session.DB("alien").C("track")

    var tracks []Track

    collection.Find(bson.M{"providers":bson.M{"$elemMatch":bson.M{"source.url": source}}}).
    Select(bson.M{"artist":1,"description":1,"lastaccess":1,"lastupdate":1,"rank":1,"title":1,"thumbnail":1,"providers.$":1}).
    All(&tracks)

    return tracks
}


