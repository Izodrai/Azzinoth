package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func Servers(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("All servers ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var strs []string
    
    for _,str := range generalConfiguration.Servers {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    fmt.Println("")
    
    for _,reference := range strs {
        fmt.Println(utils.CyanDk("   > Server Reference :  "+generalConfiguration.Servers[reference].Reference))
        fmt.Println(utils.WhitLt("      > Server DNS :     "+generalConfiguration.Servers[reference].DNS))
        
        if generalConfiguration.Servers[reference].Type == "data" {
            fmt.Println(utils.PinkLt("         > Server Type : "+generalConfiguration.Servers[reference].Type))
        } else {
            fmt.Println(utils.CyanLt("         > Server Type : "+generalConfiguration.Servers[reference].Type))
        }
        fmt.Println("")
    }
    fmt.Println("")
}

func DataServers(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("Data servers :"))
    
    var strs []string
    
    for _,str := range generalConfiguration.DataServers {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    for _,reference := range strs {
        fmt.Println(utils.CyanDk("   > "+generalConfiguration.Servers[reference].Reference))
        fmt.Println(utils.WhitLt("      > "+generalConfiguration.Servers[reference].DNS))
    }
    fmt.Println("")
}

func CalcServers(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("Calc servers :"))
    
    var strs []string
    
    for _,str := range generalConfiguration.CalcServers {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    for _,reference := range strs {
        fmt.Println(utils.CyanDk("   > "+generalConfiguration.Servers[reference].Reference))
        fmt.Println(utils.WhitLt("      > "+generalConfiguration.Servers[reference].DNS))
    }
    fmt.Println("")
}