package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecentNotebookLinks 
type RecentNotebookLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Opens the notebook in the OneNote native client if it's installed.
    oneNoteClientUrl ExternalLinkable
    // Opens the notebook in OneNote on the web.
    oneNoteWebUrl ExternalLinkable
}
// NewRecentNotebookLinks instantiates a new recentNotebookLinks and sets the default values.
func NewRecentNotebookLinks()(*RecentNotebookLinks) {
    m := &RecentNotebookLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.recentNotebookLinks";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRecentNotebookLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecentNotebookLinksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecentNotebookLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecentNotebookLinks) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecentNotebookLinks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["oneNoteClientUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteClientUrl(val.(ExternalLinkable))
        }
        return nil
    }
    res["oneNoteWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteWebUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RecentNotebookLinks) GetOdataType()(*string) {
    return m.odataType
}
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *RecentNotebookLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    return m.oneNoteClientUrl
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *RecentNotebookLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    return m.oneNoteWebUrl
}
// Serialize serializes information the current object
func (m *RecentNotebookLinks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteClientUrl", m.GetOneNoteClientUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteWebUrl", m.GetOneNoteWebUrl())
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
func (m *RecentNotebookLinks) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RecentNotebookLinks) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *RecentNotebookLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    m.oneNoteClientUrl = value
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *RecentNotebookLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    m.oneNoteWebUrl = value
}
