package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	h = make(map[rune]int, 0)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	for _, r := range d {
		switch rune(r) {
		case 'A':
			h['A'] = h['A'] + 1
		case 'C':
			h['C'] = h['C'] + 1
		case 'G':
			h['G'] = h['G'] + 1
		case 'T':
			h['T'] = h['T'] + 1
		default:
			return nil, errors.New("Invalid nucleotide")
		}

	}
	return h, nil
}
