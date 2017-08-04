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
    case string('s') :
    response=schedule(message)
    case string('h') :
    response="重量轉換: p數字(轉公斤),k數字(轉磅)\n長度轉換: m數字(轉吋與英尺),i數字(轉公尺)\n食物成份: f食物名稱\nf(k熱量,p蛋白質,f脂肪,c碳水)-數字(fk-100 = 熱量大於100的食物) \n1rm計算: r重量-最多可做組數(r100-8)\n基礎代謝率(BMR): b身高(公分)-體重(公斤)-年齡(整數)-性別(1男0女)\n訓練課表: 請按s-h查詢"
    case string('t') :
    response=training(message)
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
    foodIng := map[string]string{
        "k": "熱量",
        "p": "蛋白質",
        "f": "脂肪",
        "c": "碳水化合物",
    }

    file, err := os.Open("foods.csv")
    if err != nil {
        fmt.Println("Error:", err)
    }
    input := msg
    defer file.Close()
    reader := csv.NewReader(file)
    reader.Comma = ';'
    lineCount := 0
    search := ""
    response := ""
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
        }
        for i := 0; i < len(record); i++ {

            if strings.ContainsAny(input, "-") {
                data := strings.Split(input[1:len(input)], "-")
                food := strings.Split(record[0], ",")
                criteria, err := strconv.ParseFloat(strings.TrimSpace(data[1]), 64)
                foddKcal, err := strconv.ParseFloat(food[1], 64)
                foddProtein, err := strconv.ParseFloat(food[2], 64)
                foddFat, err := strconv.ParseFloat(food[3], 64)
                foddCarbon, err := strconv.ParseFloat(food[4], 64)

                if err == nil {
                    switch data[0] {
                    case string('k'):
                        if foddKcal >= criteria {
                            response = response + fmt.Sprintf("%s(%s)", food[0], food[1])
                        }

                    case string('p'):
                        if foddProtein >= criteria {
                            response = response + fmt.Sprintf("%s(%s) ", food[0], food[2])
                        }

                    case string('f'):
                        if foddFat >= criteria {
                            response = response + fmt.Sprintf("%s(%s) ", food[0], food[3])
                        }

                    case string('c'):
                        if foddCarbon >= criteria {
                            response = response + fmt.Sprintf("%s(%s) ", food[0], food[4])
                        }

                    }
                }else{
                    response="輸入錯誤或無法辨識"
                }
            } else {

                if record[0][0:strings.Index(record[0], ",")] == input {
                    food := strings.Split(record[0], ",")
                    response = fmt.Sprintf("%s 含有"+" 熱量(Kcal)%s"+" 蛋白質(g)%s"+" 脂肪(g):%s"+" 碳水化合物(g):%s", food[0], food[1], food[2], food[3], food[4])
                }
                if strings.ContainsAny(record[0], input) {
                    search = search + record[0][0:strings.Index(record[0], ",")] + ","
                }

            }
        }
        lineCount += 1
    }

    if strings.ContainsAny(input, "-") {
    response = fmt.Sprintf( foodIng[strings.ToLower(string(input[1]))]+"大於"+strings.TrimSpace(input[3:len(input)])+"的食物有:" +response)
    }else{
        if response=="" {
        response="找不到您輸入的食物"
        }
        response=fmt.Sprintf(response + "\n\n您可搜尋類似的食物:\n" + search[0:len(search)-1])
    }

    return response
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

func schedule(msg string) string{

response:="目前提供下列運動課表\n1. 5-3-1 訓練: s-fto-部位(s,b,d)-第幾週(1,2,3,4)-1rm重量, Ex: s-fto-s-1-120"
    ftoMap:=map[int]map[int]float64{
        1: map[int]float64{
            1:0.75,
            2:0.8,
            3:0.85,
            4:1.05,
            5:1.1,
            6:1.15,
        },
        2: map[int]float64{
            1:0.8,
            2:0.85,
            3:0.9,
            4:1.05,
            5:1.1,
            6:1.15,
        },
        3: map[int]float64{
            1:0.75,
            2:0.85,
            3:0.95,
            4:1.05,
            5:1.1,
            6:1.15,
        },
        4: map[int]float64{
            1:0.6,
            2:0.65,
            3:0.7,
        },
    }
    workoutMap:=map[string]string{
        "s":"Squat",
        "b":"Bench Press",
        "d":"Deadlift",
    }

data:=strings.Split(msg[1:len(msg)],"-")

if strings.ToLower(data[1])=="fto"{
    day ,err:=strconv.Atoi(data[3])
    weight ,err:=strconv.ParseFloat(data[4],64)
    if err==nil && day <5{
    switch day {
    case 1:
        response=fmt.Sprintf("%s-第%s週(每組5Reps)\n第一組%.2f\n第二組%.2f\n第三組%.2f\n第J+1組%.2f\n第J+2組%.2f\n第J+3組%.2f\n",workoutMap[data[2]],data[3],ftoMap[1][1]*weight,ftoMap[1][2]*weight,ftoMap[1][3]*weight,ftoMap[1][4]*weight*ftoMap[1][3],ftoMap[1][5]*weight*ftoMap[1][3],ftoMap[1][6]*weight*ftoMap[1][3])   
    case 2:
        response=fmt.Sprintf("%s-第%s週(每組3Reps)\n第一組%.2f\n第二組%.2f\n第三組%.2f\n第J+1組%.2f(破PR組)\n第J+2組%.2f(破PR組)\n第J+3組%.2f(破PR組)\n",workoutMap[data[2]],data[3],ftoMap[2][1]*weight,ftoMap[2][2]*weight,ftoMap[2][3]*weight,ftoMap[2][4]*weight*ftoMap[2][3],ftoMap[2][5]*weight*ftoMap[2][3],ftoMap[2][6]*weight*ftoMap[2][3])   
    case 3:
        response=fmt.Sprintf("%s-第%s週(5,3,1,1,1 Reps)\n第一組%.2f\n第二組%.2f\n第三組%.2f\n第J+1組%.2f\n第J+2組%.2f(破PR組)\n第J+3組%.2f(破PR組)\n",workoutMap[data[2]],data[3],ftoMap[3][1]*weight,ftoMap[3][2]*weight,ftoMap[3][3]*weight,ftoMap[3][4]*weight*ftoMap[3][3],ftoMap[3][5]*weight*ftoMap[3][3],ftoMap[3][6]*weight*ftoMap[3][3])    
    case 4:
        response=fmt.Sprintf("%s-第%s週(每組5Reps)\n第一組%.2f\n第二組%.2f\n第三組%.2f",workoutMap[data[2]],data[3],ftoMap[4][1]*weight,ftoMap[4][2]*weight,ftoMap[4][3]*weight) 
    }

    }

}
return response

}

func training(msg string) string{
    
    response:="https://s-media-cache-ak0.pinimg.com/736x/0c/0e/78/0c0e784806bc4f395e321ece1b720cc8--forearm-exercises-forearm-workout.jpg"+","+"https://s-media-cache-ak0.pinimg.com/736x/0c/0e/78/0c0e784806bc4f395e321ece1b720cc8--forearm-exercises-forearm-workout.jpg"
    
    return response
    
}