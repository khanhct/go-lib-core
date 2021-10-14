package response

type Responses interface {
	SetCode(int)
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses
}

type Response struct {
	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Code      int  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type HttpResponse response

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}
