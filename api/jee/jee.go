package jee

import (
	types "api/jee/process_data"
	"api/kry"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRoutes(mainRouter *gin.Engine) {
	var url_prefix string = ""
	DETA_ENV := os.Getenv("DETA_SPACE_APP") == "true"
	if !DETA_ENV {
		url_prefix = "/api"
	}
	router := mainRouter.Group(url_prefix)

	router.POST("/page_view_count", route_page_view_count)
	router.POST("/get_sample_result", route_get_sample_result)
	router.POST("/get_result", get_result_route)
}

func update_result_view_count() {
	if kry.PROD_ENV {
		counts_base := kry.Deta.Base("counts")
		var result_view_count kry.KeyValuePairInt
		counts_base.Get("result_view_count", &result_view_count)
		result_view_count.Value++
		counts_base.Put(&result_view_count)
	}
}

type get_result_route_body_type struct {
	ApplicationNumber string `json:"ApplicationNumber"`
	DOB               string `json:"DateOfBirth"`
}
type result_response_type struct {
	Result types.ResultDataType `json:"result"`
	Data   types.MainDataType   `json:"data"`
}

func get_result_route(c *gin.Context) {
	var bdy get_result_route_body_type
	c.BindJSON(&bdy)

	var user_info types.UserInfoType
	user_info_fetched := kry.Deta.Base("info").Get(bdy.ApplicationNumber, &user_info)

	if user_info_fetched != nil {
		// Application Number not found in database
		c.JSON(403, gin.H{
			"detail": &kry.ErrorInfoType{
				Error: "appl_numb_not_found",
			},
		})
		return
	}
	fmt.Println(user_info.DOB, bdy.DOB)
	if user_info.DOB != bdy.DOB {
		// DOB did not match
		c.JSON(403, gin.H{
			"detail": &kry.ErrorInfoType{
				Error: "dob_did_not_match",
			},
		})
		return
	}

	var data result_response_type
	kry.Deta.Base("data").Get(bdy.ApplicationNumber, &data)

	update_result_view_count()

	c.JSON(200, &data)
}

func route_get_sample_result(c *gin.Context) {
	var data result_response_type
	kry.Deta.Base("data").Get("sample_result", &data)
	c.JSON(200, &data)
}

func route_page_view_count(c *gin.Context) {
	counts_base := kry.Deta.Base("counts")
	var update bool = c.DefaultQuery("update_count", "true") == "true"

	var page_view_count kry.KeyValuePairInt
	var result_view_count kry.KeyValuePairInt
	counts_base.Get("page_view_count", &page_view_count)
	counts_base.Get("result_view_count", &result_view_count)

	if update {
		page_view_count.Value++
		counts_base.Put(&page_view_count)
	}

	c.JSON(200, gin.H{
		"page_view_count":   page_view_count.Value,
		"result_view_count": result_view_count.Value,
	})
}
