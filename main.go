package main

import (
	"fmt"
	"strings"
)

type production struct {
	left  string
	right []string
}

func main() {
	productions := []production{

        //Prueba #1 La que está en las diapositivas de Cristian
		{"E", []string{"T", "EP"}},
		{"EP", []string{"+", "T", "EP"}},
		{"EP", []string{"lambda"}},
		{"T", []string{"F", "TP"}},
		{"TP", []string{"*", "F", "TP"}},
		{"TP", []string{"lambda"}},
		{"F", []string{"id" }},
		{"F", []string{"(", "E", ")"}},

        /*
        Primeros:
       F: {id, (}
	   T: {id, (}
       E: {id, (}
       EP: {+, lambda}
       TP: {*, lambda}
        */

        // //Prueba #2
        // {"A", []string{"B", "C"}},
        // {"B", []string{"lambda"}},
        // {"B", []string{"m"}},
        // {"C", []string{"lambda"}},
        // {"C", []string{"s"}},

        // //Prueba #3 Sacada del vídeo del Santi : https://www.youtube.com/watch?v=neZpR3V2WNY&t=615s
        // {"E", []string{"T", "EP"}},
        // {"EP", []string{"+", "T", "EP"}},
        // {"EP", []string{"-", "T", "EP"}},
        // {"EP", []string{"lambda"}},
        // {"T", []string{"F", "TP"}},
        // {"TP", []string{"*", "F", "TP"}},
        // {"TP", []string{"/", "F", "TP"}},
        // {"TP", []string{"lambda"}},
        // {"F", []string{"(", "E", ")"}},
        // {"F", []string{"num"}},
        // {"F", []string{"id"}},
        
        /*
        Primeros:
        EP: {+, -, lambda}
        TP: {*, /, lambda}
        F: {(, num, id}
        T: {(, num, id}
        E: {(, num, id}
        */

		// //Prueba #4 La del parcial
		// {"AL", []string{"id", ":=", "P"}},
		// {"P", []string{"D", "PP"}},
		// {"PP", []string{"or","D", "PP"}},
		// {"PP", []string{"lambda"}},
		// {"D", []string{"C", "DP"}},
		// {"DP", []string{"and","C","DP"}},
		// {"DP", []string{"lambda"}},
		// {"C", []string{"S"}},
		// {"C", []string{"not", "(", "P", ")"}},
		// {"S", []string{"(", "P", ")"}},
		// {"S", []string{"OP", "REL", "OP"}},
		// {"S", []string{"true"}},
		// {"S", []string{"false"}},
		// {"REL", []string{"=",}},
		// {"REL", []string{"<","RP"}},
		// {"REL", []string{">", "EP"}},
		// {"RP", []string{"="}},
		// {"RP", []string{">"}},
		// {"RP", []string{"lambda"}},
		// {"EP", []string{"="}},
		// {"EP", []string{"lambda"}},
		// {"OP", []string{"id"}},
		// {"OP", []string{"num"}},
	}

	// Se buscan los primeros
	firsts := make(map[string][]string)
	for _, p := range productions {
		if isTerminal(p.right[0]) {
			addToSet(firsts, p.left, p.right[0])
		}
	}
	for changed := true; changed; {
		changed = false
		for _, p := range productions {
			if len(p.right) > 0 && !isTerminal(p.right[0]) {
				oldLen := len(firsts[p.left])
				for _, s := range firsts[p.right[0]] {
					addToSet(firsts, p.left, s)
				}
				if len(firsts[p.left]) > oldLen {
					changed = true
				}
			}
		}
	}

	fmt.Println("Primeros:")
	for nt, s := range firsts {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}

	// Se buscan los siguientes
	follows := make(map[string][]string)
	addToSet(follows, "E", "$")
	for changed := true; changed; {
		changed = false
		for _, p := range productions {
			for i := 0; i < len(p.right); i++ {
				if !isTerminal(p.right[i]) {
					oldLen := len(follows[p.right[i]])
					if i == len(p.right)-1 {
						// Añade los siguientes de la izquierda de la producción
						for _, s := range follows[p.left] {
							addToSet(follows, p.right[i], s)
						}
					}
					for j := i + 1; j < len(p.right); j++ {
						if isTerminal(p.right[j]) && p.right[j]!= "lambda"{
							addToSet(follows, p.right[i], p.right[j])
						}
							for _, s := range firsts[p.right[j]] {
								if s != "lambda" {
									addToSet(follows, p.right[i], s)
								} else {
									// Si lambda está en los primeros, añade los siguientes de la izquierda
									for _, t := range follows[p.left] {
										addToSet(follows, p.right[i], t)
									}
								}
							}
							if containsLambda(firsts[p.right[j]]) {
								continue
							}
							break
					}
					if len(follows[p.right[i]]) > oldLen {
						changed = true
					}
				}
			}
		}
	}

	// Imprime los siguientes
	fmt.Println("Siguientes:")
	for nt, s := range follows {
		fmt.Printf("%s: {%s}\n", nt, strings.Join(s, ", "))
	}
}

// Determina si un símbolo es un terminal
func isTerminal(symbol string) bool {
	return symbol == "lambda" || symbol[0] < 'A' || symbol[0] > 'Z'
}

// Agrega uno o más elementos a un conjunto. Si un elemento ya está en el conjunto, no se agrega de nuevo.
func addToSet(m map[string][]string, key string, values ...string) {
	for _, value := range values {
		if !contains(m[key], value) {
			m[key] = append(m[key], value)
		}
	}
}

// Determina si un elemento está en una lista de strings
func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// Determina si el símbolo "lambda" está en una lista de strings (se utiliza para verificar si "lambda" está en los primeros de un símbolo no terminal).
func containsLambda(s []string) bool {
	return contains(s, "lambda")
}
