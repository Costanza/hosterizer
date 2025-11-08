# Python Coding Standards

## Code Style

- Follow PEP 8 style guide strictly
- Use 4 spaces for indentation (never tabs)
- Maximum line length: 88 characters (Black formatter standard)
- Use double quotes for strings by default
- Add trailing commas in multi-line collections

## Type Hints

- Always use type hints for function parameters and return values
- Use `from typing import` for complex types (List, Dict, Optional, Union, etc.)
- For Python 3.9+, prefer built-in generics: `list[str]` over `List[str]`
- Use `Optional[T]` for nullable types, never `T | None` in function signatures

```python
def process_data(items: list[str], max_count: Optional[int] = None) -> dict[str, int]:
    """Process items and return counts."""
    pass
```

## Project Structure

```
project/
├── src/
│   └── package_name/
│       ├── __init__.py
│       ├── models/
│       ├── services/
│       └── utils/
├── tests/
│   └── test_*.py
├── pyproject.toml
└── README.md
```

## Dependencies

- Use `pyproject.toml` for dependency management (PEP 621)
- Pin major versions, allow minor/patch updates: `requests = "^2.31.0"`
- Separate dev dependencies from production dependencies
- Use virtual environments (venv or poetry)

## Error Handling

- Use specific exception types, avoid bare `except:`
- Create custom exceptions for domain-specific errors
- Always log exceptions with context

```python
class ValidationError(Exception):
    """Raised when data validation fails."""
    pass

try:
    validate_input(data)
except ValueError as e:
    logger.error(f"Validation failed: {e}", exc_info=True)
    raise ValidationError(f"Invalid data: {e}") from e
```

## Naming Conventions

- Classes: `PascalCase`
- Functions/methods: `snake_case`
- Constants: `UPPER_SNAKE_CASE`
- Private methods: `_leading_underscore`
- Module names: `lowercase_with_underscores`

## Documentation

- Use docstrings for all public modules, classes, and functions
- Follow Google or NumPy docstring format consistently
- Include type information in docstrings only if it adds clarity beyond type hints

```python
def calculate_total(prices: list[float], tax_rate: float = 0.0) -> float:
    """Calculate total price including tax.
    
    Args:
        prices: List of item prices
        tax_rate: Tax rate as decimal (e.g., 0.08 for 8%)
        
    Returns:
        Total price with tax applied
        
    Raises:
        ValueError: If tax_rate is negative
    """
    pass
```

## Testing

- Use `pytest` as the testing framework
- Test file names: `test_*.py` or `*_test.py`
- Use fixtures for setup/teardown
- Aim for 80%+ code coverage on business logic
- Use `pytest-cov` for coverage reports

## Imports

- Group imports: standard library, third-party, local
- Use absolute imports over relative imports
- Sort imports alphabetically within groups
- Use `isort` for automatic import sorting

```python
import os
from pathlib import Path

import requests
from fastapi import FastAPI

from myapp.models import User
from myapp.services import UserService
```

## Best Practices

- Prefer composition over inheritance
- Use list/dict comprehensions for simple transformations
- Use context managers (`with` statements) for resource management
- Avoid mutable default arguments
- Use `pathlib.Path` for file path operations
- Use f-strings for string formatting
- Keep functions small and focused (single responsibility)
- Use Pydantic models for data structures
