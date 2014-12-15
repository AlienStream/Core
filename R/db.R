source("./util.R")

library(rmongodb)
library(rjson)

host = "localhost"
username = ""
password = ""
source("./pass.R")

# Setup Our Connection On Source
db <- mongo.create(host = host, username = username, password = password, db="alien")


# Used for offline Analysis
db.useCache <- function() {
	return(list_to_data(fromJSON(file="Data/posts.json")))
}
db.updateCache <- function() {
	if (mongo.is.connected(db)) {
		posts = mongo.find.all(db, "post")
		for (i in 1:length(posts) ) {	
			posts[[i]] = lapply(posts[[i]],function(x) {
				if(typeof(x) == "character") {
					Encoding(x) <- "UTF-8"
					x = iconv(x, to="ascii",sub='')
					return(iconv(x, to="UTF-8",sub=''))
				}

				return(x)
			})
		}
		writeLines(toJSON(posts, method="C" ),"Data/posts.json")
		print("Cache Updated")
	} else {
		print("MongoDB Unreachable, Please make sure the server is accessible")
	}
}


# Get Absolutely All Posts
db.getAllPosts <- function() {
	if (mongo.is.connected(db)) {
		return(list_to_data(mongo.find.all(db, "post")))
	} else {
		return(cache)
	}
}

# Query Posts DB
db.getPosts <- function(query = "{}") {
	if (mongo.is.connected(db)) {	
		return(mongo.find.all(db, "post",query=query))
	} else {
		return(cache)
	}
}

# Update A Single Post
db.updatePost <- function(query, data) {
	if (mongo.is.connected(db)) {
		data = lapply(data,function(x) {
			if(typeof(x) == "character") {
				Encoding(x) <- "UTF-8"
				x = iconv(x, to="ascii",sub='')
				return(iconv(x, to="UTF-8",sub=''))
			}

			return(x)
		})


		return(mongo.insert(db, "posttest", query, toJSON(data)))
	} else {
		print("MongoDB Unreachable, Please make sure the server is accessible")
	}
}

# Insert A track into the process queue
db.insertTrack <- function(data) {
	if (mongo.is.connected(db)) {
		data = lapply(data,function(x) {
			if(typeof(x) == "character") {
				Encoding(x) <- "ascii"
				return(iconv(x, to="UTF-8",sub=''))
			}

			return(x)
		})

		return(mongo.insert(db, "trackqueue", data))
	} else {
		print("MongoDB Unreachable, Please make sure the server is accessible")
	}
}

if (mongo.is.connected(db)) {
	print("Mongo Connection Established, Ready For Query")	
} else {
	print("MongoDB Unreachable, Please make sure the server is accessible")
	print("Using Cached Data")
	cache <<- db.useCache()
}
