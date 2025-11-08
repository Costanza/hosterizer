# Tailwind CSS Coding Standards

## Configuration

- Use `tailwind.config.js` for customization
- Extend default theme, don't replace it
- Define custom colors, spacing, and breakpoints in config
- Use content paths to enable tree-shaking

```javascript
module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#f0f9ff',
          500: '#0ea5e9',
          900: '#0c4a6e',
        },
      },
    },
  },
}
```

## Utility-First Approach

- Use utility classes directly in HTML/JSX
- Avoid creating custom CSS classes unless absolutely necessary
- Compose utilities to build complex designs
- Use responsive modifiers: `sm:`, `md:`, `lg:`, `xl:`, `2xl:`

```jsx
<div className="flex items-center justify-between p-4 bg-white rounded-lg shadow-md">
  <h2 className="text-2xl font-bold text-gray-900">Title</h2>
  <button className="px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700">
    Action
  </button>
</div>
```

## Responsive Design

- Mobile-first by default (unprefixed utilities apply to all screen sizes)
- Add breakpoint prefixes for larger screens
- Breakpoints: `sm` (640px), `md` (768px), `lg` (1024px), `xl` (1280px), `2xl` (1536px)

```jsx
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
  {/* Responsive grid */}
</div>
```

## Component Patterns

- Extract repeated utility combinations into components
- Use `@apply` directive sparingly in CSS for component classes
- Prefer component composition over `@apply`

```css
/* Use @apply only for true component abstractions */
.btn-primary {
  @apply px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500;
}
```

## State Variants

- Use state modifiers: `hover:`, `focus:`, `active:`, `disabled:`
- Use group utilities for parent-child state relationships
- Use peer utilities for sibling state relationships

```jsx
<div className="group">
  <img className="group-hover:scale-110 transition" />
  <p className="group-hover:text-blue-600">Hover me</p>
</div>
```

## Dark Mode

- Configure dark mode in `tailwind.config.js`
- Use `dark:` modifier for dark mode styles
- Choose strategy: `media` (system preference) or `class` (manual toggle)

```jsx
<div className="bg-white dark:bg-gray-900 text-gray-900 dark:text-white">
  Content adapts to dark mode
</div>
```

## Custom Utilities

- Add custom utilities in `tailwind.config.js` using `addUtilities`
- Keep custom utilities minimal
- Document custom utilities in team documentation

## Best Practices

- Keep class lists organized: layout → spacing → sizing → colors → typography → effects
- Use Tailwind's design tokens (colors, spacing) for consistency
- Leverage JIT mode for arbitrary values: `w-[137px]`, `top-[117px]`
- Use `clsx` or `classnames` library for conditional classes
- Extract components when utility lists exceed 10-15 classes
- Use Prettier plugin for automatic class sorting
- Avoid mixing Tailwind with other CSS frameworks
