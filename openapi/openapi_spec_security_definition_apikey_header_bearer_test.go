package openapi

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAPIKeyHeaderBearerSecurityDefinition(t *testing.T) {
	Convey("Given a security definition name", t, func() {
		name := "sec_def_name"
		Convey("When newAPIKeyHeaderSecurityDefinition method is called", func() {
			specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition(name)
			Convey("Then the apiKeyHeaderAuthenticator should comply with specAPIKeyAuthenticator interface", func() {
				var _ SpecSecurityDefinition = specAPIKeyHeaderBearerSecurityDefinition
			})
		})
	})
}

func TestAPIKeyHeaderBearerSecurityDefinitionGetName(t *testing.T) {
	Convey("Given an APIKeyHeaderSecurityDefinition", t, func() {
		expectedName := "apikey_name"
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition(expectedName)
		Convey("When getTerraformConfigurationName method is called", func() {
			name := specAPIKeyHeaderBearerSecurityDefinition.getName()
			Convey("Then the result should match the original name", func() {
				So(name, ShouldEqual, expectedName)
			})
		})
	})
}

func TestAPIKeyHeaderBearerSecurityDefinitionGetType(t *testing.T) {
	Convey("Given an APIKeyHeaderSecurityDefinition", t, func() {
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition("apikey_name")
		Convey("When getType method is called", func() {
			secDefType := specAPIKeyHeaderBearerSecurityDefinition.getType()
			Convey("Then the result should be securityDefinitionAPIKey", func() {
				So(secDefType, ShouldEqual, securityDefinitionAPIKey)
			})
		})
	})
}

func TestAPIKeyHeaderBearerSecurityDefinitionGetTerraformConfigurationName(t *testing.T) {
	Convey("Given an APIKeyHeaderSecurityDefinition with a compliant name", t, func() {
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition("apikey_name")
		Convey("When getTerraformConfigurationName method is called", func() {
			secDefTfName := specAPIKeyHeaderBearerSecurityDefinition.getTerraformConfigurationName()
			Convey("Then the result should be securityDefinitionAPIKey", func() {
				So(secDefTfName, ShouldEqual, "apikey_name")
			})
		})
	})

	Convey("Given an APIKeyHeaderSecurityDefinition with a NON compliant name", t, func() {
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition("nonCompliantName")
		Convey("When getTerraformConfigurationName method is called", func() {
			secDefTfName := specAPIKeyHeaderBearerSecurityDefinition.getTerraformConfigurationName()
			Convey("Then the result should be securityDefinitionAPIKey", func() {
				So(secDefTfName, ShouldEqual, "non_compliant_name")
			})
		})
	})
}

func TestAPIKeyHeaderBearerSecurityDefinitionGetAPIKey(t *testing.T) {
	Convey("Given an APIKeyHeaderSecurityDefinition", t, func() {
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition("apikey_name")
		Convey("When getTerraformConfigurationName method is called", func() {
			apiKey := specAPIKeyHeaderBearerSecurityDefinition.getAPIKey()
			Convey("Then the result should contain the right apikey name and location", func() {
				So(apiKey.Name, ShouldEqual, "Authorization")
				So(apiKey.In, ShouldEqual, inHeader)
			})
		})
	})
}

func TestAPIKeyHeaderBearerSecurityDefinitionBuildValue(t *testing.T) {
	Convey("Given an APIKeyHeaderSecurityDefinition", t, func() {
		specAPIKeyHeaderBearerSecurityDefinition := newAPIKeyHeaderBearerSecurityDefinition("apikey_name")
		Convey("When getTerraformConfigurationName method is called", func() {
			value := "jwtToken"
			returnedValue := specAPIKeyHeaderBearerSecurityDefinition.buildValue(value)
			Convey("Then the value should be the expected value with Bearer included", func() {
				So(returnedValue, ShouldEqual, fmt.Sprintf("Bearer %s", value))
			})
		})
	})
}
