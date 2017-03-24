package main

import (
    "fmt"
    "./load"
    "errors"
    "./utils"
    "./display"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Clients (azzinothConfig utils.AzzinothConfiguration) error {
    
    var err error
    var db *sql.DB
    var generalConfiguration utils.GeneralConfiguration
    
    if db, err = sql.Open("mysql", azzinothConfig.AzzinothDatabase.DSN()); err != nil {
        return errors.New(utils.RedDk_("ERROR: open sql connection -> "+err.Error()))
    }
    defer db.Close()
    
    //////////////////////////////////////////
    
    fmt.Println(utils.GreeLt("Clients..."))
    if err = load.Clients(db, &generalConfiguration); err != nil {
        return errors.New(utils.RedDk_("ERROR: load Clients -> "+err.Error()))
    }
    fmt.Println(utils.GreeLt("loaded !"))
    
    display.Clients(&generalConfiguration)
    
    return nil
}