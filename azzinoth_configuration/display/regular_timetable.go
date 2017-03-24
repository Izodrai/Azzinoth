package display

import(
    "fmt"
    "sort"
    "time"
    "strconv"
    "../utils"
)

var days = []string{"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}

func RegularsTimetable(generalConfiguration *utils.GeneralConfiguration) {
    
    fmt.Println("")
    fmt.Println(utils.YellLt("regular timetable ("+time.Now().Format("2006-01-02 15:04:05")+"):"))
    
    var strs []string
    
    for processId,_ := range generalConfiguration.ProcessRegularsTimetable {
        strs = append(strs,processId)
    }
    
    sort.Strings(strs)
    
    fmt.Println("")
    
    for _,processId := range strs {
        fmt.Println(utils.CyanDk("   > Process Id :  "+processId))
        
        process := generalConfiguration.ProcessRegularsTimetable[processId]
        
        var regulars []string
        var reg = make(map[string]bool)
        
        ////////////////////
        // Sort regulars by need
        
        for len(regulars) != len(process) {
            for _,regular := range process {
                
                if len(regular.RegularNeed) == 0 && reg[regular.Id]==false {
                    regulars = append(regulars,regular.Id)
                    reg[regular.Id]=true
                    continue
                }
                
                if reg[regular.Id]==false{
                    _add:=true
                    for _,needId:=range regular.RegularNeed{
                        
                        if _, ok := process[needId]; !ok {
                            regulars = append(regulars,regular.Id)
                            reg[regular.Id]=true
                            continue
                        }
                        
                        if _, ok := reg[needId]; !ok {
                            _add=false
                        }
                    }
                    
                    if _add {
                        regulars = append(regulars,regular.Id)
                        reg[regular.Id]=true
                        continue
                    }
                }
            }
        }
        
        ////////////////
        
        for _,regularId := range regulars {
            regular := process[regularId]
            PrintOneRegular(generalConfiguration, regular, process)
        }
    }
    fmt.Println("")
}


func PrintOneRegular(
    generalConfiguration *utils.GeneralConfiguration, 
    regular utils.RegularTimetableStruct, 
    process map[string]utils.RegularTimetableStruct) {
    
    fmt.Println(utils.WhitLt("      > Service Reference : "+regular.ServiceReference+" (regId:"+regular.Id+")"),utils.CyanLt("-> "+regular.ConfigReference))
            
    var yes int
    
    fmt.Println(utils.PinkLt("            > Regular days :"))
    for _,day := range days {
        if d, ok := regular.RegularDays[day]; ok {
            fmt.Println(utils.WhitLt("                > "+day+" : ")+utils.GreeLt("yes"))
            if d.MinimalStartTimeSet() {
                fmt.Println(utils.WhitLt("\t\t\t\t -> minimal start time set : ")+utils.CyanLt(d.GetMinimalStartTime()))
            } else {
                fmt.Println(utils.WhitLt("\t\t\t\t -> minimal start time set : ")+utils.CyanLt("no"))
            }
            if d.MaximalStartTimeSet() {
                fmt.Println(utils.WhitLt("\t\t\t\t -> maximal start time set : ")+utils.CyanLt(d.GetMaximalStartTime()))
            } else {
                fmt.Println(utils.WhitLt("\t\t\t\t -> maximal start time set : ")+utils.CyanLt("no"))
            }
            
            fmt.Println(utils.WhitLt("\t\t\t\t -> attemps max : ")+utils.CyanLt(d.GetAttempsString()))
            
            yes++
        } else {
            fmt.Println(utils.WhitLt("                > "+day+" : ")+utils.RedLt_("no"))
        }
    }
    fmt.Println(utils.CyanLt("                -> "+strconv.Itoa(yes)+"j/7"))
    
    fmt.Println(utils.PinkLt("            > Regular needed :  "+strconv.Itoa(len(regular.RegularNeed))+""))
    for _,need := range regular.RegularNeed {
        
        fmt.Println(utils.WhitLt("                -> Service Reference : "+generalConfiguration.RegularsTimetable[need].ServiceReference+" ("+generalConfiguration.RegularsTimetable[need].Id+")"),utils.CyanLt("-> "+generalConfiguration.RegularsTimetable[need].ConfigReference))
        
        if _, ok := process[need]; !ok {
            fmt.Println(utils.YellLt("                    -> Be careful, this regular is from an another process !!! (cf -> "+generalConfiguration.RegularsTimetable[need].ProcessId+")"))
        }
    }
    
    fmt.Println("\n")
}