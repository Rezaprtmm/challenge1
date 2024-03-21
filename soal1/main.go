package main

import "fmt"

func averageScores(scores []int) float64 {
	total := 0
	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}
	return float64(total) / float64(len(scores))
}

func main() {
	scoreLumba := []int{97, 112, 101}
	scoreKoala := []int{109, 95, 106}

	avgLumba := averageScores(scoreLumba)
	avgKoala := averageScores(scoreKoala)

	minimumScore := 100

	if avgLumba >= float64(minimumScore) && avgKoala >= float64(minimumScore) {
		if avgLumba > avgKoala {
			fmt.Println("Tim lumba menang dengan mendapatkan score rata-rata", avgLumba)
		} else if avgLumba < avgKoala {
			fmt.Println("Tim koala menang dengan mendapatkan score rata-rata", avgKoala)
		} else {
			fmt.Println("Hasil seri! Kedua tim memiliki skor rata-rata yang sama dengan rata-rata", avgLumba)
		}
	} else if avgLumba >= float64(minimumScore) {
		fmt.Println("Tim Lumba-lumba memenangkan trofi dengan skor rata-rata", avgLumba)
	} else if avgKoala >= float64(minimumScore) {
		fmt.Println("Tim Koala memenangkan trofi dengan skor rata-rata", avgKoala)
	} else {
		fmt.Println("Tidak ada pemenang, kedua tim tidak mencapai skor minimum", minimumScore)
	}
}
