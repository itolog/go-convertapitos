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
  ApiMeta,
  ApiResponseData,
  ApiResponseError,
  UserCreateRequest,
  UserUpdateRequest,
  UserUser,
} from './data-contracts.ts'
import { ContentType, HttpClient, RequestParams } from './http-client.ts'

export class User<SecurityDataType = unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description Returns a list of all users with pagination and sorting options
   *
   * @tags User
   * @name UserList
   * @summary Get all users
   * @request GET:/user
   */
  userList = (
    query?: {
      /**
       * Number of records per page
       * @default 10
       */
      limit?: number
      /**
       * Page number
       * @min 1
       * @default 1
       */
      page?: number
      /**
       * Field to order by
       * @default "updated_at"
       */
      order_by?: string
      /**
       * Sort in descending order
       * @default false
       */
      desc?: boolean
    },
    params: RequestParams = {},
  ) =>
    this.request<
      ApiResponseData & {
        data?: UserUser[]
        meta?: ApiMeta
      },
      ApiResponseError & {
        error?: string
      }
    >({
      path: `/user`,
      method: 'GET',
      query: query,
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Creates a new user with provided data
   *
   * @tags User
   * @name UserCreate
   * @summary Create new user
   * @request POST:/user
   */
  userCreate = (user: UserCreateRequest, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: UserUser
      },
      ApiResponseError
    >({
      path: `/user`,
      method: 'POST',
      body: user,
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Returns user data by email
   *
   * @tags User
   * @name ByEmailDetail
   * @summary Get user by email
   * @request GET:/user/by_email/{email}
   */
  byEmailDetail = (email: string, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: UserUser
      },
      ApiResponseError
    >({
      path: `/user/by_email/${email}`,
      method: 'GET',
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Returns user data by ID
   *
   * @tags User
   * @name UserDetail
   * @summary Get user by ID
   * @request GET:/user/{id}
   */
  userDetail = (id: string, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: UserUser
      },
      ApiResponseError
    >({
      path: `/user/${id}`,
      method: 'GET',
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Deletes a user by ID
   *
   * @tags User
   * @name UserDelete
   * @summary Delete user
   * @request DELETE:/user/{id}
   */
  userDelete = (id: string, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: string
      },
      ApiResponseError
    >({
      path: `/user/${id}`,
      method: 'DELETE',
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
  /**
   * @description Updates existing user data
   *
   * @tags User
   * @name UserPartialUpdate
   * @summary Update user
   * @request PATCH:/user/{id}
   */
  userPartialUpdate = (id: string, user: UserUpdateRequest, params: RequestParams = {}) =>
    this.request<
      ApiResponseData & {
        data?: UserUser
      },
      ApiResponseError
    >({
      path: `/user/${id}`,
      method: 'PATCH',
      body: user,
      type: ContentType.Json,
      format: 'json',
      ...params,
    })
}
