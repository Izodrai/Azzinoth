package utils

import(
    "io/ioutil"
    "encoding/json"
)

const(
    Stop      = "\x1b[0m"
    C1RedDk_  = "\x1b[31m"
    C2GreeDk  = "\x1b[32m"
    C3MarrDk  = "\x1b[33m"
    C4BlueDk  = "\x1b[34m"
    C5ViolDk  = "\x1b[35m"
    C6CyanDk  = "\x1b[36m"
    C30GrayL  = "\x1b[30;1m"
    C31RedLt  = "\x1b[31;1m"
    C32GreeL  = "\x1b[32;1m"
    C33YellL  = "\x1b[33;1m"
    C34BlueL  = "\x1b[34;1m"
    C35PinkL  = "\x1b[35;1m"
    C36CyanL  = "\x1b[36;1m"
    C37WhitL  = "\x1b[37;1m"
)

func RedDk_(s string) string {return C1RedDk_+s+Stop}
func GreeDk(s string) string {return C2GreeDk+s+Stop}
func MarrDk(s string) string {return C3MarrDk+s+Stop}
func BlueDk(s string) string {return C4BlueDk+s+Stop}
func ViolDk(s string) string {return C5ViolDk+s+Stop}
func CyanDk(s string) string {return C6CyanDk+s+Stop}

func GrayLt(s string) string {return C30GrayL+s+Stop}
func RedLt_(s string) string {return C31RedLt+s+Stop}
func GreeLt(s string) string {return C32GreeL+s+Stop}
func YellLt(s string) string {return C33YellL+s+Stop}
func BlueLt(s string) string {return C34BlueL+s+Stop}
func PinkLt(s string) string {return C35PinkL+s+Stop}
func CyanLt(s string) string {return C36CyanL+s+Stop}
func WhitLt(s string) string {return C37WhitL+s+Stop}


type AzzinothConfiguration struct {
    AzzinothDatabase Database `json:"azzinothDatabase"`
}
    
func (c *AzzinothConfiguration) LoadConfig(filename string) error {
    
    raw, err := ioutil.ReadFile(filename)
    if err != nil {
            return err
    }

    err = json.Unmarshal(raw, c)
    if err != nil {
            return err
    }

    return nil
}

type Database struct {
    Protocol string `json:"protocol"`
    Host     string `json:"host"`
    Port     string `json:"port"`
    Login    string `json:"login"`
    Password string `json:"password"`
    Db       string `json:"db"`
}

func (d Database) DSN() string {
    return d.Login+":"+d.Password+"@"+d.Protocol+"("+d.Host+":"+d.Port+")/"+d.Db
}
