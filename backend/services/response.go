package services

type JsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface {} `json:"data.omitresponse"` // the data itself, can be of type interface since data can be of any kind of structure
}