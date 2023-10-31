//go:build no_runtime_type_checking

package easycdn

// Building without runtime type checking enabled, so all the below just return nil

func validateCdn_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdnParameters(scope constructs.Construct, id *string, props *CdnProps) error {
	return nil
}

