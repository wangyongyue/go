package main

    import (
        "fmt"
        "net/http"
        "strings"
        "log"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "reflect"
        "encoding/json"
        "io"
    

    )
    //数据库配置
    const (
        userName = "root"
        password = "123456"
        ip = "39.106.180.44"
        port = "3306"
        dbName = "vsk"
    )
    type RequestType int32

    const (
        RSelect     RequestType = 0
        RSelectRow  RequestType = 1
        RInsert     RequestType = 2
        RUpdate     RequestType = 3
        RDelete     RequestType = 4
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

        http.Handle("/b",Controller{b:Book{},r:RSelect})
        http.Handle("/p",Controller{b:Pen{},r:RSelect})

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

    
    type Controller struct{
        b DataInterface
        r RequestType

    }
    func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request){

        r.ParseForm()  
        form := make(map[string]string)
        for k, v := range r.Form {
            form[k] = strings.Join(v,"")
        }
        fmt.Println(form)


        serveHttps(c.b,c.r ,form,w)

        

    }

    type HttpResult struct{

        Code int32 `json:"code"`
        Data interface{} `json:"data"`
        Message string `json:"message"`

    }
   
    func (h HttpResult) success() string{

       h.Code = 1
       h.Message = "success"
       ret_json, _ := json.Marshal(h)
       fmt.Println(string(ret_json))
       return string(ret_json)
    }
    func (h HttpResult) failure() string{
       h.Code = 100
       h.Data = nil
       ret_json, _ := json.Marshal(h)
       fmt.Println(string(ret_json))
       return string(ret_json)
    }

    func serveHttps(m DataInterface,r RequestType,param map[string]string,w http.ResponseWriter) {

        h :=  HttpResult{}
        if r  == RSelect {

            sqlStr,sqlParam,err:= m.selectSql(param)
            if len(err) > 0{
                fmt.Fprint(w, err)
                return
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
            io.WriteString(w, h.success())

            
        }else if r == RSelectRow{

            sqlStr,sqlParam,err:= m.selectRowSql(param)
            if len(err) > 0{
                fmt.Fprint(w, err)
                return
            }
            row := DB.QueryRow(sqlStr,sqlParam ...)
            data := m.selectRow(row)
            fmt.Println(data)

            h.Data = data
            fmt.Fprint(w, h.success())


        }else if r == RInsert {
            
            sqlStr,sqlParam,err:= m.insertSql(param)
            if len(err) > 0{
                fmt.Fprint(w, err)
                return
            }
            result, _:= DB.Exec(sqlStr,sqlParam ...)
            fmt.Println(result)

            h.Data = result
            fmt.Fprint(w, h.success())


        }else if r == RUpdate {

            sqlStr,sqlParam,err:= m.updateSql(param)
            if len(err) > 0{
                fmt.Fprint(w, err)
                return
            }
            result, _:= DB.Exec(sqlStr,sqlParam ...)
            fmt.Println(result)
            h.Data = result
            fmt.Fprint(w, h.success())

            
        }else if r == RDelete {


            sqlStr,sqlParam,err:= m.deleteSql(param)
            if len(err) > 0{
                fmt.Fprint(w, err)
                return
            }
            result, _:= DB.Exec(sqlStr,sqlParam ...)
            fmt.Println(result)
            h.Data = result
            fmt.Fprint(w, h.success())

            
        }






        
    }























    