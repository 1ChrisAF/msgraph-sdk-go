package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0b63cd45a82a75012712ec218f245ffaa3541c1be89762c5b12114393c6e09be "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets/item/tasks/item/assignedtotaskboardformat"
    i0cc1127e54624c213444e0ed912e9f6e573eca5b01c61e0255b0c2de1a231be1 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets/item/tasks/item/progresstaskboardformat"
    ia37f1dc3e53c9a3902fa069b5b6cb2af8ebaaf76c14f6643a9c371ecf02f1414 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets/item/tasks/item/details"
    if38a9a944cfdc240dddf08af402c99d988e7ab50367452b22964ead342475b92 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets/item/tasks/item/buckettaskboardformat"
)

// Builds and executes requests for operations under \me\planner\plans\{plannerPlan-id}\buckets\{plannerBucket-id}\tasks\{plannerTask-id}
type PlannerTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PlannerTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PlannerTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only. Nullable. The collection of tasks in the bucket.
type PlannerTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PlannerTaskRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat()(*i0b63cd45a82a75012712ec218f245ffaa3541c1be89762c5b12114393c6e09be.AssignedToTaskBoardFormatRequestBuilder) {
    return i0b63cd45a82a75012712ec218f245ffaa3541c1be89762c5b12114393c6e09be.NewAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerTaskRequestBuilder) BucketTaskBoardFormat()(*if38a9a944cfdc240dddf08af402c99d988e7ab50367452b22964ead342475b92.BucketTaskBoardFormatRequestBuilder) {
    return if38a9a944cfdc240dddf08af402c99d988e7ab50367452b22964ead342475b92.NewBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PlannerTaskRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    m := &PlannerTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner/plans/{plannerPlan_id}/buckets/{plannerBucket_id}/tasks/{plannerTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PlannerTaskRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) CreateDeleteRequestInformation(options *PlannerTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) CreateGetRequestInformation(options *PlannerTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) CreatePatchRequestInformation(options *PlannerTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) Delete(options *PlannerTaskRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerTaskRequestBuilder) Details()(*ia37f1dc3e53c9a3902fa069b5b6cb2af8ebaaf76c14f6643a9c371ecf02f1414.DetailsRequestBuilder) {
    return ia37f1dc3e53c9a3902fa069b5b6cb2af8ebaaf76c14f6643a9c371ecf02f1414.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) Get(options *PlannerTaskRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlannerTask() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask), nil
}
// Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - options : Options for the request
func (m *PlannerTaskRequestBuilder) Patch(options *PlannerTaskRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerTaskRequestBuilder) ProgressTaskBoardFormat()(*i0cc1127e54624c213444e0ed912e9f6e573eca5b01c61e0255b0c2de1a231be1.ProgressTaskBoardFormatRequestBuilder) {
    return i0cc1127e54624c213444e0ed912e9f6e573eca5b01c61e0255b0c2de1a231be1.NewProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
