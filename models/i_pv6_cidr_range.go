package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IPv6CidrRange 
type IPv6CidrRange struct {
    IpRange
    // IPv6 address in CIDR notation. Not nullable.
    cidrAddress *string
}
// NewIPv6CidrRange instantiates a new IPv6CidrRange and sets the default values.
func NewIPv6CidrRange()(*IPv6CidrRange) {
    m := &IPv6CidrRange{
        IpRange: *NewIpRange(),
    }
    return m
}
// CreateIPv6CidrRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIPv6CidrRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIPv6CidrRange(), nil
}
// GetCidrAddress gets the cidrAddress property value. IPv6 address in CIDR notation. Not nullable.
func (m *IPv6CidrRange) GetCidrAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cidrAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IPv6CidrRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IpRange.GetFieldDeserializers()
    res["cidrAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCidrAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IPv6CidrRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IpRange.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cidrAddress", m.GetCidrAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCidrAddress sets the cidrAddress property value. IPv6 address in CIDR notation. Not nullable.
func (m *IPv6CidrRange) SetCidrAddress(value *string)() {
    if m != nil {
        m.cidrAddress = value
    }
}
