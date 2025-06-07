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

export enum ApiStatusType {
  StatusSuccess = "success",
  StatusError = "error",
}

export interface ApiErrorResponse {
  code?: number;
  details?: string;
  fields?: ValidationErrorFields[];
  message?: string;
}

export interface ApiMeta {
  items?: number;
  pages?: number;
}

export interface ApiResponseData<T> {
  data?: T;
  meta?: ApiMeta;
  status?: ApiStatusType;
}

export interface ApiResponseError {
  error?: ApiErrorResponse;
  status?: ApiStatusType;
}

export interface AuthLoginRequest {
  email: string;
  /**
   * @minLength 6
   * @maxLength 128
   */
  password: string;
}

export interface AuthRegisterRequest {
  email: string;
  /** @maxLength 70 */
  name: string;
  /**
   * @minLength 6
   * @maxLength 128
   */
  password: string;
}

export interface CommonAuthResponse {
  accessToken?: string;
  user?: UserUser;
}

export interface CommonRefreshResponse {
  accessToken?: string;
}

export interface UserBatchDeleteRequest {
  /** @minItems 1 */
  ids: string[];
}

export interface UserCreateRequest {
  email: string;
  /** @maxLength 70 */
  name: string;
  /**
   * @minLength 6
   * @maxLength 128
   */
  password: string;
  picture?: string;
  verifiedEmail?: boolean;
}

export interface UserUpdateRequest {
  email?: string;
  /** @maxLength 70 */
  name?: string;
  /**
   * @minLength 6
   * @maxLength 128
   */
  password?: string;
  picture?: string;
  verifiedEmail?: boolean;
  role?: string;
}

export interface UserUser {
  createdAt?: string;
  email?: string;
  id?: string;
  name?: string;
  password?: string;
  picture?: string;
  updatedAt?: string;
  verifiedEmail?: boolean;
  role?: string;
}

export interface ValidationErrorFields {
  field?: string;
  param?: string;
  tag?: string;
}
