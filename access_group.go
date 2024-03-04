package tc

type AccessGroup struct {
	Id            string     `json:"id" db:"id"`
	Router_Name   string     `json:"router_name" db:"router_name"`
	Iface_host    string     `json:"iface_host" db:"iface_host"`
	Ip            string     `json:"ip" db:"ip"`
	Iface_name    string     `json:"iface_name" db:"iface_name"`
	Iface_desc    string     `json:"iface_desc" db:"iface_desc"`
	Client_status int        `json:"client_status" db:"client_status"`
	In_policy     string     `json:"in_policy" db:"in_policy"`
	Out_policy    string     `json:"out_policy" db:"out_policy"`
	Access_group  string     `json:"access_group" db:"access_group"`
	Dognum        string     `json:"dognum" db:"dognum"`
	Clsrv         string     `json:"clsrv" db:"clsrv"`
	Created_at    string     `json:"created_at" db:"created_at"`
	Access_status string     `json:"access_status" db:"access_status"`
	Updated_at    NullString `json:"updated_at" db:"updated_at" swaggertype:"string"`
}
type UpdateAccessGroupInput struct {
	Id                string `json:"id" db:"id" binding:"required"`
	Access_status     string `json:"access_status" db:"access_status" binding:"required"`
	Access_status_old string `json:"access_status_old" db:"_"`
	Msg               string `json:"msg" db:"msg" `
}
type FilterAccessGroup struct {
	Id          string `json:"_" db:"id"`
	Router_name string `json:"router_name" db:"router_name"  binding:"required"`
	Filter_type string `json:"filter_type" db:"filter_type"  binding:"required"`
	Filter_val  string `json:"filter_val" db:"filter_val"  binding:"required"`
	Filter_desc string `json:"filter_desc" db:"filter_desc"`
	User_id     string `json:"user_id" db:"user_id"`
	Created_at  string `json:"__" db:"created_at"`
}
type FilterAccessGroupSearch struct {
	Id          string `json:"id" db:"id"`
	Router_name string `json:"router_name" db:"router_name"`
	Filter_type string `json:"filter_type" db:"filter_type"`
	Filter_val  string `json:"filter_val" db:"filter_val"`
	Filter_desc string `json:"filter_desc" db:"filter_desc"`
	User_id     string `json:"user_id" db:"user_id"`
	User_name   string `json:"user_name" db:"name"`
	Created_at  string `json:"created_at" db:"created_at"`
}

type AccessGroupHistory struct {
	Id              string `json:"-" db:"id"`
	Access_group_id string `json:"access_group_id" db:"access_group_id"  binding:"required"`
	User_id         string `json:"user_id" db:"user_id"`
	Old_val         string `json:"old_val" db:"old_val"`
	New_val         string `json:"new_val" db:"new_val"`
	Msg             string `json:"msg" db:"msg"`
	Created_at      string `json:"created_at" db:"created_at"`
	User_name       string `json:"user_name" db:"name"`
}

type GetAccessGroupHistory struct {
	Id  string   `json:"id" db:"id"`
	Ids []string `json:"ids" db:"id"`
}
