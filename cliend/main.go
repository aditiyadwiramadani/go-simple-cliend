package cliend

import (
	"github.com/ddliu/go-httpclient"
	"bytes"
	"io/ioutil"
	"log"
)


func GetData(url string) string { 
    resp, err :=httpclient.Get(url)
    if err != nil { 
        log.Fatal(err)
    }
    stringfy, _ := ioutil.ReadAll(resp.Body)
    return string(stringfy)

}

func PostData(url string, data []byte) string { 
    resp, err := httpclient.Post(url,  bytes.NewBuffer(data))
    if err != nil { 
        log.Fatal(err)
    }
    stringnify, _ := ioutil.ReadAll(resp.Body) 
    return string(stringnify)
}
func PutData(url string, data []byte) string { 

    resp, err :=httpclient.Put(url, bytes.NewBuffer(data))
    if err != nil { 
        log.Fatal(err)
    }
    stringnify, _ := ioutil.ReadAll(resp.Body) 
    return string(stringnify)
}




func DeleteData(url string) string { 
    resp, err := httpclient.Delete(url) 
    if err != nil { 
        log.Fatal(err)
    }  
    stringnify,err := ioutil.ReadAll(resp.Body)
    return string(stringnify)

}