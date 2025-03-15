// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-15 14:37:06
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

// Placeholders for undefined types
type FileHeader = any;

/**
 * UserRequest represents a request to create or update a user
 */
export interface UserRequest {
  name: string;
  email: string;
  address: Address;
}

/**
 * UserResponse represents an API response with user data
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
 */
export interface LoginForm {
  username: string;
  password: string;
  remember_me: boolean;
}

/**
 * RegisterForm represents a user registration form
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
  user_id: number;
  user_name: string;
  user_email: string;
  json_only: string;
  form_only: string;
  NoTags: string;
}

/**
 * FileUploadForm represents a form with file uploads
 */
export interface FileUploadForm {
  user_id: number;
  title: string;
  description: string;
  file?: FileHeader;
  images: FileHeader[];
}

/**
 * RouteParams represents URL parameters in a route
 */
export interface RouteParams {
  user_id: number;
  post_id: number;
  comment_id: string;
  category_id: string;
}

/**
 * MixedParamStruct demonstrates priority between param and json tags
 */
export interface MixedParamStruct {
  user_id: number;
  user_name: string;
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
  q: string;
  /**
   * @validation
   *   - validate: omitempty,dive,max=50
   */
  categories: string[];
  /**
   * @validation
   *   - validate: omitempty,min=0
   */
  min_price?: number;
  /**
   * @validation
   *   - validate: omitempty,gtfield=MinPrice
   */
  max_price?: number;
  /**
   * @validation
   *   - validate: omitempty,oneof=price
   */
  sort_by: string;
  /**
   * @validation
   *   - validate: omitempty,oneof=asc
   */
  sort_order: string;
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

