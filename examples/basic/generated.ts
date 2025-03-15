// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-15 12:34:56
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

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
  category?: Category; // Pointer type without omitempty is still optional
  createdAt: string /* RFC3339 */;
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
export interface UnexportedType {
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