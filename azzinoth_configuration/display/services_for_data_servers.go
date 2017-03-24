package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func ServicesHaveDataServers(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("Services have data servers ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var services []string
    
    for service,_ := range generalConfiguration.ServicesHaveDataServers {
        services = append(services,service)
    }
    
    sort.Strings(services)
    
    fmt.Println("")
    
    for _,reference := range services {
        
        fmt.Println(utils.CyanDk("   > Name service :  "+reference))
        servers := generalConfiguration.ServicesHaveDataServers[reference]
        
        sort.Strings(servers)
        
        for _,server := range servers {
            fmt.Println(utils.WhitLt("      > Data Server :     "+server))
        }
        fmt.Println("")
    }
}

func DataServersHaveServices(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("Data servers have services ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var servers []string
    
    for server,_ := range generalConfiguration.DataServersHaveServices {
        servers = append(servers,server)
    }
    
    sort.Strings(servers)
    
    fmt.Println("")
    
    for _,reference := range servers {
        fmt.Println(utils.CyanDk("   > Name server :  "+reference))
        services := generalConfiguration.DataServersHaveServices[reference]
        
        sort.Strings(services)
        
        for _,service := range services {
            fmt.Println(utils.WhitLt("      > Service :     "+service))
        }
        fmt.Println("")
    }
}