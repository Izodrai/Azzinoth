package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Services(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT reference, service_type_reference, calc_server_reference, last_activity FROM services`)
    if err != nil {
        return err
    }
    
    generalConfiguration.Services = make(map[string]utils.ServiceStruct)
    
    for rows.Next() {
        var (
            referenceB, serviceTypeReferenceB, calcServerReferenceB, lastActivityB []byte
            str utils.ServiceStruct
        )
        err = rows.Scan(&referenceB, &serviceTypeReferenceB, &calcServerReferenceB, &lastActivityB)
        if err != nil {
            return err
        }
        
        if len(referenceB) == 0 {
            return errors.New("Error with this referenceB services ! -> "+string(referenceB))
        }
        str.Reference = string(referenceB)
        
        if len(serviceTypeReferenceB) == 0 {
            return errors.New("Error with this serviceTypeReferenceB services ! -> "+string(serviceTypeReferenceB))
        }
        str.ServiceType = string(serviceTypeReferenceB)
        
        if len(calcServerReferenceB) == 0 {
            return errors.New("Error with this calcServerReferenceB services ! -> "+string(calcServerReferenceB))
        }
        str.ServerCalc = string(calcServerReferenceB)
        
        if len(lastActivityB) == 0 {
            return errors.New("Error with this lastActivityB services ! -> "+string(lastActivityB))
        }
        
        if err = str.SetLastActivity(string(lastActivityB)); err != nil {
            return errors.New("Error with this lastActivityB, it is not parsable !-> "+string(lastActivityB))
        }
        
        generalConfiguration.Services[str.Reference]=str
        
    }
    return nil
}