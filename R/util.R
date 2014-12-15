as.numeric.factor <- function(x) {as.numeric(levels(x))[x]}

printf <- function(...) cat(sprintf(...))

list_to_data <- function(list) {
	posts = lapply(list,function(x) {t(as.matrix(x))})
	data = do.call( rbind , posts )
	return(data)
} 