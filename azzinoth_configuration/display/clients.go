package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func Clients(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("All clients ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var clients []string
    
    for _,client := range generalConfiguration.Clients {
        clients = append(clients,client.Reference)
    }
    
    sort.Strings(clients)
    
    fmt.Println("")
    
    for _,reference := range clients {
        fmt.Println(utils.CyanDk("   > "+generalConfiguration.Clients[reference].Reference))
        fmt.Println(utils.WhitLt("     > "+generalConfiguration.Clients[reference].Name))
    }
    fmt.Println("")
}
