package tc

type RouterOnyma struct {
	Id                    string     `json:"-" db:"id"`
	Branch_service        string     `json:"branch_service" db:"branch_service"`
	Router_name           string     `json:"router_name" db:"router_name"`
	Interface_name        string     `json:"interface_name" db:"interface_name"`
	Interface_description string     `json:"interface_description" db:"interface_description"`
	In_policy_router      string     `json:"in_policy_router" db:"in_policy_router"`
	In_speed_router       string     `json:"in_speed_router" db:"in_speed_router"`
	Out_policy_router     string     `json:"out_policy_router" db:"out_policy_router"`
	Out_speed_router      string     `json:"out_speed_router" db:"out_speed_router"`
	Ip_interface          string     `json:"ip_interface" db:"ip_interface"`
	Branch_contract       string     `json:"branch_contract" db:"branch_contract"`
	Dognum                string     `json:"dognum" db:"dognum"`
	Clsrv                 string     `json:"clsrv" db:"clsrv"`
	Company_name          string     `json:"company_name" db:"company_name"`
	In_speed_onyma        string     `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma       string     `json:"out_speed_onyma" db:"out_speed_onyma"`
	Insert_datetime       string     `json:"insert_datetime" db:"insert_datetime"`
	Iface_shutdown_router int        `json:"iface_shutdown_router" db:"iface_shutdown_router"`
	Client_status_onyma   int        `json:"client_status_onyma" db:"client_status_onyma"`
	DPPP_name             NullString `json:"dppp_name" db:"dppp_name" swaggertype:"string"`
}
type RouterOnymaGroupByRouter struct {
	Router_name           	string     `json:"router_name" db:"router_name"`
	Insert_date       		string     `json:"insert_date" db:"insert_date"`
	Count			      	int         `json:"count" db:"count"`
}
type RouterOnymaGroupByInsertDate struct {
	Insert_date       		string     `json:"insert_date" db:"insert_date"`
	Count			      	int         `json:"count" db:"count"`
}
type ControlTimePauseHistory struct {
	Id                 string `json:"id" db:"id"`
	ControlTimePauseId string `json:"control_time_pause_id" db:"control_time_pause_id"`
	UserId             string `json:"user_id" db:"user_id"`
	Msg                string `json:"msg" db:"msg"`
	Created_at         string `json:"created_at" db:"created_at"`
}

type ControlTimePauseInsert struct {
	Id                    string `json:"id" db:"id"`
	Control_status        int    `json:"control_status" db:"control_status"`
	Created_at            string `json:"created_at" db:"created_at"`
	Router_onyma_speed_id string `json:"router_onyma_speed_id" db:"router_onyma_speed_id"`
}
type ControlTimePause struct {
	Id             string `json:"id" db:"id"`
	Control_status int    `json:"control_status" db:"control_status"`
	Created_at     string `json:"created_at" db:"created_at"`

	Router_onyma_speed_id string `json:"router_onyma_speed_id" db:"router_onyma_speed_id"`
	Branch_service        string `json:"branch_service" db:"branch_service"`
	Router_name           string `json:"router_name" db:"router_name"`
	Interface_name        string `json:"interface_name" db:"interface_name"`
	Interface_description string `json:"interface_description" db:"interface_description"`
	In_policy_router      string `json:"in_policy_router" db:"in_policy_router"`
	In_speed_router       string `json:"in_speed_router" db:"in_speed_router"`
	Out_policy_router     string `json:"out_policy_router" db:"out_policy_router"`
	Out_speed_router      string `json:"out_speed_router" db:"out_speed_router"`
	Ip_interface          string `json:"ip_interface" db:"ip_interface"`
	Branch_contract       string `json:"branch_contract" db:"branch_contract"`
	Dognum                string `json:"dognum" db:"dognum"`
	Clsrv                 string `json:"clsrv" db:"clsrv"`
	Company_name          string `json:"company_name" db:"company_name"`
	In_speed_onyma        string `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma       string `json:"out_speed_onyma" db:"out_speed_onyma"`
	Insert_datetime       string `json:"insert_datetime" db:"insert_datetime"`
	Iface_shutdown_router int    `json:"iface_shutdown_router" db:"iface_shutdown_router"`
	Client_status_onyma   int    `json:"client_status_onyma" db:"client_status_onyma"`
}
type ProblemRouterOnyma struct {
	Id                    string     `json:"id" db:"id"`
	Router_onyma_speed_id string     `json:"router_onyma_speed_id" db:"router_onyma_speed_id"`
	Branch_service        string     `json:"branch_service" db:"branch_service"`
	Router_name           string     `json:"router_name" db:"router_name"`
	Interface_name        string     `json:"interface_name" db:"interface_name"`
	Interface_description string     `json:"interface_description" db:"interface_description"`
	In_policy_router      string     `json:"in_policy_router" db:"in_policy_router"`
	In_speed_router       string     `json:"in_speed_router" db:"in_speed_router"`
	Out_policy_router     string     `json:"out_policy_router" db:"out_policy_router"`
	Out_speed_router      string     `json:"out_speed_router" db:"out_speed_router"`
	Ip_interface          string     `json:"ip_interface" db:"ip_interface"`
	Branch_contract       string     `json:"branch_contract" db:"branch_contract"`
	Dognum                string     `json:"dognum" db:"dognum"`
	Clsrv                 string     `json:"clsrv" db:"clsrv"`
	Company_name          string     `json:"company_name" db:"company_name"`
	In_speed_onyma        string     `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma       string     `json:"out_speed_onyma" db:"out_speed_onyma"`
	Problem_status        string     `json:"problem_status" db:"problem_status"`
	Insert_datetime       string     `json:"insert_datetime" db:"insert_datetime"`
	Updated_at            NullString `json:"updated_at" db:"updated_at" swaggertype:"string"`
	Client_status_onyma   int        `json:"client_status_onyma" db:"client_status_onyma"`
	Problem_status_old    NullString `json:"problem_status_old" db:"problem_status_old" swaggertype:"string"`
}
type ProblemRouterOnymaQuery struct {
	Id                    string     `json:"id" db:"id"`
	Router_onyma_speed_id string     `json:"router_onyma_speed_id" db:"router_onyma_speed_id"`
	Branch_service        string     `json:"branch_service" db:"branch_service"`
	Router_name           string     `json:"router_name" db:"router_name"`
	Interface_name        string     `json:"interface_name" db:"interface_name"`
	Interface_description string     `json:"interface_description" db:"interface_description"`
	In_policy_router      string     `json:"in_policy_router" db:"in_policy_router"`
	In_speed_router       string     `json:"in_speed_router" db:"in_speed_router"`
	Out_policy_router     string     `json:"out_policy_router" db:"out_policy_router"`
	Out_speed_router      string     `json:"out_speed_router" db:"out_speed_router"`
	Ip_interface          string     `json:"ip_interface" db:"ip_interface"`
	Branch_contract       string     `json:"branch_contract" db:"branch_contract"`
	Dognum                string     `json:"dognum" db:"dognum"`
	Clsrv                 string     `json:"clsrv" db:"clsrv"`
	Company_name          string     `json:"company_name" db:"company_name"`
	In_speed_onyma        string     `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma       string     `json:"out_speed_onyma" db:"out_speed_onyma"`
	Problem_status        string     `json:"problem_status" db:"problem_status"`
	Insert_datetime       string     `json:"insert_datetime" db:"insert_datetime"`
	Updated_at            NullString `json:"updated_at" db:"updated_at" swaggertype:"string"`
	Client_status_onyma   int        `json:"client_status_onyma" db:"client_status_onyma"`
	Problem_status_old    NullString `json:"problem_status_old" db:"problem_status_old" swaggertype:"string"`
	//Problem_router_onyma_history_id NullString `json:"problem_router_onyma_history_id" db:"id" swaggertype:"string"`
	Problem_router_onyma_speed_id NullString `json:"problem_router_onyma_speed_id" db:"problem_router_onyma_speed_id"   swaggertype:"string"`
	User_id                       NullString `json:"user_id" db:"user_id" swaggertype:"string"`
	Old_val                       NullString `json:"old_val" db:"old_val" swaggertype:"string"`
	New_val                       NullString `json:"new_val" db:"new_val" swaggertype:"string"`
	Msg                           NullString `json:"msg" db:"msg" swaggertype:"string"`
	Created_at                    NullString `json:"created_at" db:"created_at" swaggertype:"string"`
	UserName                      NullString `json:"name" db:"name" swaggertype:"string"`
	DPPP_name                     NullString `json:"dppp_name" db:"dppp_name" swaggertype:"string"`
}
type UpdateProblemRouterOnymaSpeedInput struct {
	Id               string `json:"id" db:"id" binding:"required"`
	Dognum           string `json:"dognum" db:"dognum"`
	Clsrv            string `json:"clsrv" db:"clsrv"`
	Company_name     string `json:"company_name" db:"company_name"`
	In_speed_onyma   string `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma  string `json:"out_speed_onyma" db:"out_speed_onyma"`
	In_speed_router  string `json:"in_speed_router" db:"in_speed_router"`
	Out_speed_router string `json:"out_speed_router" db:"out_speed_router"`
	Interface_name   string `json:"interface_name" db:"interface_name"`
	Branch_contract  string `json:"branch_contract" db:"branch_contract"`

	Problem_status     string `json:"problem_status" db:"problem_status" binding:"required"`
	Problem_status_old string `json:"problem_status_old" db:"_"`
	Msg                string `json:"msg" db:"msg" `
}

type ProblemRouterOnymaHistory struct {
	Id                            string `json:"-" db:"id"`
	Problem_router_onyma_speed_id string `json:"problem_router_onyma_speed_id" db:"problem_router_onyma_speed_id"  binding:"required"`
	User_id                       string `json:"user_id" db:"user_id"`
	Old_val                       string `json:"old_val" db:"old_val"`
	New_val                       string `json:"new_val" db:"new_val"`
	Msg                           string `json:"msg" db:"msg"`
	Created_at                    string `json:"_" db:"created_at"`
}
type ProblemRouterOnymaHistorySearch struct {
	Id                            string `json:"id" db:"id"`
	Problem_router_onyma_speed_id string `json:"problem_router_onyma_speed_id" db:"problem_router_onyma_speed_id"`
	User_id                       string `json:"user_id" db:"user_id"`
	Old_val                       string `json:"old_val" db:"old_val"`
	New_val                       string `json:"new_val" db:"new_val"`
	Msg                           string `json:"msg" db:"msg"`
	User_name                     string `json:"user_name" db:"name"`
	Created_at                    string `json:"created_at" db:"created_at"`

	ProblemRouterOnymaId  string     `json:"problem_router_onyma_id" db:"problem_router_onyma_id"`
	Router_onyma_speed_id string     `json:"router_onyma_speed_id" db:"router_onyma_speed_id"`
	Branch_service        string     `json:"branch_service" db:"branch_service"`
	Router_name           string     `json:"router_name" db:"router_name"`
	Interface_name        string     `json:"interface_name" db:"interface_name"`
	Interface_description string     `json:"interface_description" db:"interface_description"`
	In_policy_router      string     `json:"in_policy_router" db:"in_policy_router"`
	In_speed_router       string     `json:"in_speed_router" db:"in_speed_router"`
	Out_policy_router     string     `json:"out_policy_router" db:"out_policy_router"`
	Out_speed_router      string     `json:"out_speed_router" db:"out_speed_router"`
	Ip_interface          string     `json:"ip_interface" db:"ip_interface"`
	Branch_contract       string     `json:"branch_contract" db:"branch_contract"`
	Dognum                string     `json:"dognum" db:"dognum"`
	Clsrv                 string     `json:"clsrv" db:"clsrv"`
	Company_name          string     `json:"company_name" db:"company_name"`
	In_speed_onyma        string     `json:"in_speed_onyma" db:"in_speed_onyma"`
	Out_speed_onyma       string     `json:"out_speed_onyma" db:"out_speed_onyma"`
	Problem_status        string     `json:"problem_status" db:"problem_status"`
	Insert_datetime       string     `json:"insert_datetime" db:"insert_datetime"`
	Updated_at            NullString `json:"updated_at" db:"updated_at" swaggertype:"string"`
}
type FilterRouterOnyma struct {
	Id          string `json:"_" db:"id"`
	Router_name string `json:"router_name" db:"router_name"  binding:"required"`
	Filter_type string `json:"filter_type" db:"filter_type"  binding:"required"`
	Filter_val  string `json:"filter_val" db:"filter_val"  binding:"required"`
	Filter_desc string `json:"filter_desc" db:"filter_desc"`
	User_id     string `json:"user_id" db:"user_id"`
	Created_at  string `json:"__" db:"created_at"`
}
type FilterRouterOnymaSearch struct {
	Id          string `json:"id" db:"id"`
	Router_name string `json:"router_name" db:"router_name"`
	Filter_type string `json:"filter_type" db:"filter_type"`
	Filter_val  string `json:"filter_val" db:"filter_val"`
	Filter_desc string `json:"filter_desc" db:"filter_desc"`
	User_id     string `json:"user_id" db:"user_id"`
	User_name   string `json:"user_name" db:"name"`
	Created_at  string `json:"created_at" db:"created_at"`
}
