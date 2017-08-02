package gymtool

import (
	"strconv"
    "strings"
    "math"
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func Gt(message string) string{
    toolCode := strings.ToLower(string(message[0]))
    response :="看不懂你輸入的東西QQ,輸入h有好康的"

    switch toolCode {
    case string('k'),string('p'),string('m'),string('i') :
    response=weight(toolCode, message)
    case string('f') :
    response=food(message)
    case string('r') :
    response=rm(message)
    case string('b') :
    response=bodyinfo(message)
    case string('h') :
    response="重量轉換: p數字(轉公斤),k數字(轉磅)\n長度轉換: m數字(轉吋與英尺),i數字(轉公尺)\n食物成份: f食物名稱\n1rm計算: r重量-最多可做組數(r100-8)\n基礎代謝率(BMR): b身高(公分)-體重(公斤)-年齡(整數)-性別(1男0女)"
    }
	return response

}
    
func weight(code string, msg string) string{

    if _, err := strconv.ParseFloat(msg[1:len(msg)], 64); err == nil {
        value, err := strconv.ParseFloat(msg[1:len(msg)], 64)
        if err != nil {
        // do something sensible
        }
        amount := float64(value)
        if code=="k"{
        return strconv.FormatFloat(amount * 0.45, 'f', 2, 64) + "Kg" 
        }else if code=="p"{
        return strconv.FormatFloat(amount * 2.2, 'f', 2, 64) + "Lbs"
        }else if code=="m"{
        return strconv.FormatFloat(amount * 39.37, 'f', 2, 64) + "inches " + "or " + strconv.FormatFloat(amount * 3.28, 'f', 2, 64) +"feet"
        }else if code=="i"{
        return strconv.FormatFloat(amount * 0.0254, 'f', 2, 64) + "m"
        }
    }
    return "Oops, 你輸入的格式可能有誤,請參考h內的說明"
}

func food(msg string) string{
    ///response :="No this food." 
     //var response string
    // other:=""
     response:="找不到您輸入的食物,"
    // foods := map[string]string{
    // "雞蛋":"碳水化合物(g): 0 ,蛋白質(g): 6",
    // "鮭魚":"碳水化合物(g): 0 ,蛋白質(g): 20",
    // "雞胸":"碳水化合物(g): 0 ,蛋白質(g): 22",
    // "牛奶":"碳水化合物(g): 4 ,蛋白質(g): 3",
    // "鯛魚":"碳水化合物(g): 2.5 ,蛋白質(g): 18",
    // "地瓜":"碳水化合物(g): 25 ,蛋白質(g): 1.3",
    // "藜麥":"碳水化合物(g): 68 ,蛋白質(g): 12",
    // "奇亞籽":"碳水化合物(g): 44 ,蛋白質(g): 16",
    // "糙米":"碳水化合物(g): 73 ,蛋白質(g): 9",
    // "鮪魚罐頭":"碳水化合物(g):  ,蛋白質(g): 30",
    // "藍莓":"碳水化合物(g): 13 ,蛋白質(g): 0",
    // "燕麥":"碳水化合物(g): 76 ,蛋白質(g): 0",
    // }
     var input = msg[1:len(msg)]
    // if f, ok := foods[input]; ok {
    // response= "每100克 "+input+" 含: "+f
    // }else{
    // response="找不到: "+input
    // }

    // for k := range foods {
    //     if strings.ContainsAny(k,input){
    //     other=other+k+","
    //     }

    // }

    // return response+"\n您可搜尋類似食物:"+other

    file, err := os.Open("foods.csv")
    if err != nil {
        fmt.Println("Error:", err)
        //return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    reader.Comma = ';'
    lineCount := 0
    search:=""
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            //return
        }
        for i := 0; i < len(record); i++ {
            if record[0][0:strings.Index(record[0], ",")]==input{
                food:=strings.Split(record[0],",")
                response=fmt.Sprintf("%s 含有"+" 熱量(Kcal)%s"+" 蛋白質(g)%s"+" 脂肪(g):%s"+" 碳水化合物(g):%s",food[0],food[1],food[2],food[3],food[4])
            }
            if strings.ContainsAny(record[0],input){
                search=search+record[0][0:strings.Index(record[0], ",")]+","
            }
        }
        lineCount += 1
    }
    return response+"\n\n您可搜尋類似的食物:\n"+ search
}

func rm(msg string) string{
    //var weight, reps
    response:="Oops, 你輸入的格式可能有誤,請參考h內的說明"
    if strings.ContainsAny(msg,"-"){
        data:=strings.Split(msg[1:len(msg)],"-")

        if weight, err:=strconv.ParseFloat(data[0], 64);err==nil{
           if reps, err:=strconv.ParseFloat(data[1], 64);err==nil{
             response = strconv.FormatFloat( (100 * weight) / (52.2 + (41.9 * math.Exp(-0.055 * reps))) , 'f', 2, 64)
           }
        } 
    }
    return response
}


func bodyinfo(msg string) string{
    response:="輸入資料有誤,格式為: b身高(公分)-體重(公斤)-年齡(整數)-性別(1男0女),ex: b180-70-25-1"
    if strings.ContainsAny(msg,"-"){
        data:=strings.Split(msg[1:len(msg)],"-")

        if len(data)==4{
                    if sex, err:=strconv.ParseFloat(data[3],64);err==nil{
            response=strconv.FormatFloat(sex, 'f', 2, 64)
        }else{
            response="error"
        }
        if height, err:=strconv.ParseFloat(data[0], 64);err==nil{
           if weight, err:=strconv.ParseFloat(data[1], 64);err==nil{
              if age, err:=strconv.ParseFloat(data[2],64);err==nil{
                 if sex, err:=strconv.ParseFloat(data[3],64);err==nil{
                    if sex==1{
                         response = "您的基礎代謝率(BMR): "+strconv.FormatFloat( 13.7*weight+5*height-6.8*age+66 , 'f', 2, 64)+" 大卡(Kcal)"
                    }else if sex==0{
                         response = "您的基礎代謝量(BMR): "+strconv.FormatFloat( 9.6*weight+1.8*height-4.7*age+655 , 'f', 2, 64)+" 大卡(Kcal)"
                    }
                 }
               }
            }
        } 
        }

    }
    return response
}

