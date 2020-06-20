package organisation

import (
	"testing"
)

func TestCEO(t *testing.T) {
	ceo := Employee{
		id:   "1",
		name: "Bill",
		emp:  nil,
	}
	allEmp := setupEmp()
	got := commonManager(allEmp, ceo, ceo)

	if got.id != ceo.id {
		t.Errorf("TestCEO() = %q, want %q", got.name, ceo.name)
	}
}

func TestSecondLevel(t *testing.T) {

	mark := Employee{id: "2", name: "Mark", emp: nil}
	steve := Employee{id: "3", name: "Steve", emp: nil}

	ceo := setupEmp()
	got := commonManager(ceo, mark, steve)

	if got.id != ceo.id {
		t.Errorf("TestSecondLevel() = %q, want %q", got.name, ceo.name)
	}
}

func TestDiferentLevel(t *testing.T) {

	jose := Employee{id: "4", name: "Jose", emp: nil}
	steve := Employee{id: "3", name: "Steve", emp: nil}

	ceo := setupEmp()
	got := commonManager(ceo, jose, steve)

	if got.id != ceo.id {
		t.Errorf("TestDiferentLevel() = %q, want %q", got.name, ceo.name)
	}
}

func TestSameBranch(t *testing.T) {

	pablo := Employee{id: "5", name: "Pablo", emp: nil}
	pedro := Employee{id: "6", name: "Pedro", emp: nil}

	mark := Employee{id: "2", name: "Mark", emp: nil}

	ceo := setupEmp()
	got := commonManager(ceo, pablo, pedro)

	if got.id != mark.id {
		t.Errorf("TestSameBranch() = %q, want %q", got.name, mark.name)
	}
}

func TestLastLevelDiffBranch(t *testing.T) {

	jacob := Employee{id: "8", name: "Jacob", emp: nil}
	pedro := Employee{id: "6", name: "Pedro", emp: nil}

	ceo := setupEmp()
	got := commonManager(ceo, pedro, jacob)

	if got.id != ceo.id {
		t.Errorf("TestSameBranch() = %q, want %q", got.name, ceo.name)
	}
}

func setupEmp() Employee {
	//level 4
	pedro := Employee{id: "6", name: "Pedro", emp: nil}
	jacob := Employee{id: "8", name: "Jacob", emp: nil}

	//level 3
	jose := Employee{id: "4", name: "Jose", emp: nil}
	pablo := Employee{id: "5", name: "Pablo", emp: nil}
	pablo.emp = append(pablo.emp, &pedro)
	juan := Employee{id: "7", name: "Juan", emp: nil}
	juan.emp = append(juan.emp, &jacob)

	//level 2
	mark := Employee{id: "2", name: "Mark", emp: nil}
	mark.emp = append(mark.emp, &jose, &pablo)
	steve := Employee{id: "3", name: "Steve", emp: nil}
	steve.emp = append(steve.emp, &juan)

	//level 1
	ceo := Employee{id: "1", name: "Bill", emp: nil}
	ceo.emp = append(ceo.emp, &mark, &steve)

	return ceo
}
