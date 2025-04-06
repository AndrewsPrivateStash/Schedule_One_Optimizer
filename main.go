/*
	find optimal mixes by ingredient count, tier and type

	ToDO:
	- consider json ingestion for the data sets (TRANS, PRODS, etc)

	Notes:
	- mixes can have a maximum of 8 effects; ingredients added past this limit can only transform effects
	- if transformation effect already present then don't transform
		> so both the ingr effect and any transformation must not already be present to be added

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ingCntF  = flag.Int("c", 8, "the upper bound of ingredients to optimize")
	lvlF     = flag.String("r", "Kingpin I", "rank limit to apply")
	prodF    = flag.String("p", "", "product to optimize (default all)")
	outFileF = flag.String("f", "optimized_mix.txt", "file to write to in current directory")
	makeDBF  = flag.Bool("db", false, "generate db of results and write to file")
	manProcF = flag.Bool("man", false, "bypass the DB and recalculate the optimization")
)

// #############################

func main() {
	flag.Parse()

	// make DB here
	if *makeDBF {
		makeDB(8)
		return
	}

	ingCount := *ingCntF

	var rankLimit string
	if *lvlF == "" {
		rankLimit = "Kingpin I"
	} else if _, ok := RANKS[*lvlF]; ok {
		rankLimit = *lvlF
	} else {
		fmt.Printf("did not find rank: %s\n", *lvlF)
		printRanks()
		os.Exit(1)
	}

	// read DB here
	// Db : product -> rank -> [8]Mix
	var db map[string]map[string][]Mix
	if !*manProcF {
		if _, err := os.Stat("db.dat"); err == nil {
			log.Println("loading Db..")
			readGob("db.dat", &db)
		} else {
			log.Println("DB not found, run with -db flag to generate")
		}
	}

	var product Product
	if *prodF != "" {
		if p, err := get_prod_by_name(*prodF); err != nil {
			fmt.Printf("%s\n", err)
			printProducts()
			os.Exit(1)
		} else {
			product = p
		}

		checkRank(product, rankLimit) // ensure the product is available at this rank

		var opt_mix []*Mix

		// check DB first
		if db != nil && !*manProcF {
			m, err := dbLookUp(product, rankLimit, db)
			if err != nil {
				log.Printf("%s\ncalculating manually..\n", err)
				opt_mix = opt_mix_exh(product, ingCount, rankLimit)
			} else {
				opt_mix = m
			}

		} else {
			opt_mix = opt_mix_exh(product, ingCount, rankLimit)
		}

		res := makeOutputString(opt_mix, ingCount)
		fmt.Printf("%s", res)
		writeResults(res, *outFileF)
		return
	}

	// optimze all products
	opt_mixes := optAllProducts(rankLimit, ingCount, db)
	res := ""
	for _, m_o := range opt_mixes {
		res += fmt.Sprintf("#### %s ####\n", m_o[0].Prod.Name)
		res += makeOutputString(m_o, ingCount) + "\n"
	}

	fmt.Printf("%s", res)
	writeResults(res, *outFileF)

}

// make output string
func makeOutputString(ms []*Mix, ing_cnt int) string {
	res := ""
	for i := range len(ms) - 1 {
		res += fmt.Sprintf("%d  %s\n", i+1, ms[i].print_comp())
	}
	res += "\n" + ms[ing_cnt-1].print()
	return res
}

// make sure product is available at the rank
func checkRank(p Product, rank string) {
	lvl, rnk := get_item_level(p.Name)
	if lvl > RANKS[rank].Level {
		log.Printf("warning: %s is rank: %s, limit of %s specified\n\n", p.Name, rnk, rank)
	}
}

// optimize ALL products
func optAllProducts(rank string, ing_cnt int, db map[string]map[string][]Mix) [][]*Mix {

	opt_mixes := [][]*Mix{}
	for _, p := range PRODS {

		if db != nil {
			ms, err := dbLookUp(p, rank, db)
			if err != nil {
				log.Printf("%s\ncalculating manually..\n", err)
				ms = opt_mix_exh(p, ing_cnt, rank)
			}
			opt_mixes = append(opt_mixes, ms)

		} else {
			opt_mixes = append(opt_mixes, opt_mix_exh(p, ing_cnt, rank))
		}

	}

	return opt_mixes
}

// print ranks in order to stdout
func printRanks() {
	ranks := genSigRanks()

	fmt.Println("Significant Ranks:")
	for i, k := range ranks {
		if i == len(ranks)-1 {
			fmt.Printf("%s\n", k.Name)
		} else {
			fmt.Printf("%s, ", k.Name)
		}
	}
}

// print products to stdout
func printProducts() {
	for i, k := range PRODS {
		if i == len(PRODS)-1 {
			fmt.Printf("%s\n", k.Name)
		} else {
			fmt.Printf("%s, ", k.Name)
		}
	}
}

// write file with results
func writeResults(str string, dest string) {
	output := []byte(str)

	f, err := os.Create(dest)
	if err != nil {
		log.Printf("Could not open file for writing: %s\n", dest)
		return
	}
	defer f.Close()

	if _, err := f.Write(output); err != nil {
		log.Printf("Could not write to file: %s\n", dest)
		return
	}

}
