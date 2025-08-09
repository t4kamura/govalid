package rules

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/codegen"

	"github.com/sivchari/govalid/internal/validator"
)

type alphaValidator struct {
	pass       *codegen.Pass
	field      *ast.Field
	structName string
	ruleName   string
	parentPath string
}

var _ validator.Validator = (*alphaValidator)(nil)

const alphaKey = "%s-alpha"

func (v *alphaValidator) Validate() string {
	// Use external helper function for better maintainability
	return fmt.Sprintf(`!validationhelper.IsValidAlpha(t.%s)`, v.FieldName())
}

func (v *alphaValidator) FieldName() string {
	return v.field.Names[0].Name
}

func (v *alphaValidator) FieldPath() validator.FieldPath {
	if v.parentPath == "" {
		return validator.NewFieldPath(v.structName, v.FieldName())
	}
	return validator.NewFieldPath(v.structName, v.parentPath, v.FieldName())
}

func (v *alphaValidator) Err() string {
	key := fmt.Sprintf(alphaKey, v.FieldPath().WithoutDots())
	
	if validator.GeneratorMemory[key] {
		return ""
	}

	validator.GeneratorMemory[key] = true

	const errTemplate = `
		// [@ERRVARIABLE] is the error returned when field [@FIELD] is not alphabetic.
		[@ERRVARIABLE] = govaliderrors.ValidationError{Reason:"field [@FIELD] must be alphabetic",Path:"[@PATH]",Type:"[@TYPE]"}
	`

	replacer := strings.NewReplacer(
		"[@ERRVARIABLE]", v.ErrVariable(),
		"[@FIELD]", v.FieldName(),
		"[@PATH]", v.FieldPath().String(),
		"[@TYPE]", v.ruleName,
	)

	return replacer.Replace(errTemplate)
}

func (v *alphaValidator) ErrVariable() string {
	return strings.ReplaceAll("Err[@PATH]AlphaValidation", "[@PATH]", v.FieldPath().WithoutDots())
}

func (v *alphaValidator) Imports() []string {
	// Import validation helper package
	return []string{"github.com/sivchari/govalid/validation/validationhelper"}
}

// ValidateAlpha creates a new alphaValidator for string types.
func ValidateAlpha(pass *codegen.Pass, field *ast.Field, _ map[string]string, structName, ruleName string, parentPath string) validator.Validator {
	fieldName := field.Names[0].Name
	var fieldPath validator.FieldPath
	if parentPath != "" {
		fieldPath = validator.NewFieldPath(structName, parentPath, fieldName)
	} else {
		fieldPath = validator.NewFieldPath(structName, fieldName)
	}
	validator.GeneratorMemory[fmt.Sprintf(alphaKey, fieldPath.WithoutDots())] = false
	
	typ := pass.TypesInfo.TypeOf(field.Type)

	// Check if it's a string type
	basic, ok := typ.Underlying().(*types.Basic)

	if !ok || basic.Kind() != types.String {
		return nil
	}

	return &alphaValidator{
		pass:       pass,
		field:      field,
		structName: structName,
		ruleName:   ruleName,
		parentPath: parentPath,
	}
}
