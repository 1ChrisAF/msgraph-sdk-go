package models
import (
    "errors"
)
// Provides operations to manage the collection of agreement entities.
type BookingStaffRole int

const (
    GUEST_BOOKINGSTAFFROLE BookingStaffRole = iota
    ADMINISTRATOR_BOOKINGSTAFFROLE
    VIEWER_BOOKINGSTAFFROLE
    EXTERNALGUEST_BOOKINGSTAFFROLE
    UNKNOWNFUTUREVALUE_BOOKINGSTAFFROLE
)

func (i BookingStaffRole) String() string {
    return []string{"guest", "administrator", "viewer", "externalGuest", "unknownFutureValue"}[i]
}
func ParseBookingStaffRole(v string) (interface{}, error) {
    result := GUEST_BOOKINGSTAFFROLE
    switch v {
        case "guest":
            result = GUEST_BOOKINGSTAFFROLE
        case "administrator":
            result = ADMINISTRATOR_BOOKINGSTAFFROLE
        case "viewer":
            result = VIEWER_BOOKINGSTAFFROLE
        case "externalGuest":
            result = EXTERNALGUEST_BOOKINGSTAFFROLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BOOKINGSTAFFROLE
        default:
            return 0, errors.New("Unknown BookingStaffRole value: " + v)
    }
    return &result, nil
}
func SerializeBookingStaffRole(values []BookingStaffRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
