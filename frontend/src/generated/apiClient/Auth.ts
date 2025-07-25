/* eslint-disable */
/* tslint:disable */
// @ts-nocheck
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

import {
  ApiResponseData,
  ApiResponseError,
  AuthLoginRequest,
  AuthRegisterRequest,
  CommonAuthResponse,
  CommonRefreshResponse,
} from "./data-contracts";
import { ContentType, HttpClient, RequestParams } from "./http-client";

export class Auth<
  SecurityDataType = unknown,
> extends HttpClient<SecurityDataType> {
  /**
   * @description Redirects the user to the Google OAuth consent page.
   *
   * @tags Auth Google
   * @name GoogleList
   * @summary Google Auth Login
   * @request GET:/auth/google
   */
  googleList = (params: RequestParams = {}) =>
    this.request<any, string>({
      path: `/auth/google`,
      method: "GET",
      ...params,
    });
  /**
   * @description Handles OAuth callback and authenticates/creates user account using Google data.
   *
   * @tags Auth Google
   * @name GoogleCallbackList
   * @summary Google Auth Callback
   * @request GET:/auth/google/callback
   */
  googleCallbackList = (
    query: {
      /** OAuth authorization code from Google */
      code: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      ApiResponseData & {
        data?: CommonAuthResponse;
      },
      ApiResponseError
    >({
      path: `/auth/google/callback`,
      method: "GET",
      query: query,
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Authenticate user with email and password
   *
   * @tags Auth
   * @name LoginCreate
   * @summary User login
   * @request POST:/auth/login
   */
  loginCreate = (payload: AuthLoginRequest, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: CommonAuthResponse;
      },
      ApiResponseError
    >({
      path: `/auth/login`,
      method: "POST",
      body: payload,
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Performs logout by invalidating user's authentication (such as token or session)
   *
   * @tags Auth
   * @name LogoutCreate
   * @summary Logout user
   * @request POST:/auth/logout
   */
  logoutCreate = (params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: string;
      },
      ApiResponseError
    >({
      path: `/auth/logout`,
      method: "POST",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Refresh access token using refresh token cookie
   *
   * @tags Auth
   * @name RefreshTokenCreate
   * @summary Refresh JWT token
   * @request POST:/auth/refresh-token
   */
  refreshTokenCreate = (params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: CommonRefreshResponse;
      },
      ApiResponseError
    >({
      path: `/auth/refresh-token`,
      method: "POST",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Register a new user with email and password
   *
   * @tags Auth
   * @name RegisterCreate
   * @summary User registration
   * @request POST:/auth/register
   */
  registerCreate = (payload: AuthRegisterRequest, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: CommonAuthResponse;
      },
      ApiResponseError
    >({
      path: `/auth/register`,
      method: "POST",
      body: payload,
      type: ContentType.Json,
      format: "json",
      ...params,
    });
}
