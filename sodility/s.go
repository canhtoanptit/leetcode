package sodility

func solutionA(id int, employeeInfos []*EmployeeInfo) int {
	//totalImportantValue := 0

	// convert from array to map
	mapIdAndEmployee := make(map[int]*EmployeeInfo, len(employeeInfos))
	for _, employeeInfo := range employeeInfos {
		mapIdAndEmployee[employeeInfo.Id] = employeeInfo
	}

	// increase
	//employee, ok := mapIdAndEmployee[id]
	//if !ok {
	//	return 0
	//}
	//totalImportantValue += employee.ImportantValue
	//for _, subID := range employee.SubIDS {
	//	temp := solutionA(subID, employeeInfos)
	//	totalImportantValue += temp
	//}
	return getEmployeeImportantValue(id, mapIdAndEmployee)
}

func getEmployeeImportantValue(id int, mapIdAndEmployee map[int]*EmployeeInfo) int {
	totalImportantValue := 0
	employee, ok := mapIdAndEmployee[id]
	if !ok {
		return 0
	}

	totalImportantValue += employee.ImportantValue
	for _, subID := range employee.SubIDS {
		temp := getEmployeeImportantValue(subID, mapIdAndEmployee)
		totalImportantValue += temp
	}

	return totalImportantValue
}

type EmployeeInfo struct {
	Id             int
	ImportantValue int
	SubIDS         []int
}
