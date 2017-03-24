package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func ServicesTypes(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("All services types ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var strs []string
    
    for _,str := range generalConfiguration.ServicesTypes {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    fmt.Println("")
    
    for _,reference := range strs {
        fmt.Println(utils.CyanDk("   > "+generalConfiguration.ServicesTypes[reference].Reference))
        fmt.Println(utils.WhitLt("      > "+generalConfiguration.ServicesTypes[reference].Name))
        fmt.Println("")
    }
    fmt.Println("")
}