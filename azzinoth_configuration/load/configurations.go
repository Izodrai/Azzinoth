package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Configurations(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT reference, client_reference, path, stream_name FROM configurations`)
    if err != nil {
        return err
    }
    
    generalConfiguration.Configurations = make(map[string]utils.ConfigurationStruct)
    
    for rows.Next() {
        var (
            referenceB, clientReferenceB, pathB, streamNameB []byte
            str utils.ConfigurationStruct
        )
        err = rows.Scan(&referenceB, &clientReferenceB, &pathB, &streamNameB)
        if err != nil {
            return err
        }
        
        if len(referenceB) == 0 {
            return errors.New("Error with this referenceB services ! -> "+string(referenceB))
        }
        str.Reference = string(referenceB)
        
        if len(clientReferenceB) == 0 {
            return errors.New("Error with this clientReferenceB services ! -> "+string(clientReferenceB))
        }
        str.Client = string(clientReferenceB)
        
        if len(pathB) == 0 {
            return errors.New("Error with this pathB services ! -> "+string(pathB))
        }
        str.Path = string(pathB)
        
        if len(streamNameB) != 0 {
            str.StreamName = string(streamNameB)
        }
        
        generalConfiguration.Configurations[str.Reference]=str
        
    }
    return nil
}