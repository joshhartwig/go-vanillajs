export const API = {
  baseUrl: "/api/",
  getTopMovies: async ()=> {
    const response = await fetch(`${API.baseUrl}movies/top`);
    const result = await response.json()
    return result
  },
  getRandomMovies: async ()=> {
    const response = await fetch(`${API.baseUrl}movies/random`);
    const result = await response.json()
    return result
  },
  getMoviesById: async (id)=> {
    const response = await fetch(`${API.baseUrl}movies/${id}`);
    const result = await response.json()
    return result
  },
  searchMovies: async (q, order, genre)=> {
    const response = await fetch(`${API.baseUrl}movies/search/?q=${encodeURIComponent(q)}`);
    const result = await response.json()
    return result
  },
  fetch: async (serviceName, args) => {
    try {
      const queryString = args ? new URLSearchParams(args).toString() : '';
      const response = await fetch(`${API.baseUrl}${serviceName}/?${queryString}`);
      if (!response.ok) {
        throw new Error(`API request failed with status ${response.status}`);
      }
      return await response.json();
    } catch (error) {
      console.error(error);
      throw error;
    }
  }

}