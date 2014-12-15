package logging

type Error struct{
	Status string
	Severity string
	Response string
}


func Error_500(string response) models.Error {
    error := models.Error{};
    error.Status = "Failure"
    error.Severity = "Critical"
    error.Response = response

    fmt.Print("CRITICAL ERROR")
}

func Error_404(string response) models.Error {
    error := models.Error{};
    error.Status = "Failure"
    error.Severity = "Notice"
    error.Response = response
}

func Error_200(string response) models.Error {
    error := models.Error{};
    error.Status = "Success"
    error.Severity = "None"
    error.Response = response
}