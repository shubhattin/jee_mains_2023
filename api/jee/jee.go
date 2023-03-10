package jee

import (
	types "api/jee/process_data"
	"api/kry"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(mainRouter *fiber.App) {
	var url_prefix string = ""
	DETA_ENV := os.Getenv("DETA_SPACE_APP") == "true"
	if !DETA_ENV {
		url_prefix = "/api"
	}
	router := mainRouter.Group(url_prefix)
	fmt.Println("JEE API ROUTES: ", url_prefix)

	router.Post("/get_result", route_get_result)
	router.Post("/get_sample_result", route_get_sample_result)
	router.Post("/submit_result_data", route_sumbit_result_data)
	router.Post("/page_view_count", route_page_view_count)
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

type result_response_type_with_key struct {
	// This type is used to store the result in the database
	ApplicationNumber string               `json:"key"`
	Result            types.ResultDataType `json:"result"`
	Data              types.MainDataType   `json:"data"`
}

func route_get_result(c *fiber.Ctx) error {
	var bdy get_result_route_body_type
	c.BodyParser(&bdy)

	var user_info types.MetaDataType
	user_info_fetched := kry.Deta.Base("info").Get(bdy.ApplicationNumber, &user_info)

	if user_info_fetched != nil {
		// Application Number not found in database
		return c.Status(403).JSON(fiber.Map{
			"detail": &kry.ErrorInfoType{
				Error: "appl_numb_not_found",
			},
		})
	}
	if user_info.DOB != bdy.DOB {
		// DOB did not match
		return c.Status(403).JSON(&fiber.Map{
			"detail": &kry.ErrorInfoType{
				Error: "dob_did_not_match",
			},
		})
	}

	var data result_response_type
	kry.Deta.Base("data").Get(bdy.ApplicationNumber, &data)

	update_result_view_count()

	return c.Status(200).JSON(&data)
}

func route_get_sample_result(c *fiber.Ctx) error {
	var data result_response_type
	kry.Deta.Base("data").Get("sample_result", &data)
	return c.Status(200).JSON(&data)
}

type submit_result_data_route_body_type struct {
	ResponsePageData string `json:"ResponsePageData"`
	AnswerKeyData    string `json:"AnswerKeyData"`
	DateOfBirth      string `json:"DateOfBirth"`
}

func route_sumbit_result_data(c *fiber.Ctx) error {
	var bdy submit_result_data_route_body_type
	c.BodyParser(&bdy)
	// detecting if the ResponsePageData is a URL
	if len(bdy.ResponsePageData) < 300 {
		line_count := strings.Count(bdy.ResponsePageData, "\n")
		if line_count <= 3 {
			URL_RE_PATTERN := regexp.MustCompile(`(http|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`)
			match_url := URL_RE_PATTERN.FindString(bdy.ResponsePageData)
			if match_url != "" {
				page_data, err := make_GET_request(match_url)
				if err == nil {
					bdy.ResponsePageData = page_data
				} else {
					return c.Status(403).JSON(&fiber.Map{
						"detail": &kry.ErrorInfoType{
							Error: "invalid_response_sheet_url",
						},
					})
				}
			}
		}
	}
	data, err := types.GetData(bdy.AnswerKeyData, bdy.ResponsePageData)
	if err != nil {
		return c.Status(403).JSON(&fiber.Map{
			"detail": &kry.ErrorInfoType{
				Error: "invalid_data",
			},
		})
	}
	result := types.GetResult(&data)
	// storing the result in the database
	meta_data, err := types.GetMetaData(bdy.ResponsePageData)
	if err != nil {
		return c.Status(403).JSON(&fiber.Map{
			"detail": &kry.ErrorInfoType{
				Error: "invalid_meta_data",
			},
		})
	}
	return_data := &result_response_type_with_key{
		ApplicationNumber: meta_data.ApplicationNumber,
		Result:            result,
		Data:              data,
	}
	{
		// storing the result in the database
		meta_data.DOB = bdy.DateOfBirth
		kry.Deta.Base("info").Put(&meta_data)
		kry.Deta.Base("data").Put(&return_data)
		update_result_view_count()
	}
	return c.Status(200).JSON(&return_data)
}

func make_GET_request(URL string) (string, error) {
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func route_page_view_count(c *fiber.Ctx) error {
	counts_base := kry.Deta.Base("counts")
	var update bool = c.Query("update_count") == "true"

	var page_view_count kry.KeyValuePairInt
	var result_view_count kry.KeyValuePairInt
	counts_base.Get("page_view_count", &page_view_count)
	counts_base.Get("result_view_count", &result_view_count)

	if update {
		page_view_count.Value++
		counts_base.Put(&page_view_count)
	}

	return c.Status(200).JSON(&fiber.Map{
		"page_view_count":   page_view_count.Value,
		"result_view_count": result_view_count.Value,
	})
}
