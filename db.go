/*
	serialize the results of optimization to quickly return results
	DB should be: map[string]map[rank][]Mix

	TODO
*/

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// generate all optimizations and store in map
// serialize the result and store in current dir
func makeDB(limit int) {
	db := optAllDB(limit)
	if err := writeGob("db.dat", db); err != nil {
		log.Fatal("error serializing db", err)
	}
}

func optAllDB(limit int) map[string]map[string][]Mix {
	db := map[string]map[string][]Mix{}

	for _, p := range PRODS {
		for k, _ := range RANKS {
			log.Printf("building: %s, %s, %d\n\n", p.Name, k, limit)
			opt := opt_mix_exh(p, limit, k)

			if _, ok := db[p.Name]; ok {
				db[p.Name][k] = derefMixArr(opt)
			} else {
				db[p.Name] = map[string][]Mix{}
				db[p.Name][k] = derefMixArr(opt)
			}

		}
	}

	return db
}

// remove product from db
func deleteProdDB(p_nm string, db map[string]map[string][]Mix) {
	if _, ok := db[p_nm]; ok {
		delete(db, p_nm)
	} else {
		log.Fatalf("%s not in DB to remove\n", p_nm)
	}
}

// update the DB on product dimension
func updateProdDB(p Product, limit int, db map[string]map[string][]Mix) {
	for k := range RANKS {
		log.Printf("building: %s, %s, %d\n\n", p.Name, k, limit)
		opt := opt_mix_exh(p, limit, k)

		if _, ok := db[p.Name]; ok {
			db[p.Name][k] = derefMixArr(opt)
		} else {
			db[p.Name] = map[string][]Mix{}
			db[p.Name][k] = derefMixArr(opt)
		}

	}
}

func derefMixArr(arr []*Mix) []Mix {
	out_arr := make([]Mix, 0, 8)
	for _, m := range arr {
		out_arr = append(out_arr, *m)
	}
	return out_arr
}

func writeGob(filePath string, data map[string]map[string][]Mix) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	return err
}

func readGob(filePath string, data *map[string]map[string][]Mix) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(data)
	if err != nil {
		log.Fatal("Error decoding data:", err)
	}
}

// look up mix in db
func dbLookUp(p Product, r string, db map[string]map[string][]Mix) ([]*Mix, error) {
	var mixes []*Mix
	if m, ok := db[p.Name][r]; ok {
		mixes = makeMixPointerArr(m)
	} else {
		return nil, fmt.Errorf("%s, %s not found in DB", p.Name, r)
	}

	return mixes, nil

}
