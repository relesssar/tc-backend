package tc

type ADLogLine struct {
	Datetime string
	EventID  int
	String   string
}
type ADLog struct {
	Id         string     `json:"id" db:"id"`
	Ad_login   string     `json:"ad_login" db:"ad_login"`
	DateTime   string     `json:"date_time" db:"date_time"`
	Eventid    int        `json:"eventid" db:"eventid"`
	Ip         string     `json:"ip" db:"ip"`
	Str        string     `json:"str" db:"str"`
	Fullname   NullString `json:"full_name" db:"full_name" swaggertype:"string"`
	Department NullString `json:"department" db:"department" swaggertype:"string"`
}
