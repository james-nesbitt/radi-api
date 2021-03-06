package security

import (
	log "github.com/Sirupsen/logrus"

	api_operation "github.com/wunderkraut/radi-api/operation"
	api_property "github.com/wunderkraut/radi-api/property"
	api_usage "github.com/wunderkraut/radi-api/usage"
)

/**
 * Security specifc properties
 */

const (
	SECURITY_AUTHENTICATION_SUCCEEDED_PROPERTY_KEY = "security.authentication.success"
	SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY  = "security.authorization.success"
	SECURITY_AUTHORIZATION_OPERATION_PROPERTY_KEY  = "security.authorization.operation"
	SECURITY_AUTHORIZATION_RULERESULT_PROPERTY_KEY = "security.authorization.ruleresult"
	SECURITY_USER_PROPERTY_KEY                     = "security.user"
)

// A boolean property for if an authentication succeeded
type SecurityAuthenticationSucceededProperty struct {
	api_property.BooleanProperty
}

// ID returns string unique property Identifier
func (authSuccessProp *SecurityAuthenticationSucceededProperty) Id() string {
	return SECURITY_AUTHENTICATION_SUCCEEDED_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (authSuccessProp *SecurityAuthenticationSucceededProperty) Label() string {
	return "Authentication success"
}

// Description provides a longer multi-line string description of what the property does
func (authSuccessProp *SecurityAuthenticationSucceededProperty) Description() string {
	return "Indicated whether or not an authentication operation has succeeded"
}

// Mark a property as being for internal use only (no shown to users)
func (authSuccessProp *SecurityAuthenticationSucceededProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Copy the property
func (authSuccessProp *SecurityAuthenticationSucceededProperty) Copy() api_property.Property {
	prop := &SecurityAuthenticationSucceededProperty{}
	prop.Set(authSuccessProp.Get())
	return api_property.Property(prop)
}

// A boolean property for if an authorization succeeded (was granted)
type SecurityAuthorizationSucceededProperty struct {
	api_property.BooleanProperty
}

// ID returns string unique property Identifier
func (authSuccessProp *SecurityAuthorizationSucceededProperty) Id() string {
	return SECURITY_AUTHORIZATION_SUCCEEDED_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (authSuccessProp *SecurityAuthorizationSucceededProperty) Label() string {
	return "Authorization success"
}

// Description provides a longer multi-line string description of what the property does
func (authSuccessProp *SecurityAuthorizationSucceededProperty) Description() string {
	return "Indicated whether or not an autheorization operation has succeeded"
}

// Mark a property as being for internal use only (no shown to users)
func (authSuccessProp *SecurityAuthorizationSucceededProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Copy the property
func (authSuccessProp *SecurityAuthorizationSucceededProperty) Copy() api_property.Property {
	prop := &SecurityAuthorizationSucceededProperty{}
	prop.Set(authSuccessProp.Get())
	return api_property.Property(prop)
}

// An Operation property
type SecurityAuthorizationOperationProperty struct {
	api_operation.OperationProperty
}

// ID returns string unique property Identifier
func (operationProp *SecurityAuthorizationOperationProperty) Id() string {
	return SECURITY_AUTHORIZATION_OPERATION_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (operationProp *SecurityAuthorizationOperationProperty) Label() string {
	return "Operation to authorize"
}

// Description provides a longer multi-line string description of what the property does
func (operationProp *SecurityAuthorizationOperationProperty) Description() string {
	return "Operation which needs to be authorized"
}

// Mark a property as being for internal use only (no shown to users)
func (operationProp *SecurityAuthorizationOperationProperty) Usage() api_usage.Usage {
	return api_property.Usage_Internal()
}

// Copy the property
func (operationProp *SecurityAuthorizationOperationProperty) Copy() api_property.Property {
	prop := &SecurityAuthorizationOperationProperty{}
	prop.Set(operationProp.Get())
	return api_property.Property(prop)
}

// An Operation property
type SecurityAuthorizationRuleResultProperty struct {
	value RuleResult
}

// ID returns string unique property Identifier
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Id() string {
	return SECURITY_AUTHORIZATION_RULERESULT_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Label() string {
	return "Authorization rule result"
}

// Description provides a longer multi-line string description of what the property does
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Description() string {
	return "Operation authorization rule result"
}

// Mark a property as being for internal use only (no shown to users)
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Usage() api_usage.Usage {
	return api_property.Usage_Internal()
}

// Give an idea of what type of value the property consumes
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Type() string {
	return "github.com/wunderkraut/radi-api/operation/security.RuleResult"
}

// Retrieve the context, or retrieve a Background context by default
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Get() interface{} {
	return interface{}(ruleresultProp.value)
}
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Set(value interface{}) bool {
	if converted, ok := value.(RuleResult); ok {
		ruleresultProp.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected a RuleResult")
		return false
	}
}

// Copy the property
func (ruleresultProp *SecurityAuthorizationRuleResultProperty) Copy() api_property.Property {
	prop := &SecurityAuthorizationRuleResultProperty{}
	prop.Set(ruleresultProp.Get())
	return api_property.Property(prop)
}

// An Operation property
type SecurityUserProperty struct {
	value SecurityUser
}

// ID returns string unique property Identifier
func (userProp *SecurityUserProperty) Id() string {
	return SECURITY_USER_PROPERTY_KEY
}

// Label returns a short user readable label for the property
func (userProp *SecurityUserProperty) Label() string {
	return "User"
}

// Description provides a longer multi-line string description of what the property does
func (userProp *SecurityUserProperty) Description() string {
	return "User account information"
}

// Mark a property as being for internal use only (no shown to users)
func (userProp *SecurityUserProperty) Usage() api_usage.Usage {
	return api_property.Usage_ReadOnly()
}

// Give an idea of what type of value the property consumes
func (userProp *SecurityUserProperty) Type() string {
	return "github.com/wunderkraut/radi-api/operation/security.SecurityUser"
}

// Retrieve the context, or retrieve a Background context by default
func (userProp *SecurityUserProperty) Get() interface{} {
	return interface{}(userProp.value)
}
func (userProp *SecurityUserProperty) Set(value interface{}) bool {
	if converted, ok := value.(SecurityUser); ok {
		userProp.value = converted
		return true
	} else {
		log.WithFields(log.Fields{"value": value}).Error("Could not assign Property value, because the passed parameter was the wrong type. Expected a SecurityUser")
		return false
	}
}

// Copy the property
func (userProp *SecurityUserProperty) Copy() api_property.Property {
	prop := &SecurityUserProperty{}
	prop.Set(userProp.Get())
	return api_property.Property(prop)
}
