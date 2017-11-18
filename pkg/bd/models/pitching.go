package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Pitching is a model that maps the CSV to a DB Table
type Pitching struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   int `json:"gs"  csv:"GS"  db:"GS"`
   Cg   int `json:"cg"  csv:"CG"  db:"CG"`
   Sho   int `json:"sho"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sv"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"ipouts"  csv:"IPouts"  db:"IPouts"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Er   int `json:"er"  csv:"ER"  db:"ER"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   int `json:"so"  csv:"SO"  db:"SO"`
   Baopp   string `json:"baopp"  csv:"BAOpp"  db:"BAOpp"`
   Era   string `json:"era"  csv:"ERA"  db:"ERA"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Wp   int `json:"wp"  csv:"WP"  db:"WP"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Bk   int `json:"bk"  csv:"BK"  db:"BK"`
   Bfp   string `json:"bfp"  csv:"BFP"  db:"BFP"`
   Gf   int `json:"gf"  csv:"GF"  db:"GF"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Pitching) GetTableName() string {
  return "pitching"
}

// GetFileName returns the name of the source file the model was created from
func (m *Pitching) GetFileName() string {
  return "Pitching.csv"
}

// GetFilePath returns the path of the source file
func (m *Pitching) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Pitching.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Pitching) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Pitching, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Pitching ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Pitching ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Pitching ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}