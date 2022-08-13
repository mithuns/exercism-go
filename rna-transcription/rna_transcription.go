package strand

func ToRNA(dna string) string {
	if len(dna) == 0 {
		return dna
	}
	var output []byte
	for i := range dna {
		switch dna[i] {
		case 'G':
			output = append(output, 'C')
		case 'C':
			output = append(output, 'G')
		case 'T':
			output = append(output, 'A')
		case 'A':
			output = append(output, 'U')
		default:
			return "invalid string"
		}
	}
	return string(output)
}
