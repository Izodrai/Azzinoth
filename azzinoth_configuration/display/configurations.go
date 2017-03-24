package display

import(
    "fmt"
    "sort"
    "time"
    "../utils"
)

func Configurations(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("All configurations ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var strs []string
    
    for _,str := range generalConfiguration.Configurations {
        strs = append(strs,str.Reference)
    }
    
    sort.Strings(strs)
    
    fmt.Println("")
    
    
    for _,reference := range strs {
        
        configurations := generalConfiguration.Configurations[reference]
        
        if configurations.StreamName == "" {
            fmt.Println(utils.CyanDk("   > Configuration Reference (prod):            "+configurations.Reference))
        } else {
            fmt.Println(utils.CyanDk("   > Configuration Reference (data):            "+configurations.Reference))
        }
        fmt.Println(utils.WhitLt("      > Client Reference :                      "+configurations.Client))
        fmt.Println(utils.CyanLt("         > Configuration Path :                 "+configurations.Path))
        if configurations.StreamName != "" {
            fmt.Println(utils.PinkLt("            > Configuration Stream Name :       "+configurations.StreamName))
        }
        fmt.Println("")
    }
    fmt.Println("")
}