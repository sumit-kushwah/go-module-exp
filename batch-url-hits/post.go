package main

import (
	"fmt"
	"os"
	"time"
)

var bliss_data = `[{
	"DISPLAY_ID": "8346234",
		"BL_APP_12_MONTHS": 2,
		"BL_APP_3_MONTHS": 1,
		"BL_APP_6_MONTHS": 2,
		"BL_GEN_12_MONTHS": 2,
		"BL_GEN_3_MONTHS": 1,
		"BL_GEN_6_MONTHS": 2,
		"BMCAT1": "43468",
		"BROWSE_COUNT_24_HOUR": 1,
		"BROWSE_COUNT_48_HOUR": 1,
		"BUYER_TYPE": "REST",
		"CALL_COUNT_24_HOUR": 0,
		"CALL_COUNT_48_HOUR": 0,
		"EMAIL_AVAILABLE": 1,
		"EMAIL_VERIFIED": 0,
		"INTENT_TYPE": "66",
		"MOBILE_VERIFIED": 0,
		"PLATFORM": "IMOB",
		"PRE_ISQ_FILLED": 43468,
		"PRE_MCAT_ID": "1",
		"PRE_MCITY_APP_CNT_12M": 1,
		"PRE_MCITY_APP_CNT_3M": 0,
		"PRE_MCITY_APP_CNT_6M": 1,
		"PRE_MCITY_PSELLER_CNT": 0,
		"PRE_MCITY_TRANS_CNT_12M": 3,
		"PRE_MCITY_TRANS_CNT_3M": 0,
		"PRE_MCITY_TRANS_CNT_6M": 3,
		"PRE_MCITY_UNQSLD_CNT_12M": 1,
		"PRE_MCITY_UNQSLD_CNT_3M": 0,
		"PRE_MCITY_UNQSLD_CNT_6M": 1,
		"SEARCH_COUNT_24_HOUR": 0,
		"SEARCH_COUNT_48_HOUR": 0,
		"PRE_MCAT_UNQSLD_CNT_3M": 34,
		"PRE_MCAT_TRANS_CNT_3M": 22,
		"PRE_MCAT_APP_CNT_3M": 40,
		"MCAT_SOLD_3M": 64.71,
		"MCAT_TRXN_3M": 117.65,
		"PRE_MDIST_TRANS_CNT_3M": 0,
		"PRE_MDIST_UNQSLD_CNT_3M": 0,
		"PRE_MDIST_APP_CNT_3M": 0,
		"MDIST_TRXN_3M": 0.0,
		"MDIST_SOLD_3M": 0.0
}]`

func main() {
	requests := 50
	start := time.Now()
	var urls []string
	var url string
	var json_data string

	property := "product-non-pharma"

	if len(os.Args) > 1 {
		property = os.Args[1]
	}

	ingress_ip := "34.36.64.248"

	if property == "product-non-pharma" {
		url = fmt.Sprintf("http://%s/product/fasttext/predict", ingress_ip)
		json_data = `{"msg":"Sr Agarbatti Sr Agarbatti","item_name":"Sr Agarbatti"}`
	} else if property == "buylead-non-pharma" {
		url = fmt.Sprintf("http://%s/buylead/fasttext/predict", ingress_ip)
		json_data = `{"msg":"Bidi Bomb Cracker Bidi Bomb Cracker 1 pack Noise Maker Khaki"}`
	} else if property == "product-pharma" {
		url = fmt.Sprintf("http://%s/product/cnn-transe/predict", ingress_ip)
		json_data = `{"msg": "Librium 10 Mg Tablets Librium 10 Mg Tablets  10 mg 10*10 Capsules Abbott"}`
	} else if property == "buylead-pharma" {
		url = fmt.Sprintf("http://%s/buylead/cnn-transe/predict", ingress_ip)
		json_data = `{"msg": "Cilacar T Tablet Cilacar T Tablet 10 stripe 1*10 per strip 1*10"}`
	} else {
		print("Please provide a valid property")
		os.Exit(1)
	}

	for i := 0; i < requests; i++ {
		urls = append(urls, url)
	}

	checkURLs("POST", urls, json_data)
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}
