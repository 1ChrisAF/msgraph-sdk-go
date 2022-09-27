package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationUserStatus contains properties, inherited properties and actions for an MDM mobile app configuration status for a user.
type ManagedDeviceMobileAppConfigurationUserStatus struct {
    Entity
    // Devices count for that user.
    devicesCount *int32
    // Last modified date time of the policy report.
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *ComplianceStatus
    // User name of the DevicePolicyStatus.
    userDisplayName *string
    // UserPrincipalName.
    userPrincipalName *string
}
// NewManagedDeviceMobileAppConfigurationUserStatus instantiates a new managedDeviceMobileAppConfigurationUserStatus and sets the default values.
func NewManagedDeviceMobileAppConfigurationUserStatus()(*ManagedDeviceMobileAppConfigurationUserStatus) {
    m := &ManagedDeviceMobileAppConfigurationUserStatus{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedDeviceMobileAppConfigurationUserStatus";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceMobileAppConfigurationUserStatus(), nil
}
// GetDevicesCount gets the devicesCount property value. Devices count for that user.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetDevicesCount()(*int32) {
    return m.devicesCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["devicesCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesCount)
    res["lastReportedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastReportedDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceStatus , m.SetStatus)
    res["userDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDisplayName)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetStatus gets the status property value. The status property
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetStatus()(*ComplianceStatus) {
    return m.status
}
// GetUserDisplayName gets the userDisplayName property value. User name of the DevicePolicyStatus.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationUserStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("devicesCount", m.GetDevicesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDevicesCount sets the devicesCount property value. Devices count for that user.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) SetDevicesCount(value *int32)() {
    m.devicesCount = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *ManagedDeviceMobileAppConfigurationUserStatus) SetStatus(value *ComplianceStatus)() {
    m.status = value
}
// SetUserDisplayName sets the userDisplayName property value. User name of the DevicePolicyStatus.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationUserStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
