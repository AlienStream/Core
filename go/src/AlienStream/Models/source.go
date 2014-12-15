package models

type Source struct {
  Channel           string // The content owner
	Description       string
	Title             string
	Thumbnail         string
  Published_at      string
	Type              string
	URL               string
}


func (source Source) ByString(sourcestring string) Source {
   	source.URL = sourcestring

   	// TODO: Lookup rest of info

    return source
}

func (s Source) ByStrings(sourcestrings []string) []Source {
	var sources []Source
    for _,sourcestring := range sourcestrings{
    	sources = append(sources,Source{}.ByString(sourcestring))
    }
    return sources
}
    