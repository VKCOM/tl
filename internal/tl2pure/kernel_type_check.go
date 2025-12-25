package tl2pure

import (
	"fmt"

	"github.com/vkcom/tl/internal/tlast"
)

func (k *Kernel) typeCheck(tip tlast.TL2TypeDefinition, leftArguments []tlast.TL2TypeTemplate) error {
	if tip.IsUnionType {
		for _, v := range tip.UnionType.Variants {
			if err := k.typeCheckAliasFields(v.IsTypeAlias, v.TypeAlias, v.Fields, leftArguments); err != nil {
				return err
			}
		}
		return nil
	}
	if tip.IsAlias() {
		return k.typeCheckTypeRef(tip.TypeAlias, leftArguments)
	}
	return k.typeCheckAliasFields(false, tlast.TL2TypeRef{},
		tip.ConstructorFields, leftArguments)
}

func (k *Kernel) typeCheckAliasFields(isTypeAlias bool, typeAlias tlast.TL2TypeRef,
	fields []tlast.TL2Field, leftArguments []tlast.TL2TypeTemplate) error {
	if isTypeAlias {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: typeAlias}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("type alias cannot be number")
		}
		return nil
	}
	for _, f := range fields {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: f.Type}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("field type %s cannot be number", f.Name)
		}
	}
	return nil
}

func (k *Kernel) typeCheckTypeRef(tr tlast.TL2TypeRef, leftArguments []tlast.TL2TypeTemplate) error {
	if tr.IsBracket {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: tr.BracketType.ArrayType}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("bracket element type cannot be number")
		}
		if tr.BracketType.HasIndex {
			_, err = k.typeCheckArgument(tr.BracketType.IndexType, leftArguments)
			if err != nil {
				return err
			}
			// can be both type (map) and number (tuple)
		}
		return nil
	}
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name]
	if !ok {
		return fmt.Errorf("type %s does not exist", someType.Name)
	}
	if kt.comb.IsFunction {
		return fmt.Errorf("cannot reference function %s", someType.Name)
	}
	if len(someType.Arguments) != len(kt.comb.TypeDecl.TemplateArguments) {
		return fmt.Errorf("typeref %s must have %d template arguments, has %d", someType.String(), len(kt.comb.TypeDecl.TemplateArguments), len(someType.Arguments))
	}
	for i, targ := range kt.comb.TypeDecl.TemplateArguments {
		cat, err := k.typeCheckArgument(someType.Arguments[i], leftArguments)
		if err != nil {
			return err
		}
		if targ.Category != cat {
			return fmt.Errorf("typeref %s argument %s wrong category, must be %s", someType.String(), targ.Name, targ.Category)
		}
	}
	return nil
}

func (k *Kernel) typeCheckArgument(arg tlast.TL2TypeArgument, leftArguments []tlast.TL2TypeTemplate) (tlast.TL2TypeCategory, error) {
	if arg.IsNumber {
		return tlast.TL2TypeCategoryNat, nil
	}
	if !arg.Type.IsBracket && arg.Type.SomeType.Name.Namespace == "" {
		for _, la := range leftArguments {
			if arg.Type.SomeType.Name.Name == la.Name {
				if len(arg.Type.SomeType.Arguments) != 0 {
					return "", fmt.Errorf("reference to template argument %s cannot have arguments", la.Name)
				}
				return la.Category, nil
			}
		}
		// reference to global type
	}
	if err := k.typeCheckTypeRef(arg.Type, leftArguments); err != nil {
		return "", err
	}
	return tlast.TL2TypeCategoryType, nil
}
