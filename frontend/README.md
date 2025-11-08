# Hosterizer Frontend Applications

React-based frontend applications for the Hosterizer platform.

## Applications

### Admin Portal (Port 3000)
Web interface for platform administrators to manage customers, sites, costs, and policies.

**Features:**
- Customer management
- Site monitoring and management
- Cost reporting and analytics
- Policy management
- System observability integration

### Customer Portal (Port 3001)
Web interface for customers to manage their sites and deployments.

**Features:**
- Site creation and management
- Deployment tracking
- Cost dashboard
- Ecommerce platform integration
- White-label theming support

## Technology Stack

- **Framework**: React 18+ with TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **Routing**: React Router v6
- **State Management**: React Query + Context API
- **HTTP Client**: Axios

## Development Setup

### Prerequisites
- Node.js 18+
- npm or yarn

### Installation

```bash
# Install dependencies for admin portal
cd admin-portal
npm install

# Install dependencies for customer portal
cd ../customer-portal
npm install
```

### Running Development Servers

```bash
# Admin Portal (http://localhost:3000)
cd admin-portal
npm run dev

# Customer Portal (http://localhost:3001)
cd customer-portal
npm run dev
```

### Building for Production

```bash
# Build admin portal
cd admin-portal
npm run build

# Build customer portal
cd customer-portal
npm run build
```

### Running Tests

```bash
# Run tests
npm test

# Run tests with coverage
npm run test -- --coverage
```

## Project Structure

Each application follows this structure:

```
portal-name/
├── public/              # Static assets
├── src/
│   ├── components/      # React components
│   ├── pages/          # Page components
│   ├── hooks/          # Custom React hooks
│   ├── services/       # API services
│   ├── types/          # TypeScript types
│   ├── utils/          # Utility functions
│   ├── App.tsx         # Root component
│   ├── main.tsx        # Entry point
│   └── index.css       # Global styles
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── tailwind.config.js
```

## Environment Variables

Create a `.env` file in each application directory:

```
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=Hosterizer
```

## Code Standards

- Follow TypeScript strict mode
- Use functional components with hooks
- Use Tailwind utility classes for styling
- Implement proper error handling
- Write tests for critical functionality
- Use React Query for server state management
- Keep components small and focused
