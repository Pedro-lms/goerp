// Copyright 2019 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package generate

import (
	"text/template"
)

var poolModelsTemplate = template.Must(template.New("").Parse(`
// This file is autogenerated by erp-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package {{ .ModelsPackageName }}

import (
	"github.com/Pedro-lmso-erp/erp/src/models"
{{- if ne .ModelType "Mixin" }}	
	"github.com/Pedro-lmso-erp/pool/{{ .QueryPackageName }}"
{{- end }}
	"github.com/Pedro-lmso-erp/pool/{{ .ModelsPackageName }}/{{ .SnakeName }}"
    "github.com/Pedro-lmso-erp/pool/{{ .InterfacesPackageName }}"
)

// ------- MODEL ---------

// {{ .Name }}Model is a strongly typed model definition that is used
// to extend the {{ .Name }} model or to get a {{ .Name }}Set through
// its NewSet() function.
//
// To get the unique instance of this type, call {{ .Name }}().
type {{ .Name }}Model struct {
	*models.Model
}

{{ if eq .ModelType "Mixin" }}
// NewSet returns a new {{ .Name }}Set instance wrapping the given model in the given Environment
func (md {{ .Name }}Model) NewSet(env models.Environment, modelName string) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: env.Pool(modelName),
	}
}

{{ else }}
// NewSet returns a new {{ .Name }}Set instance in the given Environment
func (md {{ .Name }}Model) NewSet(env models.Environment) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: env.Pool("{{ .Name }}"),
	}
}

// Create creates a new {{ .Name }} record and returns the newly created
// {{ .Name }}Set instance.
func (md {{ .Name }}Model) Create(env models.Environment, data {{ .InterfacesPackageName }}.{{ .Name }}Data) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new {{ .Name }}Set instance
// with the records found.
func (md {{ .Name }}Model) Search(env models.Environment, cond {{ $.QueryPackageName }}.{{ .Name }}Condition) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md {{ .Name }}Model) Browse(env models.Environment, ids []int64) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md {{ .Name }}Model) BrowseOne(env models.Environment, id int64) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return {{ .SnakeName }}.{{ .Name }}Set{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

{{ end }}

// NewData returns a pointer to a new empty {{ .Name }}Data instance.
//
// Optional field maps if given will be used to populate the data.
func (md {{ .Name }}Model) NewData(fm ...models.FieldMap) {{ .InterfacesPackageName }}.{{ .Name }}Data {
	return &{{ .SnakeName }}.{{ .Name }}Data{
		ModelData: models.NewModelData({{ .Name }}(), fm...),
	}
}

// Fields returns the Field Collection of the {{ .Name }} Model
func (md {{ .Name }}Model) Fields() {{ .SnakeName }}.FieldsCollection {
	return {{ .SnakeName }}.FieldsCollection {
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the {{ .Name }} Model
func (md {{ .Name }}Model) Methods() {{ .SnakeName }}.MethodsCollection {
	return {{ .SnakeName }}.MethodsCollection {
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md {{ .Name }}Model) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = {{ .Name }}Model{}

// Coalesce takes a list of {{ .Name }}Set and return the first non-empty one
// if every record set is empty, it will return the last given
func (md {{ .Name }}Model) Coalesce(lst ...{{ .InterfacesPackageName }}.{{ .Name }}Set) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	var last {{ .InterfacesPackageName }}.{{ .Name }}Set
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// {{ .Name }} returns the unique instance of the {{ .Name }}Model type
// which is used to extend the {{ .Name }} model or to get a {{ .Name }}Set through
// its NewSet() function.
func {{ .Name }}() {{ .Name }}Model {
	return {{ .Name }}Model{
		Model: models.Registry.MustGet("{{ .Name }}"),
	}
}
`))

var poolModelsDirTemplate = template.Must(template.New("").Parse(`
// This file is autogenerated by erp-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package {{ .SnakeName }}

import (
    "github.com/Pedro-lmso-erp/pool/{{ .QueryPackageName }}"
    "github.com/Pedro-lmso-erp/pool/{{ .InterfacesPackageName }}"
{{ range .Deps }} 	"{{ . }}"
{{ end }}
)

// ------- FIELD COLLECTION ----------

// A FieldsCollection is the collection of fields
// of the {{ .Name }} model.
type FieldsCollection struct {
	*models.FieldsCollection
}

{{ range .Fields }}
// {{ .Name }} returns a pointer to the {{ .Name }} Field.
func (c FieldsCollection) {{ .Name }}() *models.Field {
	return c.MustGet("{{ .Name }}")
}
{{ end }}

// ------- METHOD COLLECTION ----------

// A MethodsCollection is the collection of methods
// of the {{ .Name }} model.
type MethodsCollection struct {
	*models.MethodsCollection
}

{{ range .AllMethods }}
// p{{ .Name }} holds the metadata of the {{ $.Name }}.{{ .Name }}() method
type p{{ .Name }} struct {
	*models.Method
}

// Extend adds the given fnct function as a new layer on this method.
func (m p{{ .Name }}) Extend(fnct func({{ $.InterfacesPackageName }}.{{ $.Name }}Set{{ if ne .ParamsTypes "" }}, {{ .ParamsTypes }}{{ end }}) ({{ .ReturnString }})) p{{ .Name }} {
	return p{{ .Name }} {
		Method: m.Method.Extend(fnct),
	}
}

// Underlying returns a pointer to the underlying Method data object.
func (m p{{ .Name }}) Underlying() *models.Method {
	return m.Method
}

var _ models.Methoder = p{{ .Name }}{}

// {{ .Name }} returns a pointer to the {{ .Name }} Method.
func (c MethodsCollection) {{ .Name }}() p{{ .Name }} {
	return p{{ .Name }} {
		Method: c.MustGet("{{ .Name }}"),
	}
}
{{ end }}

// ------- DATA STRUCT ---------

// {{ .Name }}Data is used to hold values of an {{ .Name }} object instance
// when creating or updating a {{ .Name }}Set.
type {{ .Name }}Data struct {
	*models.ModelData
}

// Set sets the given field with the given value.
// If the field already exists, then it is updated with value.
// Otherwise, a new entry is inserted.
//
// It returns the given {{ .Name }}Data so that calls can be chained
func (d {{ .Name }}Data) Set(field models.FieldName, value interface{}) {{ .InterfacesPackageName }}.{{ .Name }}Data {
	return &{{ $.Name }}Data{
		d.ModelData.Set(field, value),
	}
}

// Unset removes the value of the given field if it exists.
//
// It returns the given ModelData so that calls can be chained
func (d {{ .Name }}Data) Unset(field models.FieldName) {{ .InterfacesPackageName }}.{{ .Name }}Data {
	return &{{ $.Name }}Data{
		d.ModelData.Unset(field),
	}
}

// Copy returns a copy of this {{ $.Name }}Data
func (d {{ $.Name }}Data) Copy() {{ .InterfacesPackageName }}.{{ $.Name }}Data {
	return &{{ $.Name }}Data{
		d.ModelData.Copy(),
	}
}

// MergeWith updates this {{ $.Name }}Data with the given other {{ $.Name }}Data
// If a field of the other {{ $.Name }}Data already exists here, the value is overridden,
// otherwise, the field is inserted.
func (d {{ $.Name }}Data) MergeWith(other {{ .InterfacesPackageName }}.{{ $.Name }}Data) {
	d.ModelData.MergeWith(other.Underlying())
}

{{ range .Fields }}
// {{ .Name }} returns the value of the {{ .Name }} field.
// If this {{ .Name }} is not set in this {{ $.Name }}Data, then
// the Go zero value for the type is returned.
func (d {{ $.Name }}Data) {{ .Name }}() {{ .Type }} {
	val := d.ModelData.Get(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"))
{{- if .IsRS }}	
	if !d.Has(models.NewFieldName("{{ .Name }}", "{{ .JSON }}")) || val == nil || val == (*interface{})(nil) {
		val = models.InvalidRecordCollection("{{ .RelModel }}")
	}
	return val.(models.RecordSet).Collection().Wrap().({{ .Type }})
{{- else }}
	if !d.Has(models.NewFieldName("{{ .Name }}", "{{ .JSON }}")) {
		return *new({{ .Type }})
	}
	return val.({{ .Type }})
{{- end }}
}

// Has{{ .Name }} returns true if {{ .Name }} is set in this {{ $.Name }}Data
func (d {{ $.Name }}Data) Has{{ .Name }}() bool {
	return d.ModelData.Has(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"))
}

// Set{{ .Name }} sets the {{ .Name }} field with the given value.
// It returns this {{ $.Name }}Data so that calls can be chained.
func (d {{ $.Name }}Data) Set{{ .Name }}(value {{ .Type }}) {{ $.InterfacesPackageName }}.{{ $.Name }}Data {
	d.ModelData.Set(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"), value)
	return d
}

// Unset{{ .Name }} removes the value of the {{ .Name }} field if it exists.
// It returns this {{ $.Name }}Data so that calls can be chained.
func (d {{ $.Name }}Data) Unset{{ .Name }}() {{ $.InterfacesPackageName }}.{{ $.Name }}Data {
	d.ModelData.Unset(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"))
	return d
}

{{- if .IsRS }}
// Create{{ .Name }} stores the related {{ .RelModel }}Data to be used to create
// a related record on the fly for {{ .Name }}.
//
// This method can be called multiple times to create multiple records
func (d {{ $.Name }}Data) Create{{ .Name }}(related {{ $.InterfacesPackageName }}.{{ .RelModel }}Data) {{ $.InterfacesPackageName }}.{{ $.Name }}Data {
	d.ModelData.Create(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"), related.Underlying())
	return d
}
{{- end }}
{{ end }}

var _ {{ .InterfacesPackageName }}.{{ $.Name }}Data = new({{ .Name }}Data)
var _ {{ .InterfacesPackageName }}.{{ $.Name }}Data = {{ .Name }}Data{}

// ------ AGGREGATE ROW --------

// A {{ .Name }}GroupAggregateRow holds a row of results of a query with a group by clause
// - Values holds the values of the actual query
// - Count is the number of lines aggregated into this one
// - Condition can be used to query the aggregated rows separately if needed
type {{ .Name }}GroupAggregateRow struct {
	values    {{ .InterfacesPackageName }}.{{ .Name }}Data
	count     int
	condition {{ $.QueryPackageName }}.{{ .Name }}Condition
}

// Values returns the values of the actual query
func (a {{ .Name }}GroupAggregateRow) Values() {{ $.InterfacesPackageName }}.{{ $.Name }}Data {
	return a.values 
}

// Count returns the number of lines aggregated into this one
func (a {{ .Name }}GroupAggregateRow) Count() int {
	return a.count
}

// Condition can be used to query the aggregated rows separately if needed
func (a {{ .Name }}GroupAggregateRow) Condition() {{ $.QueryPackageName }}.{{ .Name }}Condition {
	return a.condition
}

// ------- RECORD SET ---------

// {{ .Name }}Set is an autogenerated type to handle {{ .Name }} objects.
type {{ .Name }}Set struct {
	*models.RecordCollection
}

var _ models.RecordSet = {{ .Name }}Set{}

// {{ .Name }}SeterpFunc is a dummy function to uniquely match interfaces.
func (s {{ .Name }}Set) {{ .Name }}SeterpFunc() {}

// IsValid returns true if this RecordSet has been initialized.
func (s {{ .Name }}Set) IsValid() bool {
	if s.RecordCollection == nil {
		return false
	}
	return s.RecordCollection.IsValid()
}

// ForceLoad reloads the cache for the given fields and updates the ids of this {{ .Name }}Set.
//
// If no fields are given, all DB columns of the {{ .Name }} model are retrieved.
//
// It also returns this {{ .Name }}Set.
func (s {{ .Name }}Set) ForceLoad(fields ...models.FieldName) {{ .InterfacesPackageName }}.{{ .Name }}Set {
	s.RecordCollection.ForceLoad(fields...)
	return s
}

// Records returns a slice with all the records of this RecordSet, as singleton
// RecordSets
func (s {{ .Name }}Set) Records() []{{ .InterfacesPackageName }}.{{ .Name }}Set {
	recs := s.RecordCollection.Records()
	res := make([]{{ .InterfacesPackageName }}.{{ .Name }}Set, len(recs))
	for i, rec := range recs {
		res[i] = rec.Wrap("{{ .Name }}").({{ .InterfacesPackageName }}.{{ .Name }}Set)
	}
	return res
}

// CartesianProduct returns the cartesian product of this {{ .Name }}Set with others.
func (s {{ .Name }}Set) CartesianProduct(others ...{{ .InterfacesPackageName }}.{{ .Name }}Set) []{{ .InterfacesPackageName }}.{{ .Name }}Set {
	otherSet := make([]models.RecordSet, len(others))
	for i, o := range others {
		otherSet[i] = o
	}
	recs := s.RecordCollection.CartesianProduct(otherSet...)
	res := make([]{{ .InterfacesPackageName }}.{{ .Name }}Set, len(recs))
	for i, rec := range recs {
		res[i] = rec.Wrap("{{ .Name }}").({{ .InterfacesPackageName }}.{{ .Name }}Set)
	}
	return res
}

// First returns the values of the first Record of the RecordSet as a pointer to a {{ .Name }}Data.
//
// If this RecordSet is empty, it returns an empty {{ .Name }}Data.
func (s {{ .Name }}Set) First() {{ .InterfacesPackageName }}.{{ .Name }}Data {
	return &{{ .Name }}Data {
		s.RecordCollection.First(),
	}
}

// All returns the values of all Records of the RecordCollection as a slice of {{ .Name }}Data pointers.
func (s {{ .Name }}Set) All() []{{ .InterfacesPackageName }}.{{ .Name }}Data {
	allSlice := s.RecordCollection.All()
	res := make([]{{ .InterfacesPackageName }}.{{ .Name }}Data, len(allSlice))
	for i, v := range allSlice {
		res[i] = &{{ .Name }}Data{v}
	}
	return res
}

// Sorted returns a new {{ .Name}}Set sorted according to the given less function.
//
// The less function should return true if rs1 < rs2
func (s {{ .Name}}Set) Sorted(less func(rs1, rs2 {{ .InterfacesPackageName }}.{{ .Name}}Set) bool) {{ .InterfacesPackageName }}.{{ .Name}}Set {
	res := s.RecordCollection.Sorted(func(rc1 models.RecordSet, rc2 models.RecordSet) bool {
		return less({{ .Name }}Set{RecordCollection: rc1.Collection()}, {{ .Name }}Set{RecordCollection: rc2.Collection()})
	})
	return res.Wrap("{{ .Name }}").({{ .InterfacesPackageName }}.{{ .Name}}Set)
}

// Filtered returns a new {{ .Name }}Set with only the elements of this record set
// for which test is true.
//
// Note that if this {{ .Name }}Set is not fully loaded, this function will call the database
// to load the fields before doing the filtering. In this case, it might be more efficient
// to search the database directly with the filter condition.
func (s {{ .Name}}Set) Filtered(test func(rs {{ .InterfacesPackageName }}.{{ .Name}}Set) bool) {{ .InterfacesPackageName }}.{{ .Name}}Set {
	res := s.RecordCollection.Filtered(func(rc models.RecordSet) bool {
		return test({{ .Name }}Set{RecordCollection: rc.Collection()})
	})
	return res.Wrap("{{ .Name }}").({{ .InterfacesPackageName }}.{{ .Name}}Set)
}

{{ range .Fields }}
// {{ .Name }} is a getter for the value of the "{{ .Name }}" field of the first
// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
func (s {{ $.Name }}Set) {{ .Name }}() {{ .Type }} {
{{- if .IsRS }}
	res, _ := s.RecordCollection.Get(models.NewFieldName("{{ .Name }}", "{{ .JSON }}")).(models.RecordSet).Collection().Wrap("{{ .RelModel }}").({{ .Type }})
{{- else }}
	res, _ := s.RecordCollection.Get(models.NewFieldName("{{ .Name }}", "{{ .JSON }}")).({{ .Type }}) 
{{- end }}
	return res 
}

// Set{{ .Name }} is a setter for the value of the "{{ .Name }}" field of this
// RecordSet. All Records of this RecordSet will be updated. Each call to this
// method makes an update query in the database.
//
// Set{{ .Name }} panics if the RecordSet is empty.
func (s {{ $.Name }}Set) Set{{ .Name }}(value {{ .Type }}) {
	s.RecordCollection.Set(models.NewFieldName("{{ .Name }}", "{{ .JSON }}"), value)
}
{{ end }}

// Super returns a RecordSet with a modified callstack so that call to the current
// method will execute the next method layer.
//
// This method is meant to be used inside a method layer function to call its parent,
// such as:
//
//    func (rs h.MyRecordSet) MyMethod() string {
//        res := rs.Super().MyMethod()
//        res += " ok!"
//        return res
//    }
//
// Calls to a different method than the current method will call its next layer only
// if the current method has been called from a layer of the other method. Otherwise,
// it will be the same as calling the other method directly.
func (s {{ .Name }}Set) Super() {{ .InterfacesPackageName }}.{{ .Name }}Set {
	return s.RecordCollection.Super().Wrap("{{ .Name }}").({{ .InterfacesPackageName }}.{{ .Name }}Set)
}

// ModelData returns a new {{ .Name }}Data object populated with the values
// of the given FieldMap. 
func (s {{ .Name }}Set) ModelData(fMap models.FieldMap) {{ .InterfacesPackageName }}.{{ .Name }}Data {
	res := &{{ .Name }}Data{
		models.NewModelData(models.Registry.MustGet("{{ .Name }}")),
	}
	for k, v := range fMap {
		res.Set(models.Registry.MustGet("{{ $.Name }}").FieldName(k), v)
	}
	return res
}

{{ range .Methods }}
{{ .Doc }}
func (s {{ $.Name }}Set) {{ .Name }}({{ .ParamsWithType }}) ({{ .ReturnString }}) {
{{- if eq .Returns "" }}
	s.Collection().Call("{{ .Name }}", {{ .Params}})
{{- else }}
	res := s.Collection().{{ .Call }}("{{ .Name }}", {{ .Params}})
	{{ .ReturnAsserts }}
	return {{ .Returns }}
{{- end }}
}

{{ end }}

{{- if not .IsModelMixin }}
// Aggregates returns the result of this RecordSet query, which must by a grouped query.
func m_{{ $.Name }}_Aggregates(rs {{ .Name }}Set, fieldNames ...models.FieldName) []{{ .InterfacesPackageName }}.{{ .Name }}GroupAggregateRow {
	lines := rs.RecordCollection.Aggregates(fieldNames...)
	res := make([]{{ .InterfacesPackageName }}.{{ .Name }}GroupAggregateRow, len(lines))
	for i, l := range lines {
		res[i] = {{ .Name }}GroupAggregateRow {
			values:    l.Values.Wrap().({{ .InterfacesPackageName }}.{{ .Name }}Data), 
			count:     l.Count,
			condition: {{ $.QueryPackageName }}.{{ .Name }}Condition {
				Condition: l.Condition,
			},
		}
	}
	return res
}
{{- end }}

func init() {
{{- if not .IsModelMixin }}
{{- if eq .ModelType "" }}
	models.CreateModel("{{ .Name }}", 0)
{{- else }}
	models.CreateModel("{{ .Name }}", models.{{ .ModelType }}Model)
{{- end }}
{{- end }}
	models.Registry.MustGet("{{ $.Name }}").AddFields(map[string]models.FieldDefinition{
{{- range .Fields }}
{{- if or .MixinField .EmbedField}}
		"{{ .Name }}": models.DummyField{},
{{- end }}
{{- end }}
	})
{{- range .Methods }}
{{- if .ToDeclare }}
	models.Registry.MustGet("{{ $.Name }}").AddEmptyMethod("{{ .Name }}")
{{- end }}
{{- end }}
{{- if not .IsModelMixin }}
	models.Registry.MustGet("{{ $.Name }}").NewMethod("Aggregates", m_{{ $.Name }}_Aggregates)
{{- end }}
	models.RegisterRecordSetWrapper("{{ .Name }}", {{ .Name }}Set{})
	models.RegisterModelDataWrapper("{{ .Name }}", {{ .Name }}Data{})
}
`))
