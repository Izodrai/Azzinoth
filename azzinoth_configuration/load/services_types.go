package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ServicesTypes(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT reference, name FROM services_types`)
    if err != nil {
        return err
    }
    
    generalConfiguration.ServicesTypes = make(map[string]utils.ServiceTypeStruct)
    
    for rows.Next() {
        var (
            referenceB, nameB []byte
            str utils.ServiceTypeStruct
        )
        err = rows.Scan(&referenceB, &nameB)
        if err != nil {
            return err
        }
        
        if len(referenceB) == 0 {
            return errors.New("Error with this reference services types ! -> "+string(referenceB))
        }
        str.Reference = string(referenceB)
        
        if len(nameB) == 0 {
            return errors.New("Error with this name services types ! -> "+string(nameB))
        }
        str.Name = string(nameB)
        
        generalConfiguration.ServicesTypes[str.Reference]=str
        
    }
    return nil
}