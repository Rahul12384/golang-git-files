package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()

	type par struct {
		ID     int   `json:"id"`
		Verifier  string  `json:"ve"`
		Challenger  string  `json:"ch"`
	}

	type authcode struct{
		Code string `json:"code"`
	}

	ma:=make(map[string]string)

	// fmt.Println(par)
	var p1 par
	var s string
	s=p1.Challenger
	
	if val,ok :=ma["challenger"]; ok{
		fmt.Println(val)
	}
	router.POST("/rahu",func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		var p par
        // ctx.BindJSON(&p)
		// fmt.Printf(p.Title[0])
		ctx.JSON(200, gin.H {
			"message":"OK!!",
		})
		err := json.NewDecoder(ctx.Request.Body).Decode(&p)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}
    // Do something with the Person struct...
    	log.Printf("Per: %+v", p)

		if ctx.Request.Method == "OPTIONS" {
            ctx.AbortWithStatus(204)
            return
        }
		ctx.String(http.StatusOK,"rahu challeger"+p.Challenger)
		ctx.String(http.StatusOK,"rahu challeger"+p.Verifier)
		ma["verifier"]=p.Verifier
		ma["challenger"]=p.Challenger

		ctx.Next()
	})
	router.GET("/rah",func(ctx *gin.Context) {
		var p par
		ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.JSON(200, gin.H {
			"message":"OK!!",
		})
		ctx.String(http.StatusOK,p.Challenger)
		fmt.Printf("s"+s)
		ctx.String(http.StatusOK,ma["challenger"])
		ctx.String(http.StatusOK, "Hello1"+s+"s")

		// ctx.String(http.StatusOK,"get")
		// res, _ := json.Marshal(p)
		// ctx.String(http.StatusOK,string(res))
	})
	router.GET("/rahul",func(ctx *gin.Context) {
		var p par
		ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.JSON(200, gin.H {
			"message":"OK!!",
		})
		ctx.String(http.StatusOK,p.Challenger)
		fmt.Printf("s"+s)
		ctx.String(http.StatusOK,ma["verifier"])
		ctx.String(http.StatusOK, "Hello1"+s+"s")
		// base , err:=url.Parse("https://dev-12u2388w.us.auth0.com/authorize")
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }	
		// params := url.Values{}
		// params.Add("response_type", "code")
		// params.Add("code_challenge", ma["challenger"])
		// params.Add("code_challenge_method", "S256")
		// params.Add("client_id", "UeOlj5xtXsESZ2PGctxyOPoz22ybghi3")
		// params.Add("redirect_uri", "http://localhost:3000/welcome")
		// params.Add("scope", "openid%20profile")
		// params.Add("state", "xyzABC123")
		// // "https://dev-12u2388w.us.auth0.com/authorize?response_type=code&code_challenge={ma["challenger"]}&code_challenge_method=S256&client_id=UeOlj5xtXsESZ2PGctxyOPoz22ybghi3&redirect_uri=http://localhost:3000/welcome&scope=openid%20profile&state=xyzABC123"
		// // var ur="https://dev-12u2388w.us.auth0.com/authorize?response_type=code&code_challenge={ma["challenger"]}&code_challenge_method=S256&client_id=UeOlj5xtXsESZ2PGctxyOPoz22ybghi3&redirect_uri=http://localhost:3000/welcome&scope=openid%20profile&state=xyzABC123"
		// base.RawQuery = params.Encode()
		// log.Println(base.String())
		// resp, err := http.Get(base.String())
		// if err != nil {
		//    log.Fatalln(err)
		// }
	
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		//    log.Fatalln(err)
		// }

		// sb := string(body)
		// log.Printf(sb)
		// log.Println("hello in rahul")
		// ctx.String(http.StatusOK,"get")
		// res, _ := json.Marshal(p)
		// ctx.String(http.StatusOK,string(res))
	})


	router.POST("/rahultok",func(ctx *gin.Context) {
		var c authcode

		ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		err := json.NewDecoder(ctx.Request.Body).Decode(&c)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(ma)
		log.Println("code_verifier"+ma["verifier"])
		log.Println("code "+c.Code)
		log.Println(len(c.Code))
		// var v=ma["verifier"]
		if len(c.Code)>0{
			log.Println("Code in len if"+c.Code)
			ur := "https://dev-12u2388w.us.auth0.com/oauth/token"

			// payload := strings.NewReader("grant_type=authorization_code&client_id=%24%7Baccount.clientId%7D")
			payload := url.Values{}
			payload.Add("grant_type","authorization_code")
			// payload.Add("client_id","UeOlj5xtXsESZ2PGctxyOPoz22ybghi3")
			payload.Add("client_id","wplaaHT4M6bWFG1wmKdvKU6hPbFfWlOy")
			payload.Add("code_verifier",ma["verifier"])
			payload.Add("code",c.Code)
			payload.Add("redirect_uri", "http://localhost:3000/welcome")
			
			client:=&http.Client{}

			log.Println(ur)
			log.Println(payload)

			req, _ := http.NewRequest("POST", ur, strings.NewReader(payload.Encode()))

			req.Header.Add("content-type", "application/x-www-form-urlencoded")

			// res, err := http.DefaultClient.Do(req)

			res, err := client.Do(req)

			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)

			if err != nil {
				log.Fatal(err)
			}
			// log.Println(string(body))

			// fmt.Println(res)
			fmt.Println(string(body))
		}

		// Machine to Machine token using client credentials 
		//checking if we are able to get refresh token 

			urlclient := "https://dev-12u2388w.us.auth0.com/oauth/token"

			payloadclient := strings.NewReader("grant_type=client_credentials&client_id=wplaaHT4M6bWFG1wmKdvKU6hPbFfWlOy&client_secret=MjQDjUKXnqPJx0WNV12mDnOBm8wf8ca-b4mvnel4_lz8FHeHRc3BAi9DC3D3_pp0&audience=htpps://apiCloneHotstar")

			reqclient, _ := http.NewRequest("POST", urlclient, payloadclient)

			reqclient.Header.Add("content-type", "application/x-www-form-urlencoded")

			resclient, err := http.DefaultClient.Do(reqclient)

			if err != nil {
				log.Println("error in client")
				log.Fatal(err)
			}

			defer resclient.Body.Close()
			bodyclient, _ := ioutil.ReadAll(resclient.Body)

			log.Println("client credentials access token")

			fmt.Println(resclient)
			fmt.Println(string(bodyclient))


		// getting refresh token 

		urlrefresh := "https://dev-12u2388w.us.auth0.com/oauth/token"

		//payloadrefresh := strings.NewReader("grant_type=authorization_code&client_id=%24%7Baccount.clientId%7D&client_secret=YOUR_CLIENT_SECRET&code=YOUR_AUTHORIZATION_CODE&redirect_uri=%24%7Baccount.callback%7D")

		payloadrefresh := url.Values{}
			payloadrefresh.Add("grant_type","authorization_code")
			payloadrefresh.Add("client_id","wplaaHT4M6bWFG1wmKdvKU6hPbFfWlOy")
			payloadrefresh.Add("client_secret","MjQDjUKXnqPJx0WNV12mDnOBm8wf8ca-b4mvnel4_lz8FHeHRc3BAi9DC3D3_pp0")
			payloadrefresh.Add("code_verifier",ma["verifier"])
			payloadrefresh.Add("code",c.Code)
			payloadrefresh.Add("redirect_uri", "http://localhost:3000/welcome")
		
		reqrefresh, _ := http.NewRequest("POST", urlrefresh, strings.NewReader(payloadrefresh.Encode()))

		reqrefresh.Header.Add("content-type", "application/x-www-form-urlencoded")

		resrefresh, _ := http.DefaultClient.Do(reqrefresh)

		if err != nil {
				log.Println("error in client")
				log.Fatal(err)
			}

		defer resrefresh.Body.Close()
		bodyrefresh, _ := ioutil.ReadAll(resrefresh.Body)

		log.Println("access + refresh token")

		fmt.Println(resrefresh)
		fmt.Println(string(bodyrefresh))
		


		//  diff method

		// base , err:=url.Parse("https://dev-12u2388w.us.auth0.com/oauth/token")
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }	
		// params := url.Values{}
		// params.Add("grant_type","authorization_code")
		// params.Add("client_id","%24%7Baccount.clientId%7D")
		// params.Add("code_verifier",ma["challenger"])
		// params.Add("code",c.Code)
		// params.Add("redirect_uri", "%24%7Baccount.callback%7D")


		// base.RawQuery = params.Encode()
		// log.Println(base.String())
		// resp, err := http.Post(base.String(),"application/x-www-form-urlencoded",)
		// if err != nil {
		//    log.Fatalln(err)
		// }
	
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		//    log.Fatalln(err)
		// }

		// sb := string(body)
		// log.Printf(sb)
		// log.Println("rahultok hello")

	})
	router.Run(":8080")

}