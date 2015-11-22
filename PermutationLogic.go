package main

import (
//  "fmt"
)

func removeFromArray(array []string, index int) []string{
  // Removes element at given index from array
  removed := make([]string, len(array[:index]) + len(array[index+1:]))
  copy(removed, array[:index])
  copy(removed[len(array[:index]):], array[index+1:])
  return removed
}

func permutations(input []string) [][]string{
  // Returns a list of all permutations
  perms := [][]string{}
  if len(input) == 1{
    perms = append(perms, input)
  } else {

    for i:=0;i<len(input);i++{
      arr_ex := removeFromArray(input, i)
      permutations_of_subset := permutations(arr_ex)
      for j:=0;j<len(permutations_of_subset);j++{
      
        p := append(permutations_of_subset[j], input[i])
        perms = append(perms, p)
      }
    }

  }
  return perms
  
}



