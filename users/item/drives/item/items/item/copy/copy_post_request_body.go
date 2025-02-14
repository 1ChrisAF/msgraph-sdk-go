package copy

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CopyPostRequestBody provides operations to call the copy method.
type CopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name property
    name *string
    // The parentReference property
    parentReference iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable
}
// NewCopyPostRequestBody instantiates a new copyPostRequestBody and sets the default values.
func NewCopyPostRequestBody()(*CopyPostRequestBody) {
    m := &CopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parentReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *CopyPostRequestBody) GetName()(*string) {
    return m.name
}
// GetParentReference gets the parentReference property value. The parentReference property
func (m *CopyPostRequestBody) GetParentReference()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable) {
    return m.parentReference
}
// Serialize serializes information the current object
func (m *CopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentReference", m.GetParentReference())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *CopyPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentReference sets the parentReference property value. The parentReference property
func (m *CopyPostRequestBody) SetParentReference(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable)() {
    m.parentReference = value
}
