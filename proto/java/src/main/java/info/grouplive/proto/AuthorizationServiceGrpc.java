package info.grouplive.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.38.0)",
    comments = "Source: api.proto")
public final class AuthorizationServiceGrpc {

  private AuthorizationServiceGrpc() {}

  public static final String SERVICE_NAME = "pb.AuthorizationService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<info.grouplive.proto.Api.HasPermissionRequest,
      info.grouplive.proto.Api.HasPermissionResponse> getHasPermissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "HasPermission",
      requestType = info.grouplive.proto.Api.HasPermissionRequest.class,
      responseType = info.grouplive.proto.Api.HasPermissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<info.grouplive.proto.Api.HasPermissionRequest,
      info.grouplive.proto.Api.HasPermissionResponse> getHasPermissionMethod() {
    io.grpc.MethodDescriptor<info.grouplive.proto.Api.HasPermissionRequest, info.grouplive.proto.Api.HasPermissionResponse> getHasPermissionMethod;
    if ((getHasPermissionMethod = AuthorizationServiceGrpc.getHasPermissionMethod) == null) {
      synchronized (AuthorizationServiceGrpc.class) {
        if ((getHasPermissionMethod = AuthorizationServiceGrpc.getHasPermissionMethod) == null) {
          AuthorizationServiceGrpc.getHasPermissionMethod = getHasPermissionMethod =
              io.grpc.MethodDescriptor.<info.grouplive.proto.Api.HasPermissionRequest, info.grouplive.proto.Api.HasPermissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "HasPermission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  info.grouplive.proto.Api.HasPermissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  info.grouplive.proto.Api.HasPermissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthorizationServiceMethodDescriptorSupplier("HasPermission"))
              .build();
        }
      }
    }
    return getHasPermissionMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AuthorizationServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceStub>() {
        @java.lang.Override
        public AuthorizationServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthorizationServiceStub(channel, callOptions);
        }
      };
    return AuthorizationServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AuthorizationServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceBlockingStub>() {
        @java.lang.Override
        public AuthorizationServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthorizationServiceBlockingStub(channel, callOptions);
        }
      };
    return AuthorizationServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AuthorizationServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthorizationServiceFutureStub>() {
        @java.lang.Override
        public AuthorizationServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthorizationServiceFutureStub(channel, callOptions);
        }
      };
    return AuthorizationServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class AuthorizationServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void hasPermission(info.grouplive.proto.Api.HasPermissionRequest request,
        io.grpc.stub.StreamObserver<info.grouplive.proto.Api.HasPermissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getHasPermissionMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getHasPermissionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                info.grouplive.proto.Api.HasPermissionRequest,
                info.grouplive.proto.Api.HasPermissionResponse>(
                  this, METHODID_HAS_PERMISSION)))
          .build();
    }
  }

  /**
   */
  public static final class AuthorizationServiceStub extends io.grpc.stub.AbstractAsyncStub<AuthorizationServiceStub> {
    private AuthorizationServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthorizationServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthorizationServiceStub(channel, callOptions);
    }

    /**
     */
    public void hasPermission(info.grouplive.proto.Api.HasPermissionRequest request,
        io.grpc.stub.StreamObserver<info.grouplive.proto.Api.HasPermissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getHasPermissionMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class AuthorizationServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<AuthorizationServiceBlockingStub> {
    private AuthorizationServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthorizationServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthorizationServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public info.grouplive.proto.Api.HasPermissionResponse hasPermission(info.grouplive.proto.Api.HasPermissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getHasPermissionMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class AuthorizationServiceFutureStub extends io.grpc.stub.AbstractFutureStub<AuthorizationServiceFutureStub> {
    private AuthorizationServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthorizationServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthorizationServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<info.grouplive.proto.Api.HasPermissionResponse> hasPermission(
        info.grouplive.proto.Api.HasPermissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getHasPermissionMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_HAS_PERMISSION = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AuthorizationServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(AuthorizationServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_HAS_PERMISSION:
          serviceImpl.hasPermission((info.grouplive.proto.Api.HasPermissionRequest) request,
              (io.grpc.stub.StreamObserver<info.grouplive.proto.Api.HasPermissionResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class AuthorizationServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AuthorizationServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return info.grouplive.proto.Api.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AuthorizationService");
    }
  }

  private static final class AuthorizationServiceFileDescriptorSupplier
      extends AuthorizationServiceBaseDescriptorSupplier {
    AuthorizationServiceFileDescriptorSupplier() {}
  }

  private static final class AuthorizationServiceMethodDescriptorSupplier
      extends AuthorizationServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    AuthorizationServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AuthorizationServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AuthorizationServiceFileDescriptorSupplier())
              .addMethod(getHasPermissionMethod())
              .build();
        }
      }
    }
    return result;
  }
}
