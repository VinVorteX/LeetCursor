interface LeaderboardParams {
  platform: string;
  sortBy: string;
  page: number;
}

export const getLeaderboard = async (params: LeaderboardParams) => {
  const response = await fetch(`${process.env.REACT_APP_API_URL}/api/leaderboard?` + 
    new URLSearchParams({
      platform: params.platform,
      sortBy: params.sortBy,
      page: params.page.toString()
    })
  );
  return response.json();
}; 