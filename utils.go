/*
	general utility functions
*/

package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
)

// translate the map of maps for transformations into a grid
func make_grid() {
	all_effs := get_all_effects()

	fmt.Printf("All Effects\t")
	for _, eff := range all_effs {
		fmt.Printf("%s\t", eff)
	}
	fmt.Printf("\n")

	for _, ing := range INGRS {
		fmt.Printf("%s_%s\t", ing.Name, ing.Effect)
		for _, eff := range all_effs {
			fmt.Printf("%s\t", TRANS[ing.Effect][eff])
		}
		fmt.Printf("\n")
	}

}

// get all the effects possible
func get_all_effects() []string {

	is_in := func(s string, effs []string) bool {
		for _, e := range effs {
			if s == e {
				return true
			}
		}
		return false
	}

	effects := []string{}

	for k, v := range TRANS {
		if !is_in(k, effects) {
			effects = append(effects, k)
		}
		for k2, v2 := range v {
			if !is_in(v2, effects) {
				effects = append(effects, v2)
			}
			if !is_in(k2, effects) {
				effects = append(effects, k2)
			}
		}
	}

	return effects
}

// are two mixes equevelent in terms of effect permutations
func effects_equal(m1 *Mix, m2 *Mix) bool {
	if len(m1.Effects) != len(m2.Effects) {
		return false
	}

	for _, e := range m1.Effects {
		if !m2.contains_effect(e) {
			return false
		}
	}

	return true
}

// get all the ranks that unlock something relavent
func genSigRanks() []Rank {
	ranks := make([]string, 0, 10)

	// grab sig products
	for _, p := range PRODS {
		for k, v := range RANKS {
			if slices.Contains(v.Unlocks, p.Name) {
				if !slices.Contains(ranks, k) {
					ranks = append(ranks, k)
				}
			}
		}
	}

	// grab sig ingredients
	for _, ing := range INGRS {
		for k, v := range RANKS {
			if slices.Contains(v.Unlocks, ing.Name) {
				if !slices.Contains(ranks, k) {
					ranks = append(ranks, k)
				}
			}
		}
	}

	// make slice of Ranks to sort on level
	rankSlice := []Rank{}
	for _, rs := range ranks {
		rankSlice = append(rankSlice, RANKS[rs])
	}

	sort.Slice(rankSlice, func(i, j int) bool {
		return rankSlice[i].Level < rankSlice[j].Level
	})

	return rankSlice

}

// make array of mix pointers
func makeMixPointerArr(ms []Mix) []*Mix {
	out_mixes := []*Mix{}
	for _, m := range ms {
		out_mixes = append(out_mixes, &m)
	}
	return out_mixes
}

// get product by name
func get_prod_by_name(pr_name string) (Product, error) {
	for i, p := range PRODS {
		if pr_name == p.Name {
			return PRODS[i], nil
		}
	}
	return Product{}, errors.New(pr_name + " product not found")
}

// get ingredient by name
func get_ingr_by_name(ingr_name string) (Ingredient, error) {
	for i, p := range INGRS {
		if ingr_name == p.Name {
			return INGRS[i], nil
		}
	}
	return Ingredient{}, errors.New(ingr_name + " ingredient not found")
}

// get the items level
func get_item_level(item string) (int, string) {
	for key, val := range RANKS {
		if slices.Contains(val.Unlocks, item) {
			return val.Level, key
		}
	}
	return -1, "" // not found
}
