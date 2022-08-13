package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {	  
	if len(rhyme) == 0 {
      return nil
  }
	results:= make([]string, 1)
	
	for index:=0; index<len(rhyme)-1;index++ {
		solution :=  fmt.Sprintf("For want of a %s the %s was lost.",rhyme[index],rhyme[index+1])
		if len(results[0]) == 0 {
			results[0] = solution
		}else{
			results = append(results, solution)
		}		
	}
	
	solution :=  fmt.Sprintf("And all for the want of a %s.",rhyme[0])
	if len(results[0]) == 0 {
		results[0] = solution
	}else{
		results = append(results, solution)
	}

	return results
}
