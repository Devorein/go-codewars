package codewars

func DirReduction1(paths []string) []string {
	final_paths := []string{}
	last_path := ""

	for _, path := range paths {
		final_path_len := len(final_paths)
		if (last_path == "NORTH" && path == "SOUTH") || (path == "NORTH" && last_path == "SOUTH") || (last_path == "WEST" && path == "EAST") || (path == "WEST" && last_path == "EAST") {
			if final_path_len != 0 {
				final_paths = final_paths[:final_path_len-1]
			}
			new_final_path_len := len(final_paths)
			if new_final_path_len == 0 {
				last_path = ""
			} else {
				last_path = final_paths[new_final_path_len-1]
			}
			continue
		}
		last_path = path
		final_paths = append(final_paths, path)
	}

	return final_paths
}

func DirReduction2(paths []string) []string {
	final_paths := []string{}

	directionMap := map[string]string{
		"EAST":  "WEST",
		"WEST":  "EAST",
		"SOUTH": "NORTH",
		"NORTH": "SOUTH",
	}

	for _, path := range paths {
		l := len(final_paths)
		if l > 0 && directionMap[final_paths[l-1]] == path {
			final_paths = final_paths[:l-1]
		} else {
			final_paths = append(final_paths, path)
		}
	}

	return final_paths
}
