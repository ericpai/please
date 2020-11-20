package adapter

type Response struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Tasks []Task `json:"tasks"`
}

type Request struct {
	Task `json:"task"`
}
