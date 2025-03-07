import React, { useState, useEffect } from 'react';
import { useQuery } from '@tanstack/react-query';
import { getLeaderboard } from '../services/api';

export function Leaderboard() {
    const [platform, setPlatform] = useState('all');
    const [sortBy, setSortBy] = useState('totalSolved');
    const [page, setPage] = useState(1);
    
    const { data, isLoading, error } = useQuery({
        queryKey: ['leaderboard', platform, sortBy, page],
        queryFn: () => getLeaderboard({ platform, sortBy, page })
    });

    if (isLoading) return <div>Loading...</div>;
    if (error) return <div>Error loading leaderboard</div>;

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <h1 className="text-3xl font-bold">Leaderboard</h1>
                <div className="space-x-4">
                    <select
                        value={platform}
                        onChange={(e) => setPlatform(e.target.value)}
                        className="rounded-md border-gray-300"
                    >
                        <option value="all">All Platforms</option>
                        <option value="leetcode">LeetCode</option>
                        <option value="codeforces">Codeforces</option>
                        <option value="codechef">CodeChef</option>
                    </select>
                    <select
                        value={sortBy}
                        onChange={(e) => setSortBy(e.target.value)}
                        className="rounded-md border-gray-300"
                    >
                        <option value="totalSolved">Total Solved</option>
                        <option value="rating">Rating</option>
                    </select>
                </div>
            </div>

            <div className="bg-white shadow-md rounded-lg overflow-hidden">
                <table className="min-w-full">
                    <thead className="bg-gray-50">
                        <tr>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                Rank
                            </th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                User
                            </th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                Total Solved
                            </th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                Easy
                            </th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                Medium
                            </th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                Hard
                            </th>
                        </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                        {data?.users.map((user, index) => (
                            <tr key={user.id}>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {(page - 1) * 20 + index + 1}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    <div className="flex items-center">
                                        <img
                                            className="h-8 w-8 rounded-full"
                                            src={user.picture}
                                            alt=""
                                        />
                                        <div className="ml-4">
                                            <div className="text-sm font-medium text-gray-900">
                                                {user.name}
                                            </div>
                                        </div>
                                    </div>
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {user.totalSolved}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {user.easySolved}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {user.mediumSolved}
                                </td>
                                <td className="px-6 py-4 whitespace-nowrap">
                                    {user.hardSolved}
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>

            <div className="flex justify-center space-x-4">
                <button
                    onClick={() => setPage(p => Math.max(1, p - 1))}
                    disabled={page === 1}
                    className="px-4 py-2 border rounded-md disabled:opacity-50"
                >
                    Previous
                </button>
                <button
                    onClick={() => setPage(p => p + 1)}
                    disabled={!data?.users.length}
                    className="px-4 py-2 border rounded-md disabled:opacity-50"
                >
                    Next
                </button>
            </div>
        </div>
    );
} 