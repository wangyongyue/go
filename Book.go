package main

    import (
        "fmt"
        "database/sql"

    )

    type Word struct{
        Id int       `json:"id"`
        Data string  `json:"data"`
    }    
    type Book struct{
        title string
    }
    
    func (b Book) selectRow(row *sql.Row) interface{}{
        var w Word
        row.Scan(&w.Id,&w.Data)
        return w
    }
    func (b Book) selectRows(rows *sql.Rows) interface{}{
        var w Word
        rows.Scan(&w.Id,&w.Data)
        return w
    }
    func (b Book) selectRowSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        sqlParam = append(sqlParam,5)
        sqlStr = "SELECT * FROM word_w  WHERE id = ?"
        return
    }

    
    func (b Book) selectSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        sqlStr = "SELECT * FROM word_w"
        fmt.Println(sqlStr)
        return

    }
    func (b Book) insertSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
    func (b Book) updateSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
    func (b Book) deleteSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
   