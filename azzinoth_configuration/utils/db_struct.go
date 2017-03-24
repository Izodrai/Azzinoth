package utils

import (
//     "fmt"
    "time"
    "strconv"
)

const (
    StepLastActivity0 = "5m"
    StepLastActivity1 = "15m"
)
type GeneralConfiguration struct {
    
    Clients                         map[string]ClientStruct
    
    Servers                         map[string]ServerStruct
    DataServers                     map[string]ServerStruct
    CalcServers                     map[string]ServerStruct
    
    ServicesTypes                   map[string]ServiceTypeStruct
    Services                        map[string]ServiceStruct
    Configurations                  map[string]ConfigurationStruct
    
    ServicesHaveDataServers         map[string][]string
    DataServersHaveServices         map[string][]string
    
    ProcessRegularsTimetable        map[string]map[string]RegularTimetableStruct //key1 : processId | key2 : regular_id
//     ProcessRegularsTimetableSorted  map[string][]string                          //key1 : processId | regular_id sorted
    RegularsTimetable               map[string]RegularTimetableStruct            //key1 : regular_id
    
}

type ClientStruct struct {
    Reference   string
    Name        string
}

type ServerStruct struct {
    Reference   string
    ShortDNS    string
    DNS         string
    Type        string
}

func (d *ServerStruct) SetDNS() {
    d.DNS = d.Reference+d.ShortDNS
}

type ServiceTypeStruct struct {
    Reference   string
    Name        string
}

type ServiceStruct struct {
    Reference       string
    ServiceType     string
    ServerCalc      string
    LastActivity    time.Time
}

func (d *ServiceStruct) ErrorLastActivity() int {
    
    step0, _ := time.ParseDuration(StepLastActivity0)
    step1, _ := time.ParseDuration(StepLastActivity1)
    
    if time.Now().After(d.LastActivity.Add(step1)) {return 2}
    
    if time.Now().After(d.LastActivity.Add(step0)) {return 1}
    
    return 0
}

func (d *ServiceStruct) GetLastActivity() string {
    return d.LastActivity.Format("2006-01-02 15:04:05")
}

func (d *ServiceStruct) SetLastActivity(t string) error {
    var err error
    if d.LastActivity, err = time.Parse("2006-01-02 15:04:05",t); err != nil {return err}
    return nil
}

type ConfigurationStruct struct {
    Reference       string
    Client          string
    Path            string
    StreamName      string
}

type RegularTimetableStruct struct {
    Id                  string
    ServiceReference    string                          // reference of the service for this regular
    ConfigReference     string                          // reference of the configuration for this regular
    ProcessId           string                          // process id of this regular
    RegularNeed         []string                        // regular need for exec
    RegularDays         map[string]RegularDayStruct     // key = day | value = RegularDayStruct
}

type RegularDayStruct struct {
    RegularId           string
    Day                 string
    Attemps             int
    MinimalStartTime    time.Time 
    MaximalStartTime    time.Time
}

func (d *RegularDayStruct) GetAttempsString() string {
    return strconv.Itoa(d.Attemps)
}

func (d *RegularDayStruct) MinimalStartTimeSet() bool {
    if d.GetMinimalStartTime() != "00:00:00" {
        return true
    }
    return false
}

func (d *RegularDayStruct) GetMinimalStartTime() string {
    return d.MinimalStartTime.Format("15:04:05")
}

func (d *RegularDayStruct) SetMinimalStartTime(t string) error {
    var err error
    if d.MinimalStartTime, err = time.Parse("15:04:05",t); err != nil {return err}
    return nil
}

func (d *RegularDayStruct) MaximalStartTimeSet() bool {
    if d.GetMaximalStartTime() != "23:59:59" {
        return true
    }
    return false
}

func (d *RegularDayStruct) GetMaximalStartTime() string {
    return d.MaximalStartTime.Format("15:04:05")
}

func (d *RegularDayStruct) SetMaximalStartTime(t string) error {
    var err error
    if d.MaximalStartTime, err = time.Parse("15:04:05",t); err != nil {return err}
    return nil
}