func intToRoman(num int) string {
    m:=map[string]string{
        "1":"I",
        "4":"IV",
        "5":"V",
        "9":"IX",
        "10":"X",
        "40":"XL",
        "50":"L",
        "90":"XC",
        "100":"C",
        "400":"CD",
        "500":"D",
        "900":"CM",
        "1000":"M",
    }

    res:=""
    for num>0 {
        if num>=1000 {
            res=res+m["1000"]
            num-=1000
        } else if num>=900 {
            res=res+m["900"]
            num-=900
        } else if num>=500 {
            res=res+m["500"]
            num-=500
        } else if num>=400 {
            res=res+m["400"]
            num-=400
        } else if num>=100 {
            res=res+m["100"]
            num-=100
        } else if num>=90 {
            res=res+m["90"]
            num-=90
        } else if num>=50 {
            res=res+m["50"]
            num-=50
        } else if num>=40 {
            res=res+m["40"]
            num-=40
        } else if num>=10 {
            res=res+m["10"]
            num-=10
        } else if num>=9 {
            res=res+m["9"]
            num-=9
        } else if num>=5 {
            res=res+m["5"]
            num-=5
        } else if num>=4 {
            res=res+m["4"]
            num-=4
        } else if num>=1 {
            res=res+m["1"]
            num-=1
        }
        // fmt.Println(num)
    }

    return res
}