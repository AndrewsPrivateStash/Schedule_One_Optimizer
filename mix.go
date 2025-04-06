/*
	The mix class and constructors
*/

package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"
)

// object to optimize
type Mix struct {
	Prod        Product
	Ingredients []Ingredient
	Effects     []string
	Cost        float64
	Price       float64
}

// Mix constructor
func newMix(prd Product) *Mix {
	m := new(Mix)
	m.Prod = prd
	m.Cost = 0
	if prd.Base_effect != "" {
		m.Effects = append(m.Effects, prd.Base_effect)
	}
	m.Price = m.calc_price()

	return m
}

// make a new mix from an array of ingredients
func mixFromArray(prd Product, ings []string) *Mix {

	m := newMix(prd)

	for _, ing := range ings {
		ii, ei := get_ingr_by_name(ing)
		if ei != nil {
			log.Printf("%s\n", ei)
			continue
		}
		m.add_ingr(ii)
	}

	return m
}

// ############ Mix methods ############

func (m *Mix) copyMix() *Mix {
	newMix := new(Mix)
	newMix.Prod = m.Prod
	newMix.Ingredients = append(newMix.Ingredients, m.Ingredients...)
	newMix.Effects = append(newMix.Effects, m.Effects...)
	newMix.Cost = m.Cost
	newMix.Price = m.Price

	return newMix
}

// add ingredient to mix
func (m *Mix) add_ingr(ing Ingredient) {
	m.Ingredients = append(m.Ingredients, ing)
	new_effect := ing.Effect

	// loop over the existing effects and check mapping to see if any need to be replaced
	for i, eff := range m.Effects {
		if v, ok := TRANS[new_effect][eff]; ok {
			if !m.contains_effect(v) {
				m.Effects[i] = v
			}
		}
	}

	if !m.contains_effect(new_effect) && len(m.Effects) < 8 {
		m.Effects = append(m.Effects, new_effect)
	}

	m.update()

}

// calculate the mix cost
func (m *Mix) calc_cost() float64 {
	var c float64 = 0
	for _, ing := range m.Ingredients {
		c += ing.Cost
	}

	return c
}

// calculate the sell price for the mix
func (m *Mix) calc_price() float64 {

	base_price := m.Prod.Base_price

	var mult_sum float64 = 0

	add_mult := func(eff string, sum *float64) {
		if m, ok := MULT[eff]; ok {
			*sum += m
		}
	}

	// sum multipliers
	for _, eff := range m.Effects {
		add_mult(eff, &mult_sum)
	}

	return math.Round(base_price * (1 + mult_sum))
}

// calculate the profit
func (m *Mix) calc_profit() float64 {
	return m.calc_price() - m.calc_cost()
}

// check if an effect exists in effects member
func (m *Mix) contains_effect(search string) bool {
	return slices.Contains(m.Effects, search)
}

// update mix
func (m *Mix) update() {
	m.Price = m.calc_price()
	m.Cost = m.calc_cost()

	// sort the effects
	slices.Sort(m.Effects)
}

func (m *Mix) stringifyEffects() string {
	return strings.Join(m.Effects, "")
}

// find rank of mix
func (m *Mix) get_rank() string {
	var max_level int = 0
	var max_rank string = ""

	// Product rank
	if l, m := get_item_level(m.Prod.Name); l != -1 {
		max_level, max_rank = l, m
	}

	// inspect each ingredient and swap if greater
	for _, ing := range m.Ingredients {
		if v, ok := RANKS[ing.Rank]; ok {
			if v.Level > max_level {
				max_level = v.Level
				max_rank = v.Name
			}
		}
	}

	return max_rank
}

// pretty print the Mix
func (m *Mix) print() string {

	outStr := ""

	_, r := get_item_level(m.Prod.Name)
	outStr += fmt.Sprintf("\nMix for: %s\t(%s)\n", m.Prod.Name, r)
	outStr += "Ingredients:\n"

	// Ingredients
	for i, ing := range m.Ingredients {
		outStr += fmt.Sprintf("  %d: %-20s cost: %-6.2f rank: %-20s effect: %-30s\n", i+1, ing.Name, ing.Cost, ing.Rank, ing.Effect)
	}

	// Effects
	outStr += "\nEffects:\n{ "
	for i, eff := range m.Effects {
		if i == len(m.Effects)-1 {
			outStr += eff
		} else {
			outStr += fmt.Sprintf("%s, ", eff)
		}
	}
	outStr += " }\n"

	outStr += fmt.Sprintf("\nTotal Ingredient Cost: %.2f\n", m.Cost)
	outStr += fmt.Sprintf("Total Price: %.2f\n", m.Price)
	outStr += fmt.Sprintf("Total Gross Profit: %.2f\n\n", m.calc_profit())
	outStr += fmt.Sprintf("Required Rank: %s\n\n", m.get_rank())

	return outStr
}

// compact 1-line print mix
func (m *Mix) print_comp() string {
	out_str := fmt.Sprintf("%s: { ", m.Prod.Name)

	// Ingredients
	for i, ing := range m.Ingredients {
		if i < len(m.Ingredients)-1 {
			out_str += fmt.Sprintf("%s -> ", ing.Name)
		} else {
			out_str += fmt.Sprintf("%s } ", ing.Name)
		}
	}

	out_str += fmt.Sprintf("%.2f", m.calc_profit())

	return out_str

}

// ############ ############ ############
