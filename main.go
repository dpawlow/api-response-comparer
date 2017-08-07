package main

import (
	"fmt"
	"time"
	"log"
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func main() {
	defer un(trace("DURATION:"))

	urls := make([][]string, 32)

	for i := range urls {
		urls[i] = make([]string, 2)
	}

	urls[0][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla4239"
	urls[0][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla4239"

	urls[1][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mco5607"
	urls[1][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mco5607"

	urls[2][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlb1055"
	urls[2][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlb1055"

	urls[3][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla1645"
	urls[3][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla1645"

	urls[4][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla1649"
	urls[4][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla1649"

	urls[5][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla4611"
	urls[5][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla4611"

	urls[6][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mpe1910"
	urls[6][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mpe1910"

	urls[7][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mrd3530"
	urls[7][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mrd3530"

	urls[8][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mpe1714"
	urls[8][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mpe1714"

	urls[9][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla86838"
	urls[9][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla86838"

	urls[10][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mpe1712"
	urls[10][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mpe1712"

	urls[11][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mpe4660"
	urls[11][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mpe4660"

	urls[12][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlc3937"
	urls[12][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlc3937"

	urls[13][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlu9959"
	urls[13][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlu9959"

	urls[14][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlm6472"
	urls[14][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlm6472"

	urls[15][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla8803"
	urls[15][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla8803"

	urls[16][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla1651"
	urls[16][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla1651"

	urls[17][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlc1384"
	urls[17][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlc1384"

	urls[18][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla1656"
	urls[18][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla1656"

	urls[19][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlv1574"
	urls[19][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlv1574"

	urls[20][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlv112138"
	urls[20][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlv112138"

	urls[21][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlc1246"
	urls[21][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlc1246"

	urls[22][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mpt1743"
	urls[22][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mpt1743"

	urls[23][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla4624"
	urls[23][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla4624"

	urls[24][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlb1196"
	urls[24][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlb1196"

	urls[25][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlm4887"
	urls[25][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlm4887"

	urls[26][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla5892"
	urls[26][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla5892"

	urls[27][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlv2968"
	urls[27][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlv2968"

	urls[28][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlu5998"
	urls[28][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlu5998"

	urls[29][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlv1582"
	urls[29][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlv1582"

	urls[30][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mla8814"
	urls[30][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mla8814"

	urls[31][0] = "https://api.mercadolibre.com/sites/MLA/search_configurations/mlm8626"
	urls[31][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLA/search_configurations/mlm8626"
/*
	urls[1][0] = "https://api.mercadolibre.com/sites/MLB/search_configurations/dump"
	urls[1][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLB/search_configurations/dump"

	urls[2][0] = "https://api.mercadolibre.com/sites/MLM/search_configurations/dump"
	urls[2][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLM/search_configurations/dump"

	urls[3][0] = "https://api.mercadolibre.com/sites/MLV/search_configurations/dump"
	urls[3][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLV/search_configurations/dump"

	urls[4][0] = "https://api.mercadolibre.com/sites/MLU/search_configurations/dump"
	urls[4][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLU/search_configurations/dump"

	urls[5][0] = "https://api.mercadolibre.com/sites/MCO/search_configurations/dump"
	urls[5][1] = "http://furytest.search-config-api.melifrontends.com/sites/MCO/search_configurations/dump"

	urls[6][0] = "https://api.mercadolibre.com/sites/MLC/search_configurations/dump"
	urls[6][1] = "http://furytest.search-config-api.melifrontends.com/sites/MLC/search_configurations/dump"

	urls[7][0] = "https://api.mercadolibre.com/sites/MPA/search_configurations/dump"
	urls[7][1] = "http://furytest.search-config-api.melifrontends.com/sites/MPA/search_configurations/dump"

*/
	var diffResponses []bool = compareResponses(urls)
	fmt.Println(diffResponses)
}
