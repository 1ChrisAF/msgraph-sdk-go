package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharepointIds 
type SharepointIds struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier (guid) for the item's list in SharePoint.
    listId *string
    // An integer identifier for the item within the containing list.
    listItemId *string
    // The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
    listItemUniqueId *string
    // The OdataType property
    odataType *string
    // The unique identifier (guid) for the item's site collection (SPSite).
    siteId *string
    // The SharePoint URL for the site that contains the item.
    siteUrl *string
    // The unique identifier (guid) for the tenancy.
    tenantId *string
    // The unique identifier (guid) for the item's site (SPWeb).
    webId *string
}
// NewSharepointIds instantiates a new sharepointIds and sets the default values.
func NewSharepointIds()(*SharepointIds) {
    m := &SharepointIds{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.sharepointIds";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSharepointIdsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharepointIdsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharepointIds(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharepointIds) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharepointIds) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["listId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetListId)
    res["listItemId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetListItemId)
    res["listItemUniqueId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetListItemUniqueId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["siteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteId)
    res["siteUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteUrl)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["webId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebId)
    return res
}
// GetListId gets the listId property value. The unique identifier (guid) for the item's list in SharePoint.
func (m *SharepointIds) GetListId()(*string) {
    return m.listId
}
// GetListItemId gets the listItemId property value. An integer identifier for the item within the containing list.
func (m *SharepointIds) GetListItemId()(*string) {
    return m.listItemId
}
// GetListItemUniqueId gets the listItemUniqueId property value. The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
func (m *SharepointIds) GetListItemUniqueId()(*string) {
    return m.listItemUniqueId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SharepointIds) GetOdataType()(*string) {
    return m.odataType
}
// GetSiteId gets the siteId property value. The unique identifier (guid) for the item's site collection (SPSite).
func (m *SharepointIds) GetSiteId()(*string) {
    return m.siteId
}
// GetSiteUrl gets the siteUrl property value. The SharePoint URL for the site that contains the item.
func (m *SharepointIds) GetSiteUrl()(*string) {
    return m.siteUrl
}
// GetTenantId gets the tenantId property value. The unique identifier (guid) for the tenancy.
func (m *SharepointIds) GetTenantId()(*string) {
    return m.tenantId
}
// GetWebId gets the webId property value. The unique identifier (guid) for the item's site (SPWeb).
func (m *SharepointIds) GetWebId()(*string) {
    return m.webId
}
// Serialize serializes information the current object
func (m *SharepointIds) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("listId", m.GetListId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listItemId", m.GetListItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listItemUniqueId", m.GetListItemUniqueId())
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
        err := writer.WriteStringValue("siteId", m.GetSiteId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteUrl", m.GetSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webId", m.GetWebId())
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
func (m *SharepointIds) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetListId sets the listId property value. The unique identifier (guid) for the item's list in SharePoint.
func (m *SharepointIds) SetListId(value *string)() {
    m.listId = value
}
// SetListItemId sets the listItemId property value. An integer identifier for the item within the containing list.
func (m *SharepointIds) SetListItemId(value *string)() {
    m.listItemId = value
}
// SetListItemUniqueId sets the listItemUniqueId property value. The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
func (m *SharepointIds) SetListItemUniqueId(value *string)() {
    m.listItemUniqueId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharepointIds) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSiteId sets the siteId property value. The unique identifier (guid) for the item's site collection (SPSite).
func (m *SharepointIds) SetSiteId(value *string)() {
    m.siteId = value
}
// SetSiteUrl sets the siteUrl property value. The SharePoint URL for the site that contains the item.
func (m *SharepointIds) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// SetTenantId sets the tenantId property value. The unique identifier (guid) for the tenancy.
func (m *SharepointIds) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetWebId sets the webId property value. The unique identifier (guid) for the item's site (SPWeb).
func (m *SharepointIds) SetWebId(value *string)() {
    m.webId = value
}
