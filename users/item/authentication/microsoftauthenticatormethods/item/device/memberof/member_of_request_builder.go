package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i26266f5b76fe466d8d8cb25ad164ed593675efeee90a3c7e7d954d6794e67d6d "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/user"
    i6fe77b8e51c8d56cfc93b100b8a445739adb04fe6b4ff66c6cb91f67925f5bf1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/count"
    i7b2491adf31e27cb3e7783667e34599a82491adc658d5985cf305eb7b07261f6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/group"
    i833178419ceb62fd01f6dfdb8f27956aeb46cb397bdc95a0fd301c08dd902d03 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/serviceprincipal"
    ic849c1bfde6b7b0360192353a77d0052163a319c48c86d95443f478153ebecb1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/application"
    icee1bc4f1ea0cd6113ed9449d530b660b6793161e973dd8647ddd094ddfeaa53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/orgcontact"
    ie32a3b49692cd69d8a7edaeffa4d1dc1cb51464e17113621556cbadebc47342c "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/device"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MemberOfRequestBuilderGetQueryParameters
}
// Application the application property
func (m *MemberOfRequestBuilder) Application()(*ic849c1bfde6b7b0360192353a77d0052163a319c48c86d95443f478153ebecb1.ApplicationRequestBuilder) {
    return ic849c1bfde6b7b0360192353a77d0052163a319c48c86d95443f478153ebecb1.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *MemberOfRequestBuilder) Count()(*i6fe77b8e51c8d56cfc93b100b8a445739adb04fe6b4ff66c6cb91f67925f5bf1.CountRequestBuilder) {
    return i6fe77b8e51c8d56cfc93b100b8a445739adb04fe6b4ff66c6cb91f67925f5bf1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *MemberOfRequestBuilder) Device()(*ie32a3b49692cd69d8a7edaeffa4d1dc1cb51464e17113621556cbadebc47342c.DeviceRequestBuilder) {
    return ie32a3b49692cd69d8a7edaeffa4d1dc1cb51464e17113621556cbadebc47342c.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *MemberOfRequestBuilder) Group()(*i7b2491adf31e27cb3e7783667e34599a82491adc658d5985cf305eb7b07261f6.GroupRequestBuilder) {
    return i7b2491adf31e27cb3e7783667e34599a82491adc658d5985cf305eb7b07261f6.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*icee1bc4f1ea0cd6113ed9449d530b660b6793161e973dd8647ddd094ddfeaa53.OrgContactRequestBuilder) {
    return icee1bc4f1ea0cd6113ed9449d530b660b6793161e973dd8647ddd094ddfeaa53.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i833178419ceb62fd01f6dfdb8f27956aeb46cb397bdc95a0fd301c08dd902d03.ServicePrincipalRequestBuilder) {
    return i833178419ceb62fd01f6dfdb8f27956aeb46cb397bdc95a0fd301c08dd902d03.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i26266f5b76fe466d8d8cb25ad164ed593675efeee90a3c7e7d954d6794e67d6d.UserRequestBuilder) {
    return i26266f5b76fe466d8d8cb25ad164ed593675efeee90a3c7e7d954d6794e67d6d.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
