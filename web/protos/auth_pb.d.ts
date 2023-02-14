import * as jspb from 'google-protobuf';

export class SignupRequest extends jspb.Message {
    getUsername(): string;
    setUsername(value: string): SignupRequest;

    getEmail(): string;
    setEmail(value: string): SignupRequest;

    getPassword(): string;
    setPassword(value: string): SignupRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignupRequest.AsObject;
    static toObject(
        includeInstance: boolean,
        msg: SignupRequest
    ): SignupRequest.AsObject;
    static serializeBinaryToWriter(
        message: SignupRequest,
        writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): SignupRequest;
    static deserializeBinaryFromReader(
        message: SignupRequest,
        reader: jspb.BinaryReader
    ): SignupRequest;
}

export namespace SignupRequest {
    export type AsObject = {
        username: string;
        email: string;
        password: string;
    };
}

export class SignupResponse extends jspb.Message {
    getUserId(): string;
    setUserId(value: string): SignupResponse;

    getTokenType(): string;
    setTokenType(value: string): SignupResponse;

    getAccessToken(): string;
    setAccessToken(value: string): SignupResponse;

    getRefreshToken(): string;
    setRefreshToken(value: string): SignupResponse;

    getExpiresIn(): number;
    setExpiresIn(value: number): SignupResponse;

    getScope(): string;
    setScope(value: string): SignupResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignupResponse.AsObject;
    static toObject(
        includeInstance: boolean,
        msg: SignupResponse
    ): SignupResponse.AsObject;
    static serializeBinaryToWriter(
        message: SignupResponse,
        writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): SignupResponse;
    static deserializeBinaryFromReader(
        message: SignupResponse,
        reader: jspb.BinaryReader
    ): SignupResponse;
}

export namespace SignupResponse {
    export type AsObject = {
        userId: string;
        tokenType: string;
        accessToken: string;
        refreshToken: string;
        expiresIn: number;
        scope: string;
    };
}

export class SigninRequest extends jspb.Message {
    getEmail(): string;
    setEmail(value: string): SigninRequest;

    getPassword(): string;
    setPassword(value: string): SigninRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SigninRequest.AsObject;
    static toObject(
        includeInstance: boolean,
        msg: SigninRequest
    ): SigninRequest.AsObject;
    static serializeBinaryToWriter(
        message: SigninRequest,
        writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): SigninRequest;
    static deserializeBinaryFromReader(
        message: SigninRequest,
        reader: jspb.BinaryReader
    ): SigninRequest;
}

export namespace SigninRequest {
    export type AsObject = {
        email: string;
        password: string;
    };
}

export class SigninResponse extends jspb.Message {
    getUserId(): string;
    setUserId(value: string): SigninResponse;

    getTokenType(): string;
    setTokenType(value: string): SigninResponse;

    getAccessToken(): string;
    setAccessToken(value: string): SigninResponse;

    getRefreshToken(): string;
    setRefreshToken(value: string): SigninResponse;

    getExpiresIn(): number;
    setExpiresIn(value: number): SigninResponse;

    getScope(): string;
    setScope(value: string): SigninResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SigninResponse.AsObject;
    static toObject(
        includeInstance: boolean,
        msg: SigninResponse
    ): SigninResponse.AsObject;
    static serializeBinaryToWriter(
        message: SigninResponse,
        writer: jspb.BinaryWriter
    ): void;
    static deserializeBinary(bytes: Uint8Array): SigninResponse;
    static deserializeBinaryFromReader(
        message: SigninResponse,
        reader: jspb.BinaryReader
    ): SigninResponse;
}

export namespace SigninResponse {
    export type AsObject = {
        userId: string;
        tokenType: string;
        accessToken: string;
        refreshToken: string;
        expiresIn: number;
        scope: string;
    };
}
