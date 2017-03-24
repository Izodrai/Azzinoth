package load

import(
    "errors"
    "../utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ServicesHaveDataServers(db *sql.DB, generalConfiguration *utils.GeneralConfiguration) error {
    
    rows, err := db.Query(`SELECT service_reference, server_reference FROM services_have_data_servers`)
    if err != nil {
        return err
    }
    
    generalConfiguration.ServicesHaveDataServers = make(map[string][]string)
    generalConfiguration.DataServersHaveServices = make(map[string][]string)
    
    for rows.Next() {
        var (
            serviceReferenceB, serverReferenceB []byte
            serviceReference, serverReference string
        )
        err = rows.Scan(&serviceReferenceB, &serverReferenceB)
        if err != nil {
            return err
        }
        
        if len(serviceReferenceB) == 0 {
            return errors.New("Error with this serviceReferenceB ! -> "+string(serviceReferenceB))
        }
        serviceReference = string(serviceReferenceB)
        
        if len(serverReferenceB) == 0 {
            return errors.New("Error with this serverReferenceB ! -> "+string(serverReferenceB))
        }
        serverReference = string(serverReferenceB)
        
        if dataServers, ok := generalConfiguration.ServicesHaveDataServers[serviceReference]; ok {
            dataServers = append(dataServers,serverReference)
            generalConfiguration.ServicesHaveDataServers[serviceReference] = dataServers
        } else {
            generalConfiguration.ServicesHaveDataServers[serviceReference] = []string{serverReference,}
        }
        
        if services, ok := generalConfiguration.DataServersHaveServices[serverReference]; ok {
            services = append(services,serviceReference)
            generalConfiguration.DataServersHaveServices[serverReference] = services
        } else {
            generalConfiguration.DataServersHaveServices[serverReference] = []string{serviceReference,}
        }
    }
    return nil
}