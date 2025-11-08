import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
    return (
        <BrowserRouter>
            <div className="min-h-screen bg-gray-50">
                <Routes>
                    <Route path="/" element={<HomePage />} />
                </Routes>
            </div>
        </BrowserRouter>
    );
}

function HomePage() {
    return (
        <div className="flex items-center justify-center min-h-screen">
            <div className="text-center">
                <h1 className="text-4xl font-bold text-gray-900 mb-4">
                    Hosterizer Customer Portal
                </h1>
                <p className="text-lg text-gray-600">
                    Manage your cloud hosting deployments
                </p>
            </div>
        </div>
    );
}

export default App;
