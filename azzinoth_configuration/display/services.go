package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func Services(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("All services ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var strs []string
    
    for _,str := range generalConfiguration.Services {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    fmt.Println("")
    
    for _,reference := range strs {
        
        service := generalConfiguration.Services[reference]
        
        fmt.Println(utils.CyanDk("   > Service Reference :        "+service.Reference))
        fmt.Println(utils.WhitLt("      > Service Type :          "+service.ServiceType))
        fmt.Println(utils.CyanLt("         > Calculation Server : "+service.ServerCalc))
        
        switch service.ErrorLastActivity() {
            case 2: 
                fmt.Println(utils.RedLt_("             > Last Activity (>"+utils.StepLastActivity1+"): "+service.GetLastActivity()))
            case 1: 
                fmt.Println(utils.YellLt("             > Last Activity ("+utils.StepLastActivity0+">|<"+utils.StepLastActivity1+"): "+service.GetLastActivity()))
            case 0: 
                fmt.Println(utils.GreeLt("             > Last Activity (<"+utils.StepLastActivity0+"): "+service.GetLastActivity()))
        }
        fmt.Println("")
    }
    fmt.Println("")
}