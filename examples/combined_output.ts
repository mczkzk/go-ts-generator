// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-16 13:14:25
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

// Placeholders for undefined types
type FileHeader = any;

/**
 * Address represents a physical address in API requests/responses
 */
export interface Address {
  street_line1: string;
  street_line2?: string;
  city: string;
  state: string;
  postal_code: string;
  country: string;
}

/**
 * FileUploadForm represents a form with file uploads
 */
export interface FileUploadForm {
  user_id: number;
  title: string;
  description: string;
  file?: FileHeader;
  images?: FileHeader[];
}

/**
 * RouteParams represents URL parameters in a route
 */
export interface RouteParams {
  id: number;
  postId: number;
  commentId: string;
  category_id: string;
}

/**
 * UserRequest represents a request to create or update a user
 *
 * @api Used in the following endpoints:
 * - post /users (Request)
 * - put /users/{id} (Request)
 */
export interface UserRequest {
  name: string;
  email: string;
  address: Address;
}

/**
 * UserResponse represents an API response with user data
 *
 * @api Used in the following endpoints:
 * - post /users (Response)
 * - put /users/{id} (Response)
 * - get /users/search (Response)
 * - post /auth/login (Response)
 * - post /auth/register (Response)
 * - get /users (Response)
 * - get /users/{id} (Response)
 */
export interface UserResponse {
  user_id: number;
  first_name: string;
  last_name: string;
  email: string;
  created_at: string /* RFC3339 */;
  updated_at: string /* RFC3339 */;
}

/**
 * SearchParams represents query parameters for search endpoints
 */
export interface SearchParams {
  q: string;
  page: number;
  limit: number;
  sort_by: string;
  sort_order: string;
  filters: string[];
}

/**
 * LoginForm represents a login form submission
 *
 * @api Used in the following endpoints:
 * - post /auth/login (Request)
 */
export interface LoginForm {
  user: string;
  pass: string;
  remember: boolean;
}

/**
 * RegisterForm represents a user registration form
 *
 * @api Used in the following endpoints:
 * - post /auth/register (Request)
 */
export interface RegisterForm {
  /**
   * @validation
   *   - binding: required
   *   - validate: min=3,max=50
   */
  username: string;
  /**
   * @validation
   *   - binding: required
   *   - validate: email
   */
  email: string;
  /**
   * @validation
   *   - binding: required
   *   - validate: min=8,containsAny=!@#$%^&*
   */
  password: string;
  /**
   * @validation
   *   - binding: required
   *   - validate: eqfield=Password
   */
  confirm_password: string;
  /**
   * @validation
   *   - binding: required
   *   - validate: eq=true
   */
  accept_terms: boolean;
}

/**
 * MixedTagsStruct demonstrates priority between JSON and form tags
 */
export interface MixedTagsStruct {
  id: number;
  name: string;
  email: string;
  json_only: string;
  form_only: string;
  NoTags: string;
}

/**
 * MixedParamStruct demonstrates priority between param and json tags
 */
export interface MixedParamStruct {
  id: number;
  name: string;
  json_only: string;
  param_only: string;
}

/**
 * SearchForm represents a search form with various filters
 */
export interface SearchForm {
  /**
   * @validation
   *   - validate: omitempty,max=100
   */
  query: string;
  /**
   * @validation
   *   - validate: omitempty,dive,max=50
   */
  categories: string[];
  /**
   * @validation
   *   - validate: omitempty,min=0
   */
  minPrice?: number;
  /**
   * @validation
   *   - validate: omitempty,gtfield=MinPrice
   */
  maxPrice?: number;
  /**
   * @validation
   *   - validate: omitempty,oneof=price
   */
  sortBy: string;
  /**
   * @validation
   *   - validate: omitempty,oneof=asc
   */
  sortOrder: string;
  /**
   * @validation
   *   - validate: min=1
   */
  page: number;
  /**
   * @validation
   *   - validate: min=1,max=100
   */
  limit: number;
}

