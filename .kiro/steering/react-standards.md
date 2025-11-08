# React Coding Standards

## Component Structure

- Use functional components with hooks (no class components)
- One component per file
- Use named exports for components
- Co-locate related files (component, styles, tests)

```
components/
├── UserProfile/
│   ├── UserProfile.tsx
│   ├── UserProfile.module.css
│   ├── UserProfile.test.tsx
│   └── index.ts
```

## Component Patterns

```tsx
import { useState, useEffect } from 'react';
import styles from './UserProfile.module.css';

interface UserProfileProps {
  userId: string;
  onUpdate?: (user: User) => void;
}

export function UserProfile({ userId, onUpdate }: UserProfileProps) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchUser(userId).then(setUser).finally(() => setLoading(false));
  }, [userId]);

  if (loading) return <LoadingSpinner />;
  if (!user) return <ErrorMessage message="User not found" />;

  return (
    <div className={styles.container}>
      <h2>{user.name}</h2>
      <p>{user.email}</p>
    </div>
  );
}
```

## Hooks

- Use built-in hooks: `useState`, `useEffect`, `useContext`, `useCallback`, `useMemo`
- Create custom hooks for reusable logic (prefix with `use`)
- Keep hooks at the top level (never in conditionals or loops)
- List all dependencies in dependency arrays

```tsx
// Custom hook
function useUser(userId: string) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    let cancelled = false;
    
    fetchUser(userId)
      .then(data => !cancelled && setUser(data))
      .catch(err => !cancelled && setError(err))
      .finally(() => !cancelled && setLoading(false));

    return () => { cancelled = true; };
  }, [userId]);

  return { user, loading, error };
}
```

## State Management

- Use `useState` for local component state
- Use `useContext` for shared state across components
- Use `useReducer` for complex state logic
- Consider Zustand or Redux Toolkit for global state
- Keep state as close to where it's used as possible

```tsx
// Context pattern
interface AppContextType {
  user: User | null;
  setUser: (user: User | null) => void;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

export function AppProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  
  return (
    <AppContext.Provider value={{ user, setUser }}>
      {children}
    </AppContext.Provider>
  );
}

export function useApp() {
  const context = useContext(AppContext);
  if (!context) throw new Error('useApp must be used within AppProvider');
  return context;
}
```

## Props

- Use TypeScript interfaces for prop types
- Destructure props in function parameters
- Use optional chaining and nullish coalescing
- Avoid prop drilling (use context or composition)

```tsx
interface ButtonProps {
  variant?: 'primary' | 'secondary';
  size?: 'small' | 'medium' | 'large';
  disabled?: boolean;
  onClick: () => void;
  children: React.ReactNode;
}

export function Button({ 
  variant = 'primary', 
  size = 'medium',
  disabled = false,
  onClick,
  children 
}: ButtonProps) {
  return (
    <button
      className={`btn btn-${variant} btn-${size}`}
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </button>
  );
}
```

## Performance

- Use `React.memo()` for expensive components that re-render often
- Use `useMemo()` for expensive calculations
- Use `useCallback()` for functions passed as props
- Lazy load components with `React.lazy()` and `Suspense`
- Avoid inline object/array creation in render

```tsx
const MemoizedList = React.memo(function UserList({ users }: { users: User[] }) {
  return (
    <ul>
      {users.map(user => <li key={user.id}>{user.name}</li>)}
    </ul>
  );
});

// Lazy loading
const Dashboard = React.lazy(() => import('./Dashboard'));

function App() {
  return (
    <Suspense fallback={<LoadingSpinner />}>
      <Dashboard />
    </Suspense>
  );
}
```

## Event Handlers

- Name handlers with `handle` prefix: `handleClick`, `handleSubmit`
- Use arrow functions for inline handlers only when necessary
- Extract complex logic into separate functions

```tsx
function Form() {
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    // process form data
  };

  return <form onSubmit={handleSubmit}>...</form>;
}
```

## Styling

- Use CSS Modules for component-scoped styles
- Use consistent naming: `styles.container`, `styles.button`
- Consider Tailwind CSS for utility-first approach
- Avoid inline styles except for dynamic values

## Testing

- Use React Testing Library (not Enzyme)
- Test user behavior, not implementation details
- Use `screen` queries: `getByRole`, `getByLabelText`, `getByText`
- Use `userEvent` for simulating user interactions

```tsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

test('submits form with user input', async () => {
  const handleSubmit = jest.fn();
  render(<LoginForm onSubmit={handleSubmit} />);
  
  await userEvent.type(screen.getByLabelText(/username/i), 'john');
  await userEvent.type(screen.getByLabelText(/password/i), 'secret');
  await userEvent.click(screen.getByRole('button', { name: /submit/i }));
  
  expect(handleSubmit).toHaveBeenCalledWith({ username: 'john', password: 'secret' });
});
```

## Best Practices

- Keep components small and focused
- Extract reusable logic into custom hooks
- Use composition over prop drilling
- Avoid premature optimization
- Use semantic HTML elements
- Ensure accessibility (ARIA labels, keyboard navigation)
- Handle loading and error states
- Clean up side effects in `useEffect` return functions
