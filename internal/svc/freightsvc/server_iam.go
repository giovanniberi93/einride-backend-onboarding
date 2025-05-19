package freightsvc

import (
    "context"

    "google.golang.org/genproto/googleapis/iam/v1"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// SetIamPolicy implements iam.IAMPolicyServer.
func (s *Server) SetIamPolicy(
    ctx context.Context,
    request *iam.SetIamPolicyRequest,
) (*iam.Policy, error) {
    return nil, status.Error(codes.Unimplemented, "TODO: implement me")
}

// GetIamPolicy implements iam.IAMPolicyServer.
func (s *Server) GetIamPolicy(
    ctx context.Context,
    request *iam.GetIamPolicyRequest,
) (*iam.Policy, error) {
    return nil, status.Error(codes.Unimplemented, "TODO: implement me")
}

// TestIamPermissions implements iam.IAMPolicyServer.
func (s *Server) TestIamPermissions(
    ctx context.Context,
    request *iam.TestIamPermissionsRequest,
) (*iam.TestIamPermissionsResponse, error) {
    return nil, status.Error(codes.Unimplemented, "TODO: implement me")
}
