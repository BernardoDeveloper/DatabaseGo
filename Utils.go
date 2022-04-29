package main

func isElementExist(s []Syntax, str int) Syntax {
	for _, v := range s {
		if v.Id == str {
			var data = Syntax{
				Id:    v.Id,
				Key:   v.Key,
				Value: v.Value,
			}

			return data
		}
	}

	return Syntax{}
}
