package gymtool

import (
	"strconv"
)

func Gt(message string) string{
    toolCode := string(message[0])
    var response string

    switch toolCode {
    case string('k'),string('p') :
    response=weight(toolCode, message)
    case string('f') :
    response=food(message)

    }



	return response

}
    
func weight(code string, msg string) string{

    if _, err := strconv.Atoi(msg[1:len(msg)]); err == nil {
        value, err := strconv.ParseFloat(msg[1:len(msg)], 64)
        if err != nil {
        // do something sensible
        }
        amount := float64(value)
        if code=="k"{
        return strconv.FormatFloat(amount * 0.45, 'f', 2, 64) + "Kg"
        }else{
        return strconv.FormatFloat(amount * 2.2, 'f', 2, 64) + "Lbs"
        }
    }
    return "invalid input"
}

func food(msg string) string{
    foods := map[string]string{
    "雞蛋":"碳水(g): 0 蛋白質(g): 6",
    "鮭魚":"碳水(g): 0 蛋白質(g): 20",
    "雞胸":"碳水(g): 0 蛋白質(g): 22",
    "牛奶":"碳水(g): 4 蛋白質(g): 3",
    "鯛魚":"碳水(g): 2.5 蛋白質(g): 18",
    "地瓜":"碳水(g): 25 蛋白質(g): 1.3",
    "藜麥":"碳水(g): 68 蛋白質(g): 12",
    "奇亞籽":"碳水(g): 44 蛋白質(g): 16",
    "糙米":"碳水(g): 73 蛋白質(g): 9",
    "鮪魚罐頭":"碳水(g):  蛋白質(g): 30",
    "藍莓":"碳水(g): 13 蛋白質(g): 0",
    "燕麥":"碳水(g): 76 蛋白質(g): 0",
    }
    var input = msg[1:len(msg)]
    if f, ok := foods[input]; ok {
    return f
    }
    return "No this food."
}
