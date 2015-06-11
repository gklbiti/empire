package empire

import (
	"reflect"
	"testing"
)

func TestConfigsQuery(t *testing.T) {
	id := "1234"
	app := &App{ID: "4321"}

	tests := scopeTests{
		{ConfigsQuery{}, "", []interface{}{}},
		{ConfigsQuery{ID: &id}, "WHERE (id = $1)", []interface{}{id}},
		{ConfigsQuery{App: app}, "WHERE (app_id = $1)", []interface{}{app.ID}},
	}

	tests.Run(t)
}

func TestMergeVars(t *testing.T) {
	var (
		PRODUCTION   = "production"
		STAGING      = "staging"
		EMPTY        = ""
		DATABASE_URL = "postgres://localhost"
	)

	// Old vars
	vars := Vars{
		"RAILS_ENV":    &PRODUCTION,
		"DATABASE_URL": &DATABASE_URL,
	}

	tests := []struct {
		in  Vars
		out Vars
	}{
		// Removing a variable
		{
			Vars{
				"RAILS_ENV": nil,
			},
			Vars{
				"DATABASE_URL": &DATABASE_URL,
			},
		},

		// Setting an empty variable
		{
			Vars{
				"RAILS_ENV": &EMPTY,
			},
			Vars{
				"RAILS_ENV":    &EMPTY,
				"DATABASE_URL": &DATABASE_URL,
			},
		},

		// Updating a variable
		{
			Vars{
				"RAILS_ENV": &STAGING,
			},
			Vars{
				"RAILS_ENV":    &STAGING,
				"DATABASE_URL": &DATABASE_URL,
			},
		},
	}

	for _, tt := range tests {
		v := mergeVars(vars, tt.in)

		if got, want := v, tt.out; !reflect.DeepEqual(got, want) {
			t.Errorf("mergeVars => want %v; got %v", want, got)
		}
	}
}
