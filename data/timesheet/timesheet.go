package timesheet

// Define the structure matching the JSON format
type EntryDetail struct {
	Key                int      `json:"key"`
	DateRange          []string `json:"dateRange"`
	ValidationValue    int      `json:"validationValue"`
	ValidationType     int      `json:"validationType"`
	HdnSelTask         string   `json:"hdn_seltask"`
	HdnSelActivity     string   `json:"hdn_selactivity"`
	HdnSelProject      string   `json:"hdn_selproject"`
	HdnSelProjectPhase string   `json:"hdn_selprojectphase"`
	TaskDesc           string   `json:"task_desc"`
	Remark             string   `json:"remark"`
	MinuteDuration     int      `json:"minuteDuration"`
	HourDuration       int      `json:"hourDuration"`
	StartTime          string   `json:"starttime"`
	EndTime            string   `json:"endtime"`
	FlagStartDate      int      `json:"flag_startdate"`
	FlagEndDate        int      `json:"flag_enddate"`
	ChkActivityType    *string  `json:"chkactivitytype"`
	ChkBillable        *string  `json:"chkbillable"`
	FieldList          []string `json:"fieldList"`
	SrcFlag            string   `json:"srcflag"`
	IDTask             string   `json:"id_task"`
	DetailID           *string  `json:"detail_id"`
	ProjectCode        string   `json:"projectCode"`
	Location           *string  `json:"location"`
	Attachment         []string `json:"attachment"`
	UniqueID           int      `json:"uniqueId"`
	IsDisabled         bool     `json:"isDisabled"`
	IsNotComplete      bool     `json:"isNotComplete"`
	IsDisabledButton   bool     `json:"isDisabledButton"`
	CostCenterCode     string   `json:"costcenter_code"`
}

type DateEntry struct {
	HdnRowCount     int           `json:"HDN_ROWCOUNT"`
	Status          string        `json:"STATUS"`
	ActualStartTime string        `json:"ACTUALSTARTTIME"`
	ActualEndTime   string        `json:"ACTUALENDTIME"`
	ShiftCode       string        `json:"SHIFT_CODE"`
	Detail          []EntryDetail `json:"DETAIL"`
}

type Template struct {
	HdnRowCount int           `json:"HDN_ROWCOUNT"`
	Detail      []EntryDetail `json:"DETAIL"`
}

type Entries struct {
	RequestNo  string               `json:"REQUEST_NO"`
	ActRequest string               `json:"ACTREQUEST"`
	StartDate  string               `json:"STARTDATE"`
	EndDate    string               `json:"ENDDATE"`
	IsMax      bool                 `json:"ISMAX"`
	RequestBy  string               `json:"REQUESTBY"`
	RequestFor string               `json:"REQUESTFOR"`
	RefDoc     string               `json:"REFDOC"`
	Remark     string               `json:"REMARK"`
	Template   Template             `json:"TEMPLATE"`
	HdnField   []string             `json:"HDN_FIELD"`
	Dates      map[string]DateEntry `json:"DATES"`
}

type JSONStructure struct {
	Entries Entries `json:"ENTRIES"`
}
