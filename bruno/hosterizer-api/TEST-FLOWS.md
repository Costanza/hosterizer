# API Test Flows

Visual guide to testing workflows in the Hosterizer API collection.

## Standard Authentication Flow

```mermaid
graph TD
    A[Health Check] -->|200 OK| B[Login - Success]
    B -->|Save tokens| C[Get Current User]
    C -->|Use access_token| D[Refresh Token]
    D -->|Get new tokens| E[Logout]
    
    style A fill:#e1f5ff
    style B fill:#c8e6c9
    style C fill:#fff9c4
    style D fill:#ffe0b2
    style E fill:#f8bbd0
```

## MFA Setup Flow

```mermaid
graph TD
    A[Login - Success] -->|Get authenticated| B[MFA Setup]
    B -->|Get QR code & secret| C[Scan QR Code]
    C -->|Use authenticator app| D[MFA Verify]
    D -->|Enter TOTP code| E{Code Valid?}
    E -->|Yes| F[MFA Enabled]
    E -->|No| D
    F --> G[Logout]
    G --> H[Login - With MFA]
    H -->|Requires TOTP| I[Authenticated]
    
    style A fill:#c8e6c9
    style B fill:#fff9c4
    style C fill:#e1bee7
    style D fill:#ffe0b2
    style E fill:#ffccbc
    style F fill:#c8e6c9
    style G fill:#f8bbd0
    style H fill:#fff9c4
    style I fill:#c8e6c9
```

## Error Testing Flow

```mermaid
graph TD
    A[Login - Invalid Credentials] -->|401 Unauthorized| B[Login - Missing Fields]
    B -->|400 Bad Request| C[Get Current User - Unauthorized]
    C -->|401 Unauthorized| D[Refresh Token - Invalid]
    D -->|401 Unauthorized| E[All Error Cases Tested]
    
    style A fill:#ffccbc
    style B fill:#ffccbc
    style C fill:#ffccbc
    style D fill:#ffccbc
    style E fill:#c8e6c9
```

## Complete Test Suite Flow

```mermaid
graph TD
    Start[Start Testing] --> Health[Health Check]
    
    Health -->|Service OK| Auth[Authentication Tests]
    Auth --> Login1[Login - Success]
    Auth --> Login2[Login - Invalid]
    Auth --> Login3[Login - Missing Fields]
    
    Login1 -->|Tokens Saved| Protected[Protected Endpoints]
    Protected --> Me1[Get Current User]
    Protected --> Me2[Get Current User - Unauthorized]
    
    Me1 --> Token[Token Management]
    Token --> Refresh1[Refresh Token]
    Token --> Refresh2[Refresh Token - Invalid]
    
    Refresh1 --> MFA[MFA Tests]
    MFA --> Setup[MFA Setup]
    MFA --> Verify[MFA Verify]
    MFA --> LoginMFA[Login - With MFA]
    
    LoginMFA --> Logout[Logout]
    Logout --> End[Testing Complete]
    
    style Start fill:#e1f5ff
    style Health fill:#e1f5ff
    style Auth fill:#fff9c4
    style Protected fill:#c8e6c9
    style Token fill:#ffe0b2
    style MFA fill:#e1bee7
    style Logout fill:#f8bbd0
    style End fill:#c8e6c9
```

## Request Dependencies

### No Dependencies (Can run anytime)
- Health Check
- Login - Invalid Credentials
- Login - Missing Fields
- Get Current User - Unauthorized
- Refresh Token - Invalid

### Requires Login
- Get Current User
- MFA Setup
- MFA Verify
- Logout

### Requires Refresh Token
- Refresh Token

### Requires MFA Setup
- MFA Verify
- Login - With MFA (after MFA enabled)

## Execution Order

### Recommended Order for First Run
1. Health Check
2. Login - Success
3. Get Current User
4. MFA Setup
5. MFA Verify (manual: enter code from app)
6. Refresh Token
7. Logout

### Quick Smoke Test
1. Health Check
2. Login - Success
3. Get Current User
4. Logout

### Error Testing Only
1. Login - Invalid Credentials
2. Login - Missing Fields
3. Get Current User - Unauthorized
4. Refresh Token - Invalid

### MFA Testing Only
1. Login - Success
2. MFA Setup
3. MFA Verify
4. Logout
5. Login - With MFA

## Token Lifecycle

```mermaid
sequenceDiagram
    participant Client
    participant Auth Service
    participant Database
    participant Redis
    
    Client->>Auth Service: POST /login
    Auth Service->>Database: Verify credentials
    Database-->>Auth Service: User valid
    Auth Service->>Auth Service: Generate tokens
    Auth Service->>Redis: Store session
    Auth Service-->>Client: access_token + refresh_token
    
    Note over Client: Use access_token for requests
    
    Client->>Auth Service: GET /me (with access_token)
    Auth Service->>Auth Service: Validate token
    Auth Service-->>Client: User info
    
    Note over Client: Token expires after 15 min
    
    Client->>Auth Service: POST /refresh (with refresh_token)
    Auth Service->>Auth Service: Validate refresh token
    Auth Service->>Auth Service: Generate new tokens
    Auth Service-->>Client: New access_token + refresh_token
    
    Client->>Auth Service: POST /logout
    Auth Service->>Redis: Delete session
    Auth Service-->>Client: Success
```

## MFA Flow

```mermaid
sequenceDiagram
    participant Client
    participant Auth Service
    participant Database
    participant Authenticator App
    
    Client->>Auth Service: POST /mfa/setup
    Auth Service->>Auth Service: Generate TOTP secret
    Auth Service->>Database: Store secret (disabled)
    Auth Service-->>Client: QR code + secret
    
    Client->>Authenticator App: Scan QR code
    Authenticator App-->>Client: Generate TOTP code
    
    Client->>Auth Service: POST /mfa/verify (with code)
    Auth Service->>Auth Service: Validate TOTP code
    Auth Service->>Database: Enable MFA
    Auth Service-->>Client: MFA enabled
    
    Note over Client: Future logins require MFA
    
    Client->>Auth Service: POST /login (email + password)
    Auth Service->>Database: Verify credentials
    Auth Service-->>Client: requires_mfa: true
    
    Client->>Authenticator App: Get current code
    Authenticator App-->>Client: TOTP code
    
    Client->>Auth Service: POST /login (with MFA code)
    Auth Service->>Auth Service: Validate TOTP
    Auth Service-->>Client: Tokens + user info
```

## Environment Variable Flow

```mermaid
graph LR
    A[Login - Success] -->|Sets| B[access_token]
    A -->|Sets| C[refresh_token]
    A -->|Sets| D[user_id]
    A -->|Sets| E[user_uuid]
    
    F[MFA Setup] -->|Sets| G[mfa_secret]
    F -->|Sets| H[mfa_qr_code]
    
    I[Refresh Token] -->|Updates| B
    I -->|Updates| C
    
    B -->|Used by| J[Get Current User]
    B -->|Used by| K[MFA Setup]
    B -->|Used by| L[MFA Verify]
    B -->|Used by| M[Logout]
    
    C -->|Used by| I
    
    style A fill:#c8e6c9
    style F fill:#fff9c4
    style I fill:#ffe0b2
    style B fill:#e1f5ff
    style C fill:#e1f5ff
```

## Testing Strategies

### Smoke Testing (2 minutes)
Quick verification that core functionality works:
1. Health Check
2. Login - Success
3. Get Current User

### Regression Testing (5 minutes)
Verify all functionality after changes:
1. Run entire "Auth Service" folder
2. Review all test results
3. Verify no regressions

### Integration Testing (10 minutes)
Test complete user journeys:
1. Standard auth flow
2. MFA setup flow
3. Token refresh flow
4. Error handling flow

### Security Testing (15 minutes)
Verify security controls:
1. Test invalid credentials
2. Test missing authentication
3. Test invalid tokens
4. Test MFA enforcement
5. Test account lockout (manual)

## Tips for Efficient Testing

1. **Use Folder Runner**: Run entire folders instead of individual requests
2. **Check Test Results**: Review the "Tests" tab after each request
3. **Monitor Logs**: Keep auth service logs visible during testing
4. **Use Variables**: Let the collection manage tokens automatically
5. **Test Errors**: Don't just test happy paths
6. **Document Issues**: Note any unexpected behavior
7. **Clean State**: Logout between test runs for clean state

## Troubleshooting Flows

### Token Issues
```
Login fails → Check credentials
Token expired → Run Refresh Token
Refresh fails → Run Login - Success again
```

### MFA Issues
```
Setup fails → Check authentication
Verify fails → Check TOTP code is current
Login fails → Ensure MFA is enabled
```

### Connection Issues
```
Health Check fails → Check service is running
All requests fail → Check base_url in environment
Timeout → Check service logs for errors
```
