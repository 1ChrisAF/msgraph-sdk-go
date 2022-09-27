package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallMediaState 
type CallMediaState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The audio media state. Possible values are: active, inactive, unknownFutureValue.
    audio *MediaState
    // The OdataType property
    odataType *string
}
// NewCallMediaState instantiates a new callMediaState and sets the default values.
func NewCallMediaState()(*CallMediaState) {
    m := &CallMediaState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.callMediaState";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCallMediaStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallMediaStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallMediaState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallMediaState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAudio gets the audio property value. The audio media state. Possible values are: active, inactive, unknownFutureValue.
func (m *CallMediaState) GetAudio()(*MediaState) {
    return m.audio
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallMediaState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["audio"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMediaState , m.SetAudio)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CallMediaState) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CallMediaState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAudio() != nil {
        cast := (*m.GetAudio()).String()
        err := writer.WriteStringValue("audio", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *CallMediaState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAudio sets the audio property value. The audio media state. Possible values are: active, inactive, unknownFutureValue.
func (m *CallMediaState) SetAudio(value *MediaState)() {
    m.audio = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CallMediaState) SetOdataType(value *string)() {
    m.odataType = value
}
