package gymtool

import (
	"strconv"
    "strings"
    "math"
)

func Gt(message string) string{
    toolCode := strings.ToLower(string(message[0]))
    response :="輸入錯誤,請輸入h查看功能列表"

    switch toolCode {
    case string('k'),string('p') :
    response=weight(toolCode, message)
    case string('m'),string('i') :
    response=food(message)
    case string('r') :
    response=rm(message)

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
        return strconv.FormatFloat(amount * 39.37, 'f', 2, 64) + "inches" + "or " + strconv.FormatFloat(amount * 3.28, 'f', 2, 64) +"feet"
        }else if code=="i"{
        return strconv.FormatFloat(amount * 0.0254, 'f', 2, 64) + "m"
        }
    }
    return "invalid input"
}

func food(msg string) string{
    foods := map[string]string{
    "雞蛋":"碳水化合物(g): 0 ,蛋白質(g): 6",
    "鮭魚":"碳水化合物(g): 0 ,蛋白質(g): 20",
    "雞胸":"碳水化合物(g): 0 ,蛋白質(g): 22",
    "牛奶":"碳水化合物(g): 4 ,蛋白質(g): 3",
    "鯛魚":"碳水化合物(g): 2.5 ,蛋白質(g): 18",
    "地瓜":"碳水化合物(g): 25 ,蛋白質(g): 1.3",
    "藜麥":"碳水化合物(g): 68 ,蛋白質(g): 12",
    "奇亞籽":"碳水化合物(g): 44 ,蛋白質(g): 16",
    "糙米":"碳水化合物(g): 73 ,蛋白質(g): 9",
    "鮪魚罐頭":"碳水化合物(g):  ,蛋白質(g): 30",
    "藍莓":"碳水化合物(g): 13 ,蛋白質(g): 0",
    "燕麥":"碳水化合物(g): 76 ,蛋白質(g): 0",
    }
    var input = msg[1:len(msg)]
    if f, ok := foods[input]; ok {
    return "每100克"+input+"含"+f
    }
    return "No this food."
}

func rm(msg string) string{
    //var weight, reps
    var response string
    if strings.ContainsAny(msg,"-"){
        data:=strings.Split(msg[1:len(msg)],"-")

        if weight, err:=strconv.ParseFloat(data[0], 64);err==nil{
           if reps, err:=strconv.ParseFloat(data[1], 64);err==nil{
             response = strconv.FormatFloat(100 * weight / 52.2 + 41.9 * math.Exp(-0.055 * reps), 'f', 2, 64)
           }
        } 
    }
    return response
}
// func schdule(msg string) {
//     traningCode := msg(1:3)
//     switch traningCode {
//     case "531" :

//     }
// }

