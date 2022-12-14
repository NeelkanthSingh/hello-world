package models

type TemplateData struct {
	StringMap  map[string]string
	IntegerMap map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	error      string
}
