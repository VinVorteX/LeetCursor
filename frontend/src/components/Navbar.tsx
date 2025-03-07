import { Link } from 'react-router-dom';

export function Navbar() {
  return (
    <nav className="bg-white shadow-md">
      <div className="container mx-auto px-4">
        <div className="flex justify-between h-16">
          <div className="flex">
            <Link to="/" className="flex items-center px-2 py-2 text-gray-700 hover:text-gray-900">
              LeetCursor
            </Link>
            <Link to="/leaderboard" className="flex items-center px-2 py-2 text-gray-700 hover:text-gray-900">
              Leaderboard
            </Link>
            <Link to="/profile" className="flex items-center px-2 py-2 text-gray-700 hover:text-gray-900">
              Profile
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
} 