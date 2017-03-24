package load

import(
    "errors"
    "strconv"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func RegularsTimetable(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT id, service_reference, config_reference, process_id FROM regular_timetable`)
    if err != nil {
        return err
    }
    
    generalConfiguration.ProcessRegularsTimetable = make(map[string]map[string]utils.RegularTimetableStruct)
    generalConfiguration.RegularsTimetable = make(map[string]utils.RegularTimetableStruct)
    
    for rows.Next() {
        var (
            idB, serviceReferenceB, configReferenceB, processIdB []byte
            str utils.RegularTimetableStruct
        )
        err = rows.Scan(&idB, &serviceReferenceB, &configReferenceB, &processIdB)
        if err != nil {
            return err
        }
        
        if len(idB) == 0 {
            return errors.New("Error with this idB services ! -> "+string(idB))
        }
        str.Id = string(idB)
        
        if len(serviceReferenceB) == 0 {
            return errors.New("Error with this serviceReferenceB services ! -> "+string(serviceReferenceB))
        }
        str.ServiceReference = string(serviceReferenceB)
        
        if len(configReferenceB) == 0 {
            return errors.New("Error with this configReferenceB services ! -> "+string(configReferenceB))
        }
        str.ConfigReference = string(configReferenceB)
        
        if len(processIdB) == 0 {
            return errors.New("Error with this processIdB services ! -> "+string(processIdB))
        }
        str.ProcessId = string(processIdB)
        
        if process, ok := generalConfiguration.ProcessRegularsTimetable[str.ProcessId]; ok {
            process[str.Id]=str 
            generalConfiguration.ProcessRegularsTimetable[str.ProcessId]=process
        } else {
            processTimetable := make(map[string]utils.RegularTimetableStruct)
            processTimetable[str.Id]=str 
            generalConfiguration.ProcessRegularsTimetable[str.ProcessId]=processTimetable
        }
    }
    
    for process_id, process := range generalConfiguration.ProcessRegularsTimetable {
        for _, str := range process {
            
            rows, err = db.Query(`SELECT need_id FROM regular_need WHERE regular_id = `+str.Id)
            if err != nil {
                return err
            }
            
            for rows.Next() {
                var (
                    needIdB []byte
                )
                err = rows.Scan(&needIdB)
                if err != nil {
                    return err
                }
                
                if len(needIdB) == 0 {
                    return errors.New("Error with this needIdB services ! -> "+string(needIdB))
                }
                str.RegularNeed = append(str.RegularNeed,string(needIdB))
            }
            
            str.RegularDays = make(map[string]utils.RegularDayStruct)
                        
            rows, err = db.Query(`SELECT day, minimal_start_time, maximal_start_time, attemps FROM regular_days WHERE regular_id = `+str.Id)
            if err != nil {
                return err
            }
            
            for rows.Next() {
                var (
                    dayB, minimalStartTimeB, maximalStartTimeB, attempsB []byte
                    regularDay utils.RegularDayStruct
                )
                err = rows.Scan(&dayB, &minimalStartTimeB, &maximalStartTimeB, &attempsB)
                if err != nil {
                    return err
                }
                
                regularDay.RegularId = str.Id
                
                if len(dayB) == 0 {
                    return errors.New("Error with this dayB services ! -> "+string(dayB))
                }
                regularDay.Day = string(dayB)
                
                if err = regularDay.SetMinimalStartTime(string(minimalStartTimeB)); err != nil {
                    return errors.New("Error with this minimalStartTimeB, it is not parsable !-> "+string(minimalStartTimeB))
                }
                
                if err = regularDay.SetMaximalStartTime(string(maximalStartTimeB)); err != nil {
                    return errors.New("Error with this maximalStartTimeB, it is not parsable !-> "+string(maximalStartTimeB))
                }
                
                if len(attempsB) == 0 {
                    return errors.New("Error with this attempsB services ! -> "+string(attempsB))
                }
                
                regularDay.Attemps,_ = strconv.Atoi(string(attempsB))
                
                str.RegularDays[regularDay.Day] = regularDay
            }
            process[str.Id]=str
            generalConfiguration.RegularsTimetable[str.Id]=str
        }
        generalConfiguration.ProcessRegularsTimetable[process_id]=process
    }
    
    return nil
}