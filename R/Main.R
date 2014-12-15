# vendor libraries
library("cluster")
library("tm")
library("SnowballC")
library("foreach")
library("slam")


# includes
source("./db.R")
source("./util.R")

similarityscore <- function(TrackA, TrackB) {
	termsimilarity = 0
	for(i in 1:length(TrackA$Terms)) {
		termsimilarity = termsimilarity + is.element(names(TrackA$Terms)[[i]], names(TrackB$Terms))
	} 
	termsimilarity = termsimilarity / length(TrackA$Terms)

	titlesimiarity = 1 - sum(adist(TrackA$Posts[[1]]$provider_title,TrackB$Posts[[1]]$provider_title)) / nchar(TrackA$Posts[[1]]$provider_title)

	# special cases

	if(grepl("remix",TrackA$Posts[[1]]$provider_title) && !grepl("remix",TrackB$Posts[[1]]$provider_title)) {
		titlesimiarity = titlesimiarity*.5
	} else if(!grepl("remix",TrackA$Posts[[1]]$provider_title) && grepl("remix",TrackB$Posts[[1]]$provider_title)) {
		titlesimiarity = titlesimiarity*.5
	}

	if(grepl("cover",TrackA$Posts[[1]]$provider_title) && !grepl("cover",TrackB$Posts[[1]]$provider_title)) {
		titlesimiarity = titlesimiarity*.5
	} else if(!grepl("remix",TrackA$Posts[[1]]$provider_title) && grepl("cover",TrackB$Posts[[1]]$provider_title)) {
		titlesimiarity = titlesimiarity*.5
	}
	return((0.8*titlesimiarity) + (0.2*termsimilarity))
}

remove_outside_terms <- function (x, lowfreq = 0, highfreq = Inf) {
    stopifnot(inherits(x, c("DocumentTermMatrix", "TermDocumentMatrix")), 
        is.numeric(lowfreq), is.numeric(highfreq))
    if (inherits(x, "DocumentTermMatrix")) 
        x <- t(x)
    rs <- slam::row_sums(x)
    names(rs[rs > lowfreq & rs <= highfreq])
    x[names(rs[rs >= lowfreq & rs < highfreq]),]
}

duplicate_prediction <- function() {
	print("duplicate prediction")

	min_word_length = 5
	max_word_length = Inf
	max_term_occurance = 100
	min_term_occurance = 2

	# TODO: Switch to pulling only sys flagged as music

	# Get All The Posts
	Posts = db.getAllPosts()
	posts.corpus = Corpus(VectorSource(Posts[,"provider_title"]))
	posts.corpus.copy = posts.corpus

	# remove things we don't want
	posts.corpus = tm_map(posts.corpus,content_transformer(tolower))
	posts.corpus = tm_map(posts.corpus, removePunctuation)
	posts.corpus = tm_map(posts.corpus, removeNumbers)
	posts.corpus = tm_map(posts.corpus, removeWords, stopwords('english'))

	# grab any words longer than 5 characters
	posts.TmDoc = TermDocumentMatrix(posts.corpus, control=list(wordLengths=c(min_word_length, max_word_length)))

	# remove all but the cluster terms
	posts.TmDoc = remove_outside_terms(posts.TmDoc, min_term_occurance, max_term_occurance)
	
	Tracks = list()
	Tracks = foreach(i = 1:length(Posts[,1])) %dopar% {
		Track = list()
		Track$Posts = list()
		Track$Posts[[1]] = Posts[i,]
		sums = slam::row_sums(posts.TmDoc[,i])
		Track$Terms = names(which(sums != 0))
		names(Track$Terms) = names(which(sums != 0))
		return(Track)
	}

	Clusters = list()

	# find all clusters
	Clusters = foreach(i = 1:posts.TmDoc$nrow) %dopar% {
		sums = slam::col_sums(posts.TmDoc[i,])
		return(which(sums != 0))
	}
	names(Clusters) = names(slam::row_sums(posts.TmDoc[,1]))

	# Mega Merge

	mergethresh = .8

	for(i in 1:length(Tracks)) {
		Track = Tracks[[i]]

		# Skip if track is pure or merged
		if(!is.null(Track$merged) || length(Track$Terms) == 0) {
			next
		}

		# For Each Term Check all matches and merge
		for(q in 1:length(Track$Terms)) {
			term = Track$Terms[[q]]

			for(w in 1:length(Clusters[[term]])) {
				match = Tracks[[Clusters[[term]][[w]]]]
				
				#skip if i match myself
				if(Track$Posts[[1]][["_id"]] == match$Posts[[1]][["_id"]]) {
					next
				}
				# If the track matches
				if(similarityscore(Track, match) < mergethresh) {
					next
				}

				# Check Where it's from
				sourceflag = FALSE
				for(j in 1:length(Track$Posts)){
					for(k in 1:length(Track$Posts)){
						if(Track$Posts[[j]]$source == match$Posts[[1]]$source) {
							sourceflag = TRUE
						}
					}
				}


				# If In Same Subreddit we need to filter out duplicates
				if(sourceflag) {
					# We Found A Duplicate! remove it
					duplicate = match

					if(is.null(duplicate$merged)) {
						Tracks[[Clusters[[term]][[w]]]]$merged = TRUE 
					}
				} 

				# If not the same source we need to add a reference in the Track Object 
				# and flag the duplicate
				else {
					duplicate =  match

					# printf("Merged! %s: Matched %s With: %s\n", duplicate$Posts[[1]]$title,Track$Posts[[1]]$title, term)
					# We Want to Save The reference to the original Post
				
					# Loop through the rest of the tracks and flag all hit
					if(is.null(duplicate$merged)) {
						Tracks[[i]]$Posts[[length(Tracks[[i]]$Posts)+1]] = duplicate$Posts[[1]]
						Tracks[[Clusters[[term]][[w]]]]$merged = TRUE 
					}

				}
			}
		}
	}

	# Update Post Prediction Values
	# Predicted_Duplicate = Posts[which(vapply(Tracks,function(Track){!is.null(Track$merged) || length(Track$Posts) > 1},FALSE)),]
	# foreach(Post = Predicted_Duplicate) %do% {
	#	Post$duplicate_sysflag = 1
	#	query = sprintf('{"url":"%s","source":"%s"}',Post$url,Post$source)
	#	db.updatePost(query, Post)
	#  }

	#Pure_Posts = Posts[which(vapply(Tracks,function(Track){length(Track$Posts) == 1},FALSE)),]
	#foreach(Post = Pure_Posts ) %do% {
	#	Post$duplicate_sysflag = 0
	#	query = sprintf('{"url":"%s","source":"%s"}',Post$url,Post$source)
	#	db.updatePost(query, Post)
	#

	# sort priority determines merge

	for(j in 1:length(Tracks)) {
		for(k in (j+1):length(Tracks)) {
			if(k > length(Tracks)) {
				next
			}

			if(is.null(Tracks[[j]]$merged)) {
				for(l in 1:length(Tracks[[j]]$Posts)) {
					Post = Tracks[[j]]$Posts[[l]]
					#skip if i match myself
					if(Tracks[[k]]$Posts[[1]][["_id"]] == Post[["_id"]]) {
						next
					}
					if(Tracks[[k]]$Posts[[1]]$url == Post$url) {
						Tracks[[j]]$Posts[[length(Tracks[[j]]$Posts)+1]] = Tracks[[k]]$Posts[[1]]
						Tracks[[k]]$merged = TRUE 
					}
				}
			}
		}
	}

	# filter out merged tracks and update the queue
	Tracks = Tracks[which(vapply(Tracks,function(Track){is.null(Track$merged)},FALSE))]
	foreach(Track = Tracks ) %do% {
		db.insertTrack(Track)
	}
	misses = list()

}

nonmusic_prediction <- function() {
	print("nonmusic prediction")
	posts = db.getPosts('{"Content_sysflag": -1}')

	# text mining

	# source
	remix = grepl("remix",temp[,"title"])
	posts = cbind(posts,title_contains_remix=remix)

	remix = grepl("interview",temp[,"title"])
	posts = cbind(posts,title_contains_interview=remix)


	# provider
	provider_remix = grepl("remix",temp[,"provider_title"])
	posts = cbind(posts,provider_contains_remix=remix)

	provider_remix = grepl("interview",temp[,"provider_title"])
	posts = cbind(posts,provider_contains_interview=remix)


}

Run <- function() {
	print("Prediction Running")
	
	while(TRUE) {
		Sys.sleep(60*60)
		duplicate_prediction()
		#nonmusic_prediction()
		
	}	

}

# Run The Loop
Run()
