package main

import(
    "os"
    "strconv"
    "time"
    )

func main(){
    sec := 10//默认等待10秒
    if len(os.Args)>1{
        sec,_ = strconv.Atoi(os.Args[1])    
    }
	time.Sleep(time.Duration(sec)*time.Second)
}