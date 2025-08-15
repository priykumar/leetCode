func intToRoman(num int) string {
    m:=map[string]string{
        "1":"I",
        "5":"V",
        "10":"X",
        "50":"L",
        "100":"C",
        "500":"D",
        "1000":"M",
    }

    res:=""
    for num>0 {
        if num>=1000 {
            res=res+m["1000"]
            num-=1000
        } else if num>=500 {
            if num>=900 && num <=999 {
                res=res+"CM"
                num-=900
            } else {
                res=res+m["500"]
                num-=500
            }
        } else if num>=100 {
            if num>=400 && num <=499 {
                res=res+"CD"
                num-=400
            } else {
                res=res+m["100"]
                num-=100
            }
        } else if num>=50 {
            if num>=90 && num <=99 {
                res=res+"XC"
                num-=90
            } else {
                res=res+m["50"]
                num-=50
            }
        } else if num>=10 {
            if num>=40 && num <=49 {
                res=res+"XL"
                num-=40
            } else{
                res=res+m["10"]
                num-=10
            }
        } else if num>=5 {
            if num==9 {
                res=res+"IX"
                num-=9
            } else{
                res=res+m["5"]
                num-=5
            }
        } else if num>=1 {
            if num == 4 {
                res=res+"IV"
                num-=4
            } else{
                res=res+m["1"]
                num-=1
            }
        }
        fmt.Println(num)
    }

    return res
}