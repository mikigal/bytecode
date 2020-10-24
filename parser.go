package bytecode

func countParameters(descriptor string) uint16 {
	parameters := 0
	foundReference := false

	for _, char := range descriptor {
		if char == '(' || char == '[' {
			continue
		}

		if char == ')' {
			break
		}

		if char == 'L' {
			foundReference = true
		}

		if foundReference {
			if char == ';' {
				foundReference = false
				parameters++
			}

			continue
		}

		parameters++
	}

	return uint16(parameters)
}