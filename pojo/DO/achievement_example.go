package DO

import (
	"reflect"
)

type Criterion struct {
	Condition    string
	Value        any
	SecondValue  any
	NoValue      bool
	SingleValue  bool
	BetweenValue bool
	ListValue    bool
	TypeHandler  string
}

func NewCriterion1(condition string) *Criterion {
	return &Criterion{
		Condition: condition,
		NoValue:   true,
	}
}

func NewCriterion2(condition, typeHandler string, value any) *Criterion {

	var listValue bool
	kind := reflect.TypeOf(value).Kind()
	//只校验slice
	if kind == reflect.Slice {
		listValue = true
	}

	return &Criterion{
		Condition:   condition,
		Value:       value,
		TypeHandler: typeHandler,
		ListValue:   listValue,
		SingleValue: !listValue,
	}
}

func NewCriterion3(condition string, value any) *Criterion {
	return NewCriterion2(condition, "", value)
}

func NewCriterion4(condition, typeHandler string, value, secondValue any) *Criterion {
	return &Criterion{
		Condition:    condition,
		Value:        value,
		SecondValue:  secondValue,
		TypeHandler:  typeHandler,
		BetweenValue: true,
	}
}

func NewCriterion5(condition string, value, secondValue any) *Criterion {
	return NewCriterion4(condition, "", value, secondValue)
}

type GeneratedCriteria struct {
	Criteria []*Criterion
}

func (g *GeneratedCriteria) IsValid() bool {
	return len(g.Criteria) > 0
}

func (g *GeneratedCriteria) AddCriterion(condition string) {
	if condition == "" {
		panic("condition can't be empty")
	}
	g.Criteria = append(g.Criteria, &Criterion{Condition: condition})
}

func (g *GeneratedCriteria) AddCriterionWithValue(condition, property string, value any) {
	if value == nil {
		panic("Between values for " + property + " cannot be null")
	}
	g.Criteria = append(g.Criteria, &Criterion{Condition: condition, Value: value})
}

func (g *GeneratedCriteria) AddCriterionWithValue2(condition, property string, value1, value2 any) {
	if value1 == nil || value2 == nil {
		panic("Between values for " + property + " cannot be null")
	}
	g.Criteria = append(g.Criteria, &Criterion{Condition: condition, Value: value1, SecondValue: value2})
}

func (g *GeneratedCriteria) AndUserIdEqualTo(value string) *GeneratedCriteria {
	g.AddCriterionWithValue("user_id =", "userId", value)
	return g
}

type Criteria struct {
	GeneratedCriteria
}

type AchievementExample struct {
	OrderByClause string
	Distinct      bool
	OredCriteria  []*Criteria
}

//func NewAchievementExample()  {
//
//}

func (a *AchievementExample) OrCriteria(criteria *Criteria) {
	a.OredCriteria = append(a.OredCriteria, criteria)
}

func (a *AchievementExample) Or() *Criteria {
	criteria := createCriteriaInternal()
	a.OredCriteria = append(a.OredCriteria, criteria)
	return criteria
}

func (a *AchievementExample) CreateCriteria() *Criteria {
	criteria := createCriteriaInternal()
	if a.OredCriteria == nil {
		a.OredCriteria = append(a.OredCriteria, criteria)
	}
	return criteria
}

func createCriteriaInternal() *Criteria {
	return &Criteria{}
}

func (a *AchievementExample) Clear() {
	a.OredCriteria = nil
	a.OrderByClause = ""
	a.Distinct = false
}
