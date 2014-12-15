
# Build a Test Set
test = db.getAllPosts()[1:1000,]
train = db.getAllPosts()[1001:2000,]






# Clustering Tests

	# Title
	
		http://beyondvalence.blogspot.com/2014/01/text-mining-converting-tweet-text-list.html

		library(tm)
		posts.corpus = Corpus(VectorSource(cache[,"provider_title"]))

		# remove things we don't want

		posts.corpus = tm_map(posts.corpus,content_transformer(tolower))

		posts.corpus = tm_map(posts.corpus, removePunctuation)

		posts.corpus = tm_map(posts.corpus, removeNumbers)

		http://beyondvalence.blogspot.com/2014/01/text-mining-3-stemming-text-and.html

		library(SnowballC)

		posts.corpus = tm_map(posts.corpus, removeWords, stopwords('english'))

		posts.TmDoc = TermDocumentMatrix(posts.corpus, control=list(wordLengths=c(5, Inf)))

		# max occurance
		max_occur = max(apply(posts.TmDoc, 1, sum))
		which(apply(posts.TmDoc,1,sum) == max_occur)
		which(apply(posts.TmDoc, 1, sum) > 500)

		http://beyondvalence.blogspot.com/2014/01/text-mining-4-performing-term.html
		posts.TmDoc2 = removeSparseTerms(posts.TmDoc, sparse=0.99)
		posts.matrix2 = as.matrix(posts.TmDoc2)
		distMatrix = dist(scale(posts.matrix2))
		posts.fit = hclust(distMatrix, method="ward.D")
		plot(posts.fit, cex=.9, hang=-1, main="Word Cluster Dendrogram")



	# Description

		http://beyondvalence.blogspot.com/2014/01/text-mining-converting-tweet-text-list.html

		library(tm)
		posts.corpus = Corpus(VectorSource(cache[,"provider_description"]))

		# remove things we don't want

		posts.corpus = tm_map(posts.corpus,content_transformer(tolower))

		posts.corpus = tm_map(posts.corpus, removePunctuation)

		posts.corpus = tm_map(posts.corpus, removeNumbers)

		removeURLS = function(x) {gsub("http[[:alnum:]]*","",x)}

		posts.corpus = tm_map(posts.corpus, content_transformer(removeURLS))

		# maybe remove stopwords



		http://beyondvalence.blogspot.com/2014/01/text-mining-3-stemming-text-and.html

		library(SnowballC)

		posts.corpus = tm_map(posts.corpus, stemDocument)

		posts.TmDoc = TermDocumentMatrix(posts.corpus, control=list(wordLengths=c(5, Inf)))

		# max occurance
		which(apply(posts.TmDoc,1,sum) == max(apply(posts.TmDoc, 1, sum)))

		which(apply(posts.TmDoc, 1, sum) > 500)

		         a      about      after      again       aint         al      album 
		         4        128        606        631        764        875        919 
		       all       also      alway         am     amazon         an        and 
		      1078       1220       1274       1283       1319       1477       1527 
		       ani      anoth        are     around     arrang        art     artist 
		      1630       1712       2127       2265       2279       2313       2343 
		        as         at      audio      avail       away          b       back 
		      2400       2604       2753       2925       3014       3111       3157 
		      band       bass         be       beat     becaus      becom       been 
		      3405       3684       3886       3928       3992       4013       4040 
		     befor       best     better        big      black      break        but 
		      4068       4389       4428       4556       4859       5869       6561 
		       buy         by          c       call        can       cant       caus 
		      6595       6624       6667       6797       6920       6988       7425 
		        cd    channel      check       citi    classic       club       come 
		      7486       7770       7916       8466       8552       8741       9058 
		   comment     compos    contact  copyright      could      cover      creat 
		      9113       9220       9508       9725       9881       9970      10099 
		         d       danc       dark       date        day         de       dead 
		     10649      10810      10935      11008      11080      11162      11164 
		     death      debut       deep        did        die      digit     direct 
		     11208      11273      11384      12105      12114      12178      12299 
		  director         dj         do       dont       down   download      dream 
		     12313      12666      12798      12996      13193      13206      13332 
		      drum        dub    dubstep          e       edit        end      enjoy 
		     13512      13609      13663      13964      14160      14812      14916 
		        ep       even       ever      everi        eye       face   facebook 
		     15053      15521      15532      15550      15862      15922      15925 
		      fall       feat     featur       feel       film      final       find 
		     16225      16509      16516      16554      16852      16894      16909 
		     first     follow        for       free     friend       from         ft 
		     17009      17475      17529      17865      17991      18041      18105 
		      fuck       full       game       genr        get       girl       give 
		     18125      18190      18520      18901      19023      19239      19268 
		        go        god       good        got      great     guitar        guy 
		     19484      19499      19637      19735      19975      20396      20513 
		       had       hand        has       have         he       head       hear 
		     20623      20828      21077      21160      21228      21229      21283 
		     heart       help        her       here       high        him        his 
		     21296      21514      21579      21596      21766      21843      21915 
		       hit       home       hope       hous        how          i         if 
		     21935      22167      22293      22440      22465      22844      22988 
		       ill         im         in     includ  instagram       into         is 
		     23076      23135      23294      23347      23755      23997      24168 
		        it       itun        ive       jazz       john       just       keep 
		     24251      24301      24330      24694      25077      25412      25862 
		      know         la      label       last        let       life      light 
		     26524      27175      27185      27513      27973      28174      28211 
		       lik       like       link     listen      littl       live       long 
		     28243      28245      28366      28435      28476      28490      28727 
		      look       lost       love      lyric     lyrics       made       make 
		     28764      28844      28906      29259      29263      29392      29591 
		       man       mani        may         me      metal         mi       mind 
		     29715      29773      30350      30559      31024      31152      31430 
		       mix       more       most       much      music         my       name 
		     31698      32204      32321      32584      32789      32974      33164 
		      need      never        new      night         no        not        now 
		     33444      33668      33689      33862      34051      34283      34382 
		        of        off     offici         oh        old         on        one 
		     34803      34808      34831      34895      34976      35084      35100 
		      onli       open         or     origin      other        our        out 
		     35167      35253      35336      35421      35548      35588      35594 
		      over        own       page       part      peopl    perform      piano 
		     35679      35783      35886      36191      36618      36683      37101 
		     place       play      pleas     produc    product    purchas        put 
		     37406      37486      37524      38484      38513      39011      39063 
		    realli     record    records     releas      remix      right       rock 
		     39888      40015      40037      40366      40447      41063      41332 
		       run          s       said        say     second        see        set 
		     41857      41985      42101      42559      43062      43092      43485 
		     share        she       shit       show      singl         so       some 
		     43714      43768      43942      44074      44450      45317      45488 
		    someth       song       soon       soul      sound soundcloud      start 
		     45505      45541      45621      45705      45748      45762      46580 
		      stay      still     studio      style   subscrib       such        sun 
		     46626      46810      47193      47242      47351      47421      47557 
		   support       take      taken       tell         th       than      thank 
		     47706      48265      48272      48754      49007      49019      49022 
		      that        the      their       them      theme       then      there 
		     49032      49045      49098      49117      49125      49145      49182 
		     these       they        thi      thing      think       this    thought 
		     49208      49231      49238      49265      49275      49307      49375 
		   through       time       titl         to        too       tour      track 
		     49455      49637      49740      49819      49998      50143      50224 
		       tri       tune       turn    twitter        two         uk      under 
		     50536      50920      50976      51109      51305      51479      51672 
		        up     upload         us        use       veri    version      video 
		     52009      52036      52113      52132      52618      52671      52823 
		     vinyl      visit      vocal          w       want        was      watch 
		     53001      53095      53189      53434      53615      53690      53712 
		       way         we     websit       well       were       what       when 
		     53782      53807      53852      53944      53996      54079      54107 
		     where        whi      which      while        who       will       with 
		     54114      54124      54125      54132      54199      54342      54506 
		      wont       work      world      would    written          x       yeah 
		     54639      54704      54728      54771      54836      56265      56522 
		      year        you      young       your     youtub 
		     56529      56706      56722      56738      56756 


		http://beyondvalence.blogspot.com/2014/01/text-mining-4-performing-term.html

	# Get all the terms that show up at least twice but not more than 10 times


	library(tm)

		min_word_length = 5
	max_word_length = Inf

	# Get All The Posts
	Posts = db.getAllPosts()
	posts.corpus = Corpus(VectorSource(Posts[,"provider_title"]))

	# remove things we don't want
	posts.corpus = tm_map(posts.corpus,content_transformer(tolower))
	posts.corpus = tm_map(posts.corpus, removePunctuation)
	posts.corpus = tm_map(posts.corpus, removeNumbers)
	posts.corpus = tm_map(posts.corpus, removeWords, stopwords('english'))

	# grab any words longer than 5 characters
	posts.TmDoc = TermDocumentMatrix(posts.corpus, control=list(wordLengths=c(min_word_length, max_word_length)))

	# get the words that occur more than once, but not more than 10 times
	clustterms = names(which(apply(posts.TmDoc, 1, sum) > 1 & apply(posts.TmDoc, 1, sum) < 10))

	#
	Tracks = apply(Posts, 1, function(Post){
		Track = list()
		Track$Posts = list()
		Track$Posts[[1]] = Post
		Track$Terms = list()
		return(Track)
	})

	library(foreach)

	Clusters = list()
	Matches = list()

	# find all terms
	foreach(term = clustterms) %do% {
		matches = grep(term, Posts[,"provider_title"], ignore.case=TRUE)
		Matches[[term]] = matches
		# account for root words
		if(length(matches) < 10) {
			foreach(match = matches) %do% {
				Tracks[[match]]$Terms[term] = length(matches)
			}
		}
	}

	# done seperately as a 2nd pass once all terms are established
	foreach(term = clustterms) %do% {
		Clusters[[term]] = Tracks[Matches[[term]]];
	}

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

	# Mega Merge

	mergethresh = .8

	duplicates = list()
	removed_duplicates = list()

	for(i in 1:length(Tracks)) {
		Track = Tracks[[i]]

		if(is.null(Track$merged)) {
			# If Track is not pure
			if(length(Track$Terms) > 0) {
				# For Each Term
				for(j in 1:length(Track$Terms)) {
					# Get the list of matching tracks
					term = names(Track$Terms)[[j]]
					matches = Clusters[[term]]
					for(l in 1:length(matches)) {
						# If the track matches
						if(similarityscore(Track,matches[[l]]) > mergethresh) {
							# Check Where it's from

							# If In Same Subreddit we need to filter out duplicates
							if(Track$Posts[[1]]$source == matches[[l]]$Posts[[1]]$source) {
								# We Found A Duplicate! remove it
								duplicate =  matches[[l]]

								# Loop through the rest of the tracks and flag all hits
								for(k in (i+1):length(Tracks)) {
									if(k <= length(Tracks) && is.null(Tracks[[k]]$merged)) {
										if(Tracks[[k]]$Posts[[1]][["_id"]] == duplicate$Posts[[1]][["_id"]]) {
											if(duplicate$Posts[[1]]$url != Track$Posts[[1]]$url)
												printf("Removed %s: Matched %s With: %s(%f) in %s \n", duplicate$Posts[[1]]$url,Track$Posts[[1]]$url, term, similarityscore(Track,matches[[l]]), Track$Posts[[1]]$source)
											removed_duplicates = c(removed_duplicates,duplicate$Posts[[1]]$title)
											Tracks[[k]]$merged = TRUE 
										}
									}
								}
							} 

							# If not the same source we need to add a reference in the Track Object 
							# and flag the duplicate
							else {
								duplicate =  matches[[l]]

								

								# printf("Merged! %s: Matched %s With: %s\n", duplicate$Posts[[1]]$title,Track$Posts[[1]]$title, term)
								# We Want to Save The reference to the original Post
							
								# Loop through the rest of the tracks and flag all hits
								for(k in (i+1):length(Tracks)) {
									if(k <= length(Tracks) && is.null(Tracks[[k]]$merged)) {
										if(Tracks[[k]]$Posts[[1]][["_id"]] == duplicate$Posts[[1]][["_id"]]) {
											duplicates = c(duplicates,duplicate$Posts[[1]]$title)
											# printf("Merged! %s: Matched %s With: %s\n", duplicate$Posts[[1]]$url,Track$Posts[[1]]$url, term)
											Tracks[[i]]$Posts[[length(Tracks[[i]]$Posts) + 1]] = duplicate$Posts[[1]]
											Tracks[[k]]$merged = TRUE 
										}
									}
								}
							}
						} 
					}
				}

			}
		}
	}

	# filter out merged tracks
	Tracks = Tracks[which(vapply(Tracks,function(Track){is.null(Track$merged)},FALSE))]


	# filter out merged tracks
	Tracks = Tracks[which(lapply(Tracks,function(Track){is.null(Track$merged)}))]

	# For Each Track

		# Check 

