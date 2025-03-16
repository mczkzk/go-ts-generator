// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-16 13:39:00
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

// Placeholders for undefined types
type FileHeader = any;

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
 * Address represents a physical address
 */
export interface Address {
  street: string;
  city: string;
  state: string;
  zipCode: string;
  country: string;
}

/**
 * Product represents a product in the catalog
 */
export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  /**
   * Pointer type without omitempty
   */
  category?: Category;
  createdAt: string /* RFC3339 */;
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
 * User represents a user in the system
 */
export interface User {
  id: number;
  name: string;
  email: string;
  createdAt: string /* RFC3339 */;
  updatedAt: string /* RFC3339 */;
  address?: Address;
}

/**
 * UserList represents a list of users
 */
export type UserList = User[];

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
 * MixedParamStruct demonstrates priority between param and json tags
 */
export interface MixedParamStruct {
  id: number;
  name: string;
  json_only: string;
  param_only: string;
}

/**
 * CategoryMap is a map of category IDs to categories
 */
export type CategoryMap = Record<number, Category>;

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
 * StringArray is a simple string array
 */
export type StringArray = string[];

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
  user: string;
  pass: string;
  remember: boolean;
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

/**
 * Category represents a product category
 */
export interface Category {
  id: number;
  name: string;
}

/**
 * unexportedType is not exported
 */
/**
 * Note: This is an unexported type. In Go code, it's defined with a lowercase identifier.
 * It cannot be accessed directly from outside the package.
 */
export interface unexportedType {
  /**
   * Note: This is an unexported field. In Go code, it's defined with a lowercase identifier.
   * It cannot be accessed directly from outside the package.
   */
  field1: string;
  /**
   * Note: This is an unexported field. In Go code, it's defined with a lowercase identifier.
   * It cannot be accessed directly from outside the package.
   */
  field2: number;
}

