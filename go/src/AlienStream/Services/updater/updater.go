package updater

import (
    "time"
    "fmt"
    //custom packages
	"AlienStream/Models"
    "AlienStream/Services/db"
    "AlienStream/Services/requester"
    "labix.org/v2/mgo/bson"
)

type TempTrack struct {
    Posts []requester.ExtPost  `bson:"Posts"`
    Terms []string             `bson:"Terms"`
}

// Fetch New Content
func RefreshContent() {

    // Get all the temporary tracks from R and convert them to full track objects

    // Connect to our database
    session := db.Connection()
    defer session.Close()

    // Specify the TrackQueue Collection
    collection := session.DB("alien").C("trackqueue")
    // get all tracks
    temptracks := []TempTrack{}

    err := collection.Find(bson.M{}).All(&temptracks)
    if(err !=nil) {
        panic(err)
    }

    // For each track check if we already have one with the same name, if not insert
    for _,temptrack := range temptracks {
        var track models.Track
        // Add each post as a source
        for _,post := range temptrack.Posts {
            var provider models.Provider
            provider.Channel = post.Provider_Channel
            provider.Description = post.Provider_Description
            provider.Thumbnail = post.Provider_Thumbnail
            provider.Title = post.Provider_Title
            provider.Tags = []string{}
            provider.Type = post.Provider
            provider.Source = models.Source{
                URL: post.Source,
                Title: post.Source,
            }
            provider.Permalink = post.Permalink
            provider.URL = post.Url

            track.Providers = append(track.Providers, provider)


        }

        // Extract track properties from first post
        // TODO add artist lookup and time data
        track.Artist = track.Providers[0].Channel
        track.Description = track.Providers[0].Description
        track.LastAccess = time.Now().String()
        track.LastUpdate = time.Now().String()
        track.Rank = temptrack.Posts[0].Likes
        track.Title = track.Providers[0].Title
        fmt.Printf("found: %s \n", track.Title)
        track.Thumbnail = track.Providers[0].Thumbnail 

        track.Upsert()
    }

    // remove all the temp tracks when done
    collection.Remove(bson.M{})

   
    
}