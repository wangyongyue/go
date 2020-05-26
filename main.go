package main

    import (
        "fmt"
        "net/http"
        "strings"
        "log"
        _"github.com/go-sql-driver/mysql"
        "database/sql"
        "reflect"
        "encoding/json"    

    )
    //数据库配置
    const (
        userName = "root"
        password = "123456"
        ip = "39.106.180.44"
        port = "3306"
        dbName = "vsk"
    )
    

    var DB *sql.DB
    func main() {
   
        connectMySql()
        registerUrls()
        err := http.ListenAndServe(":9090", nil) //设置监听的端口
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }
    }
    func connectMySql(){

        path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
        DB, _ = sql.Open("mysql", path)
        DB.SetConnMaxLifetime(1000)
        DB.SetMaxIdleConns(10)
        if err := DB.Ping(); err != nil{
            fmt.Println("opon database fail")
            return
        }
        fmt.Println("connnect success")

    }
    func registerUrls(){

        http.Handle("/b",SelectController{b:Book{}})
        http.Handle("/p",SelectController{b:Pen{}})
       

    }
    type DataInterface interface{
        selectSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string)
        insertSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string)
        updateSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string)
        deleteSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string)
        selectRowSql(param map[string]string) (sqlStr string,sqlParam []interface{},err string)

        selectRows(rows *sql.Rows) interface{}
        selectRow(row *sql.Row) interface{}

    }

    
    type SelectController struct{
        b DataInterface

    }
    func (c SelectController) ServeHTTP(w http.ResponseWriter, r *http.Request){

        w.Header().Set("content-type","text/json")

        r.ParseForm()  
        form := make(map[string]string)
        for k, v := range r.Form {
            form[k] = strings.Join(v,"")
        }
        fmt.Println(form)
        h := selectRequest(c.b,form)
        ret_json, _ := json.Marshal(h)
        w.Write(ret_json)

    }


    type HttpResult struct{

        Code int32 `json:"code"`
        Data interface{} `json:"data"`
        Message string `json:"message"`

    }
   
    func (h HttpResult) success() HttpResult{

       h.Code = 1
       h.Message = "success"
       return h
    }
    func (h HttpResult) failure() HttpResult{
       h.Code = 100
       h.Data = nil
       return h
    }

    func selectRequest(m DataInterface,param map[string]string) HttpResult{

    	h :=  HttpResult{}
    	sqlStr,sqlParam,err:= m.selectSql(param)
        if len(err) > 0{
            h.Message = err
            return h.failure()
        }
        rows, _:= DB.Query(sqlStr,sqlParam ...)
        datas := make([]interface{}, 0)
        for rows.Next(){
            fmt.Println(reflect.TypeOf(rows)) 
            data := m.selectRows(rows)
            datas = append(datas,data)
        }
        fmt.Println(datas)
        h.Data = datas   
        return h.success()       

    }
    func selectRowRequest(m DataInterface,param map[string]string) HttpResult{

        h :=  HttpResult{}
        sqlStr,sqlParam,err:= m.selectRowSql(param)
        if len(err) > 0{
            h.Message = err
            return h.failure()
        }
        row := DB.QueryRow(sqlStr,sqlParam ...)
        data := m.selectRow(row)
        fmt.Println(data)
        h.Data = data
        return h.success()       

    }
    func insertRequest(m DataInterface,param map[string]string) HttpResult{

        h :=  HttpResult{}
        sqlStr,sqlParam,err:= m.insertSql(param)
        if len(err) > 0{
            h.Message = err
            return h.failure()
        }
        result, _:= DB.Exec(sqlStr,sqlParam ...)
        fmt.Println(result)
        h.Data = result
        return h.success()       

    }

    func updateRequest(m DataInterface,param map[string]string) HttpResult{

        h :=  HttpResult{}
        sqlStr,sqlParam,err:= m.updateSql(param)
        if len(err) > 0{
            h.Message = err
            return h.failure()
        }
        result, _:= DB.Exec(sqlStr,sqlParam ...)
        fmt.Println(result)
        h.Data = result
        return h.success()       

    }

    func deleteRequest(m DataInterface,param map[string]string) HttpResult{

        h :=  HttpResult{}
        sqlStr,sqlParam,err:= m.deleteSql(param)
        if len(err) > 0{
            h.Message = err
            return h.failure()
        }
        result, _:= DB.Exec(sqlStr,sqlParam ...)
        fmt.Println(result)
        h.Data = result
        return h.success()       

    }
