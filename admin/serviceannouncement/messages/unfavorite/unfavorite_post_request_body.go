package unfavorite

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnfavoritePostRequestBody provides operations to call the unfavorite method.
type UnfavoritePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The messageIds property
    messageIds []string
}
// NewUnfavoritePostRequestBody instantiates a new unfavoritePostRequestBody and sets the default values.
func NewUnfavoritePostRequestBody()(*UnfavoritePostRequestBody) {
    m := &UnfavoritePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnfavoritePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnfavoritePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnfavoritePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnfavoritePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnfavoritePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["messageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMessageIds(res)
        }
        return nil
    }
    return res
}
// GetMessageIds gets the messageIds property value. The messageIds property
func (m *UnfavoritePostRequestBody) GetMessageIds()([]string) {
    return m.messageIds
}
// Serialize serializes information the current object
func (m *UnfavoritePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMessageIds() != nil {
        err := writer.WriteCollectionOfStringValues("messageIds", m.GetMessageIds())
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
func (m *UnfavoritePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMessageIds sets the messageIds property value. The messageIds property
func (m *UnfavoritePostRequestBody) SetMessageIds(value []string)() {
    m.messageIds = value
}
