package chef

func Exists(chefID string) bool {
	for _, chef := range chefs {
		if chef.Id == chefID {
			return true
		}
	}

	return false
}
