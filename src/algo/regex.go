package Algo

import (
	model "cad/models"
	// "fmt"
	"regexp"
	"strconv"
	"sort"
	"strings"
)

func Regex(text string, questions []model.Question, newQuestion *model.Question, searchMethod int) (string, int) {
	/* retcode :
		-1 	: regex gagal
		1 	: regex berhasil
		2 	: tambah pertanyaan
		3	: hapus pertanyaan
	*/

	/* Periksa apakah terdapat format tanggal dalam teks */
	regex := regexp.MustCompile(`([0][1-9]|[1-2][0-9]|[3][0-1])\/(0[1-9]|1[0-2])\/[0-9]{4}`)
	return_bool, _ := regexp.MatchString(`([0][1-9]|[1-2][0-9]|[3][0-1])\/(0[1-9]|1[0-2])\/[0-9]{4}`, text)
	if return_bool {
		matches := regex.FindAllString(text, -1)
		d := Day(matches[0])
		switch d {
		case 0:
			return ("Hari Minggu"), 1 
		case 1:
			return ("Hari Senin"), 1
		case 2:
			return ("Hari Selasa"), 1
		case 3:
			return ("Hari Rabu"), 1
		case 4:
			return ("Hari Kamis"), 1
		case 5:
			return ("Hari Jumat"), 1
		case 6:
			return ("Hari Sabtu"), 1
		}
		return "day not found", 1
	}

	/* Periksa apakah ekspresi matematika yang valid terkandung dalam teks*/
	regex = regexp.MustCompile(`(?:\s|^)([\-\+\*\^\/\(\)\s?0-9.?]+)`)
	return_bool, _ = regexp.MatchString(`(?:\s|^)([\-\+\*\^\/\(\)\s?0-9.?]+)`, text)
	if return_bool {
		matches := regex.FindAllString(text, -1)
		res, _ := Calculate(matches[0])
		return strconv.FormatFloat(res, 'f', 2, 64), 1
	}

	/* Periksa apakah terkandung kata Hapus Pertanyaan xxx */
	regex = regexp.MustCompile(`(?i)^Hapus Pertanyaan\s+(.*)$`)
	return_bool, _ = regexp.MatchString(`[Hh][Aa][Pp][Uu][Ss]\s[Pp][Ee][Rr][Tt][Aa][Nn][Yy][Aa][Aa][Nn]\s`, text)
	if return_bool {
		matches := regex.FindStringSubmatch(text)
		newQuestion.Question = matches[1]
		return "", 3
	}

	/* Periksa apakah terkandung kata Tambah Pertanyaan XXX dengan jawaban YYY */
	regex = regexp.MustCompile(`(?i)^tambah pertanyaan\s+(.*?)\s+dengan jawaban\s+(.*?)$`)
	matches := regex.FindStringSubmatch(text)
	if len(matches) > 0 {
		// fmt.Println(matches[1])
		// fmt.Println(matches[2])
		newQuestion.Question = matches[1]
		newQuestion.Answer = matches[2]
		return "", 2
	}

	/* Periksa dengan KMP atau BM */
	for _, question := range questions {
		// make text and question to lowercase
		if(searchMethod == 1){
			if KMP(strings.ToLower(text), strings.ToLower(question.Question)) != -1 {
				return question.Answer, 1
			}
		} else {
			if BMMatch(strings.ToLower(text), strings.ToLower(question.Question)) != -1 {
				return question.Answer, 1
			}
		}
	}

	/* Jika tidak ada yang exact match, coba levenstein */
	ans := ""
	num := 1
	
	/* sort questions berdasarkan match ratio-nya */
	sortedQuestions := sortByMatchRatio(questions, text)
	
	/* masukkan top 3 questions yang memiliki match ratio < 0.9 */
	for i, question := range sortedQuestions {
		if i >= 3 || MatchRatio(strings.ToLower(text), strings.ToLower(question.Question)) >= 0.9 {
			break
		}
		ans += strconv.Itoa(num) + ". " + question.Question + " (" + strconv.FormatFloat(MatchRatio(strings.ToLower(text), strings.ToLower(question.Question))*100, 'f', 1, 64) + "%)\n"
		num++
	}
	
	/* jika ada yang match ratio-nya >= 0.9, return jawaban pertama */
	if num == 1 {
		return sortedQuestions[0].Answer, 1 
	}
	
	/* jika tidak, return pertanyaan yang memiliki match ratio tertinggi */
	if ans != "" {
		return "Mungkin pertanyaan yang anda maksud adalah :\n" + ans, 1
	}

	return "Error", -1
}

func sortByMatchRatio(questions []model.Question, text string) []model.Question {
	sort.Slice(questions, func(i, j int) bool {
		return MatchRatio(strings.ToLower(text), strings.ToLower(questions[i].Question)) > MatchRatio(strings.ToLower(text), strings.ToLower(questions[j].Question))
	})
	return questions
}
