package main

import (
    "fmt"
    "./utils"
)

func main() {
    
    var err error
    var configFile = "config.json"
    var azzinothConfig utils.AzzinothConfiguration
    
    if err = azzinothConfig.LoadConfig(configFile); err != nil {
        fmt.Println(utils.RedDk_("ERROR: load configuration file -> "+err.Error()))
        return
    }
    
//     if err = Clients(azzinothConfig); err != nil {
//         fmt.Println(err)
//     }
//     
//     if err = Servers(azzinothConfig); err != nil {
//         fmt.Println(err)
//     }
//     
//     if err = Services(azzinothConfig); err != nil {
//         fmt.Println(err)
//     }
//     
//     if err = ServicesHaveDataServers(azzinothConfig); err != nil {
//         fmt.Println(err)
//     }
    
//     if err = Configurations(azzinothConfig); err != nil {
//         fmt.Println(err)
//     }
//     
    if err = RegularsTimetable(azzinothConfig); err != nil {
        fmt.Println(err)
    }
}