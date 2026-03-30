export const API = {
  getTopMovies: async ()=> {
    const response = await fetch("/api/movies/top/");
    const result = await response.json()
    return result
  },

}