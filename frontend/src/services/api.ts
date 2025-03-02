import axios, { type AxiosResponse } from "axios";

const API_URL = "http://localhost:8011";

class ApiService {
  async register(data: any): Promise<AxiosResponse> {
    const response = await axios.post(`${API_URL}/user`, data, {
      headers: {
        "Content-Type": "application/json",
      },
      validateStatus: () => true,
    });

    return response;
  }
}

export default new ApiService();
