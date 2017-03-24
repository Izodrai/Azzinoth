package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Servers(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT reference, dns, type FROM servers`)
    if err != nil {
        return err
    }
    
    generalConfiguration.Servers = make(map[string]utils.ServerStruct)
    generalConfiguration.DataServers = make(map[string]utils.ServerStruct)
    generalConfiguration.CalcServers = make(map[string]utils.ServerStruct)
    
    for rows.Next() {
        var (
            referenceB, dnsB, typeB []byte
            server utils.ServerStruct
        )
        err = rows.Scan(&referenceB, &dnsB, &typeB)
        if err != nil {
            return err
        }
        
        if len(referenceB) == 0 {
            return errors.New("Error with this reference server ! -> "+string(referenceB))
        }
        server.Reference = string(referenceB)
        
        if len(dnsB) == 0 {
            return errors.New("Error with this dns server ! -> "+string(dnsB))
        }
        server.ShortDNS = string(dnsB)
        
        if len(typeB) == 0 {
            return errors.New("Error with this type server ! -> "+string(typeB))
        }
        
        if string(typeB) != "calc" && string(typeB) != "data" {
            return errors.New("Error with this type server ! -> "+string(typeB))
        }
        
        server.Type = string(typeB)
        
        server.SetDNS()
        
        
        generalConfiguration.Servers[server.Reference]=server
        
        if server.Type == "calc" {
            generalConfiguration.CalcServers[server.Reference]=server
        } else {
            generalConfiguration.DataServers[server.Reference]=server
        }
    }
    return nil
}