# Bootstrap Coding Standards

## Version

- Use Bootstrap 5.x (latest stable)
- Include via CDN or npm package
- Use the compiled CSS or customize via Sass

## Grid System

- Use responsive grid classes: `container`, `row`, `col-*`
- Mobile-first approach: start with `col-*`, add breakpoints as needed
- Breakpoints: `sm` (≥576px), `md` (≥768px), `lg` (≥992px), `xl` (≥1200px), `xxl` (≥1400px)

```html
<div class="container">
  <div class="row">
    <div class="col-12 col-md-6 col-lg-4">Column</div>
  </div>
</div>
```

## Utility Classes

- Use spacing utilities: `m-*`, `p-*`, `mt-*`, `mb-*`, `mx-*`, `my-*`
- Spacing scale: 0-5 (0, 0.25rem, 0.5rem, 1rem, 1.5rem, 3rem)
- Display utilities: `d-none`, `d-block`, `d-flex`, `d-grid`
- Flexbox utilities: `justify-content-*`, `align-items-*`, `flex-direction-*`

```html
<div class="d-flex justify-content-between align-items-center p-3 mb-4">
  <h2 class="mb-0">Title</h2>
  <button class="btn btn-primary">Action</button>
</div>
```

## Components

- Use semantic component classes: `btn`, `card`, `navbar`, `modal`
- Combine with modifier classes: `btn-primary`, `card-header`, `navbar-dark`
- Use data attributes for JavaScript components: `data-bs-toggle`, `data-bs-target`

```html
<button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#myModal">
  Launch Modal
</button>

<div class="card">
  <div class="card-header">Header</div>
  <div class="card-body">
    <h5 class="card-title">Title</h5>
    <p class="card-text">Content</p>
  </div>
</div>
```

## Customization

- Override Bootstrap variables in custom Sass file
- Import Bootstrap after variable overrides
- Use `!default` flag for Bootstrap variables

```scss
// custom.scss
$primary: #007bff;
$border-radius: 0.5rem;

@import "bootstrap/scss/bootstrap";
```

## Best Practices

- Don't modify Bootstrap source files directly
- Use utility classes before writing custom CSS
- Maintain consistent spacing using Bootstrap's scale
- Use responsive utilities to hide/show elements
- Leverage Bootstrap's color system for consistency
- Use form validation classes: `is-valid`, `is-invalid`
- Prefer Bootstrap components over custom implementations
