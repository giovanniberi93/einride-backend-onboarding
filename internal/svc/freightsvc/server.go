package freightsvc

import onboardingv1 "github.com/giovanniberi93/einride-backend-onboarding/proto/gen/go/einride/onboarding/v1"

// Server implements onboardingv1.FreightServiceServer.
type Server struct{}

var _ onboardingv1.FreightServiceServer = &Server{}
