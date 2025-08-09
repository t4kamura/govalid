package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type urlValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*urlValidator)(nil)

const urlKey = "%s-url"

func (u *urlValidator) Validate() string {
	fieldName := u.FieldName()
	// Use external helper function for better maintainability
	return fmt.Sprintf("!validationhelper.IsValidURL(t.%s)", fieldName)
}

func (u *urlValidator) FieldName() string {
	return u.field.Names[0].Name
}

func (u *urlValidator) FieldPath() validator.FieldPath {
	return validator.NewFieldPath(u.structName, u.parentPath, u.FieldName())
}

func (u *urlValidator) Err() string {
	// No need to generate inline function - using external helper
	key := fmt.Sprintf(urlKey, u.structName+u.FieldPath().WithoutDots())
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when the field is not a valid URL.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be a valid URL",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", u.ErrVariable(),
		"[@FIELD]", u.FieldName(),
		"[@PATH]", u.FieldPath().String(),
		"[@TYPE]", u.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (u *urlValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]URLValidation", `[@PATH]`, u.FieldPath().WithoutDots())
}

func (u *urlValidator) Imports() []string {
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateURL creates a new urlValidator for string types.
func ValidateURL(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName, ruleName, parentPath string) validator.Validator {
	fieldPath := validator.NewFieldPath(structName, parentPath, field.Names[0].Name)
	validator.GeneratorMemory[fmt.Sprintf(urlKey, structName+fieldPath.WithoutDots())] = false

	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)
	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &urlValidator{
		pass:       pass,
		field:      field,
		structName: structName,
		ruleName:   ruleName,
		parentPath: parentPath,
	}
}
