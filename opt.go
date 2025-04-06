/*
	Optimization functions

	ToDo
		-

*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// ############### Optimization ###############

// perform an exhaustive search, maximizing profit
func opt_mix_exh(p Product, cnt int, rank string) []*Mix {

	ret_mixes := []*Mix{}
	var max_lvl int = RANKS[rank].Level

	// filter possible ingredients by max rank
	IngrFiltered := []Ingredient{}
	for _, ing := range INGRS {
		if RANKS[ing.Rank].Level <= max_lvl {
			IngrFiltered = append(IngrFiltered, ing)
		}
	}

	priorMixes := make([]*Mix, 0, 256)
	newMixes := make([]*Mix, 0, 256)

	s1 := time.Now()

	// do first pass manually to populate with order 1 values
	fmt.Printf("working on order: 1")
	for i, ing := range IngrFiltered {
		priorMixes = append(priorMixes, newMix(p))
		priorMixes[i].add_ingr(ing)
	}

	best_idx := 0
	best_profit := 0.0

	for i, m := range priorMixes {
		if pr := m.calc_profit(); pr > best_profit {
			best_idx = i
			best_profit = pr
		}
	}

	fmt.Printf("  %-8d  %-14s  ", len(priorMixes), time.Since(s1))
	fmt.Printf("%s\n", priorMixes[best_idx].print_comp())
	ret_mixes = append(ret_mixes, priorMixes[best_idx])

	// multi ingredient case
	for c := 2; c <= cnt; c++ {
		fmt.Printf("working on order: %d", c)

		best_idx = 0
		best_profit = 0
		newMixes = newMixes[:0] // clear slice for next order

		divisor := 8
		if len(priorMixes) < 800 {
			divisor = 1
		} else {
			divisor = len(priorMixes) / 8
		}

		chunks := make_chunks(priorMixes, divisor)
		workers := len(chunks)

		ch := make(chan *Mix, 256)
		wks := make(chan int, workers)
		bst := make(chan *Mix, workers)

		// worker function
		make_mixes := func(arr []*Mix, ings []Ingredient) {

			var best_mix *Mix = nil
			var best_profit float64 = 0.0

			for _, m := range arr {
				for _, ing := range ings {

					m_i := m.copyMix()
					m_i.add_ingr(ing)

					if pr := m_i.calc_profit(); pr > best_profit {
						best_mix = m_i
						best_profit = pr
					}

					ch <- m_i
				}
			}
			bst <- best_mix
			wks <- -1

		}

		s1 = time.Now()

		// schedule workers
		for _, c := range chunks {
			go make_mixes(c, IngrFiltered)
		}

		best_results := []*Mix{}

		var cont bool = true
		for cont {
			select {

			case res, ok_ch := <-ch:
				if ok_ch {
					newMixes = append(newMixes, res)
					// newMixes = appendUnique(newMixes, res)
				} else {
					cont = false
				}

			case bst, ok_b := <-bst:
				if ok_b {
					best_results = append(best_results, bst)
				}

			case wkr, ok_wkr := <-wks:
				if ok_wkr {
					workers += wkr
					if workers == 0 {
						close(wks)
						close(ch)
						close(bst)
					}
				}
			}

		}

		best_idx = 0
		best_profit = 0
		for i, m := range best_results {
			if pr := m.calc_profit(); pr > best_profit {
				best_idx = i
				best_profit = pr
			}
		}

		// set the prior order to the current for the next order
		if c != cnt {
			priorMixes = priorMixes[:0]
			priorMixes = append(priorMixes, concDedup(newMixes)...)
		}

		fmt.Printf("  %-8d  %-14s  ", len(newMixes), time.Since(s1))
		fmt.Printf("%s\n", best_results[best_idx].print_comp())

		ret_mixes = append(ret_mixes, best_results[best_idx])

	}

	return ret_mixes

}

// take a slice and return an array of subslices
func make_chunks(input []*Mix, chunkSize int) [][]*Mix {
	var chunks [][]*Mix
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

// sort mixes on effects then cost descending
func sortMixes(arr []*Mix) {
	sort.Slice(arr, func(i, j int) bool {

		efs1 := strings.Join(arr[i].Effects, "")
		efs2 := strings.Join(arr[j].Effects, "")
		if efs1 != efs2 {
			return efs1 < efs2
		}

		return arr[i].Cost > arr[j].Cost // sort descending since we want to keep the last one (lowest cost)

	})
}

// remove dominated mixes
func dedupSortedMixes(arr []*Mix) []*Mix {
	deduped_arr := make([]*Mix, 0, len(arr))

	for i, m := range arr {
		if i == len(arr)-1 {
			deduped_arr = append(deduped_arr, m)
			break
		}

		if m.stringifyEffects() == arr[i+1].stringifyEffects() {
			continue
		}

		deduped_arr = append(deduped_arr, m)
	}

	return deduped_arr
}

// sort mixes then remove dominated
func sortRemoveDups(arr []*Mix) []*Mix {
	sortMixes(arr)
	return dedupSortedMixes(arr)
}

// concurrent dedup
func concDedup(arr []*Mix) []*Mix {

	sortMixes(arr)

	divisor := 8
	if len(arr) < 8 {
		divisor = 1
	} else {
		divisor = len(arr) / 8
	}

	chunks := make_chunks(arr, divisor)
	workers := len(chunks)

	ch := make(chan []*Mix, 8)
	wks := make(chan int, workers)

	job := func(arr []*Mix) {
		ch <- dedupSortedMixes(arr)
		wks <- -1
	}

	for _, arr := range chunks {
		go job(arr)
	}

	results := make([]*Mix, 0, len(arr)/2)

	var cont bool = true
	for cont {
		select {

		case res, ok_ch := <-ch:
			if ok_ch {
				results = append(results, res...)
			} else {
				cont = false
			}

		case wkr, ok_wkr := <-wks:
			if ok_wkr {
				workers += wkr
				if workers == 0 {
					close(wks)
					close(ch)
				}
			}
		}

	}

	return results

}
