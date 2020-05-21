package main

    import (
        "fmt"
        "database/sql"


    )
    
    type WPen struct{
        Id int       `json:"id"`
        Data string  `json:"data"`
    }    
    type Pen struct{
        title string
    }
    func (b Pen) selectRow(row *sql.Row) interface{}{
        var w WPen
        row.Scan(&w.Id,&w.Data)
        return w
    }
    func (b Pen) selectRows(rows *sql.Rows) interface{}{
        var w WPen
        rows.Scan(&w.Id,&w.Data)
        return w
    }
    func (b Pen) selectRowSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        sqlParam = append(sqlParam,5)
        sqlStr ,err = "SELECT * FROM word_w  WHERE id = ?" ,""
        return
    }

    
    func (b Pen) selectSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        sqlParam = append(sqlParam,6)
        sqlStr ,err = "SELECT * FROM word_w  WHERE id = ?" ,""
        fmt.Println(sqlStr)
        return

    }
    func (b Pen) insertSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
    func (b Pen) updateSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
    func (b Pen) deleteSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string){

        return

    }
   