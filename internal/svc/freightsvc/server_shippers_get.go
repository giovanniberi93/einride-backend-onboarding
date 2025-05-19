package freightsvc

import (
    "context"

    onboardingv1 "github.com/giovanniberi93/einride-backend-onboarding/proto/gen/go/einride/onboarding/v1"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// GetShipper implements onboardingv1.FreightServiceServer.
func (s *Server) GetShipper(
    ctx context.Context,
    request *onboardingv1.GetShipperRequest,
) (*onboardingv1.Shipper, error) {
    return nil, status.Error(codes.Unimplemented, "TODO: implement me")
}
