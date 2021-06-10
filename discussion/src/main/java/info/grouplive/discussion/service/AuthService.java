package info.grouplive.discussion.service;

import info.grouplive.proto.Api.*;
import info.grouplive.proto.AuthenticationServiceGrpc;
import info.grouplive.proto.AuthenticationServiceGrpc.AuthenticationServiceBlockingStub;
import info.grouplive.proto.AuthorizationServiceGrpc;
import info.grouplive.proto.AuthorizationServiceGrpc.AuthorizationServiceBlockingStub;
import info.grouplive.proto.UserServiceGrpc;
import info.grouplive.proto.UserServiceGrpc.UserServiceBlockingStub;
import io.grpc.Channel;
import io.grpc.ManagedChannelBuilder;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class AuthService {
    private final AuthenticationServiceBlockingStub authenticationServiceStub;
    private final AuthorizationServiceBlockingStub authorizationServiceStub;
    private final UserServiceBlockingStub userServiceStub;

    public AuthService(
            @Value("${auth.service.host}") String host,
            @Value("${auth.service.port}") int port) {
        Channel channel = ManagedChannelBuilder
                .forAddress(host, port)
                .usePlaintext()
                .build();
        authenticationServiceStub = AuthenticationServiceGrpc.newBlockingStub(channel);
        authorizationServiceStub = AuthorizationServiceGrpc.newBlockingStub(channel);
        userServiceStub = UserServiceGrpc.newBlockingStub(channel);
    }

    public info.grouplive.discussion.model.UserModel getUser(String authToken) {
        System.out.println("++++++++++++++++++++++");
        System.out.println(authToken);
        System.out.println("++++++++++++++++++++++");
        return getUserById(verifyIdentity(authToken));
    }

    public String verifyIdentity(String authToken) {
        VerifyIdentityRequest request = VerifyIdentityRequest
                .newBuilder()
                .setAuthToken(authToken)
                .build();
        VerifyIdentityResponse response = authenticationServiceStub.verifyIdentity(request);
        return response.getUserId();
    }

    public boolean hasPermission(User user, String permissionId, String resourceTypeId, String resourceId) {
        HasPermissionRequest request = HasPermissionRequest
                .newBuilder()
                .setPermissionId(permissionId)
                .setUserId(user.getId())
                .setResourceTypeId(resourceTypeId)
                .setResourceId(resourceId)
                .build();
        HasPermissionResponse response = authorizationServiceStub.hasPermission(request);
        return response.getHasPermission();
    }

    public info.grouplive.discussion.model.UserModel getUserById(String userId) {
        GetUserRequest request = GetUserRequest
                .newBuilder()
                .setUserId(userId)
                .build();
        User protoUser = userServiceStub.getUser(request);
        Unit protoUnit = protoUser.getUnit();

        return new info.grouplive.discussion.model.UserModel(
                protoUser.getId(),
                protoUser.getUsername(),
                protoUser.getFirstname(),
                protoUser.getLastname(),
                protoUser.getEmail(),
                protoUser.getPhone(),
                protoUnit.getAddress(),
                protoUnit.getAptNumber());
    }
}
