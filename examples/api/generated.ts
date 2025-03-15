// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-15 12:34:56
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

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