package Algo

import (
	"fmt"
	"regexp"
);

func Regex(flag int, text string) bool{
	if(flag == 1){
		/* Periksa apakah terdapat format tanggal dalam teks */
		regex := regexp.MustCompile(`[0-9]{2}\/[0-9]{2}\/[0-9]{4}`);
		return_bool, _ := regexp.MatchString("[0-9]{2}/[0-9]{2}/[0-9]{4}", text);
		if(return_bool){
			matches := regex.FindAllString(text, -1)
			fmt.Println(matches[0]);
			Day(matches[0]);
		}

		return return_bool;

	}else if(flag == 2){
		/* Periksa apakah ekspresi matematika yang valid terkandung dalam teks*/
		regex := regexp.MustCompile(`^\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\))\s*(([-+*/]\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\)))+)\s*$`);
		return_bool, _ := regexp.MatchString(`^\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\))\s*(([-+*/]\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\)))+)\s*$`, text);
		if(return_bool){
			matches := regex.FindAllString(text, -1)
			fmt.Println(matches[0]);
			res, err := Calculate(matches[0]);
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %v\n", res)
			}
		}
		return return_bool;
	}

	return true;
}