package bootdev

type RequestDTO struct {
	Body string `json:"body"`
}

func (req *RequestDTO) Length() int {
	return len(string(req.Body))
}
