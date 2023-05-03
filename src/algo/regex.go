package Algo

import (
	"fmt"
	"regexp"
	"strconv"
);

func Regex(text string) string{
	/* Periksa apakah terdapat format tanggal dalam teks */
	regex := regexp.MustCompile(`([0][1-9]|[1-2][0-9]|[3][0-1])\/(0[1-9]|1[0-2])\/[0-9]{4}`);
	return_bool, _ := regexp.MatchString(`([0][1-9]|[1-2][0-9]|[3][0-1])\/(0[1-9]|1[0-2])\/[0-9]{4}`, text);
	if(return_bool){
		matches := regex.FindAllString(text, -1)
		// fmt.Println(matches[0]);
		d := Day(matches[0]);
		switch d{
		case 0:
			return("Hari Minggu")
		case 1:
			return("Hari Senin")
		case 2:
			return("Hari Selasa")
		case 3:
			return("Hari Rabu")
		case 4:
			return("Hari Kamis")
		case 5:
			return("Hari Jumat")
		case 6:
			return("Hari Sabtu")
		}
		return "day not found";
	}

	/* Periksa apakah ekspresi matematika yang valid terkandung dalam teks*/
	regex = regexp.MustCompile(`([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\))\s*(([-+*/^]\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\)))+)`);
	return_bool, _ = regexp.MatchString(`([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\))\s*(([-+*/^]\s*([-+]?\s*\d+(\.\d+)?|\(\s*.+?\s*\)))+)`, text);
	if(return_bool){
		matches := regex.FindAllString(text, -1)
		res, _ := Calculate(matches[0])
		return strconv.FormatFloat(res, 'f', 2, 64);
	}

	/* Periksa apakah terkandung kata Hapus Pertanyaan xxx */
	regex = regexp.MustCompile(`(?i)^Hapus Pertanyaan\s+(.*)$`);
	return_bool, _ = regexp.MatchString(`[Hh][Aa][Pp][Uu][Ss]\s[Pp][Ee][Rr][Tt][Aa][Nn][Yy][Aa][Aa][Nn]\s`, text);
	if(return_bool){
		matches := regex.FindStringSubmatch(text)
		fmt.Println(matches[1])
		return "Sukses memasukkan ke database";
	}

	/* Periksa apakah terkandung kata Tambah Pertanyaan XXX dengan jawaban YYY */
	regex = regexp.MustCompile(`(?i)^tambah pertanyaan\s+(.*?)\s+dengan jawaban\s+(.*?)$`);
	matches := regex.FindStringSubmatch(text)
	if(len(matches) > 0){
		fmt.Println(matches[1])
		fmt.Println(matches[2])
		return "Sukses memasukkan ke database"
	}

	return "Error";
}