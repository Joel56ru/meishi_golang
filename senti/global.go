package senti

type UserJWT struct {
	StuffsID int64                  `json:"staffs_id"`
	Role     string                 `json:"role"`
	ReqP     int64                  `json:"req_p"`
	SecureIp bool                   `json:"secure_ip"`
	IP       string                 `json:"ip"`
	IsAdmin  bool                   `json:"is_admin"`
	Staff    map[string]interface{} `json:"staff,omitempty"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required,min=4,max=64"`
	Password string `json:"password" binding:"required,min=4,max=64"`
	FIO      string `json:"name" binding:"required,min=2,max=128"`
}

//LogUser Структура логирования сотрудников
type LogUser struct {
	NewVal           string `json:"newval" db:"newval"`
	OldVal           string `json:"oldval" db:"oldval"`
	ParentId         int64  `json:"parent_id" db:"parent_id"`
	ChildId          int64  `json:"child_id" db:"child_id"`
	ChildAspirantId  int64  `json:"child_aspirant_id" db:"child_aspirant_id"`
	ChildClientId    int64  `json:"child_client_id" db:"child_client_id"`
	ChildInterviewId int64  `json:"child_interview_id" db:"child_interview_id"`
	ChildOrderId     int64  `json:"child_order_id" db:"child_order_id"`
	Tip              string `json:"tip" db:"tip"`
	Ts               string `json:"ts" db:"ts"`
}

type CategoryUsersST struct {
	Tip       string `json:"tip" db:"tip" binding:"required"`
	Person    string `json:"person" db:"person" binding:"required"`
	Personend string `json:"personend" db:"personend" binding:"required"`
}

type StaffsRole struct {
	List *[]CategoryUsersST `json:"list"`
}

type AdressTextSearchInput struct {
	Input string `json:"text" binding:"required,min=1,max=150"`
}

//GeoCode Структура получения ответа от dadata
type GeoCode struct {
	FullName          string  `json:"fullName"`
	UnrestrictedValue string  `json:"unrestricted_value"`
	Metro             Metro   `json:"metro"`
	GeoLat            float64 `json:"geo_lat"`
	GeoLon            float64 `json:"geo_lon"`
	Guid              string  `json:"guid"`
	City              string  `json:"city"`
	Region            string  `json:"region"`
	Country           string  `json:"country"`
	PostalCode        string  `json:"postal_code"`
	Settlement        string  `json:"settlement"`
	Street            string  `json:"street"`
	House             string  `json:"house"`
	Block             string  `json:"block"`
}
type Metro struct {
	MetroId  int64   `json:"metro_id" db:"metro_id"`
	GeoLon   float64 `json:"geo_lon" db:"geo_lon"`
	GeoLat   float64 `json:"geo_lat" db:"geo_lat"`
	Title    string  `json:"title" db:"title"`
	LineName string  `json:"line_name" db:"line_name"`
	City     string  `json:"city" db:"city"`
	LineId   string  `json:"line_id" db:"line_id"`
	Distance float64 `json:"distance"`
}
