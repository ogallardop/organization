package organisation

type Employee struct {
	id   string
	name string
	emp  []*Employee
}

func commonManager(ceo Employee, empA Employee, empB Employee) Employee {
	var pathA, pathB []Employee
	_, pathA = getPath(ceo, empA, nil)
	_, pathB = getPath(ceo, empB, nil)

	level := len(pathA)
	if level > len(pathB) {
		level = len(pathB)
	}

	i := level - 1
	for i >= 0 {
		if pathA[i].id == pathB[i].id {
			return pathA[i]
		}
		i--
	}

	return ceo
}

func getPath(manager Employee, emp Employee, path []Employee) (bool, []Employee) {
	if manager.id == emp.id {
		return true, path
	}

	if len(manager.emp) > 0 {
		path = append(path, manager)
		for _, v := range manager.emp {
			found, path := getPath(*v, emp, path)
			if found {
				return true, path
			}
		}
	}

	return false, path
}
