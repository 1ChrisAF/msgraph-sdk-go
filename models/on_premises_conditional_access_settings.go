package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesConditionalAccessSettings 
type OnPremisesConditionalAccessSettings struct {
    Entity
    // Indicates if on premises conditional access is enabled for this organization
    enabled *bool
    // User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
    excludedGroups []string
    // User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
    includedGroups []string
    // Override the default access rule when allowing a device to ensure access is granted.
    overrideDefaultRule *bool
}
// NewOnPremisesConditionalAccessSettings instantiates a new onPremisesConditionalAccessSettings and sets the default values.
func NewOnPremisesConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    m := &OnPremisesConditionalAccessSettings{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.onPremisesConditionalAccessSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesConditionalAccessSettings(), nil
}
// GetEnabled gets the enabled property value. Indicates if on premises conditional access is enabled for this organization
func (m *OnPremisesConditionalAccessSettings) GetEnabled()(*bool) {
    return m.enabled
}
// GetExcludedGroups gets the excludedGroups property value. User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
func (m *OnPremisesConditionalAccessSettings) GetExcludedGroups()([]string) {
    return m.excludedGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesConditionalAccessSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["excludedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedGroups(res)
        }
        return nil
    }
    res["includedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludedGroups(res)
        }
        return nil
    }
    res["overrideDefaultRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverrideDefaultRule(val)
        }
        return nil
    }
    return res
}
// GetIncludedGroups gets the includedGroups property value. User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
func (m *OnPremisesConditionalAccessSettings) GetIncludedGroups()([]string) {
    return m.includedGroups
}
// GetOverrideDefaultRule gets the overrideDefaultRule property value. Override the default access rule when allowing a device to ensure access is granted.
func (m *OnPremisesConditionalAccessSettings) GetOverrideDefaultRule()(*bool) {
    return m.overrideDefaultRule
}
// Serialize serializes information the current object
func (m *OnPremisesConditionalAccessSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetExcludedGroups() != nil {
        err = writer.WriteCollectionOfStringValues("excludedGroups", m.GetExcludedGroups())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedGroups() != nil {
        err = writer.WriteCollectionOfStringValues("includedGroups", m.GetIncludedGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("overrideDefaultRule", m.GetOverrideDefaultRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnabled sets the enabled property value. Indicates if on premises conditional access is enabled for this organization
func (m *OnPremisesConditionalAccessSettings) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetExcludedGroups sets the excludedGroups property value. User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy.
func (m *OnPremisesConditionalAccessSettings) SetExcludedGroups(value []string)() {
    m.excludedGroups = value
}
// SetIncludedGroups sets the includedGroups property value. User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access.
func (m *OnPremisesConditionalAccessSettings) SetIncludedGroups(value []string)() {
    m.includedGroups = value
}
// SetOverrideDefaultRule sets the overrideDefaultRule property value. Override the default access rule when allowing a device to ensure access is granted.
func (m *OnPremisesConditionalAccessSettings) SetOverrideDefaultRule(value *bool)() {
    m.overrideDefaultRule = value
}
