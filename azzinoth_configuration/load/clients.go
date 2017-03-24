package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Clients(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT reference, name FROM clients`)
    if err != nil {
        return err
    }
    
    generalConfiguration.Clients = make(map[string]utils.ClientStruct)
    
    for rows.Next() {
        var (
            referenceB, nameB []byte
            str utils.ClientStruct
        )
        err = rows.Scan(&referenceB, &nameB)
        if err != nil {
            return err
        }
        
        if len(referenceB) == 0 {
            return errors.New("Error with this reference client ! -> "+string(referenceB))
        }
        str.Reference = string(referenceB)
        
        if len(nameB) == 0 {
            return errors.New("Error with this name client ! -> "+string(nameB))
        }
        str.Name = string(nameB)
        
        generalConfiguration.Clients[str.Reference]=str
        
    }
    return nil
}