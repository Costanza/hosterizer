# TypeScript Coding Standards

## Configuration

- Use strict mode in `tsconfig.json`
- Enable all strict type-checking options
- Set `target` to ES2020 or later
- Use `module: "ESNext"` with bundlers

```json
{
  "compilerOptions": {
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "target": "ES2020",
    "module": "ESNext",
    "moduleResolution": "bundler"
  }
}
```

## Type Definitions

- Prefer `interface` for object shapes that can be extended
- Use `type` for unions, intersections, and mapped types
- Use `enum` sparingly; prefer string literal unions
- Avoid `any`; use `unknown` when type is truly unknown

```typescript
// Interfaces for object shapes
interface User {
  id: string;
  name: string;
  email: string;
}

interface AdminUser extends User {
  permissions: string[];
}

// Types for unions and complex types
type Status = 'pending' | 'active' | 'inactive';
type Result<T> = { success: true; data: T } | { success: false; error: string };

// Avoid enums, use const objects instead
const UserRole = {
  ADMIN: 'admin',
  USER: 'user',
  GUEST: 'guest',
} as const;

type UserRole = typeof UserRole[keyof typeof UserRole];
```

## Type Annotations

- Annotate function parameters and return types
- Let TypeScript infer variable types when obvious
- Use explicit types for complex objects and arrays
- Use `readonly` for immutable data

```typescript
// Good: explicit parameter and return types
function calculateTotal(prices: number[], taxRate: number): number {
  const subtotal = prices.reduce((sum, price) => sum + price, 0);
  return subtotal * (1 + taxRate);
}

// Good: inferred type is obvious
const userName = 'John'; // string
const count = 42; // number

// Good: explicit type for complex structure
const config: AppConfig = {
  apiUrl: 'https://api.example.com',
  timeout: 5000,
  retries: 3,
};

// Readonly for immutable data
interface Point {
  readonly x: number;
  readonly y: number;
}
```

## Generics

- Use generics for reusable, type-safe functions and classes
- Use descriptive names for type parameters (not just `T`)
- Constrain generics when needed with `extends`

```typescript
// Generic function
function first<Item>(array: Item[]): Item | undefined {
  return array[0];
}

// Constrained generic
function getProperty<Obj, Key extends keyof Obj>(obj: Obj, key: Key): Obj[Key] {
  return obj[key];
}

// Generic interface
interface Repository<Entity> {
  findById(id: string): Promise<Entity | null>;
  save(entity: Entity): Promise<Entity>;
  delete(id: string): Promise<void>;
}
```

## Utility Types

- Use built-in utility types: `Partial`, `Required`, `Pick`, `Omit`, `Record`
- Create custom utility types for domain-specific transformations

```typescript
// Built-in utilities
type PartialUser = Partial<User>;
type UserWithoutEmail = Omit<User, 'email'>;
type UserIdAndName = Pick<User, 'id' | 'name'>;
type UserMap = Record<string, User>;

// Custom utility type
type Nullable<T> = { [K in keyof T]: T[K] | null };
type AsyncFunction<Args extends any[], Return> = (...args: Args) => Promise<Return>;
```

## Null Safety

- Use strict null checks
- Use optional chaining (`?.`) and nullish coalescing (`??`)
- Prefer explicit null checks over type assertions
- Use non-null assertion (`!`) sparingly and only when certain

```typescript
// Optional chaining
const userName = user?.profile?.name;

// Nullish coalescing
const displayName = userName ?? 'Anonymous';

// Type guard
function isUser(value: unknown): value is User {
  return (
    typeof value === 'object' &&
    value !== null &&
    'id' in value &&
    'name' in value
  );
}

if (isUser(data)) {
  console.log(data.name); // TypeScript knows data is User
}
```

## Type Guards

- Create type guards for runtime type checking
- Use `is` keyword for type predicate functions
- Prefer type guards over type assertions

```typescript
// Type guard function
function isError(value: unknown): value is Error {
  return value instanceof Error;
}

// Discriminated union
type ApiResponse<T> =
  | { status: 'success'; data: T }
  | { status: 'error'; error: string };

function handleResponse<T>(response: ApiResponse<T>): T {
  if (response.status === 'success') {
    return response.data; // TypeScript knows this is success case
  }
  throw new Error(response.error); // TypeScript knows this is error case
}
```

## Async/Await

- Always type Promise return values
- Use `async`/`await` over raw Promises
- Handle errors with try/catch
- Type error objects properly

```typescript
async function fetchUser(id: string): Promise<User> {
  try {
    const response = await fetch(`/api/users/${id}`);
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const data: unknown = await response.json();
    if (!isUser(data)) {
      throw new Error('Invalid user data');
    }
    return data;
  } catch (error) {
    if (error instanceof Error) {
      console.error('Failed to fetch user:', error.message);
    }
    throw error;
  }
}
```

## Modules

- Use ES modules (`import`/`export`)
- Use named exports over default exports
- Group imports: external libraries, internal modules, types
- Use path aliases for cleaner imports

```typescript
// Named exports (preferred)
export function calculateTotal(prices: number[]): number { }
export interface User { }

// Import grouping
import { useState, useEffect } from 'react';
import { format } from 'date-fns';

import { UserService } from '@/services/UserService';
import { formatCurrency } from '@/utils/format';

import type { User, Order } from '@/types';
```

## Best Practices

- Enable and fix all TypeScript errors (no `@ts-ignore`)
- Use `unknown` instead of `any` for truly unknown types
- Prefer immutability: use `readonly`, `const`, and `ReadonlyArray`
- Use discriminated unions for state machines
- Type external API responses and validate at runtime
- Use branded types for nominal typing when needed
- Keep types DRY (Don't Repeat Yourself)
- Document complex types with JSDoc comments

```typescript
// Branded type for nominal typing
type UserId = string & { readonly __brand: 'UserId' };
type OrderId = string & { readonly __brand: 'OrderId' };

function createUserId(id: string): UserId {
  return id as UserId;
}

// This prevents mixing up IDs
function getUser(id: UserId): User { }
function getOrder(id: OrderId): Order { }

const userId = createUserId('user-123');
const orderId = createOrderId('order-456');

getUser(userId); // OK
getUser(orderId); // Type error!
```
