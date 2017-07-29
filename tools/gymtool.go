package gymtool

import (
	"strconv"
)

func Gt(message string) string{

	return weight(message)

}

func weight(x string) string{
    st :=x
    s := string(st[0])
    var aaa string
    if _, err := strconv.Atoi(st[1:len(st)]); err == nil {
        value, err := strconv.ParseFloat(st[1:len(st)], 64)
        if err != nil {
        // do something sensible
        }
        amount := float64(value)
        switch s {
        case string('k'):
        aaa=weightLbs(amount)+" Lbs"
        //return weightLbs(amount)+" Lbs"
        //fmt.Println(weightLbs(amount)+" Lbs")
        case string('p'):
        aaa=weightKg(amount)+" Kg"
        //return weightKg(amount)+" Kg"
        //fmt.Println(weightKg(amount)+" Kg")
        }
    }else{
        aaa="invalid input"
    }
    return aaa
}


func weightLbs(x float64) string{
    return strconv.FormatFloat(x * 2.2, 'f', 2, 64)
}
func weightKg(x float64) string{
    return strconv.FormatFloat(x * 0.45, 'f', 2, 64)
}
