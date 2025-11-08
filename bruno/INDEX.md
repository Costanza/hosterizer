# Bruno Collection Index

Quick reference guide to all Bruno collection documentation.

## 🚀 Getting Started

**New to Bruno?** Start here:
1. [QUICKSTART.md](QUICKSTART.md) - Get testing in 5 minutes
2. [hosterizer-api/README.md](hosterizer-api/README.md) - Comprehensive guide
3. [hosterizer-api/Auth Service/README.md](hosterizer-api/Auth%20Service/README.md) - Auth endpoints

## 📚 Documentation Structure

```
📁 bruno/
│
├── 🚀 QUICKSTART.md                    ← Start here!
├── 📖 README.md                        ← Bruno overview
├── 📊 COLLECTION-SUMMARY.md            ← Statistics & metrics
├── 📑 INDEX.md                         ← This file
│
└── 📁 hosterizer-api/
    ├── 📖 README.md                    ← Main collection docs
    ├── 📊 TEST-FLOWS.md                ← Visual diagrams
    ├── 🔧 collection.bru               ← Collection metadata
    │
    ├── 📁 environments/
    │   ├── local.bru                   ← Local config
    │   └── dev.bru                     ← Dev config
    │
    └── 📁 Auth Service/
        ├── 📖 README.md                ← Service docs
        └── *.bru                       ← 12 API requests
```

## 📖 Documentation Files

### Quick Reference
| File | Purpose | Read Time |
|------|---------|-----------|
| [QUICKSTART.md](QUICKSTART.md) | Get started fast | 5 min |
| [README.md](README.md) | Bruno overview | 3 min |
| [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md) | Stats & metrics | 5 min |

### Detailed Guides
| File | Purpose | Read Time |
|------|---------|-----------|
| [hosterizer-api/README.md](hosterizer-api/README.md) | Complete collection guide | 15 min |
| [hosterizer-api/TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md) | Visual test flows | 10 min |
| [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md) | Auth endpoints | 10 min |

### External Documentation
| File | Purpose | Location |
|------|---------|----------|
| API Testing Guide | Testing best practices | [docs/api-testing.md](../docs/api-testing.md) |
| Auth Service README | Service implementation | [backend/auth-service/README.md](../backend/auth-service/README.md) |

## 🎯 Quick Links by Task

### I want to...

#### Get Started
→ [QUICKSTART.md](QUICKSTART.md)

#### Understand the Collection
→ [hosterizer-api/README.md](hosterizer-api/README.md)

#### See Test Flows
→ [hosterizer-api/TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md)

#### Learn About Auth Endpoints
→ [hosterizer-api/Auth Service/README.md](hosterizer-api/Auth%20Service/README.md)

#### View Statistics
→ [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md)

#### Setup Test Users
→ [../scripts/create-test-user.sh](../scripts/create-test-user.sh)

#### Troubleshoot Issues
→ [hosterizer-api/README.md#troubleshooting](hosterizer-api/README.md#troubleshooting)

#### Integrate with CI/CD
→ [hosterizer-api/README.md#cicd-integration](hosterizer-api/README.md#cicd-integration)

#### Add New Requests
→ [hosterizer-api/README.md#contributing](hosterizer-api/README.md#contributing)

## 🔍 Find Information By Topic

### Authentication
- Login flows: [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md#standard-authentication-flow)
- Token management: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#token-management)
- Error handling: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#error-testing-flow)

### Multi-Factor Authentication
- MFA setup: [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md#mfa-setup-flow)
- MFA testing: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#mfa-testing-only)
- Troubleshooting: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#mfa-verification-fails)

### Environment Configuration
- Variables: [README.md](hosterizer-api/README.md#environment-variables)
- Setup: [QUICKSTART.md](QUICKSTART.md#step-3-open-collection-in-bruno)
- Switching: [README.md](README.md#setup-environment)

### Testing Strategies
- Smoke testing: [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md#quick-smoke-test-2-minutes)
- Regression testing: [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md#full-regression-test-5-minutes)
- Security testing: [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md#security-testing-15-minutes)

### Request Details
- All endpoints: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#endpoints-overview)
- Request flow: [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md#request-flow)
- Dependencies: [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md#request-dependencies)

## 📊 Collection Overview

### Statistics
- **12 API requests** across 1 service
- **60+ automated tests**
- **2,500+ lines of documentation**
- **8 environment variables** (6 auto-managed)

### Coverage
- ✅ Health checks
- ✅ Authentication (login/logout)
- ✅ Token management (access/refresh)
- ✅ User information
- ✅ Multi-factor authentication
- ✅ Error handling
- ✅ Security testing

## 🎓 Learning Path

### Beginner (30 minutes)
1. Read [QUICKSTART.md](QUICKSTART.md)
2. Install Bruno
3. Open collection
4. Run "Health Check"
5. Run "Login - Success"
6. Run "Get Current User"

### Intermediate (1 hour)
1. Complete beginner path
2. Read [hosterizer-api/README.md](hosterizer-api/README.md)
3. Run all Auth Service requests
4. Review test results
5. Explore environment variables

### Advanced (2 hours)
1. Complete intermediate path
2. Read [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md)
3. Setup and test MFA flow
4. Test all error cases
5. Review [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md)
6. Explore CI/CD integration

## 🛠️ Helper Resources

### Scripts
- `../scripts/create-test-user.sh` - Create test users (Linux/Mac)
- `../scripts/create-test-user.ps1` - Create test users (Windows)

### Test Credentials
- Admin: `admin@hosterizer.com` / `AdminPass123!`
- Customer: `customer@hosterizer.com` / `AdminPass123!`

### External Links
- [Bruno Website](https://www.usebruno.com/)
- [Bruno Documentation](https://docs.usebruno.com/)
- [Bruno GitHub](https://github.com/usebruno/bruno)

## 🔄 Document Updates

### When to Update

**Add new endpoint:**
1. Create `.bru` file
2. Update [Auth Service/README.md](hosterizer-api/Auth%20Service/README.md)
3. Update [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md) statistics

**Add new service:**
1. Create service folder
2. Create service README
3. Update [hosterizer-api/README.md](hosterizer-api/README.md)
4. Update [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md)

**Add new environment:**
1. Create `.bru` file in environments/
2. Update [hosterizer-api/README.md](hosterizer-api/README.md)
3. Update [QUICKSTART.md](QUICKSTART.md)

## 📞 Getting Help

### Documentation Issues
1. Check this index for the right document
2. Use browser search (Ctrl+F) within documents
3. Check table of contents in each document

### Technical Issues
1. Review [Troubleshooting](hosterizer-api/README.md#troubleshooting)
2. Check service logs
3. Verify prerequisites
4. Review [QUICKSTART.md](QUICKSTART.md)

### Feature Requests
1. Review [Future Enhancements](COLLECTION-SUMMARY.md#-future-enhancements)
2. Check if already planned
3. Document your use case

## 🎯 Next Steps

**Ready to start?**
1. 📖 Read [QUICKSTART.md](QUICKSTART.md)
2. 💻 Install Bruno
3. 🚀 Start testing!

**Want to learn more?**
1. 📚 Read [hosterizer-api/README.md](hosterizer-api/README.md)
2. 📊 Review [TEST-FLOWS.md](hosterizer-api/TEST-FLOWS.md)
3. 📈 Check [COLLECTION-SUMMARY.md](COLLECTION-SUMMARY.md)

**Need help?**
1. 🔍 Search this index
2. 📖 Check relevant documentation
3. 🛠️ Review troubleshooting guides

---

**Last Updated**: 2024
**Collection Version**: 1.0
**Total Requests**: 12
**Total Documentation**: 2,500+ lines
