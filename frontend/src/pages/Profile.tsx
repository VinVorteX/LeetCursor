import React from 'react';

export function Profile() {
  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">Profile</h1>
      <div className="bg-white shadow-md rounded-lg p-6">
        <h2 className="text-xl font-semibold mb-4">Your Profiles</h2>
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">LeetCode Username</label>
            <input type="text" className="mt-1 block w-full rounded-md border-gray-300 shadow-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">CodeForces Username</label>
            <input type="text" className="mt-1 block w-full rounded-md border-gray-300 shadow-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">CodeChef Username</label>
            <input type="text" className="mt-1 block w-full rounded-md border-gray-300 shadow-sm" />
          </div>
          <button className="bg-blue-500 text-white px-4 py-2 rounded-md">Save Changes</button>
        </div>
      </div>
    </div>
  );
}

export default Profile; 