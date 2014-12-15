package models

// A Provider is a streamable source such as soundcloud or youtube

type Provider struct { // youtube or soundcloud only
    Channel string // The content owner
    Description string // The description from the source
    Title string // The title from the source
    Tags []string // genres, etc
    URL string // If this matches we don’t even need to bother with clustering
    Type  string
    Source Source
    Permalink string
    Thumbnail string
}